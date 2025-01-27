package seat

import (
	"PopcornMovie/config"
	"PopcornMovie/ent"
	"PopcornMovie/ent/showtime"
	"PopcornMovie/internal/utils"
	"PopcornMovie/model"
	"PopcornMovie/repository"
	"context"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Service interface {
	ListAvailableSeats(ctx context.Context, input model.ListAvailableSeatInput) ([]*ent.Seat, int, error)
}

type impl struct {
	repository repository.Registry
	logger     *zap.Logger
	appConfig  config.AppConfig
}

func (i impl) ListAvailableSeats(ctx context.Context, input model.ListAvailableSeatInput) ([]*ent.Seat, int, error) {
	var (
		showTimeID uuid.UUID
		err        error
	)
	if input.Filter != nil {
		showTimeID, err = uuid.Parse(*input.Filter.ShowTimeID)
		if err != nil {
			i.logger.Error(err.Error())
			return nil, 0, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
		}
	}

	query := i.repository.Seat().SeatQuery().WithRoom().WithTickets(
		func(ticketQuery *ent.TicketQuery) {
			ticketQuery.WithShowTime(func(showTimeQuery *ent.ShowTimeQuery) {
				showTimeQuery.Where(showtime.ID(showTimeID))
			})
		})

	count, err := i.repository.Seat().CountSeats(ctx, query)
	if err != nil {
		i.logger.Error(err.Error())
		return nil, 0, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
	}

	if input.Pagination != nil {
		offset := utils.CalculateOffset(input.Pagination.Page, input.Pagination.Limit)
		query.Limit(input.Pagination.Limit).Offset(offset)
	}

	seats, err := i.repository.Seat().GetAllSeats(ctx, query)
	if err != nil {
		i.logger.Error(err.Error())
		return nil, 0, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
	}

	return seats, *count, nil
}

func New(repository repository.Registry, logger *zap.Logger, appConfig config.AppConfig) Service {
	return &impl{
		repository: repository,
		logger:     logger,
		appConfig:  appConfig,
	}
}

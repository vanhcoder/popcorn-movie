package service

import (
	"PopcornMovie/config"
	"PopcornMovie/ent"
	"PopcornMovie/repository"
	"PopcornMovie/service/auth"
	"PopcornMovie/service/food"
	"PopcornMovie/service/movie"
	"PopcornMovie/service/room"
	"PopcornMovie/service/seat"
	"PopcornMovie/service/show_time"
	"PopcornMovie/service/theater"
	"PopcornMovie/service/transaction"
	"PopcornMovie/service/user"
	"go.uber.org/zap"
)

type Registry interface {
	User() user.Service
	Auth() auth.Service
	Theater() theater.Service
	Room() room.Service
	Food() food.Service
	Movie() movie.Service
	ShowTime() show_time.Service
	Seat() seat.Service
	Transaction() transaction.Service
}

type impl struct {
	user        user.Service
	auth        auth.Service
	theater     theater.Service
	room        room.Service
	food        food.Service
	movie       movie.Service
	showTime    show_time.Service
	seat        seat.Service
	transaction transaction.Service
}

func (i impl) Transaction() transaction.Service {
	return i.transaction
}

func (i impl) Seat() seat.Service {
	return i.seat
}

func (i impl) ShowTime() show_time.Service {
	return i.showTime
}

func (i impl) Movie() movie.Service {
	return i.movie
}

func (i impl) Food() food.Service {
	return i.food
}

func (i impl) Room() room.Service {
	return i.room
}

func (i impl) Theater() theater.Service { return i.theater }

func (i impl) User() user.Service {
	//TODO implement me
	return i.user
}

func (i impl) Auth() auth.Service {
	return i.auth
}

func New(entClient *ent.Client, logger *zap.Logger, appConfig config.AppConfig) Registry {
	repositoryRegistry := repository.New(entClient)

	return &impl{
		user:        user.New(repositoryRegistry, logger, appConfig),
		auth:        auth.New(repositoryRegistry, logger, appConfig),
		theater:     theater.New(repositoryRegistry, logger, appConfig),
		room:        room.New(repositoryRegistry, logger, appConfig),
		food:        food.New(repositoryRegistry, logger, appConfig),
		movie:       movie.New(repositoryRegistry, logger, appConfig),
		showTime:    show_time.New(repositoryRegistry, logger, appConfig),
		seat:        seat.New(repositoryRegistry, logger, appConfig),
		transaction: transaction.New(repositoryRegistry, logger, appConfig),
	}
}

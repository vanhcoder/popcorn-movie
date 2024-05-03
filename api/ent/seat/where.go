// Code generated by ent, DO NOT EDIT.

package seat

import (
	"PopcornMovie/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Seat {
	return predicate.Seat(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Seat {
	return predicate.Seat(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Seat {
	return predicate.Seat(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Seat {
	return predicate.Seat(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Seat {
	return predicate.Seat(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Seat {
	return predicate.Seat(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Seat {
	return predicate.Seat(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Seat {
	return predicate.Seat(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Seat {
	return predicate.Seat(sql.FieldLTE(FieldID, id))
}

// SeatNumber applies equality check predicate on the "seat_number" field. It's identical to SeatNumberEQ.
func SeatNumber(v string) predicate.Seat {
	return predicate.Seat(sql.FieldEQ(FieldSeatNumber, v))
}

// RoomID applies equality check predicate on the "room_id" field. It's identical to RoomIDEQ.
func RoomID(v uuid.UUID) predicate.Seat {
	return predicate.Seat(sql.FieldEQ(FieldRoomID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Seat {
	return predicate.Seat(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Seat {
	return predicate.Seat(sql.FieldEQ(FieldUpdatedAt, v))
}

// SeatNumberEQ applies the EQ predicate on the "seat_number" field.
func SeatNumberEQ(v string) predicate.Seat {
	return predicate.Seat(sql.FieldEQ(FieldSeatNumber, v))
}

// SeatNumberNEQ applies the NEQ predicate on the "seat_number" field.
func SeatNumberNEQ(v string) predicate.Seat {
	return predicate.Seat(sql.FieldNEQ(FieldSeatNumber, v))
}

// SeatNumberIn applies the In predicate on the "seat_number" field.
func SeatNumberIn(vs ...string) predicate.Seat {
	return predicate.Seat(sql.FieldIn(FieldSeatNumber, vs...))
}

// SeatNumberNotIn applies the NotIn predicate on the "seat_number" field.
func SeatNumberNotIn(vs ...string) predicate.Seat {
	return predicate.Seat(sql.FieldNotIn(FieldSeatNumber, vs...))
}

// SeatNumberGT applies the GT predicate on the "seat_number" field.
func SeatNumberGT(v string) predicate.Seat {
	return predicate.Seat(sql.FieldGT(FieldSeatNumber, v))
}

// SeatNumberGTE applies the GTE predicate on the "seat_number" field.
func SeatNumberGTE(v string) predicate.Seat {
	return predicate.Seat(sql.FieldGTE(FieldSeatNumber, v))
}

// SeatNumberLT applies the LT predicate on the "seat_number" field.
func SeatNumberLT(v string) predicate.Seat {
	return predicate.Seat(sql.FieldLT(FieldSeatNumber, v))
}

// SeatNumberLTE applies the LTE predicate on the "seat_number" field.
func SeatNumberLTE(v string) predicate.Seat {
	return predicate.Seat(sql.FieldLTE(FieldSeatNumber, v))
}

// SeatNumberContains applies the Contains predicate on the "seat_number" field.
func SeatNumberContains(v string) predicate.Seat {
	return predicate.Seat(sql.FieldContains(FieldSeatNumber, v))
}

// SeatNumberHasPrefix applies the HasPrefix predicate on the "seat_number" field.
func SeatNumberHasPrefix(v string) predicate.Seat {
	return predicate.Seat(sql.FieldHasPrefix(FieldSeatNumber, v))
}

// SeatNumberHasSuffix applies the HasSuffix predicate on the "seat_number" field.
func SeatNumberHasSuffix(v string) predicate.Seat {
	return predicate.Seat(sql.FieldHasSuffix(FieldSeatNumber, v))
}

// SeatNumberEqualFold applies the EqualFold predicate on the "seat_number" field.
func SeatNumberEqualFold(v string) predicate.Seat {
	return predicate.Seat(sql.FieldEqualFold(FieldSeatNumber, v))
}

// SeatNumberContainsFold applies the ContainsFold predicate on the "seat_number" field.
func SeatNumberContainsFold(v string) predicate.Seat {
	return predicate.Seat(sql.FieldContainsFold(FieldSeatNumber, v))
}

// RoomIDEQ applies the EQ predicate on the "room_id" field.
func RoomIDEQ(v uuid.UUID) predicate.Seat {
	return predicate.Seat(sql.FieldEQ(FieldRoomID, v))
}

// RoomIDNEQ applies the NEQ predicate on the "room_id" field.
func RoomIDNEQ(v uuid.UUID) predicate.Seat {
	return predicate.Seat(sql.FieldNEQ(FieldRoomID, v))
}

// RoomIDIn applies the In predicate on the "room_id" field.
func RoomIDIn(vs ...uuid.UUID) predicate.Seat {
	return predicate.Seat(sql.FieldIn(FieldRoomID, vs...))
}

// RoomIDNotIn applies the NotIn predicate on the "room_id" field.
func RoomIDNotIn(vs ...uuid.UUID) predicate.Seat {
	return predicate.Seat(sql.FieldNotIn(FieldRoomID, vs...))
}

// CategoryEQ applies the EQ predicate on the "category" field.
func CategoryEQ(v Category) predicate.Seat {
	return predicate.Seat(sql.FieldEQ(FieldCategory, v))
}

// CategoryNEQ applies the NEQ predicate on the "category" field.
func CategoryNEQ(v Category) predicate.Seat {
	return predicate.Seat(sql.FieldNEQ(FieldCategory, v))
}

// CategoryIn applies the In predicate on the "category" field.
func CategoryIn(vs ...Category) predicate.Seat {
	return predicate.Seat(sql.FieldIn(FieldCategory, vs...))
}

// CategoryNotIn applies the NotIn predicate on the "category" field.
func CategoryNotIn(vs ...Category) predicate.Seat {
	return predicate.Seat(sql.FieldNotIn(FieldCategory, vs...))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Seat {
	return predicate.Seat(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Seat {
	return predicate.Seat(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Seat {
	return predicate.Seat(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Seat {
	return predicate.Seat(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Seat {
	return predicate.Seat(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Seat {
	return predicate.Seat(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Seat {
	return predicate.Seat(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Seat {
	return predicate.Seat(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Seat {
	return predicate.Seat(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Seat {
	return predicate.Seat(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Seat {
	return predicate.Seat(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Seat {
	return predicate.Seat(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Seat {
	return predicate.Seat(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Seat {
	return predicate.Seat(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Seat {
	return predicate.Seat(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Seat {
	return predicate.Seat(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasRoom applies the HasEdge predicate on the "room" edge.
func HasRoom() predicate.Seat {
	return predicate.Seat(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RoomTable, RoomColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRoomWith applies the HasEdge predicate on the "room" edge with a given conditions (other predicates).
func HasRoomWith(preds ...predicate.Room) predicate.Seat {
	return predicate.Seat(func(s *sql.Selector) {
		step := newRoomStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTickets applies the HasEdge predicate on the "tickets" edge.
func HasTickets() predicate.Seat {
	return predicate.Seat(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TicketsTable, TicketsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTicketsWith applies the HasEdge predicate on the "tickets" edge with a given conditions (other predicates).
func HasTicketsWith(preds ...predicate.Ticket) predicate.Seat {
	return predicate.Seat(func(s *sql.Selector) {
		step := newTicketsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Seat) predicate.Seat {
	return predicate.Seat(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Seat) predicate.Seat {
	return predicate.Seat(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Seat) predicate.Seat {
	return predicate.Seat(sql.NotPredicates(p))
}
// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type CreateUserInput struct {
	DisplayName string `json:"displayName"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Role        *Role  `json:"role,omitempty"`
}

type Jwt struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Mutation struct {
}

type PageInfo struct {
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *string `json:"startCursor,omitempty"`
	EndCursor       *string `json:"endCursor,omitempty"`
}

type PaginationInput struct {
	After  *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	First  *int    `json:"first,omitempty"`
	Last   *int    `json:"last,omitempty"`
}

type Query struct {
}

type RegisterInput struct {
	DisplayName     string `json:"displayName"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

type User struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsLocked    bool   `json:"isLocked"`
	Role        Role   `json:"role"`
}

type OrderDirection string

const (
	OrderDirectionAsc  OrderDirection = "ASC"
	OrderDirectionDesc OrderDirection = "DESC"
)

var AllOrderDirection = []OrderDirection{
	OrderDirectionAsc,
	OrderDirectionDesc,
}

func (e OrderDirection) IsValid() bool {
	switch e {
	case OrderDirectionAsc, OrderDirectionDesc:
		return true
	}
	return false
}

func (e OrderDirection) String() string {
	return string(e)
}

func (e *OrderDirection) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrderDirection(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrderDirection", str)
	}
	return nil
}

func (e OrderDirection) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Role string

const (
	RoleCustomer      Role = "CUSTOMER"
	RoleReceptionist  Role = "RECEPTIONIST"
	RoleTicketManager Role = "TICKET_MANAGER"
	RoleAdmin         Role = "ADMIN"
)

var AllRole = []Role{
	RoleCustomer,
	RoleReceptionist,
	RoleTicketManager,
	RoleAdmin,
}

func (e Role) IsValid() bool {
	switch e {
	case RoleCustomer, RoleReceptionist, RoleTicketManager, RoleAdmin:
		return true
	}
	return false
}

func (e Role) String() string {
	return string(e)
}

func (e *Role) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Role(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}

func (e Role) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

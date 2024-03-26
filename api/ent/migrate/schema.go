// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "displayname", Type: field.TypeString, Size: 255},
		{Name: "email", Type: field.TypeString, Size: 255},
		{Name: "password", Type: field.TypeString, Size: 255},
		{Name: "is_locked", Type: field.TypeBool, Default: false},
		{Name: "role", Type: field.TypeEnum, Enums: []string{"CUSTOMER", "RECEPTIONIST", "TICKET_MANAGER", "ADMIN"}, Default: "CUSTOMER"},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		UsersTable,
	}
)

func init() {
}

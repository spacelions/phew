// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// PhonesColumns holds the columns for the "phones" table.
	PhonesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "number", Type: field.TypeString},
		{Name: "country_code", Type: field.TypeString, Default: "+86"},
		{Name: "phone_user", Type: field.TypeInt, Nullable: true},
	}
	// PhonesTable holds the schema information for the "phones" table.
	PhonesTable = &schema.Table{
		Name:       "phones",
		Columns:    PhonesColumns,
		PrimaryKey: []*schema.Column{PhonesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "phones_users_user",
				Columns:    []*schema.Column{PhonesColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "phone_country_code_number",
				Unique:  true,
				Columns: []*schema.Column{PhonesColumns[2], PhonesColumns[1]},
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "age", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "user_phone", Type: field.TypeInt, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_phones_phone",
				Columns:    []*schema.Column{UsersColumns[5]},
				RefColumns: []*schema.Column{PhonesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "user_username",
				Unique:  true,
				Columns: []*schema.Column{UsersColumns[1]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		PhonesTable,
		UsersTable,
	}
)

func init() {
	PhonesTable.ForeignKeys[0].RefTable = UsersTable
	UsersTable.ForeignKeys[0].RefTable = PhonesTable
}

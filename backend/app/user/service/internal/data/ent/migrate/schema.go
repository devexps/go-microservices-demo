// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// UserColumns holds the columns for the "user" table.
	UserColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true, Comment: "id"},
		{Name: "create_by", Type: field.TypeUint32, Nullable: true, Comment: "create_by"},
		{Name: "create_time", Type: field.TypeTime, Nullable: true, Comment: "create_time"},
		{Name: "update_time", Type: field.TypeTime, Nullable: true, Comment: "update_time"},
		{Name: "delete_time", Type: field.TypeTime, Nullable: true, Comment: "delete_time"},
		{Name: "status", Type: field.TypeEnum, Nullable: true, Comment: "status", Enums: []string{"OFF", "ON"}, Default: "ON"},
		{Name: "username", Type: field.TypeString, Unique: true, Nullable: true, Size: 50, Comment: "username"},
		{Name: "password", Type: field.TypeString, Nullable: true, Size: 255, Comment: "password"},
		{Name: "nick_name", Type: field.TypeString, Nullable: true, Size: 128, Comment: "nick_name"},
		{Name: "real_name", Type: field.TypeString, Nullable: true, Size: 128, Comment: "real_name"},
		{Name: "email", Type: field.TypeString, Nullable: true, Size: 127, Comment: "email"},
		{Name: "phone", Type: field.TypeString, Nullable: true, Size: 11, Comment: "phone"},
		{Name: "authority", Type: field.TypeEnum, Nullable: true, Comment: "authority", Enums: []string{"SYS_ADMIN", "CUSTOMER_USER", "GUEST_USER", "REFRESH_TOKEN"}, Default: "CUSTOMER_USER"},
	}
	// UserTable holds the schema information for the "user" table.
	UserTable = &schema.Table{
		Name:       "user",
		Columns:    UserColumns,
		PrimaryKey: []*schema.Column{UserColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "user_id",
				Unique:  false,
				Columns: []*schema.Column{UserColumns[0]},
			},
			{
				Name:    "user_id_username",
				Unique:  true,
				Columns: []*schema.Column{UserColumns[0], UserColumns[6]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		UserTable,
	}
)

func init() {
	UserTable.Annotation = &entsql.Annotation{
		Table:     "user",
		Charset:   "utf8mb4",
		Collation: "utf8mb4_bin",
	}
}
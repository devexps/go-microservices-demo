package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/devexps/go-utils/entgo/mixin"
	"regexp"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Annotations of the user
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "user",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").
			Comment("username").
			Unique().
			Immutable().
			Optional().
			Nillable().
			NotEmpty().
			MaxLen(50).
			Match(regexp.MustCompile("^[a-zA-Z0-9]{4,16}$")),
		field.String("password").
			Comment("password").
			Optional().
			Nillable().
			NotEmpty().
			MaxLen(255),
		field.String("nick_name").
			Comment("nick_name").
			MaxLen(128).
			Optional().
			Nillable(),
		field.String("real_name").
			Comment("real_name").
			MaxLen(128).
			Optional().
			Nillable(),
		field.String("email").
			Comment("email").
			MaxLen(127).
			Optional().
			Nillable(),
		field.String("phone").
			Comment("phone").
			MaxLen(11).
			Optional().
			Nillable(),
		field.Enum("authority").
			Comment("authority").
			Optional().
			Nillable().
			Values(
				"SYS_ADMIN",
				"CUSTOMER_USER",
				"GUEST_USER",
				"REFRESH_TOKEN",
			).
			Default("CUSTOMER_USER"),
	}
}

// Mixin of the User.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.NumberId{},
		mixin.CreateBy{},
		mixin.Time{},
		mixin.SwitchStatus{},
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

// Indexes of the User.
func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id", "username").Unique(),
	}
}

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/schema/field"
)

type UserID int

type UserState string

const (
	Enable  UserState = "enable"
	Disable UserState = "disable"
)

func (u UserState) Values() []string {
	return []string{
		string(Enable),
		string(Disable),
	}
}

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			GoType(UserID(0)),
		field.String("name").
			Default("ゲスト"),
		field.String("email").
			Optional().GoType(sql.NullString{}),
		field.Enum("state").GoType(UserState("")),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

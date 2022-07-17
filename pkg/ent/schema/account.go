package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	ent.Schema
}

// Fields of the Account.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Int64("amount").Default(0),
		field.Enum("currency_type").Values("RUB", "USD", "EUR"),
		field.Bool("frozen").Default(false),
		field.Int64("user_id"),
	}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("accounts").Unique().Required().Field("user_id"),
	}
}

package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Kyc holds the schema definition for the Kyc entity.
type Kyc struct {
	ent.Schema
}

// Fields of the Kyc.
func (Kyc) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("user_id", uuid.UUID{}),
		field.UUID("app_id", uuid.UUID{}),
		field.String("first_name"),
		field.String("last_name"),
		field.String("region"),
		field.String("card_type"),
		field.String("card_id"),
		field.String("front_card_img"),
		field.String("back_card_img"),
		field.String("user_handling_card_img"),
		field.Uint32("create_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.Uint32("update_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}).
			UpdateDefault(func() uint32 {
				return uint32(time.Now().Unix())
			}),
	}
}

// Edges of the Kyc.
func (Kyc) Edges() []ent.Edge {
	return nil
}

func (Kyc) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "app_id"),
		index.Fields("card_id", "card_type", "app_id"),
	}
}

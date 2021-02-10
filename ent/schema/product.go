package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique().NotEmpty().Immutable(),
		field.String("name").NotEmpty(),
		field.Int("net_price"),
		field.Int("sale_price"),
		field.Int("discount_rate"),
		field.String("link_url").NotEmpty(),
		field.String("image_url").NotEmpty(),
		field.String("company").NotEmpty(),
		field.Time("expired_at"),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("belongs", Category.Type).Ref("contains").Unique(),
	}
}

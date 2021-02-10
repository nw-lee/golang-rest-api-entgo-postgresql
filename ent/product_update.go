// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/quavious/golang-docker-forum/ent/category"
	"github.com/quavious/golang-docker-forum/ent/predicate"
	"github.com/quavious/golang-docker-forum/ent/product"
)

// ProductUpdate is the builder for updating Product entities.
type ProductUpdate struct {
	config
	hooks    []Hook
	mutation *ProductMutation
}

// Where adds a new predicate for the ProductUpdate builder.
func (pu *ProductUpdate) Where(ps ...predicate.Product) *ProductUpdate {
	pu.mutation.predicates = append(pu.mutation.predicates, ps...)
	return pu
}

// SetName sets the "name" field.
func (pu *ProductUpdate) SetName(s string) *ProductUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetNetPrice sets the "net_price" field.
func (pu *ProductUpdate) SetNetPrice(i int) *ProductUpdate {
	pu.mutation.ResetNetPrice()
	pu.mutation.SetNetPrice(i)
	return pu
}

// AddNetPrice adds i to the "net_price" field.
func (pu *ProductUpdate) AddNetPrice(i int) *ProductUpdate {
	pu.mutation.AddNetPrice(i)
	return pu
}

// SetSalePrice sets the "sale_price" field.
func (pu *ProductUpdate) SetSalePrice(i int) *ProductUpdate {
	pu.mutation.ResetSalePrice()
	pu.mutation.SetSalePrice(i)
	return pu
}

// AddSalePrice adds i to the "sale_price" field.
func (pu *ProductUpdate) AddSalePrice(i int) *ProductUpdate {
	pu.mutation.AddSalePrice(i)
	return pu
}

// SetDiscountRate sets the "discount_rate" field.
func (pu *ProductUpdate) SetDiscountRate(i int) *ProductUpdate {
	pu.mutation.ResetDiscountRate()
	pu.mutation.SetDiscountRate(i)
	return pu
}

// AddDiscountRate adds i to the "discount_rate" field.
func (pu *ProductUpdate) AddDiscountRate(i int) *ProductUpdate {
	pu.mutation.AddDiscountRate(i)
	return pu
}

// SetLinkURL sets the "link_url" field.
func (pu *ProductUpdate) SetLinkURL(s string) *ProductUpdate {
	pu.mutation.SetLinkURL(s)
	return pu
}

// SetImageURL sets the "image_url" field.
func (pu *ProductUpdate) SetImageURL(s string) *ProductUpdate {
	pu.mutation.SetImageURL(s)
	return pu
}

// SetCompany sets the "company" field.
func (pu *ProductUpdate) SetCompany(s string) *ProductUpdate {
	pu.mutation.SetCompany(s)
	return pu
}

// SetExpiredAt sets the "expired_at" field.
func (pu *ProductUpdate) SetExpiredAt(t time.Time) *ProductUpdate {
	pu.mutation.SetExpiredAt(t)
	return pu
}

// SetBelongsID sets the "belongs" edge to the Category entity by ID.
func (pu *ProductUpdate) SetBelongsID(id int) *ProductUpdate {
	pu.mutation.SetBelongsID(id)
	return pu
}

// SetNillableBelongsID sets the "belongs" edge to the Category entity by ID if the given value is not nil.
func (pu *ProductUpdate) SetNillableBelongsID(id *int) *ProductUpdate {
	if id != nil {
		pu = pu.SetBelongsID(*id)
	}
	return pu
}

// SetBelongs sets the "belongs" edge to the Category entity.
func (pu *ProductUpdate) SetBelongs(c *Category) *ProductUpdate {
	return pu.SetBelongsID(c.ID)
}

// Mutation returns the ProductMutation object of the builder.
func (pu *ProductUpdate) Mutation() *ProductMutation {
	return pu.mutation
}

// ClearBelongs clears the "belongs" edge to the Category entity.
func (pu *ProductUpdate) ClearBelongs() *ProductUpdate {
	pu.mutation.ClearBelongs()
	return pu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *ProductUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		if err = pu.check(); err != nil {
			return 0, err
		}
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProductMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pu.check(); err != nil {
				return 0, err
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ProductUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProductUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProductUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *ProductUpdate) check() error {
	if v, ok := pu.mutation.Name(); ok {
		if err := product.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := pu.mutation.LinkURL(); ok {
		if err := product.LinkURLValidator(v); err != nil {
			return &ValidationError{Name: "link_url", err: fmt.Errorf("ent: validator failed for field \"link_url\": %w", err)}
		}
	}
	if v, ok := pu.mutation.ImageURL(); ok {
		if err := product.ImageURLValidator(v); err != nil {
			return &ValidationError{Name: "image_url", err: fmt.Errorf("ent: validator failed for field \"image_url\": %w", err)}
		}
	}
	if v, ok := pu.mutation.Company(); ok {
		if err := product.CompanyValidator(v); err != nil {
			return &ValidationError{Name: "company", err: fmt.Errorf("ent: validator failed for field \"company\": %w", err)}
		}
	}
	return nil
}

func (pu *ProductUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   product.Table,
			Columns: product.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: product.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldName,
		})
	}
	if value, ok := pu.mutation.NetPrice(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: product.FieldNetPrice,
		})
	}
	if value, ok := pu.mutation.AddedNetPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: product.FieldNetPrice,
		})
	}
	if value, ok := pu.mutation.SalePrice(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: product.FieldSalePrice,
		})
	}
	if value, ok := pu.mutation.AddedSalePrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: product.FieldSalePrice,
		})
	}
	if value, ok := pu.mutation.DiscountRate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: product.FieldDiscountRate,
		})
	}
	if value, ok := pu.mutation.AddedDiscountRate(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: product.FieldDiscountRate,
		})
	}
	if value, ok := pu.mutation.LinkURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldLinkURL,
		})
	}
	if value, ok := pu.mutation.ImageURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldImageURL,
		})
	}
	if value, ok := pu.mutation.Company(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldCompany,
		})
	}
	if value, ok := pu.mutation.ExpiredAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: product.FieldExpiredAt,
		})
	}
	if pu.mutation.BelongsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   product.BelongsTable,
			Columns: []string{product.BelongsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.BelongsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   product.BelongsTable,
			Columns: []string{product.BelongsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{product.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ProductUpdateOne is the builder for updating a single Product entity.
type ProductUpdateOne struct {
	config
	hooks    []Hook
	mutation *ProductMutation
}

// SetName sets the "name" field.
func (puo *ProductUpdateOne) SetName(s string) *ProductUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetNetPrice sets the "net_price" field.
func (puo *ProductUpdateOne) SetNetPrice(i int) *ProductUpdateOne {
	puo.mutation.ResetNetPrice()
	puo.mutation.SetNetPrice(i)
	return puo
}

// AddNetPrice adds i to the "net_price" field.
func (puo *ProductUpdateOne) AddNetPrice(i int) *ProductUpdateOne {
	puo.mutation.AddNetPrice(i)
	return puo
}

// SetSalePrice sets the "sale_price" field.
func (puo *ProductUpdateOne) SetSalePrice(i int) *ProductUpdateOne {
	puo.mutation.ResetSalePrice()
	puo.mutation.SetSalePrice(i)
	return puo
}

// AddSalePrice adds i to the "sale_price" field.
func (puo *ProductUpdateOne) AddSalePrice(i int) *ProductUpdateOne {
	puo.mutation.AddSalePrice(i)
	return puo
}

// SetDiscountRate sets the "discount_rate" field.
func (puo *ProductUpdateOne) SetDiscountRate(i int) *ProductUpdateOne {
	puo.mutation.ResetDiscountRate()
	puo.mutation.SetDiscountRate(i)
	return puo
}

// AddDiscountRate adds i to the "discount_rate" field.
func (puo *ProductUpdateOne) AddDiscountRate(i int) *ProductUpdateOne {
	puo.mutation.AddDiscountRate(i)
	return puo
}

// SetLinkURL sets the "link_url" field.
func (puo *ProductUpdateOne) SetLinkURL(s string) *ProductUpdateOne {
	puo.mutation.SetLinkURL(s)
	return puo
}

// SetImageURL sets the "image_url" field.
func (puo *ProductUpdateOne) SetImageURL(s string) *ProductUpdateOne {
	puo.mutation.SetImageURL(s)
	return puo
}

// SetCompany sets the "company" field.
func (puo *ProductUpdateOne) SetCompany(s string) *ProductUpdateOne {
	puo.mutation.SetCompany(s)
	return puo
}

// SetExpiredAt sets the "expired_at" field.
func (puo *ProductUpdateOne) SetExpiredAt(t time.Time) *ProductUpdateOne {
	puo.mutation.SetExpiredAt(t)
	return puo
}

// SetBelongsID sets the "belongs" edge to the Category entity by ID.
func (puo *ProductUpdateOne) SetBelongsID(id int) *ProductUpdateOne {
	puo.mutation.SetBelongsID(id)
	return puo
}

// SetNillableBelongsID sets the "belongs" edge to the Category entity by ID if the given value is not nil.
func (puo *ProductUpdateOne) SetNillableBelongsID(id *int) *ProductUpdateOne {
	if id != nil {
		puo = puo.SetBelongsID(*id)
	}
	return puo
}

// SetBelongs sets the "belongs" edge to the Category entity.
func (puo *ProductUpdateOne) SetBelongs(c *Category) *ProductUpdateOne {
	return puo.SetBelongsID(c.ID)
}

// Mutation returns the ProductMutation object of the builder.
func (puo *ProductUpdateOne) Mutation() *ProductMutation {
	return puo.mutation
}

// ClearBelongs clears the "belongs" edge to the Category entity.
func (puo *ProductUpdateOne) ClearBelongs() *ProductUpdateOne {
	puo.mutation.ClearBelongs()
	return puo
}

// Save executes the query and returns the updated Product entity.
func (puo *ProductUpdateOne) Save(ctx context.Context) (*Product, error) {
	var (
		err  error
		node *Product
	)
	if len(puo.hooks) == 0 {
		if err = puo.check(); err != nil {
			return nil, err
		}
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProductMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = puo.check(); err != nil {
				return nil, err
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ProductUpdateOne) SaveX(ctx context.Context) *Product {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *ProductUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProductUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *ProductUpdateOne) check() error {
	if v, ok := puo.mutation.Name(); ok {
		if err := product.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := puo.mutation.LinkURL(); ok {
		if err := product.LinkURLValidator(v); err != nil {
			return &ValidationError{Name: "link_url", err: fmt.Errorf("ent: validator failed for field \"link_url\": %w", err)}
		}
	}
	if v, ok := puo.mutation.ImageURL(); ok {
		if err := product.ImageURLValidator(v); err != nil {
			return &ValidationError{Name: "image_url", err: fmt.Errorf("ent: validator failed for field \"image_url\": %w", err)}
		}
	}
	if v, ok := puo.mutation.Company(); ok {
		if err := product.CompanyValidator(v); err != nil {
			return &ValidationError{Name: "company", err: fmt.Errorf("ent: validator failed for field \"company\": %w", err)}
		}
	}
	return nil
}

func (puo *ProductUpdateOne) sqlSave(ctx context.Context) (_node *Product, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   product.Table,
			Columns: product.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: product.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Product.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := puo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldName,
		})
	}
	if value, ok := puo.mutation.NetPrice(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: product.FieldNetPrice,
		})
	}
	if value, ok := puo.mutation.AddedNetPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: product.FieldNetPrice,
		})
	}
	if value, ok := puo.mutation.SalePrice(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: product.FieldSalePrice,
		})
	}
	if value, ok := puo.mutation.AddedSalePrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: product.FieldSalePrice,
		})
	}
	if value, ok := puo.mutation.DiscountRate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: product.FieldDiscountRate,
		})
	}
	if value, ok := puo.mutation.AddedDiscountRate(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: product.FieldDiscountRate,
		})
	}
	if value, ok := puo.mutation.LinkURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldLinkURL,
		})
	}
	if value, ok := puo.mutation.ImageURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldImageURL,
		})
	}
	if value, ok := puo.mutation.Company(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldCompany,
		})
	}
	if value, ok := puo.mutation.ExpiredAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: product.FieldExpiredAt,
		})
	}
	if puo.mutation.BelongsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   product.BelongsTable,
			Columns: []string{product.BelongsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.BelongsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   product.BelongsTable,
			Columns: []string{product.BelongsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Product{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{product.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
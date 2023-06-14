// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"insightstream/ent/cooccurrencenetworkpool"
	"insightstream/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
)

// CooccurrenceNetworkPoolUpdate is the builder for updating CooccurrenceNetworkPool entities.
type CooccurrenceNetworkPoolUpdate struct {
	config
	hooks    []Hook
	mutation *CooccurrenceNetworkPoolMutation
}

// Where appends a list predicates to the CooccurrenceNetworkPoolUpdate builder.
func (cnpu *CooccurrenceNetworkPoolUpdate) Where(ps ...predicate.CooccurrenceNetworkPool) *CooccurrenceNetworkPoolUpdate {
	cnpu.mutation.Where(ps...)
	return cnpu
}

// SetSiteURL sets the "site_url" field.
func (cnpu *CooccurrenceNetworkPoolUpdate) SetSiteURL(s string) *CooccurrenceNetworkPoolUpdate {
	cnpu.mutation.SetSiteURL(s)
	return cnpu
}

// SetTitles sets the "titles" field.
func (cnpu *CooccurrenceNetworkPoolUpdate) SetTitles(s []string) *CooccurrenceNetworkPoolUpdate {
	cnpu.mutation.SetTitles(s)
	return cnpu
}

// AppendTitles appends s to the "titles" field.
func (cnpu *CooccurrenceNetworkPoolUpdate) AppendTitles(s []string) *CooccurrenceNetworkPoolUpdate {
	cnpu.mutation.AppendTitles(s)
	return cnpu
}

// SetDescriptions sets the "descriptions" field.
func (cnpu *CooccurrenceNetworkPoolUpdate) SetDescriptions(s []string) *CooccurrenceNetworkPoolUpdate {
	cnpu.mutation.SetDescriptions(s)
	return cnpu
}

// AppendDescriptions appends s to the "descriptions" field.
func (cnpu *CooccurrenceNetworkPoolUpdate) AppendDescriptions(s []string) *CooccurrenceNetworkPoolUpdate {
	cnpu.mutation.AppendDescriptions(s)
	return cnpu
}

// Mutation returns the CooccurrenceNetworkPoolMutation object of the builder.
func (cnpu *CooccurrenceNetworkPoolUpdate) Mutation() *CooccurrenceNetworkPoolMutation {
	return cnpu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cnpu *CooccurrenceNetworkPoolUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, CooccurrenceNetworkPoolMutation](ctx, cnpu.sqlSave, cnpu.mutation, cnpu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cnpu *CooccurrenceNetworkPoolUpdate) SaveX(ctx context.Context) int {
	affected, err := cnpu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cnpu *CooccurrenceNetworkPoolUpdate) Exec(ctx context.Context) error {
	_, err := cnpu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cnpu *CooccurrenceNetworkPoolUpdate) ExecX(ctx context.Context) {
	if err := cnpu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cnpu *CooccurrenceNetworkPoolUpdate) check() error {
	if v, ok := cnpu.mutation.SiteURL(); ok {
		if err := cooccurrencenetworkpool.SiteURLValidator(v); err != nil {
			return &ValidationError{Name: "site_url", err: fmt.Errorf(`ent: validator failed for field "CooccurrenceNetworkPool.site_url": %w`, err)}
		}
	}
	return nil
}

func (cnpu *CooccurrenceNetworkPoolUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cnpu.check(); err != nil {
		return n, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   cooccurrencenetworkpool.Table,
			Columns: cooccurrencenetworkpool.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: cooccurrencenetworkpool.FieldID,
			},
		},
	}
	if ps := cnpu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cnpu.mutation.SiteURL(); ok {
		_spec.SetField(cooccurrencenetworkpool.FieldSiteURL, field.TypeString, value)
	}
	if value, ok := cnpu.mutation.Titles(); ok {
		_spec.SetField(cooccurrencenetworkpool.FieldTitles, field.TypeJSON, value)
	}
	if value, ok := cnpu.mutation.AppendedTitles(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, cooccurrencenetworkpool.FieldTitles, value)
		})
	}
	if value, ok := cnpu.mutation.Descriptions(); ok {
		_spec.SetField(cooccurrencenetworkpool.FieldDescriptions, field.TypeJSON, value)
	}
	if value, ok := cnpu.mutation.AppendedDescriptions(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, cooccurrencenetworkpool.FieldDescriptions, value)
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cnpu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cooccurrencenetworkpool.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cnpu.mutation.done = true
	return n, nil
}

// CooccurrenceNetworkPoolUpdateOne is the builder for updating a single CooccurrenceNetworkPool entity.
type CooccurrenceNetworkPoolUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CooccurrenceNetworkPoolMutation
}

// SetSiteURL sets the "site_url" field.
func (cnpuo *CooccurrenceNetworkPoolUpdateOne) SetSiteURL(s string) *CooccurrenceNetworkPoolUpdateOne {
	cnpuo.mutation.SetSiteURL(s)
	return cnpuo
}

// SetTitles sets the "titles" field.
func (cnpuo *CooccurrenceNetworkPoolUpdateOne) SetTitles(s []string) *CooccurrenceNetworkPoolUpdateOne {
	cnpuo.mutation.SetTitles(s)
	return cnpuo
}

// AppendTitles appends s to the "titles" field.
func (cnpuo *CooccurrenceNetworkPoolUpdateOne) AppendTitles(s []string) *CooccurrenceNetworkPoolUpdateOne {
	cnpuo.mutation.AppendTitles(s)
	return cnpuo
}

// SetDescriptions sets the "descriptions" field.
func (cnpuo *CooccurrenceNetworkPoolUpdateOne) SetDescriptions(s []string) *CooccurrenceNetworkPoolUpdateOne {
	cnpuo.mutation.SetDescriptions(s)
	return cnpuo
}

// AppendDescriptions appends s to the "descriptions" field.
func (cnpuo *CooccurrenceNetworkPoolUpdateOne) AppendDescriptions(s []string) *CooccurrenceNetworkPoolUpdateOne {
	cnpuo.mutation.AppendDescriptions(s)
	return cnpuo
}

// Mutation returns the CooccurrenceNetworkPoolMutation object of the builder.
func (cnpuo *CooccurrenceNetworkPoolUpdateOne) Mutation() *CooccurrenceNetworkPoolMutation {
	return cnpuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cnpuo *CooccurrenceNetworkPoolUpdateOne) Select(field string, fields ...string) *CooccurrenceNetworkPoolUpdateOne {
	cnpuo.fields = append([]string{field}, fields...)
	return cnpuo
}

// Save executes the query and returns the updated CooccurrenceNetworkPool entity.
func (cnpuo *CooccurrenceNetworkPoolUpdateOne) Save(ctx context.Context) (*CooccurrenceNetworkPool, error) {
	return withHooks[*CooccurrenceNetworkPool, CooccurrenceNetworkPoolMutation](ctx, cnpuo.sqlSave, cnpuo.mutation, cnpuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cnpuo *CooccurrenceNetworkPoolUpdateOne) SaveX(ctx context.Context) *CooccurrenceNetworkPool {
	node, err := cnpuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cnpuo *CooccurrenceNetworkPoolUpdateOne) Exec(ctx context.Context) error {
	_, err := cnpuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cnpuo *CooccurrenceNetworkPoolUpdateOne) ExecX(ctx context.Context) {
	if err := cnpuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cnpuo *CooccurrenceNetworkPoolUpdateOne) check() error {
	if v, ok := cnpuo.mutation.SiteURL(); ok {
		if err := cooccurrencenetworkpool.SiteURLValidator(v); err != nil {
			return &ValidationError{Name: "site_url", err: fmt.Errorf(`ent: validator failed for field "CooccurrenceNetworkPool.site_url": %w`, err)}
		}
	}
	return nil
}

func (cnpuo *CooccurrenceNetworkPoolUpdateOne) sqlSave(ctx context.Context) (_node *CooccurrenceNetworkPool, err error) {
	if err := cnpuo.check(); err != nil {
		return _node, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   cooccurrencenetworkpool.Table,
			Columns: cooccurrencenetworkpool.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: cooccurrencenetworkpool.FieldID,
			},
		},
	}
	id, ok := cnpuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CooccurrenceNetworkPool.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cnpuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, cooccurrencenetworkpool.FieldID)
		for _, f := range fields {
			if !cooccurrencenetworkpool.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != cooccurrencenetworkpool.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cnpuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cnpuo.mutation.SiteURL(); ok {
		_spec.SetField(cooccurrencenetworkpool.FieldSiteURL, field.TypeString, value)
	}
	if value, ok := cnpuo.mutation.Titles(); ok {
		_spec.SetField(cooccurrencenetworkpool.FieldTitles, field.TypeJSON, value)
	}
	if value, ok := cnpuo.mutation.AppendedTitles(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, cooccurrencenetworkpool.FieldTitles, value)
		})
	}
	if value, ok := cnpuo.mutation.Descriptions(); ok {
		_spec.SetField(cooccurrencenetworkpool.FieldDescriptions, field.TypeJSON, value)
	}
	if value, ok := cnpuo.mutation.AppendedDescriptions(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, cooccurrencenetworkpool.FieldDescriptions, value)
		})
	}
	_node = &CooccurrenceNetworkPool{config: cnpuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cnpuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cooccurrencenetworkpool.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cnpuo.mutation.done = true
	return _node, nil
}

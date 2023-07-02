// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"insightstream/ent/feedaudittrailaction"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FeedAuditTrailActionCreate is the builder for creating a FeedAuditTrailAction entity.
type FeedAuditTrailActionCreate struct {
	config
	mutation *FeedAuditTrailActionMutation
	hooks    []Hook
}

// SetAction sets the "action" field.
func (fatac *FeedAuditTrailActionCreate) SetAction(s string) *FeedAuditTrailActionCreate {
	fatac.mutation.SetAction(s)
	return fatac
}

// SetID sets the "id" field.
func (fatac *FeedAuditTrailActionCreate) SetID(i int) *FeedAuditTrailActionCreate {
	fatac.mutation.SetID(i)
	return fatac
}

// Mutation returns the FeedAuditTrailActionMutation object of the builder.
func (fatac *FeedAuditTrailActionCreate) Mutation() *FeedAuditTrailActionMutation {
	return fatac.mutation
}

// Save creates the FeedAuditTrailAction in the database.
func (fatac *FeedAuditTrailActionCreate) Save(ctx context.Context) (*FeedAuditTrailAction, error) {
	return withHooks[*FeedAuditTrailAction, FeedAuditTrailActionMutation](ctx, fatac.sqlSave, fatac.mutation, fatac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (fatac *FeedAuditTrailActionCreate) SaveX(ctx context.Context) *FeedAuditTrailAction {
	v, err := fatac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fatac *FeedAuditTrailActionCreate) Exec(ctx context.Context) error {
	_, err := fatac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fatac *FeedAuditTrailActionCreate) ExecX(ctx context.Context) {
	if err := fatac.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fatac *FeedAuditTrailActionCreate) check() error {
	if _, ok := fatac.mutation.Action(); !ok {
		return &ValidationError{Name: "action", err: errors.New(`ent: missing required field "FeedAuditTrailAction.action"`)}
	}
	if v, ok := fatac.mutation.Action(); ok {
		if err := feedaudittrailaction.ActionValidator(v); err != nil {
			return &ValidationError{Name: "action", err: fmt.Errorf(`ent: validator failed for field "FeedAuditTrailAction.action": %w`, err)}
		}
	}
	return nil
}

func (fatac *FeedAuditTrailActionCreate) sqlSave(ctx context.Context) (*FeedAuditTrailAction, error) {
	if err := fatac.check(); err != nil {
		return nil, err
	}
	_node, _spec := fatac.createSpec()
	if err := sqlgraph.CreateNode(ctx, fatac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	fatac.mutation.id = &_node.ID
	fatac.mutation.done = true
	return _node, nil
}

func (fatac *FeedAuditTrailActionCreate) createSpec() (*FeedAuditTrailAction, *sqlgraph.CreateSpec) {
	var (
		_node = &FeedAuditTrailAction{config: fatac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: feedaudittrailaction.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: feedaudittrailaction.FieldID,
			},
		}
	)
	if id, ok := fatac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := fatac.mutation.Action(); ok {
		_spec.SetField(feedaudittrailaction.FieldAction, field.TypeString, value)
		_node.Action = value
	}
	return _node, _spec
}

// FeedAuditTrailActionCreateBulk is the builder for creating many FeedAuditTrailAction entities in bulk.
type FeedAuditTrailActionCreateBulk struct {
	config
	builders []*FeedAuditTrailActionCreate
}

// Save creates the FeedAuditTrailAction entities in the database.
func (fatacb *FeedAuditTrailActionCreateBulk) Save(ctx context.Context) ([]*FeedAuditTrailAction, error) {
	specs := make([]*sqlgraph.CreateSpec, len(fatacb.builders))
	nodes := make([]*FeedAuditTrailAction, len(fatacb.builders))
	mutators := make([]Mutator, len(fatacb.builders))
	for i := range fatacb.builders {
		func(i int, root context.Context) {
			builder := fatacb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FeedAuditTrailActionMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, fatacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fatacb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, fatacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fatacb *FeedAuditTrailActionCreateBulk) SaveX(ctx context.Context) []*FeedAuditTrailAction {
	v, err := fatacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fatacb *FeedAuditTrailActionCreateBulk) Exec(ctx context.Context) error {
	_, err := fatacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fatacb *FeedAuditTrailActionCreateBulk) ExecX(ctx context.Context) {
	if err := fatacb.Exec(ctx); err != nil {
		panic(err)
	}
}

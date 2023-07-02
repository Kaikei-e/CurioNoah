// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"insightstream/ent/feedaudittrailaction"
	"insightstream/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FeedAuditTrailActionDelete is the builder for deleting a FeedAuditTrailAction entity.
type FeedAuditTrailActionDelete struct {
	config
	hooks    []Hook
	mutation *FeedAuditTrailActionMutation
}

// Where appends a list predicates to the FeedAuditTrailActionDelete builder.
func (fatad *FeedAuditTrailActionDelete) Where(ps ...predicate.FeedAuditTrailAction) *FeedAuditTrailActionDelete {
	fatad.mutation.Where(ps...)
	return fatad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (fatad *FeedAuditTrailActionDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, FeedAuditTrailActionMutation](ctx, fatad.sqlExec, fatad.mutation, fatad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (fatad *FeedAuditTrailActionDelete) ExecX(ctx context.Context) int {
	n, err := fatad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (fatad *FeedAuditTrailActionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: feedaudittrailaction.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: feedaudittrailaction.FieldID,
			},
		},
	}
	if ps := fatad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, fatad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	fatad.mutation.done = true
	return affected, err
}

// FeedAuditTrailActionDeleteOne is the builder for deleting a single FeedAuditTrailAction entity.
type FeedAuditTrailActionDeleteOne struct {
	fatad *FeedAuditTrailActionDelete
}

// Exec executes the deletion query.
func (fatado *FeedAuditTrailActionDeleteOne) Exec(ctx context.Context) error {
	n, err := fatado.fatad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{feedaudittrailaction.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (fatado *FeedAuditTrailActionDeleteOne) ExecX(ctx context.Context) {
	fatado.fatad.ExecX(ctx)
}

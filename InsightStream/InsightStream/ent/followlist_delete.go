// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"insightstream/ent/followlist"
	"insightstream/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FollowListDelete is the builder for deleting a FollowList entity.
type FollowListDelete struct {
	config
	hooks    []Hook
	mutation *FollowListMutation
}

// Where appends a list predicates to the FollowListDelete builder.
func (fld *FollowListDelete) Where(ps ...predicate.FollowList) *FollowListDelete {
	fld.mutation.Where(ps...)
	return fld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (fld *FollowListDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, FollowListMutation](ctx, fld.sqlExec, fld.mutation, fld.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (fld *FollowListDelete) ExecX(ctx context.Context) int {
	n, err := fld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (fld *FollowListDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: followlist.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: followlist.FieldID,
			},
		},
	}
	if ps := fld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, fld.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	fld.mutation.done = true
	return affected, err
}

// FollowListDeleteOne is the builder for deleting a single FollowList entity.
type FollowListDeleteOne struct {
	fld *FollowListDelete
}

// Exec executes the deletion query.
func (fldo *FollowListDeleteOne) Exec(ctx context.Context) error {
	n, err := fldo.fld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{followlist.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (fldo *FollowListDeleteOne) ExecX(ctx context.Context) {
	fldo.fld.ExecX(ctx)
}
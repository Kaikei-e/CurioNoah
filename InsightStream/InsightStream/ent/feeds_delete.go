// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"insightstream/ent/feeds"
	"insightstream/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FeedsDelete is the builder for deleting a Feeds entity.
type FeedsDelete struct {
	config
	hooks    []Hook
	mutation *FeedsMutation
}

// Where appends a list predicates to the FeedsDelete builder.
func (fd *FeedsDelete) Where(ps ...predicate.Feeds) *FeedsDelete {
	fd.mutation.Where(ps...)
	return fd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (fd *FeedsDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, fd.sqlExec, fd.mutation, fd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (fd *FeedsDelete) ExecX(ctx context.Context) int {
	n, err := fd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (fd *FeedsDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(feeds.Table, sqlgraph.NewFieldSpec(feeds.FieldID, field.TypeUUID))
	if ps := fd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, fd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	fd.mutation.done = true
	return affected, err
}

// FeedsDeleteOne is the builder for deleting a single Feeds entity.
type FeedsDeleteOne struct {
	fd *FeedsDelete
}

// Where appends a list predicates to the FeedsDelete builder.
func (fdo *FeedsDeleteOne) Where(ps ...predicate.Feeds) *FeedsDeleteOne {
	fdo.fd.mutation.Where(ps...)
	return fdo
}

// Exec executes the deletion query.
func (fdo *FeedsDeleteOne) Exec(ctx context.Context) error {
	n, err := fdo.fd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{feeds.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (fdo *FeedsDeleteOne) ExecX(ctx context.Context) {
	if err := fdo.Exec(ctx); err != nil {
		panic(err)
	}
}

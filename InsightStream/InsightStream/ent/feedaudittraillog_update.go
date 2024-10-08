// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"insightstream/ent/feedaudittrailaction"
	"insightstream/ent/feedaudittraillog"
	"insightstream/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FeedAuditTrailLogUpdate is the builder for updating FeedAuditTrailLog entities.
type FeedAuditTrailLogUpdate struct {
	config
	hooks    []Hook
	mutation *FeedAuditTrailLogMutation
}

// Where appends a list predicates to the FeedAuditTrailLogUpdate builder.
func (fatlu *FeedAuditTrailLogUpdate) Where(ps ...predicate.FeedAuditTrailLog) *FeedAuditTrailLogUpdate {
	fatlu.mutation.Where(ps...)
	return fatlu
}

// SetUpdatedAt sets the "updated_at" field.
func (fatlu *FeedAuditTrailLogUpdate) SetUpdatedAt(t time.Time) *FeedAuditTrailLogUpdate {
	fatlu.mutation.SetUpdatedAt(t)
	return fatlu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (fatlu *FeedAuditTrailLogUpdate) ClearUpdatedAt() *FeedAuditTrailLogUpdate {
	fatlu.mutation.ClearUpdatedAt()
	return fatlu
}

// SetActionID sets the "action" edge to the FeedAuditTrailAction entity by ID.
func (fatlu *FeedAuditTrailLogUpdate) SetActionID(id int) *FeedAuditTrailLogUpdate {
	fatlu.mutation.SetActionID(id)
	return fatlu
}

// SetAction sets the "action" edge to the FeedAuditTrailAction entity.
func (fatlu *FeedAuditTrailLogUpdate) SetAction(f *FeedAuditTrailAction) *FeedAuditTrailLogUpdate {
	return fatlu.SetActionID(f.ID)
}

// Mutation returns the FeedAuditTrailLogMutation object of the builder.
func (fatlu *FeedAuditTrailLogUpdate) Mutation() *FeedAuditTrailLogMutation {
	return fatlu.mutation
}

// ClearAction clears the "action" edge to the FeedAuditTrailAction entity.
func (fatlu *FeedAuditTrailLogUpdate) ClearAction() *FeedAuditTrailLogUpdate {
	fatlu.mutation.ClearAction()
	return fatlu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fatlu *FeedAuditTrailLogUpdate) Save(ctx context.Context) (int, error) {
	fatlu.defaults()
	return withHooks(ctx, fatlu.sqlSave, fatlu.mutation, fatlu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fatlu *FeedAuditTrailLogUpdate) SaveX(ctx context.Context) int {
	affected, err := fatlu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fatlu *FeedAuditTrailLogUpdate) Exec(ctx context.Context) error {
	_, err := fatlu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fatlu *FeedAuditTrailLogUpdate) ExecX(ctx context.Context) {
	if err := fatlu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fatlu *FeedAuditTrailLogUpdate) defaults() {
	if _, ok := fatlu.mutation.UpdatedAt(); !ok && !fatlu.mutation.UpdatedAtCleared() {
		v := feedaudittraillog.UpdateDefaultUpdatedAt()
		fatlu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fatlu *FeedAuditTrailLogUpdate) check() error {
	if fatlu.mutation.ActionCleared() && len(fatlu.mutation.ActionIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "FeedAuditTrailLog.action"`)
	}
	return nil
}

func (fatlu *FeedAuditTrailLogUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := fatlu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(feedaudittraillog.Table, feedaudittraillog.Columns, sqlgraph.NewFieldSpec(feedaudittraillog.FieldID, field.TypeInt))
	if ps := fatlu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fatlu.mutation.UpdatedAt(); ok {
		_spec.SetField(feedaudittraillog.FieldUpdatedAt, field.TypeTime, value)
	}
	if fatlu.mutation.UpdatedAtCleared() {
		_spec.ClearField(feedaudittraillog.FieldUpdatedAt, field.TypeTime)
	}
	if fatlu.mutation.ActionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   feedaudittraillog.ActionTable,
			Columns: []string{feedaudittraillog.ActionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(feedaudittrailaction.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fatlu.mutation.ActionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   feedaudittraillog.ActionTable,
			Columns: []string{feedaudittraillog.ActionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(feedaudittrailaction.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, fatlu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{feedaudittraillog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	fatlu.mutation.done = true
	return n, nil
}

// FeedAuditTrailLogUpdateOne is the builder for updating a single FeedAuditTrailLog entity.
type FeedAuditTrailLogUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FeedAuditTrailLogMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (fatluo *FeedAuditTrailLogUpdateOne) SetUpdatedAt(t time.Time) *FeedAuditTrailLogUpdateOne {
	fatluo.mutation.SetUpdatedAt(t)
	return fatluo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (fatluo *FeedAuditTrailLogUpdateOne) ClearUpdatedAt() *FeedAuditTrailLogUpdateOne {
	fatluo.mutation.ClearUpdatedAt()
	return fatluo
}

// SetActionID sets the "action" edge to the FeedAuditTrailAction entity by ID.
func (fatluo *FeedAuditTrailLogUpdateOne) SetActionID(id int) *FeedAuditTrailLogUpdateOne {
	fatluo.mutation.SetActionID(id)
	return fatluo
}

// SetAction sets the "action" edge to the FeedAuditTrailAction entity.
func (fatluo *FeedAuditTrailLogUpdateOne) SetAction(f *FeedAuditTrailAction) *FeedAuditTrailLogUpdateOne {
	return fatluo.SetActionID(f.ID)
}

// Mutation returns the FeedAuditTrailLogMutation object of the builder.
func (fatluo *FeedAuditTrailLogUpdateOne) Mutation() *FeedAuditTrailLogMutation {
	return fatluo.mutation
}

// ClearAction clears the "action" edge to the FeedAuditTrailAction entity.
func (fatluo *FeedAuditTrailLogUpdateOne) ClearAction() *FeedAuditTrailLogUpdateOne {
	fatluo.mutation.ClearAction()
	return fatluo
}

// Where appends a list predicates to the FeedAuditTrailLogUpdate builder.
func (fatluo *FeedAuditTrailLogUpdateOne) Where(ps ...predicate.FeedAuditTrailLog) *FeedAuditTrailLogUpdateOne {
	fatluo.mutation.Where(ps...)
	return fatluo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fatluo *FeedAuditTrailLogUpdateOne) Select(field string, fields ...string) *FeedAuditTrailLogUpdateOne {
	fatluo.fields = append([]string{field}, fields...)
	return fatluo
}

// Save executes the query and returns the updated FeedAuditTrailLog entity.
func (fatluo *FeedAuditTrailLogUpdateOne) Save(ctx context.Context) (*FeedAuditTrailLog, error) {
	fatluo.defaults()
	return withHooks(ctx, fatluo.sqlSave, fatluo.mutation, fatluo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fatluo *FeedAuditTrailLogUpdateOne) SaveX(ctx context.Context) *FeedAuditTrailLog {
	node, err := fatluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fatluo *FeedAuditTrailLogUpdateOne) Exec(ctx context.Context) error {
	_, err := fatluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fatluo *FeedAuditTrailLogUpdateOne) ExecX(ctx context.Context) {
	if err := fatluo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fatluo *FeedAuditTrailLogUpdateOne) defaults() {
	if _, ok := fatluo.mutation.UpdatedAt(); !ok && !fatluo.mutation.UpdatedAtCleared() {
		v := feedaudittraillog.UpdateDefaultUpdatedAt()
		fatluo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fatluo *FeedAuditTrailLogUpdateOne) check() error {
	if fatluo.mutation.ActionCleared() && len(fatluo.mutation.ActionIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "FeedAuditTrailLog.action"`)
	}
	return nil
}

func (fatluo *FeedAuditTrailLogUpdateOne) sqlSave(ctx context.Context) (_node *FeedAuditTrailLog, err error) {
	if err := fatluo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(feedaudittraillog.Table, feedaudittraillog.Columns, sqlgraph.NewFieldSpec(feedaudittraillog.FieldID, field.TypeInt))
	id, ok := fatluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "FeedAuditTrailLog.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fatluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, feedaudittraillog.FieldID)
		for _, f := range fields {
			if !feedaudittraillog.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != feedaudittraillog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fatluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fatluo.mutation.UpdatedAt(); ok {
		_spec.SetField(feedaudittraillog.FieldUpdatedAt, field.TypeTime, value)
	}
	if fatluo.mutation.UpdatedAtCleared() {
		_spec.ClearField(feedaudittraillog.FieldUpdatedAt, field.TypeTime)
	}
	if fatluo.mutation.ActionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   feedaudittraillog.ActionTable,
			Columns: []string{feedaudittraillog.ActionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(feedaudittrailaction.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fatluo.mutation.ActionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   feedaudittraillog.ActionTable,
			Columns: []string{feedaudittraillog.ActionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(feedaudittrailaction.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &FeedAuditTrailLog{config: fatluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fatluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{feedaudittraillog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	fatluo.mutation.done = true
	return _node, nil
}

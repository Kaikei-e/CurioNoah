// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"insightstream/ent/feedaudittrailaction"
	"insightstream/ent/feedaudittraillog"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FeedAuditTrailLogCreate is the builder for creating a FeedAuditTrailLog entity.
type FeedAuditTrailLogCreate struct {
	config
	mutation *FeedAuditTrailLogMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetUpdatedAt sets the "updated_at" field.
func (fatlc *FeedAuditTrailLogCreate) SetUpdatedAt(t time.Time) *FeedAuditTrailLogCreate {
	fatlc.mutation.SetUpdatedAt(t)
	return fatlc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (fatlc *FeedAuditTrailLogCreate) SetNillableUpdatedAt(t *time.Time) *FeedAuditTrailLogCreate {
	if t != nil {
		fatlc.SetUpdatedAt(*t)
	}
	return fatlc
}

// SetID sets the "id" field.
func (fatlc *FeedAuditTrailLogCreate) SetID(i int) *FeedAuditTrailLogCreate {
	fatlc.mutation.SetID(i)
	return fatlc
}

// SetActionID sets the "action" edge to the FeedAuditTrailAction entity by ID.
func (fatlc *FeedAuditTrailLogCreate) SetActionID(id int) *FeedAuditTrailLogCreate {
	fatlc.mutation.SetActionID(id)
	return fatlc
}

// SetAction sets the "action" edge to the FeedAuditTrailAction entity.
func (fatlc *FeedAuditTrailLogCreate) SetAction(f *FeedAuditTrailAction) *FeedAuditTrailLogCreate {
	return fatlc.SetActionID(f.ID)
}

// Mutation returns the FeedAuditTrailLogMutation object of the builder.
func (fatlc *FeedAuditTrailLogCreate) Mutation() *FeedAuditTrailLogMutation {
	return fatlc.mutation
}

// Save creates the FeedAuditTrailLog in the database.
func (fatlc *FeedAuditTrailLogCreate) Save(ctx context.Context) (*FeedAuditTrailLog, error) {
	fatlc.defaults()
	return withHooks(ctx, fatlc.sqlSave, fatlc.mutation, fatlc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (fatlc *FeedAuditTrailLogCreate) SaveX(ctx context.Context) *FeedAuditTrailLog {
	v, err := fatlc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fatlc *FeedAuditTrailLogCreate) Exec(ctx context.Context) error {
	_, err := fatlc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fatlc *FeedAuditTrailLogCreate) ExecX(ctx context.Context) {
	if err := fatlc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fatlc *FeedAuditTrailLogCreate) defaults() {
	if _, ok := fatlc.mutation.UpdatedAt(); !ok {
		v := feedaudittraillog.DefaultUpdatedAt()
		fatlc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fatlc *FeedAuditTrailLogCreate) check() error {
	if _, ok := fatlc.mutation.ActionID(); !ok {
		return &ValidationError{Name: "action", err: errors.New(`ent: missing required edge "FeedAuditTrailLog.action"`)}
	}
	return nil
}

func (fatlc *FeedAuditTrailLogCreate) sqlSave(ctx context.Context) (*FeedAuditTrailLog, error) {
	if err := fatlc.check(); err != nil {
		return nil, err
	}
	_node, _spec := fatlc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fatlc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	fatlc.mutation.id = &_node.ID
	fatlc.mutation.done = true
	return _node, nil
}

func (fatlc *FeedAuditTrailLogCreate) createSpec() (*FeedAuditTrailLog, *sqlgraph.CreateSpec) {
	var (
		_node = &FeedAuditTrailLog{config: fatlc.config}
		_spec = sqlgraph.NewCreateSpec(feedaudittraillog.Table, sqlgraph.NewFieldSpec(feedaudittraillog.FieldID, field.TypeInt))
	)
	_spec.OnConflict = fatlc.conflict
	if id, ok := fatlc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := fatlc.mutation.UpdatedAt(); ok {
		_spec.SetField(feedaudittraillog.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := fatlc.mutation.ActionIDs(); len(nodes) > 0 {
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
		_node.action_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.FeedAuditTrailLog.Create().
//		SetUpdatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FeedAuditTrailLogUpsert) {
//			SetUpdatedAt(v+v).
//		}).
//		Exec(ctx)
func (fatlc *FeedAuditTrailLogCreate) OnConflict(opts ...sql.ConflictOption) *FeedAuditTrailLogUpsertOne {
	fatlc.conflict = opts
	return &FeedAuditTrailLogUpsertOne{
		create: fatlc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.FeedAuditTrailLog.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (fatlc *FeedAuditTrailLogCreate) OnConflictColumns(columns ...string) *FeedAuditTrailLogUpsertOne {
	fatlc.conflict = append(fatlc.conflict, sql.ConflictColumns(columns...))
	return &FeedAuditTrailLogUpsertOne{
		create: fatlc,
	}
}

type (
	// FeedAuditTrailLogUpsertOne is the builder for "upsert"-ing
	//  one FeedAuditTrailLog node.
	FeedAuditTrailLogUpsertOne struct {
		create *FeedAuditTrailLogCreate
	}

	// FeedAuditTrailLogUpsert is the "OnConflict" setter.
	FeedAuditTrailLogUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *FeedAuditTrailLogUpsert) SetUpdatedAt(v time.Time) *FeedAuditTrailLogUpsert {
	u.Set(feedaudittraillog.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FeedAuditTrailLogUpsert) UpdateUpdatedAt() *FeedAuditTrailLogUpsert {
	u.SetExcluded(feedaudittraillog.FieldUpdatedAt)
	return u
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *FeedAuditTrailLogUpsert) ClearUpdatedAt() *FeedAuditTrailLogUpsert {
	u.SetNull(feedaudittraillog.FieldUpdatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.FeedAuditTrailLog.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(feedaudittraillog.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *FeedAuditTrailLogUpsertOne) UpdateNewValues() *FeedAuditTrailLogUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(feedaudittraillog.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.FeedAuditTrailLog.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *FeedAuditTrailLogUpsertOne) Ignore() *FeedAuditTrailLogUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FeedAuditTrailLogUpsertOne) DoNothing() *FeedAuditTrailLogUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FeedAuditTrailLogCreate.OnConflict
// documentation for more info.
func (u *FeedAuditTrailLogUpsertOne) Update(set func(*FeedAuditTrailLogUpsert)) *FeedAuditTrailLogUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FeedAuditTrailLogUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *FeedAuditTrailLogUpsertOne) SetUpdatedAt(v time.Time) *FeedAuditTrailLogUpsertOne {
	return u.Update(func(s *FeedAuditTrailLogUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FeedAuditTrailLogUpsertOne) UpdateUpdatedAt() *FeedAuditTrailLogUpsertOne {
	return u.Update(func(s *FeedAuditTrailLogUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *FeedAuditTrailLogUpsertOne) ClearUpdatedAt() *FeedAuditTrailLogUpsertOne {
	return u.Update(func(s *FeedAuditTrailLogUpsert) {
		s.ClearUpdatedAt()
	})
}

// Exec executes the query.
func (u *FeedAuditTrailLogUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for FeedAuditTrailLogCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FeedAuditTrailLogUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *FeedAuditTrailLogUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *FeedAuditTrailLogUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// FeedAuditTrailLogCreateBulk is the builder for creating many FeedAuditTrailLog entities in bulk.
type FeedAuditTrailLogCreateBulk struct {
	config
	err      error
	builders []*FeedAuditTrailLogCreate
	conflict []sql.ConflictOption
}

// Save creates the FeedAuditTrailLog entities in the database.
func (fatlcb *FeedAuditTrailLogCreateBulk) Save(ctx context.Context) ([]*FeedAuditTrailLog, error) {
	if fatlcb.err != nil {
		return nil, fatlcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(fatlcb.builders))
	nodes := make([]*FeedAuditTrailLog, len(fatlcb.builders))
	mutators := make([]Mutator, len(fatlcb.builders))
	for i := range fatlcb.builders {
		func(i int, root context.Context) {
			builder := fatlcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FeedAuditTrailLogMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, fatlcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = fatlcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fatlcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, fatlcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fatlcb *FeedAuditTrailLogCreateBulk) SaveX(ctx context.Context) []*FeedAuditTrailLog {
	v, err := fatlcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fatlcb *FeedAuditTrailLogCreateBulk) Exec(ctx context.Context) error {
	_, err := fatlcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fatlcb *FeedAuditTrailLogCreateBulk) ExecX(ctx context.Context) {
	if err := fatlcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.FeedAuditTrailLog.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FeedAuditTrailLogUpsert) {
//			SetUpdatedAt(v+v).
//		}).
//		Exec(ctx)
func (fatlcb *FeedAuditTrailLogCreateBulk) OnConflict(opts ...sql.ConflictOption) *FeedAuditTrailLogUpsertBulk {
	fatlcb.conflict = opts
	return &FeedAuditTrailLogUpsertBulk{
		create: fatlcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.FeedAuditTrailLog.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (fatlcb *FeedAuditTrailLogCreateBulk) OnConflictColumns(columns ...string) *FeedAuditTrailLogUpsertBulk {
	fatlcb.conflict = append(fatlcb.conflict, sql.ConflictColumns(columns...))
	return &FeedAuditTrailLogUpsertBulk{
		create: fatlcb,
	}
}

// FeedAuditTrailLogUpsertBulk is the builder for "upsert"-ing
// a bulk of FeedAuditTrailLog nodes.
type FeedAuditTrailLogUpsertBulk struct {
	create *FeedAuditTrailLogCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.FeedAuditTrailLog.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(feedaudittraillog.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *FeedAuditTrailLogUpsertBulk) UpdateNewValues() *FeedAuditTrailLogUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(feedaudittraillog.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.FeedAuditTrailLog.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *FeedAuditTrailLogUpsertBulk) Ignore() *FeedAuditTrailLogUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FeedAuditTrailLogUpsertBulk) DoNothing() *FeedAuditTrailLogUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FeedAuditTrailLogCreateBulk.OnConflict
// documentation for more info.
func (u *FeedAuditTrailLogUpsertBulk) Update(set func(*FeedAuditTrailLogUpsert)) *FeedAuditTrailLogUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FeedAuditTrailLogUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *FeedAuditTrailLogUpsertBulk) SetUpdatedAt(v time.Time) *FeedAuditTrailLogUpsertBulk {
	return u.Update(func(s *FeedAuditTrailLogUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FeedAuditTrailLogUpsertBulk) UpdateUpdatedAt() *FeedAuditTrailLogUpsertBulk {
	return u.Update(func(s *FeedAuditTrailLogUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *FeedAuditTrailLogUpsertBulk) ClearUpdatedAt() *FeedAuditTrailLogUpsertBulk {
	return u.Update(func(s *FeedAuditTrailLogUpsert) {
		s.ClearUpdatedAt()
	})
}

// Exec executes the query.
func (u *FeedAuditTrailLogUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the FeedAuditTrailLogCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for FeedAuditTrailLogCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FeedAuditTrailLogUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

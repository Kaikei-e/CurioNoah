// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"insightstream/ent/cooccurrencenetworkpool"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// CooccurrenceNetworkPoolCreate is the builder for creating a CooccurrenceNetworkPool entity.
type CooccurrenceNetworkPoolCreate struct {
	config
	mutation *CooccurrenceNetworkPoolMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetSiteURL sets the "site_url" field.
func (cnpc *CooccurrenceNetworkPoolCreate) SetSiteURL(s string) *CooccurrenceNetworkPoolCreate {
	cnpc.mutation.SetSiteURL(s)
	return cnpc
}

// SetTitles sets the "titles" field.
func (cnpc *CooccurrenceNetworkPoolCreate) SetTitles(s []string) *CooccurrenceNetworkPoolCreate {
	cnpc.mutation.SetTitles(s)
	return cnpc
}

// SetDescriptions sets the "descriptions" field.
func (cnpc *CooccurrenceNetworkPoolCreate) SetDescriptions(s []string) *CooccurrenceNetworkPoolCreate {
	cnpc.mutation.SetDescriptions(s)
	return cnpc
}

// SetID sets the "id" field.
func (cnpc *CooccurrenceNetworkPoolCreate) SetID(u uuid.UUID) *CooccurrenceNetworkPoolCreate {
	cnpc.mutation.SetID(u)
	return cnpc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (cnpc *CooccurrenceNetworkPoolCreate) SetNillableID(u *uuid.UUID) *CooccurrenceNetworkPoolCreate {
	if u != nil {
		cnpc.SetID(*u)
	}
	return cnpc
}

// Mutation returns the CooccurrenceNetworkPoolMutation object of the builder.
func (cnpc *CooccurrenceNetworkPoolCreate) Mutation() *CooccurrenceNetworkPoolMutation {
	return cnpc.mutation
}

// Save creates the CooccurrenceNetworkPool in the database.
func (cnpc *CooccurrenceNetworkPoolCreate) Save(ctx context.Context) (*CooccurrenceNetworkPool, error) {
	cnpc.defaults()
	return withHooks[*CooccurrenceNetworkPool, CooccurrenceNetworkPoolMutation](ctx, cnpc.sqlSave, cnpc.mutation, cnpc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cnpc *CooccurrenceNetworkPoolCreate) SaveX(ctx context.Context) *CooccurrenceNetworkPool {
	v, err := cnpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cnpc *CooccurrenceNetworkPoolCreate) Exec(ctx context.Context) error {
	_, err := cnpc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cnpc *CooccurrenceNetworkPoolCreate) ExecX(ctx context.Context) {
	if err := cnpc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cnpc *CooccurrenceNetworkPoolCreate) defaults() {
	if _, ok := cnpc.mutation.Titles(); !ok {
		v := cooccurrencenetworkpool.DefaultTitles
		cnpc.mutation.SetTitles(v)
	}
	if _, ok := cnpc.mutation.Descriptions(); !ok {
		v := cooccurrencenetworkpool.DefaultDescriptions
		cnpc.mutation.SetDescriptions(v)
	}
	if _, ok := cnpc.mutation.ID(); !ok {
		v := cooccurrencenetworkpool.DefaultID()
		cnpc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cnpc *CooccurrenceNetworkPoolCreate) check() error {
	if _, ok := cnpc.mutation.SiteURL(); !ok {
		return &ValidationError{Name: "site_url", err: errors.New(`ent: missing required field "CooccurrenceNetworkPool.site_url"`)}
	}
	if v, ok := cnpc.mutation.SiteURL(); ok {
		if err := cooccurrencenetworkpool.SiteURLValidator(v); err != nil {
			return &ValidationError{Name: "site_url", err: fmt.Errorf(`ent: validator failed for field "CooccurrenceNetworkPool.site_url": %w`, err)}
		}
	}
	if _, ok := cnpc.mutation.Titles(); !ok {
		return &ValidationError{Name: "titles", err: errors.New(`ent: missing required field "CooccurrenceNetworkPool.titles"`)}
	}
	if _, ok := cnpc.mutation.Descriptions(); !ok {
		return &ValidationError{Name: "descriptions", err: errors.New(`ent: missing required field "CooccurrenceNetworkPool.descriptions"`)}
	}
	return nil
}

func (cnpc *CooccurrenceNetworkPoolCreate) sqlSave(ctx context.Context) (*CooccurrenceNetworkPool, error) {
	if err := cnpc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cnpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cnpc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	cnpc.mutation.id = &_node.ID
	cnpc.mutation.done = true
	return _node, nil
}

func (cnpc *CooccurrenceNetworkPoolCreate) createSpec() (*CooccurrenceNetworkPool, *sqlgraph.CreateSpec) {
	var (
		_node = &CooccurrenceNetworkPool{config: cnpc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: cooccurrencenetworkpool.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: cooccurrencenetworkpool.FieldID,
			},
		}
	)
	_spec.OnConflict = cnpc.conflict
	if id, ok := cnpc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := cnpc.mutation.SiteURL(); ok {
		_spec.SetField(cooccurrencenetworkpool.FieldSiteURL, field.TypeString, value)
		_node.SiteURL = value
	}
	if value, ok := cnpc.mutation.Titles(); ok {
		_spec.SetField(cooccurrencenetworkpool.FieldTitles, field.TypeJSON, value)
		_node.Titles = value
	}
	if value, ok := cnpc.mutation.Descriptions(); ok {
		_spec.SetField(cooccurrencenetworkpool.FieldDescriptions, field.TypeJSON, value)
		_node.Descriptions = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.CooccurrenceNetworkPool.Create().
//		SetSiteURL(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CooccurrenceNetworkPoolUpsert) {
//			SetSiteURL(v+v).
//		}).
//		Exec(ctx)
func (cnpc *CooccurrenceNetworkPoolCreate) OnConflict(opts ...sql.ConflictOption) *CooccurrenceNetworkPoolUpsertOne {
	cnpc.conflict = opts
	return &CooccurrenceNetworkPoolUpsertOne{
		create: cnpc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.CooccurrenceNetworkPool.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (cnpc *CooccurrenceNetworkPoolCreate) OnConflictColumns(columns ...string) *CooccurrenceNetworkPoolUpsertOne {
	cnpc.conflict = append(cnpc.conflict, sql.ConflictColumns(columns...))
	return &CooccurrenceNetworkPoolUpsertOne{
		create: cnpc,
	}
}

type (
	// CooccurrenceNetworkPoolUpsertOne is the builder for "upsert"-ing
	//  one CooccurrenceNetworkPool node.
	CooccurrenceNetworkPoolUpsertOne struct {
		create *CooccurrenceNetworkPoolCreate
	}

	// CooccurrenceNetworkPoolUpsert is the "OnConflict" setter.
	CooccurrenceNetworkPoolUpsert struct {
		*sql.UpdateSet
	}
)

// SetSiteURL sets the "site_url" field.
func (u *CooccurrenceNetworkPoolUpsert) SetSiteURL(v string) *CooccurrenceNetworkPoolUpsert {
	u.Set(cooccurrencenetworkpool.FieldSiteURL, v)
	return u
}

// UpdateSiteURL sets the "site_url" field to the value that was provided on create.
func (u *CooccurrenceNetworkPoolUpsert) UpdateSiteURL() *CooccurrenceNetworkPoolUpsert {
	u.SetExcluded(cooccurrencenetworkpool.FieldSiteURL)
	return u
}

// SetTitles sets the "titles" field.
func (u *CooccurrenceNetworkPoolUpsert) SetTitles(v []string) *CooccurrenceNetworkPoolUpsert {
	u.Set(cooccurrencenetworkpool.FieldTitles, v)
	return u
}

// UpdateTitles sets the "titles" field to the value that was provided on create.
func (u *CooccurrenceNetworkPoolUpsert) UpdateTitles() *CooccurrenceNetworkPoolUpsert {
	u.SetExcluded(cooccurrencenetworkpool.FieldTitles)
	return u
}

// SetDescriptions sets the "descriptions" field.
func (u *CooccurrenceNetworkPoolUpsert) SetDescriptions(v []string) *CooccurrenceNetworkPoolUpsert {
	u.Set(cooccurrencenetworkpool.FieldDescriptions, v)
	return u
}

// UpdateDescriptions sets the "descriptions" field to the value that was provided on create.
func (u *CooccurrenceNetworkPoolUpsert) UpdateDescriptions() *CooccurrenceNetworkPoolUpsert {
	u.SetExcluded(cooccurrencenetworkpool.FieldDescriptions)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.CooccurrenceNetworkPool.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(cooccurrencenetworkpool.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *CooccurrenceNetworkPoolUpsertOne) UpdateNewValues() *CooccurrenceNetworkPoolUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(cooccurrencenetworkpool.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.CooccurrenceNetworkPool.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *CooccurrenceNetworkPoolUpsertOne) Ignore() *CooccurrenceNetworkPoolUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CooccurrenceNetworkPoolUpsertOne) DoNothing() *CooccurrenceNetworkPoolUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CooccurrenceNetworkPoolCreate.OnConflict
// documentation for more info.
func (u *CooccurrenceNetworkPoolUpsertOne) Update(set func(*CooccurrenceNetworkPoolUpsert)) *CooccurrenceNetworkPoolUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CooccurrenceNetworkPoolUpsert{UpdateSet: update})
	}))
	return u
}

// SetSiteURL sets the "site_url" field.
func (u *CooccurrenceNetworkPoolUpsertOne) SetSiteURL(v string) *CooccurrenceNetworkPoolUpsertOne {
	return u.Update(func(s *CooccurrenceNetworkPoolUpsert) {
		s.SetSiteURL(v)
	})
}

// UpdateSiteURL sets the "site_url" field to the value that was provided on create.
func (u *CooccurrenceNetworkPoolUpsertOne) UpdateSiteURL() *CooccurrenceNetworkPoolUpsertOne {
	return u.Update(func(s *CooccurrenceNetworkPoolUpsert) {
		s.UpdateSiteURL()
	})
}

// SetTitles sets the "titles" field.
func (u *CooccurrenceNetworkPoolUpsertOne) SetTitles(v []string) *CooccurrenceNetworkPoolUpsertOne {
	return u.Update(func(s *CooccurrenceNetworkPoolUpsert) {
		s.SetTitles(v)
	})
}

// UpdateTitles sets the "titles" field to the value that was provided on create.
func (u *CooccurrenceNetworkPoolUpsertOne) UpdateTitles() *CooccurrenceNetworkPoolUpsertOne {
	return u.Update(func(s *CooccurrenceNetworkPoolUpsert) {
		s.UpdateTitles()
	})
}

// SetDescriptions sets the "descriptions" field.
func (u *CooccurrenceNetworkPoolUpsertOne) SetDescriptions(v []string) *CooccurrenceNetworkPoolUpsertOne {
	return u.Update(func(s *CooccurrenceNetworkPoolUpsert) {
		s.SetDescriptions(v)
	})
}

// UpdateDescriptions sets the "descriptions" field to the value that was provided on create.
func (u *CooccurrenceNetworkPoolUpsertOne) UpdateDescriptions() *CooccurrenceNetworkPoolUpsertOne {
	return u.Update(func(s *CooccurrenceNetworkPoolUpsert) {
		s.UpdateDescriptions()
	})
}

// Exec executes the query.
func (u *CooccurrenceNetworkPoolUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CooccurrenceNetworkPoolCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CooccurrenceNetworkPoolUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *CooccurrenceNetworkPoolUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: CooccurrenceNetworkPoolUpsertOne.ID is not supported by MySQL driver. Use CooccurrenceNetworkPoolUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *CooccurrenceNetworkPoolUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// CooccurrenceNetworkPoolCreateBulk is the builder for creating many CooccurrenceNetworkPool entities in bulk.
type CooccurrenceNetworkPoolCreateBulk struct {
	config
	builders []*CooccurrenceNetworkPoolCreate
	conflict []sql.ConflictOption
}

// Save creates the CooccurrenceNetworkPool entities in the database.
func (cnpcb *CooccurrenceNetworkPoolCreateBulk) Save(ctx context.Context) ([]*CooccurrenceNetworkPool, error) {
	specs := make([]*sqlgraph.CreateSpec, len(cnpcb.builders))
	nodes := make([]*CooccurrenceNetworkPool, len(cnpcb.builders))
	mutators := make([]Mutator, len(cnpcb.builders))
	for i := range cnpcb.builders {
		func(i int, root context.Context) {
			builder := cnpcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CooccurrenceNetworkPoolMutation)
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
					_, err = mutators[i+1].Mutate(root, cnpcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = cnpcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cnpcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, cnpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cnpcb *CooccurrenceNetworkPoolCreateBulk) SaveX(ctx context.Context) []*CooccurrenceNetworkPool {
	v, err := cnpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cnpcb *CooccurrenceNetworkPoolCreateBulk) Exec(ctx context.Context) error {
	_, err := cnpcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cnpcb *CooccurrenceNetworkPoolCreateBulk) ExecX(ctx context.Context) {
	if err := cnpcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.CooccurrenceNetworkPool.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CooccurrenceNetworkPoolUpsert) {
//			SetSiteURL(v+v).
//		}).
//		Exec(ctx)
func (cnpcb *CooccurrenceNetworkPoolCreateBulk) OnConflict(opts ...sql.ConflictOption) *CooccurrenceNetworkPoolUpsertBulk {
	cnpcb.conflict = opts
	return &CooccurrenceNetworkPoolUpsertBulk{
		create: cnpcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.CooccurrenceNetworkPool.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (cnpcb *CooccurrenceNetworkPoolCreateBulk) OnConflictColumns(columns ...string) *CooccurrenceNetworkPoolUpsertBulk {
	cnpcb.conflict = append(cnpcb.conflict, sql.ConflictColumns(columns...))
	return &CooccurrenceNetworkPoolUpsertBulk{
		create: cnpcb,
	}
}

// CooccurrenceNetworkPoolUpsertBulk is the builder for "upsert"-ing
// a bulk of CooccurrenceNetworkPool nodes.
type CooccurrenceNetworkPoolUpsertBulk struct {
	create *CooccurrenceNetworkPoolCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.CooccurrenceNetworkPool.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(cooccurrencenetworkpool.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *CooccurrenceNetworkPoolUpsertBulk) UpdateNewValues() *CooccurrenceNetworkPoolUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(cooccurrencenetworkpool.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.CooccurrenceNetworkPool.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *CooccurrenceNetworkPoolUpsertBulk) Ignore() *CooccurrenceNetworkPoolUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CooccurrenceNetworkPoolUpsertBulk) DoNothing() *CooccurrenceNetworkPoolUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CooccurrenceNetworkPoolCreateBulk.OnConflict
// documentation for more info.
func (u *CooccurrenceNetworkPoolUpsertBulk) Update(set func(*CooccurrenceNetworkPoolUpsert)) *CooccurrenceNetworkPoolUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CooccurrenceNetworkPoolUpsert{UpdateSet: update})
	}))
	return u
}

// SetSiteURL sets the "site_url" field.
func (u *CooccurrenceNetworkPoolUpsertBulk) SetSiteURL(v string) *CooccurrenceNetworkPoolUpsertBulk {
	return u.Update(func(s *CooccurrenceNetworkPoolUpsert) {
		s.SetSiteURL(v)
	})
}

// UpdateSiteURL sets the "site_url" field to the value that was provided on create.
func (u *CooccurrenceNetworkPoolUpsertBulk) UpdateSiteURL() *CooccurrenceNetworkPoolUpsertBulk {
	return u.Update(func(s *CooccurrenceNetworkPoolUpsert) {
		s.UpdateSiteURL()
	})
}

// SetTitles sets the "titles" field.
func (u *CooccurrenceNetworkPoolUpsertBulk) SetTitles(v []string) *CooccurrenceNetworkPoolUpsertBulk {
	return u.Update(func(s *CooccurrenceNetworkPoolUpsert) {
		s.SetTitles(v)
	})
}

// UpdateTitles sets the "titles" field to the value that was provided on create.
func (u *CooccurrenceNetworkPoolUpsertBulk) UpdateTitles() *CooccurrenceNetworkPoolUpsertBulk {
	return u.Update(func(s *CooccurrenceNetworkPoolUpsert) {
		s.UpdateTitles()
	})
}

// SetDescriptions sets the "descriptions" field.
func (u *CooccurrenceNetworkPoolUpsertBulk) SetDescriptions(v []string) *CooccurrenceNetworkPoolUpsertBulk {
	return u.Update(func(s *CooccurrenceNetworkPoolUpsert) {
		s.SetDescriptions(v)
	})
}

// UpdateDescriptions sets the "descriptions" field to the value that was provided on create.
func (u *CooccurrenceNetworkPoolUpsertBulk) UpdateDescriptions() *CooccurrenceNetworkPoolUpsertBulk {
	return u.Update(func(s *CooccurrenceNetworkPoolUpsert) {
		s.UpdateDescriptions()
	})
}

// Exec executes the query.
func (u *CooccurrenceNetworkPoolUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the CooccurrenceNetworkPoolCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CooccurrenceNetworkPoolCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CooccurrenceNetworkPoolUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

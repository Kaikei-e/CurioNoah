// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"insightstream/ent/feeds"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// FeedsCreate is the builder for creating a Feeds entity.
type FeedsCreate struct {
	config
	mutation *FeedsMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetSiteURL sets the "site_url" field.
func (fc *FeedsCreate) SetSiteURL(s string) *FeedsCreate {
	fc.mutation.SetSiteURL(s)
	return fc
}

// SetNillableSiteURL sets the "site_url" field if the given value is not nil.
func (fc *FeedsCreate) SetNillableSiteURL(s *string) *FeedsCreate {
	if s != nil {
		fc.SetSiteURL(*s)
	}
	return fc
}

// SetTitle sets the "title" field.
func (fc *FeedsCreate) SetTitle(s string) *FeedsCreate {
	fc.mutation.SetTitle(s)
	return fc
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (fc *FeedsCreate) SetNillableTitle(s *string) *FeedsCreate {
	if s != nil {
		fc.SetTitle(*s)
	}
	return fc
}

// SetDescription sets the "description" field.
func (fc *FeedsCreate) SetDescription(s string) *FeedsCreate {
	fc.mutation.SetDescription(s)
	return fc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (fc *FeedsCreate) SetNillableDescription(s *string) *FeedsCreate {
	if s != nil {
		fc.SetDescription(*s)
	}
	return fc
}

// SetFeedURL sets the "feed_url" field.
func (fc *FeedsCreate) SetFeedURL(s string) *FeedsCreate {
	fc.mutation.SetFeedURL(s)
	return fc
}

// SetNillableFeedURL sets the "feed_url" field if the given value is not nil.
func (fc *FeedsCreate) SetNillableFeedURL(s *string) *FeedsCreate {
	if s != nil {
		fc.SetFeedURL(*s)
	}
	return fc
}

// SetLanguage sets the "language" field.
func (fc *FeedsCreate) SetLanguage(s string) *FeedsCreate {
	fc.mutation.SetLanguage(s)
	return fc
}

// SetNillableLanguage sets the "language" field if the given value is not nil.
func (fc *FeedsCreate) SetNillableLanguage(s *string) *FeedsCreate {
	if s != nil {
		fc.SetLanguage(*s)
	}
	return fc
}

// SetDtCreated sets the "dt_created" field.
func (fc *FeedsCreate) SetDtCreated(t time.Time) *FeedsCreate {
	fc.mutation.SetDtCreated(t)
	return fc
}

// SetNillableDtCreated sets the "dt_created" field if the given value is not nil.
func (fc *FeedsCreate) SetNillableDtCreated(t *time.Time) *FeedsCreate {
	if t != nil {
		fc.SetDtCreated(*t)
	}
	return fc
}

// SetDtUpdated sets the "dt_updated" field.
func (fc *FeedsCreate) SetDtUpdated(t time.Time) *FeedsCreate {
	fc.mutation.SetDtUpdated(t)
	return fc
}

// SetNillableDtUpdated sets the "dt_updated" field if the given value is not nil.
func (fc *FeedsCreate) SetNillableDtUpdated(t *time.Time) *FeedsCreate {
	if t != nil {
		fc.SetDtUpdated(*t)
	}
	return fc
}

// SetFavorites sets the "favorites" field.
func (fc *FeedsCreate) SetFavorites(i int64) *FeedsCreate {
	fc.mutation.SetFavorites(i)
	return fc
}

// SetNillableFavorites sets the "favorites" field if the given value is not nil.
func (fc *FeedsCreate) SetNillableFavorites(i *int64) *FeedsCreate {
	if i != nil {
		fc.SetFavorites(*i)
	}
	return fc
}

// SetID sets the "id" field.
func (fc *FeedsCreate) SetID(u uuid.UUID) *FeedsCreate {
	fc.mutation.SetID(u)
	return fc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (fc *FeedsCreate) SetNillableID(u *uuid.UUID) *FeedsCreate {
	if u != nil {
		fc.SetID(*u)
	}
	return fc
}

// Mutation returns the FeedsMutation object of the builder.
func (fc *FeedsCreate) Mutation() *FeedsMutation {
	return fc.mutation
}

// Save creates the Feeds in the database.
func (fc *FeedsCreate) Save(ctx context.Context) (*Feeds, error) {
	fc.defaults()
	return withHooks(ctx, fc.sqlSave, fc.mutation, fc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FeedsCreate) SaveX(ctx context.Context) *Feeds {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fc *FeedsCreate) Exec(ctx context.Context) error {
	_, err := fc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fc *FeedsCreate) ExecX(ctx context.Context) {
	if err := fc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fc *FeedsCreate) defaults() {
	if _, ok := fc.mutation.SiteURL(); !ok {
		v := feeds.DefaultSiteURL
		fc.mutation.SetSiteURL(v)
	}
	if _, ok := fc.mutation.Title(); !ok {
		v := feeds.DefaultTitle
		fc.mutation.SetTitle(v)
	}
	if _, ok := fc.mutation.Description(); !ok {
		v := feeds.DefaultDescription
		fc.mutation.SetDescription(v)
	}
	if _, ok := fc.mutation.FeedURL(); !ok {
		v := feeds.DefaultFeedURL
		fc.mutation.SetFeedURL(v)
	}
	if _, ok := fc.mutation.Language(); !ok {
		v := feeds.DefaultLanguage
		fc.mutation.SetLanguage(v)
	}
	if _, ok := fc.mutation.DtCreated(); !ok {
		v := feeds.DefaultDtCreated()
		fc.mutation.SetDtCreated(v)
	}
	if _, ok := fc.mutation.Favorites(); !ok {
		v := feeds.DefaultFavorites
		fc.mutation.SetFavorites(v)
	}
	if _, ok := fc.mutation.ID(); !ok {
		v := feeds.DefaultID()
		fc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fc *FeedsCreate) check() error {
	if _, ok := fc.mutation.SiteURL(); !ok {
		return &ValidationError{Name: "site_url", err: errors.New(`ent: missing required field "Feeds.site_url"`)}
	}
	if v, ok := fc.mutation.SiteURL(); ok {
		if err := feeds.SiteURLValidator(v); err != nil {
			return &ValidationError{Name: "site_url", err: fmt.Errorf(`ent: validator failed for field "Feeds.site_url": %w`, err)}
		}
	}
	if _, ok := fc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Feeds.title"`)}
	}
	if v, ok := fc.mutation.Title(); ok {
		if err := feeds.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Feeds.title": %w`, err)}
		}
	}
	if _, ok := fc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Feeds.description"`)}
	}
	if v, ok := fc.mutation.Description(); ok {
		if err := feeds.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Feeds.description": %w`, err)}
		}
	}
	if _, ok := fc.mutation.FeedURL(); !ok {
		return &ValidationError{Name: "feed_url", err: errors.New(`ent: missing required field "Feeds.feed_url"`)}
	}
	if v, ok := fc.mutation.FeedURL(); ok {
		if err := feeds.FeedURLValidator(v); err != nil {
			return &ValidationError{Name: "feed_url", err: fmt.Errorf(`ent: validator failed for field "Feeds.feed_url": %w`, err)}
		}
	}
	if _, ok := fc.mutation.Language(); !ok {
		return &ValidationError{Name: "language", err: errors.New(`ent: missing required field "Feeds.language"`)}
	}
	if v, ok := fc.mutation.Language(); ok {
		if err := feeds.LanguageValidator(v); err != nil {
			return &ValidationError{Name: "language", err: fmt.Errorf(`ent: validator failed for field "Feeds.language": %w`, err)}
		}
	}
	if _, ok := fc.mutation.Favorites(); !ok {
		return &ValidationError{Name: "favorites", err: errors.New(`ent: missing required field "Feeds.favorites"`)}
	}
	return nil
}

func (fc *FeedsCreate) sqlSave(ctx context.Context) (*Feeds, error) {
	if err := fc.check(); err != nil {
		return nil, err
	}
	_node, _spec := fc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fc.driver, _spec); err != nil {
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
	fc.mutation.id = &_node.ID
	fc.mutation.done = true
	return _node, nil
}

func (fc *FeedsCreate) createSpec() (*Feeds, *sqlgraph.CreateSpec) {
	var (
		_node = &Feeds{config: fc.config}
		_spec = sqlgraph.NewCreateSpec(feeds.Table, sqlgraph.NewFieldSpec(feeds.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = fc.conflict
	if id, ok := fc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := fc.mutation.SiteURL(); ok {
		_spec.SetField(feeds.FieldSiteURL, field.TypeString, value)
		_node.SiteURL = value
	}
	if value, ok := fc.mutation.Title(); ok {
		_spec.SetField(feeds.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := fc.mutation.Description(); ok {
		_spec.SetField(feeds.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := fc.mutation.FeedURL(); ok {
		_spec.SetField(feeds.FieldFeedURL, field.TypeString, value)
		_node.FeedURL = value
	}
	if value, ok := fc.mutation.Language(); ok {
		_spec.SetField(feeds.FieldLanguage, field.TypeString, value)
		_node.Language = value
	}
	if value, ok := fc.mutation.DtCreated(); ok {
		_spec.SetField(feeds.FieldDtCreated, field.TypeTime, value)
		_node.DtCreated = value
	}
	if value, ok := fc.mutation.DtUpdated(); ok {
		_spec.SetField(feeds.FieldDtUpdated, field.TypeTime, value)
		_node.DtUpdated = value
	}
	if value, ok := fc.mutation.Favorites(); ok {
		_spec.SetField(feeds.FieldFavorites, field.TypeInt64, value)
		_node.Favorites = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Feeds.Create().
//		SetSiteURL(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FeedsUpsert) {
//			SetSiteURL(v+v).
//		}).
//		Exec(ctx)
func (fc *FeedsCreate) OnConflict(opts ...sql.ConflictOption) *FeedsUpsertOne {
	fc.conflict = opts
	return &FeedsUpsertOne{
		create: fc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Feeds.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (fc *FeedsCreate) OnConflictColumns(columns ...string) *FeedsUpsertOne {
	fc.conflict = append(fc.conflict, sql.ConflictColumns(columns...))
	return &FeedsUpsertOne{
		create: fc,
	}
}

type (
	// FeedsUpsertOne is the builder for "upsert"-ing
	//  one Feeds node.
	FeedsUpsertOne struct {
		create *FeedsCreate
	}

	// FeedsUpsert is the "OnConflict" setter.
	FeedsUpsert struct {
		*sql.UpdateSet
	}
)

// SetSiteURL sets the "site_url" field.
func (u *FeedsUpsert) SetSiteURL(v string) *FeedsUpsert {
	u.Set(feeds.FieldSiteURL, v)
	return u
}

// UpdateSiteURL sets the "site_url" field to the value that was provided on create.
func (u *FeedsUpsert) UpdateSiteURL() *FeedsUpsert {
	u.SetExcluded(feeds.FieldSiteURL)
	return u
}

// SetTitle sets the "title" field.
func (u *FeedsUpsert) SetTitle(v string) *FeedsUpsert {
	u.Set(feeds.FieldTitle, v)
	return u
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *FeedsUpsert) UpdateTitle() *FeedsUpsert {
	u.SetExcluded(feeds.FieldTitle)
	return u
}

// SetDescription sets the "description" field.
func (u *FeedsUpsert) SetDescription(v string) *FeedsUpsert {
	u.Set(feeds.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *FeedsUpsert) UpdateDescription() *FeedsUpsert {
	u.SetExcluded(feeds.FieldDescription)
	return u
}

// SetFeedURL sets the "feed_url" field.
func (u *FeedsUpsert) SetFeedURL(v string) *FeedsUpsert {
	u.Set(feeds.FieldFeedURL, v)
	return u
}

// UpdateFeedURL sets the "feed_url" field to the value that was provided on create.
func (u *FeedsUpsert) UpdateFeedURL() *FeedsUpsert {
	u.SetExcluded(feeds.FieldFeedURL)
	return u
}

// SetLanguage sets the "language" field.
func (u *FeedsUpsert) SetLanguage(v string) *FeedsUpsert {
	u.Set(feeds.FieldLanguage, v)
	return u
}

// UpdateLanguage sets the "language" field to the value that was provided on create.
func (u *FeedsUpsert) UpdateLanguage() *FeedsUpsert {
	u.SetExcluded(feeds.FieldLanguage)
	return u
}

// SetDtUpdated sets the "dt_updated" field.
func (u *FeedsUpsert) SetDtUpdated(v time.Time) *FeedsUpsert {
	u.Set(feeds.FieldDtUpdated, v)
	return u
}

// UpdateDtUpdated sets the "dt_updated" field to the value that was provided on create.
func (u *FeedsUpsert) UpdateDtUpdated() *FeedsUpsert {
	u.SetExcluded(feeds.FieldDtUpdated)
	return u
}

// ClearDtUpdated clears the value of the "dt_updated" field.
func (u *FeedsUpsert) ClearDtUpdated() *FeedsUpsert {
	u.SetNull(feeds.FieldDtUpdated)
	return u
}

// SetFavorites sets the "favorites" field.
func (u *FeedsUpsert) SetFavorites(v int64) *FeedsUpsert {
	u.Set(feeds.FieldFavorites, v)
	return u
}

// UpdateFavorites sets the "favorites" field to the value that was provided on create.
func (u *FeedsUpsert) UpdateFavorites() *FeedsUpsert {
	u.SetExcluded(feeds.FieldFavorites)
	return u
}

// AddFavorites adds v to the "favorites" field.
func (u *FeedsUpsert) AddFavorites(v int64) *FeedsUpsert {
	u.Add(feeds.FieldFavorites, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Feeds.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(feeds.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *FeedsUpsertOne) UpdateNewValues() *FeedsUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(feeds.FieldID)
		}
		if _, exists := u.create.mutation.DtCreated(); exists {
			s.SetIgnore(feeds.FieldDtCreated)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Feeds.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *FeedsUpsertOne) Ignore() *FeedsUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FeedsUpsertOne) DoNothing() *FeedsUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FeedsCreate.OnConflict
// documentation for more info.
func (u *FeedsUpsertOne) Update(set func(*FeedsUpsert)) *FeedsUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FeedsUpsert{UpdateSet: update})
	}))
	return u
}

// SetSiteURL sets the "site_url" field.
func (u *FeedsUpsertOne) SetSiteURL(v string) *FeedsUpsertOne {
	return u.Update(func(s *FeedsUpsert) {
		s.SetSiteURL(v)
	})
}

// UpdateSiteURL sets the "site_url" field to the value that was provided on create.
func (u *FeedsUpsertOne) UpdateSiteURL() *FeedsUpsertOne {
	return u.Update(func(s *FeedsUpsert) {
		s.UpdateSiteURL()
	})
}

// SetTitle sets the "title" field.
func (u *FeedsUpsertOne) SetTitle(v string) *FeedsUpsertOne {
	return u.Update(func(s *FeedsUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *FeedsUpsertOne) UpdateTitle() *FeedsUpsertOne {
	return u.Update(func(s *FeedsUpsert) {
		s.UpdateTitle()
	})
}

// SetDescription sets the "description" field.
func (u *FeedsUpsertOne) SetDescription(v string) *FeedsUpsertOne {
	return u.Update(func(s *FeedsUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *FeedsUpsertOne) UpdateDescription() *FeedsUpsertOne {
	return u.Update(func(s *FeedsUpsert) {
		s.UpdateDescription()
	})
}

// SetFeedURL sets the "feed_url" field.
func (u *FeedsUpsertOne) SetFeedURL(v string) *FeedsUpsertOne {
	return u.Update(func(s *FeedsUpsert) {
		s.SetFeedURL(v)
	})
}

// UpdateFeedURL sets the "feed_url" field to the value that was provided on create.
func (u *FeedsUpsertOne) UpdateFeedURL() *FeedsUpsertOne {
	return u.Update(func(s *FeedsUpsert) {
		s.UpdateFeedURL()
	})
}

// SetLanguage sets the "language" field.
func (u *FeedsUpsertOne) SetLanguage(v string) *FeedsUpsertOne {
	return u.Update(func(s *FeedsUpsert) {
		s.SetLanguage(v)
	})
}

// UpdateLanguage sets the "language" field to the value that was provided on create.
func (u *FeedsUpsertOne) UpdateLanguage() *FeedsUpsertOne {
	return u.Update(func(s *FeedsUpsert) {
		s.UpdateLanguage()
	})
}

// SetDtUpdated sets the "dt_updated" field.
func (u *FeedsUpsertOne) SetDtUpdated(v time.Time) *FeedsUpsertOne {
	return u.Update(func(s *FeedsUpsert) {
		s.SetDtUpdated(v)
	})
}

// UpdateDtUpdated sets the "dt_updated" field to the value that was provided on create.
func (u *FeedsUpsertOne) UpdateDtUpdated() *FeedsUpsertOne {
	return u.Update(func(s *FeedsUpsert) {
		s.UpdateDtUpdated()
	})
}

// ClearDtUpdated clears the value of the "dt_updated" field.
func (u *FeedsUpsertOne) ClearDtUpdated() *FeedsUpsertOne {
	return u.Update(func(s *FeedsUpsert) {
		s.ClearDtUpdated()
	})
}

// SetFavorites sets the "favorites" field.
func (u *FeedsUpsertOne) SetFavorites(v int64) *FeedsUpsertOne {
	return u.Update(func(s *FeedsUpsert) {
		s.SetFavorites(v)
	})
}

// AddFavorites adds v to the "favorites" field.
func (u *FeedsUpsertOne) AddFavorites(v int64) *FeedsUpsertOne {
	return u.Update(func(s *FeedsUpsert) {
		s.AddFavorites(v)
	})
}

// UpdateFavorites sets the "favorites" field to the value that was provided on create.
func (u *FeedsUpsertOne) UpdateFavorites() *FeedsUpsertOne {
	return u.Update(func(s *FeedsUpsert) {
		s.UpdateFavorites()
	})
}

// Exec executes the query.
func (u *FeedsUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for FeedsCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FeedsUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *FeedsUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: FeedsUpsertOne.ID is not supported by MySQL driver. Use FeedsUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *FeedsUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// FeedsCreateBulk is the builder for creating many Feeds entities in bulk.
type FeedsCreateBulk struct {
	config
	err      error
	builders []*FeedsCreate
	conflict []sql.ConflictOption
}

// Save creates the Feeds entities in the database.
func (fcb *FeedsCreateBulk) Save(ctx context.Context) ([]*Feeds, error) {
	if fcb.err != nil {
		return nil, fcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(fcb.builders))
	nodes := make([]*Feeds, len(fcb.builders))
	mutators := make([]Mutator, len(fcb.builders))
	for i := range fcb.builders {
		func(i int, root context.Context) {
			builder := fcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FeedsMutation)
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
					_, err = mutators[i+1].Mutate(root, fcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = fcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, fcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fcb *FeedsCreateBulk) SaveX(ctx context.Context) []*Feeds {
	v, err := fcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fcb *FeedsCreateBulk) Exec(ctx context.Context) error {
	_, err := fcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fcb *FeedsCreateBulk) ExecX(ctx context.Context) {
	if err := fcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Feeds.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FeedsUpsert) {
//			SetSiteURL(v+v).
//		}).
//		Exec(ctx)
func (fcb *FeedsCreateBulk) OnConflict(opts ...sql.ConflictOption) *FeedsUpsertBulk {
	fcb.conflict = opts
	return &FeedsUpsertBulk{
		create: fcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Feeds.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (fcb *FeedsCreateBulk) OnConflictColumns(columns ...string) *FeedsUpsertBulk {
	fcb.conflict = append(fcb.conflict, sql.ConflictColumns(columns...))
	return &FeedsUpsertBulk{
		create: fcb,
	}
}

// FeedsUpsertBulk is the builder for "upsert"-ing
// a bulk of Feeds nodes.
type FeedsUpsertBulk struct {
	create *FeedsCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Feeds.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(feeds.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *FeedsUpsertBulk) UpdateNewValues() *FeedsUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(feeds.FieldID)
			}
			if _, exists := b.mutation.DtCreated(); exists {
				s.SetIgnore(feeds.FieldDtCreated)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Feeds.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *FeedsUpsertBulk) Ignore() *FeedsUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FeedsUpsertBulk) DoNothing() *FeedsUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FeedsCreateBulk.OnConflict
// documentation for more info.
func (u *FeedsUpsertBulk) Update(set func(*FeedsUpsert)) *FeedsUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FeedsUpsert{UpdateSet: update})
	}))
	return u
}

// SetSiteURL sets the "site_url" field.
func (u *FeedsUpsertBulk) SetSiteURL(v string) *FeedsUpsertBulk {
	return u.Update(func(s *FeedsUpsert) {
		s.SetSiteURL(v)
	})
}

// UpdateSiteURL sets the "site_url" field to the value that was provided on create.
func (u *FeedsUpsertBulk) UpdateSiteURL() *FeedsUpsertBulk {
	return u.Update(func(s *FeedsUpsert) {
		s.UpdateSiteURL()
	})
}

// SetTitle sets the "title" field.
func (u *FeedsUpsertBulk) SetTitle(v string) *FeedsUpsertBulk {
	return u.Update(func(s *FeedsUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *FeedsUpsertBulk) UpdateTitle() *FeedsUpsertBulk {
	return u.Update(func(s *FeedsUpsert) {
		s.UpdateTitle()
	})
}

// SetDescription sets the "description" field.
func (u *FeedsUpsertBulk) SetDescription(v string) *FeedsUpsertBulk {
	return u.Update(func(s *FeedsUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *FeedsUpsertBulk) UpdateDescription() *FeedsUpsertBulk {
	return u.Update(func(s *FeedsUpsert) {
		s.UpdateDescription()
	})
}

// SetFeedURL sets the "feed_url" field.
func (u *FeedsUpsertBulk) SetFeedURL(v string) *FeedsUpsertBulk {
	return u.Update(func(s *FeedsUpsert) {
		s.SetFeedURL(v)
	})
}

// UpdateFeedURL sets the "feed_url" field to the value that was provided on create.
func (u *FeedsUpsertBulk) UpdateFeedURL() *FeedsUpsertBulk {
	return u.Update(func(s *FeedsUpsert) {
		s.UpdateFeedURL()
	})
}

// SetLanguage sets the "language" field.
func (u *FeedsUpsertBulk) SetLanguage(v string) *FeedsUpsertBulk {
	return u.Update(func(s *FeedsUpsert) {
		s.SetLanguage(v)
	})
}

// UpdateLanguage sets the "language" field to the value that was provided on create.
func (u *FeedsUpsertBulk) UpdateLanguage() *FeedsUpsertBulk {
	return u.Update(func(s *FeedsUpsert) {
		s.UpdateLanguage()
	})
}

// SetDtUpdated sets the "dt_updated" field.
func (u *FeedsUpsertBulk) SetDtUpdated(v time.Time) *FeedsUpsertBulk {
	return u.Update(func(s *FeedsUpsert) {
		s.SetDtUpdated(v)
	})
}

// UpdateDtUpdated sets the "dt_updated" field to the value that was provided on create.
func (u *FeedsUpsertBulk) UpdateDtUpdated() *FeedsUpsertBulk {
	return u.Update(func(s *FeedsUpsert) {
		s.UpdateDtUpdated()
	})
}

// ClearDtUpdated clears the value of the "dt_updated" field.
func (u *FeedsUpsertBulk) ClearDtUpdated() *FeedsUpsertBulk {
	return u.Update(func(s *FeedsUpsert) {
		s.ClearDtUpdated()
	})
}

// SetFavorites sets the "favorites" field.
func (u *FeedsUpsertBulk) SetFavorites(v int64) *FeedsUpsertBulk {
	return u.Update(func(s *FeedsUpsert) {
		s.SetFavorites(v)
	})
}

// AddFavorites adds v to the "favorites" field.
func (u *FeedsUpsertBulk) AddFavorites(v int64) *FeedsUpsertBulk {
	return u.Update(func(s *FeedsUpsert) {
		s.AddFavorites(v)
	})
}

// UpdateFavorites sets the "favorites" field to the value that was provided on create.
func (u *FeedsUpsertBulk) UpdateFavorites() *FeedsUpsertBulk {
	return u.Update(func(s *FeedsUpsert) {
		s.UpdateFavorites()
	})
}

// Exec executes the query.
func (u *FeedsUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the FeedsCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for FeedsCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FeedsUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

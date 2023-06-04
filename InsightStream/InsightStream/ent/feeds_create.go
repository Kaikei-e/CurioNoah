// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"

	entfeeds "insightstream/ent/feeds"
)

// FeedsCreate is the builder for creating a Feeds entity.
type FeedsCreate struct {
	config
	mutation *FeedsMutation
	hooks    []Hook
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
	return withHooks[*Feeds, FeedsMutation](ctx, fc.sqlSave, fc.mutation, fc.hooks)
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
		v := entfeeds.DefaultSiteURL
		fc.mutation.SetSiteURL(v)
	}
	if _, ok := fc.mutation.Title(); !ok {
		v := entfeeds.DefaultTitle
		fc.mutation.SetTitle(v)
	}
	if _, ok := fc.mutation.Description(); !ok {
		v := entfeeds.DefaultDescription
		fc.mutation.SetDescription(v)
	}
	if _, ok := fc.mutation.FeedURL(); !ok {
		v := entfeeds.DefaultFeedURL
		fc.mutation.SetFeedURL(v)
	}
	if _, ok := fc.mutation.Language(); !ok {
		v := entfeeds.DefaultLanguage
		fc.mutation.SetLanguage(v)
	}
	if _, ok := fc.mutation.DtCreated(); !ok {
		v := entfeeds.DefaultDtCreated()
		fc.mutation.SetDtCreated(v)
	}
	if _, ok := fc.mutation.Favorites(); !ok {
		v := entfeeds.DefaultFavorites
		fc.mutation.SetFavorites(v)
	}
	if _, ok := fc.mutation.ID(); !ok {
		v := entfeeds.DefaultID()
		fc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fc *FeedsCreate) check() error {
	if _, ok := fc.mutation.SiteURL(); !ok {
		return &ValidationError{Name: "site_url", err: errors.New(`ent: missing required field "Feeds.site_url"`)}
	}
	if v, ok := fc.mutation.SiteURL(); ok {
		if err := entfeeds.SiteURLValidator(v); err != nil {
			return &ValidationError{Name: "site_url", err: fmt.Errorf(`ent: validator failed for field "Feeds.site_url": %w`, err)}
		}
	}
	if _, ok := fc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Feeds.title"`)}
	}
	if v, ok := fc.mutation.Title(); ok {
		if err := entfeeds.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Feeds.title": %w`, err)}
		}
	}
	if _, ok := fc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Feeds.description"`)}
	}
	if v, ok := fc.mutation.Description(); ok {
		if err := entfeeds.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Feeds.description": %w`, err)}
		}
	}
	if _, ok := fc.mutation.FeedURL(); !ok {
		return &ValidationError{Name: "feed_url", err: errors.New(`ent: missing required field "Feeds.feed_url"`)}
	}
	if v, ok := fc.mutation.FeedURL(); ok {
		if err := entfeeds.FeedURLValidator(v); err != nil {
			return &ValidationError{Name: "feed_url", err: fmt.Errorf(`ent: validator failed for field "Feeds.feed_url": %w`, err)}
		}
	}
	if _, ok := fc.mutation.Language(); !ok {
		return &ValidationError{Name: "language", err: errors.New(`ent: missing required field "Feeds.language"`)}
	}
	if v, ok := fc.mutation.Language(); ok {
		if err := entfeeds.LanguageValidator(v); err != nil {
			return &ValidationError{Name: "language", err: fmt.Errorf(`ent: validator failed for field "Feeds.language": %w`, err)}
		}
	}
	if _, ok := fc.mutation.DtCreated(); !ok {
		return &ValidationError{Name: "dt_created", err: errors.New(`ent: missing required field "Feeds.dt_created"`)}
	}
	if _, ok := fc.mutation.DtUpdated(); !ok {
		return &ValidationError{Name: "dt_updated", err: errors.New(`ent: missing required field "Feeds.dt_updated"`)}
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
		_spec = &sqlgraph.CreateSpec{
			Table: entfeeds.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: entfeeds.FieldID,
			},
		}
	)
	if id, ok := fc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := fc.mutation.SiteURL(); ok {
		_spec.SetField(entfeeds.FieldSiteURL, field.TypeString, value)
		_node.SiteURL = value
	}
	if value, ok := fc.mutation.Title(); ok {
		_spec.SetField(entfeeds.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := fc.mutation.Description(); ok {
		_spec.SetField(entfeeds.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := fc.mutation.FeedURL(); ok {
		_spec.SetField(entfeeds.FieldFeedURL, field.TypeString, value)
		_node.FeedURL = value
	}
	if value, ok := fc.mutation.Language(); ok {
		_spec.SetField(entfeeds.FieldLanguage, field.TypeString, value)
		_node.Language = value
	}
	if value, ok := fc.mutation.DtCreated(); ok {
		_spec.SetField(entfeeds.FieldDtCreated, field.TypeTime, value)
		_node.DtCreated = value
	}
	if value, ok := fc.mutation.DtUpdated(); ok {
		_spec.SetField(entfeeds.FieldDtUpdated, field.TypeTime, value)
		_node.DtUpdated = value
	}
	if value, ok := fc.mutation.Favorites(); ok {
		_spec.SetField(entfeeds.FieldFavorites, field.TypeInt64, value)
		_node.Favorites = value
	}
	return _node, _spec
}

// FeedsCreateBulk is the builder for creating many Feeds entities in bulk.
type FeedsCreateBulk struct {
	config
	builders []*FeedsCreate
}

// Save creates the Feeds entities in the database.
func (fcb *FeedsCreateBulk) Save(ctx context.Context) ([]*Feeds, error) {
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
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, fcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
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
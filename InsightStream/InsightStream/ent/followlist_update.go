// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"insightstream/ent/followlist"
	"insightstream/ent/predicate"
	"insightstream/models/feeds"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// FollowListUpdate is the builder for updating FollowList entities.
type FollowListUpdate struct {
	config
	hooks    []Hook
	mutation *FollowListMutation
}

// Where appends a list predicates to the FollowListUpdate builder.
func (flu *FollowListUpdate) Where(ps ...predicate.FollowList) *FollowListUpdate {
	flu.mutation.Where(ps...)
	return flu
}

// SetUUID sets the "uuid" field.
func (flu *FollowListUpdate) SetUUID(u uuid.UUID) *FollowListUpdate {
	flu.mutation.SetUUID(u)
	return flu
}

// SetNillableUUID sets the "uuid" field if the given value is not nil.
func (flu *FollowListUpdate) SetNillableUUID(u *uuid.UUID) *FollowListUpdate {
	if u != nil {
		flu.SetUUID(*u)
	}
	return flu
}

// SetXMLVersion sets the "xml_version" field.
func (flu *FollowListUpdate) SetXMLVersion(i int8) *FollowListUpdate {
	flu.mutation.ResetXMLVersion()
	flu.mutation.SetXMLVersion(i)
	return flu
}

// SetNillableXMLVersion sets the "xml_version" field if the given value is not nil.
func (flu *FollowListUpdate) SetNillableXMLVersion(i *int8) *FollowListUpdate {
	if i != nil {
		flu.SetXMLVersion(*i)
	}
	return flu
}

// AddXMLVersion adds i to the "xml_version" field.
func (flu *FollowListUpdate) AddXMLVersion(i int8) *FollowListUpdate {
	flu.mutation.AddXMLVersion(i)
	return flu
}

// SetRssVersion sets the "rss_version" field.
func (flu *FollowListUpdate) SetRssVersion(i int8) *FollowListUpdate {
	flu.mutation.ResetRssVersion()
	flu.mutation.SetRssVersion(i)
	return flu
}

// SetNillableRssVersion sets the "rss_version" field if the given value is not nil.
func (flu *FollowListUpdate) SetNillableRssVersion(i *int8) *FollowListUpdate {
	if i != nil {
		flu.SetRssVersion(*i)
	}
	return flu
}

// AddRssVersion adds i to the "rss_version" field.
func (flu *FollowListUpdate) AddRssVersion(i int8) *FollowListUpdate {
	flu.mutation.AddRssVersion(i)
	return flu
}

// SetURL sets the "url" field.
func (flu *FollowListUpdate) SetURL(s string) *FollowListUpdate {
	flu.mutation.SetURL(s)
	return flu
}

// SetTitle sets the "title" field.
func (flu *FollowListUpdate) SetTitle(s string) *FollowListUpdate {
	flu.mutation.SetTitle(s)
	return flu
}

// SetDescription sets the "description" field.
func (flu *FollowListUpdate) SetDescription(s string) *FollowListUpdate {
	flu.mutation.SetDescription(s)
	return flu
}

// SetLink sets the "link" field.
func (flu *FollowListUpdate) SetLink(s string) *FollowListUpdate {
	flu.mutation.SetLink(s)
	return flu
}

// SetLinks sets the "links" field.
func (flu *FollowListUpdate) SetLinks(s string) *FollowListUpdate {
	flu.mutation.SetLinks(s)
	return flu
}

// SetItemDescription sets the "item_description" field.
func (flu *FollowListUpdate) SetItemDescription(fi []feeds.FeedItem) *FollowListUpdate {
	flu.mutation.SetItemDescription(fi)
	return flu
}

// AppendItemDescription appends fi to the "item_description" field.
func (flu *FollowListUpdate) AppendItemDescription(fi []feeds.FeedItem) *FollowListUpdate {
	flu.mutation.AppendItemDescription(fi)
	return flu
}

// SetLanguage sets the "language" field.
func (flu *FollowListUpdate) SetLanguage(s string) *FollowListUpdate {
	flu.mutation.SetLanguage(s)
	return flu
}

// SetDtCreated sets the "dt_created" field.
func (flu *FollowListUpdate) SetDtCreated(t time.Time) *FollowListUpdate {
	flu.mutation.SetDtCreated(t)
	return flu
}

// SetNillableDtCreated sets the "dt_created" field if the given value is not nil.
func (flu *FollowListUpdate) SetNillableDtCreated(t *time.Time) *FollowListUpdate {
	if t != nil {
		flu.SetDtCreated(*t)
	}
	return flu
}

// SetDtUpdated sets the "dt_updated" field.
func (flu *FollowListUpdate) SetDtUpdated(t time.Time) *FollowListUpdate {
	flu.mutation.SetDtUpdated(t)
	return flu
}

// SetNillableDtUpdated sets the "dt_updated" field if the given value is not nil.
func (flu *FollowListUpdate) SetNillableDtUpdated(t *time.Time) *FollowListUpdate {
	if t != nil {
		flu.SetDtUpdated(*t)
	}
	return flu
}

// SetFeedCategory sets the "feed_category" field.
func (flu *FollowListUpdate) SetFeedCategory(i int) *FollowListUpdate {
	flu.mutation.ResetFeedCategory()
	flu.mutation.SetFeedCategory(i)
	return flu
}

// SetNillableFeedCategory sets the "feed_category" field if the given value is not nil.
func (flu *FollowListUpdate) SetNillableFeedCategory(i *int) *FollowListUpdate {
	if i != nil {
		flu.SetFeedCategory(*i)
	}
	return flu
}

// AddFeedCategory adds i to the "feed_category" field.
func (flu *FollowListUpdate) AddFeedCategory(i int) *FollowListUpdate {
	flu.mutation.AddFeedCategory(i)
	return flu
}

// SetIsActive sets the "is_active" field.
func (flu *FollowListUpdate) SetIsActive(b bool) *FollowListUpdate {
	flu.mutation.SetIsActive(b)
	return flu
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (flu *FollowListUpdate) SetNillableIsActive(b *bool) *FollowListUpdate {
	if b != nil {
		flu.SetIsActive(*b)
	}
	return flu
}

// SetIsFavorite sets the "is_favorite" field.
func (flu *FollowListUpdate) SetIsFavorite(b bool) *FollowListUpdate {
	flu.mutation.SetIsFavorite(b)
	return flu
}

// SetNillableIsFavorite sets the "is_favorite" field if the given value is not nil.
func (flu *FollowListUpdate) SetNillableIsFavorite(b *bool) *FollowListUpdate {
	if b != nil {
		flu.SetIsFavorite(*b)
	}
	return flu
}

// SetIsRead sets the "is_read" field.
func (flu *FollowListUpdate) SetIsRead(b bool) *FollowListUpdate {
	flu.mutation.SetIsRead(b)
	return flu
}

// SetNillableIsRead sets the "is_read" field if the given value is not nil.
func (flu *FollowListUpdate) SetNillableIsRead(b *bool) *FollowListUpdate {
	if b != nil {
		flu.SetIsRead(*b)
	}
	return flu
}

// SetIsUpdated sets the "is_updated" field.
func (flu *FollowListUpdate) SetIsUpdated(b bool) *FollowListUpdate {
	flu.mutation.SetIsUpdated(b)
	return flu
}

// SetNillableIsUpdated sets the "is_updated" field if the given value is not nil.
func (flu *FollowListUpdate) SetNillableIsUpdated(b *bool) *FollowListUpdate {
	if b != nil {
		flu.SetIsUpdated(*b)
	}
	return flu
}

// Mutation returns the FollowListMutation object of the builder.
func (flu *FollowListUpdate) Mutation() *FollowListMutation {
	return flu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (flu *FollowListUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, FollowListMutation](ctx, flu.sqlSave, flu.mutation, flu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (flu *FollowListUpdate) SaveX(ctx context.Context) int {
	affected, err := flu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (flu *FollowListUpdate) Exec(ctx context.Context) error {
	_, err := flu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (flu *FollowListUpdate) ExecX(ctx context.Context) {
	if err := flu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (flu *FollowListUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   followlist.Table,
			Columns: followlist.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: followlist.FieldID,
			},
		},
	}
	if ps := flu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := flu.mutation.UUID(); ok {
		_spec.SetField(followlist.FieldUUID, field.TypeUUID, value)
	}
	if value, ok := flu.mutation.XMLVersion(); ok {
		_spec.SetField(followlist.FieldXMLVersion, field.TypeInt8, value)
	}
	if value, ok := flu.mutation.AddedXMLVersion(); ok {
		_spec.AddField(followlist.FieldXMLVersion, field.TypeInt8, value)
	}
	if value, ok := flu.mutation.RssVersion(); ok {
		_spec.SetField(followlist.FieldRssVersion, field.TypeInt8, value)
	}
	if value, ok := flu.mutation.AddedRssVersion(); ok {
		_spec.AddField(followlist.FieldRssVersion, field.TypeInt8, value)
	}
	if value, ok := flu.mutation.URL(); ok {
		_spec.SetField(followlist.FieldURL, field.TypeString, value)
	}
	if value, ok := flu.mutation.Title(); ok {
		_spec.SetField(followlist.FieldTitle, field.TypeString, value)
	}
	if value, ok := flu.mutation.Description(); ok {
		_spec.SetField(followlist.FieldDescription, field.TypeString, value)
	}
	if value, ok := flu.mutation.Link(); ok {
		_spec.SetField(followlist.FieldLink, field.TypeString, value)
	}
	if value, ok := flu.mutation.Links(); ok {
		_spec.SetField(followlist.FieldLinks, field.TypeString, value)
	}
	if value, ok := flu.mutation.ItemDescription(); ok {
		_spec.SetField(followlist.FieldItemDescription, field.TypeJSON, value)
	}
	if value, ok := flu.mutation.AppendedItemDescription(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, followlist.FieldItemDescription, value)
		})
	}
	if value, ok := flu.mutation.Language(); ok {
		_spec.SetField(followlist.FieldLanguage, field.TypeString, value)
	}
	if value, ok := flu.mutation.DtCreated(); ok {
		_spec.SetField(followlist.FieldDtCreated, field.TypeTime, value)
	}
	if value, ok := flu.mutation.DtUpdated(); ok {
		_spec.SetField(followlist.FieldDtUpdated, field.TypeTime, value)
	}
	if value, ok := flu.mutation.FeedCategory(); ok {
		_spec.SetField(followlist.FieldFeedCategory, field.TypeInt, value)
	}
	if value, ok := flu.mutation.AddedFeedCategory(); ok {
		_spec.AddField(followlist.FieldFeedCategory, field.TypeInt, value)
	}
	if value, ok := flu.mutation.IsActive(); ok {
		_spec.SetField(followlist.FieldIsActive, field.TypeBool, value)
	}
	if value, ok := flu.mutation.IsFavorite(); ok {
		_spec.SetField(followlist.FieldIsFavorite, field.TypeBool, value)
	}
	if value, ok := flu.mutation.IsRead(); ok {
		_spec.SetField(followlist.FieldIsRead, field.TypeBool, value)
	}
	if value, ok := flu.mutation.IsUpdated(); ok {
		_spec.SetField(followlist.FieldIsUpdated, field.TypeBool, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, flu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{followlist.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	flu.mutation.done = true
	return n, nil
}

// FollowListUpdateOne is the builder for updating a single FollowList entity.
type FollowListUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FollowListMutation
}

// SetUUID sets the "uuid" field.
func (fluo *FollowListUpdateOne) SetUUID(u uuid.UUID) *FollowListUpdateOne {
	fluo.mutation.SetUUID(u)
	return fluo
}

// SetNillableUUID sets the "uuid" field if the given value is not nil.
func (fluo *FollowListUpdateOne) SetNillableUUID(u *uuid.UUID) *FollowListUpdateOne {
	if u != nil {
		fluo.SetUUID(*u)
	}
	return fluo
}

// SetXMLVersion sets the "xml_version" field.
func (fluo *FollowListUpdateOne) SetXMLVersion(i int8) *FollowListUpdateOne {
	fluo.mutation.ResetXMLVersion()
	fluo.mutation.SetXMLVersion(i)
	return fluo
}

// SetNillableXMLVersion sets the "xml_version" field if the given value is not nil.
func (fluo *FollowListUpdateOne) SetNillableXMLVersion(i *int8) *FollowListUpdateOne {
	if i != nil {
		fluo.SetXMLVersion(*i)
	}
	return fluo
}

// AddXMLVersion adds i to the "xml_version" field.
func (fluo *FollowListUpdateOne) AddXMLVersion(i int8) *FollowListUpdateOne {
	fluo.mutation.AddXMLVersion(i)
	return fluo
}

// SetRssVersion sets the "rss_version" field.
func (fluo *FollowListUpdateOne) SetRssVersion(i int8) *FollowListUpdateOne {
	fluo.mutation.ResetRssVersion()
	fluo.mutation.SetRssVersion(i)
	return fluo
}

// SetNillableRssVersion sets the "rss_version" field if the given value is not nil.
func (fluo *FollowListUpdateOne) SetNillableRssVersion(i *int8) *FollowListUpdateOne {
	if i != nil {
		fluo.SetRssVersion(*i)
	}
	return fluo
}

// AddRssVersion adds i to the "rss_version" field.
func (fluo *FollowListUpdateOne) AddRssVersion(i int8) *FollowListUpdateOne {
	fluo.mutation.AddRssVersion(i)
	return fluo
}

// SetURL sets the "url" field.
func (fluo *FollowListUpdateOne) SetURL(s string) *FollowListUpdateOne {
	fluo.mutation.SetURL(s)
	return fluo
}

// SetTitle sets the "title" field.
func (fluo *FollowListUpdateOne) SetTitle(s string) *FollowListUpdateOne {
	fluo.mutation.SetTitle(s)
	return fluo
}

// SetDescription sets the "description" field.
func (fluo *FollowListUpdateOne) SetDescription(s string) *FollowListUpdateOne {
	fluo.mutation.SetDescription(s)
	return fluo
}

// SetLink sets the "link" field.
func (fluo *FollowListUpdateOne) SetLink(s string) *FollowListUpdateOne {
	fluo.mutation.SetLink(s)
	return fluo
}

// SetLinks sets the "links" field.
func (fluo *FollowListUpdateOne) SetLinks(s string) *FollowListUpdateOne {
	fluo.mutation.SetLinks(s)
	return fluo
}

// SetItemDescription sets the "item_description" field.
func (fluo *FollowListUpdateOne) SetItemDescription(fi []feeds.FeedItem) *FollowListUpdateOne {
	fluo.mutation.SetItemDescription(fi)
	return fluo
}

// AppendItemDescription appends fi to the "item_description" field.
func (fluo *FollowListUpdateOne) AppendItemDescription(fi []feeds.FeedItem) *FollowListUpdateOne {
	fluo.mutation.AppendItemDescription(fi)
	return fluo
}

// SetLanguage sets the "language" field.
func (fluo *FollowListUpdateOne) SetLanguage(s string) *FollowListUpdateOne {
	fluo.mutation.SetLanguage(s)
	return fluo
}

// SetDtCreated sets the "dt_created" field.
func (fluo *FollowListUpdateOne) SetDtCreated(t time.Time) *FollowListUpdateOne {
	fluo.mutation.SetDtCreated(t)
	return fluo
}

// SetNillableDtCreated sets the "dt_created" field if the given value is not nil.
func (fluo *FollowListUpdateOne) SetNillableDtCreated(t *time.Time) *FollowListUpdateOne {
	if t != nil {
		fluo.SetDtCreated(*t)
	}
	return fluo
}

// SetDtUpdated sets the "dt_updated" field.
func (fluo *FollowListUpdateOne) SetDtUpdated(t time.Time) *FollowListUpdateOne {
	fluo.mutation.SetDtUpdated(t)
	return fluo
}

// SetNillableDtUpdated sets the "dt_updated" field if the given value is not nil.
func (fluo *FollowListUpdateOne) SetNillableDtUpdated(t *time.Time) *FollowListUpdateOne {
	if t != nil {
		fluo.SetDtUpdated(*t)
	}
	return fluo
}

// SetFeedCategory sets the "feed_category" field.
func (fluo *FollowListUpdateOne) SetFeedCategory(i int) *FollowListUpdateOne {
	fluo.mutation.ResetFeedCategory()
	fluo.mutation.SetFeedCategory(i)
	return fluo
}

// SetNillableFeedCategory sets the "feed_category" field if the given value is not nil.
func (fluo *FollowListUpdateOne) SetNillableFeedCategory(i *int) *FollowListUpdateOne {
	if i != nil {
		fluo.SetFeedCategory(*i)
	}
	return fluo
}

// AddFeedCategory adds i to the "feed_category" field.
func (fluo *FollowListUpdateOne) AddFeedCategory(i int) *FollowListUpdateOne {
	fluo.mutation.AddFeedCategory(i)
	return fluo
}

// SetIsActive sets the "is_active" field.
func (fluo *FollowListUpdateOne) SetIsActive(b bool) *FollowListUpdateOne {
	fluo.mutation.SetIsActive(b)
	return fluo
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (fluo *FollowListUpdateOne) SetNillableIsActive(b *bool) *FollowListUpdateOne {
	if b != nil {
		fluo.SetIsActive(*b)
	}
	return fluo
}

// SetIsFavorite sets the "is_favorite" field.
func (fluo *FollowListUpdateOne) SetIsFavorite(b bool) *FollowListUpdateOne {
	fluo.mutation.SetIsFavorite(b)
	return fluo
}

// SetNillableIsFavorite sets the "is_favorite" field if the given value is not nil.
func (fluo *FollowListUpdateOne) SetNillableIsFavorite(b *bool) *FollowListUpdateOne {
	if b != nil {
		fluo.SetIsFavorite(*b)
	}
	return fluo
}

// SetIsRead sets the "is_read" field.
func (fluo *FollowListUpdateOne) SetIsRead(b bool) *FollowListUpdateOne {
	fluo.mutation.SetIsRead(b)
	return fluo
}

// SetNillableIsRead sets the "is_read" field if the given value is not nil.
func (fluo *FollowListUpdateOne) SetNillableIsRead(b *bool) *FollowListUpdateOne {
	if b != nil {
		fluo.SetIsRead(*b)
	}
	return fluo
}

// SetIsUpdated sets the "is_updated" field.
func (fluo *FollowListUpdateOne) SetIsUpdated(b bool) *FollowListUpdateOne {
	fluo.mutation.SetIsUpdated(b)
	return fluo
}

// SetNillableIsUpdated sets the "is_updated" field if the given value is not nil.
func (fluo *FollowListUpdateOne) SetNillableIsUpdated(b *bool) *FollowListUpdateOne {
	if b != nil {
		fluo.SetIsUpdated(*b)
	}
	return fluo
}

// Mutation returns the FollowListMutation object of the builder.
func (fluo *FollowListUpdateOne) Mutation() *FollowListMutation {
	return fluo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fluo *FollowListUpdateOne) Select(field string, fields ...string) *FollowListUpdateOne {
	fluo.fields = append([]string{field}, fields...)
	return fluo
}

// Save executes the query and returns the updated FollowList entity.
func (fluo *FollowListUpdateOne) Save(ctx context.Context) (*FollowList, error) {
	return withHooks[*FollowList, FollowListMutation](ctx, fluo.sqlSave, fluo.mutation, fluo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fluo *FollowListUpdateOne) SaveX(ctx context.Context) *FollowList {
	node, err := fluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fluo *FollowListUpdateOne) Exec(ctx context.Context) error {
	_, err := fluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fluo *FollowListUpdateOne) ExecX(ctx context.Context) {
	if err := fluo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fluo *FollowListUpdateOne) sqlSave(ctx context.Context) (_node *FollowList, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   followlist.Table,
			Columns: followlist.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: followlist.FieldID,
			},
		},
	}
	id, ok := fluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "FollowList.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, followlist.FieldID)
		for _, f := range fields {
			if !followlist.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != followlist.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fluo.mutation.UUID(); ok {
		_spec.SetField(followlist.FieldUUID, field.TypeUUID, value)
	}
	if value, ok := fluo.mutation.XMLVersion(); ok {
		_spec.SetField(followlist.FieldXMLVersion, field.TypeInt8, value)
	}
	if value, ok := fluo.mutation.AddedXMLVersion(); ok {
		_spec.AddField(followlist.FieldXMLVersion, field.TypeInt8, value)
	}
	if value, ok := fluo.mutation.RssVersion(); ok {
		_spec.SetField(followlist.FieldRssVersion, field.TypeInt8, value)
	}
	if value, ok := fluo.mutation.AddedRssVersion(); ok {
		_spec.AddField(followlist.FieldRssVersion, field.TypeInt8, value)
	}
	if value, ok := fluo.mutation.URL(); ok {
		_spec.SetField(followlist.FieldURL, field.TypeString, value)
	}
	if value, ok := fluo.mutation.Title(); ok {
		_spec.SetField(followlist.FieldTitle, field.TypeString, value)
	}
	if value, ok := fluo.mutation.Description(); ok {
		_spec.SetField(followlist.FieldDescription, field.TypeString, value)
	}
	if value, ok := fluo.mutation.Link(); ok {
		_spec.SetField(followlist.FieldLink, field.TypeString, value)
	}
	if value, ok := fluo.mutation.Links(); ok {
		_spec.SetField(followlist.FieldLinks, field.TypeString, value)
	}
	if value, ok := fluo.mutation.ItemDescription(); ok {
		_spec.SetField(followlist.FieldItemDescription, field.TypeJSON, value)
	}
	if value, ok := fluo.mutation.AppendedItemDescription(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, followlist.FieldItemDescription, value)
		})
	}
	if value, ok := fluo.mutation.Language(); ok {
		_spec.SetField(followlist.FieldLanguage, field.TypeString, value)
	}
	if value, ok := fluo.mutation.DtCreated(); ok {
		_spec.SetField(followlist.FieldDtCreated, field.TypeTime, value)
	}
	if value, ok := fluo.mutation.DtUpdated(); ok {
		_spec.SetField(followlist.FieldDtUpdated, field.TypeTime, value)
	}
	if value, ok := fluo.mutation.FeedCategory(); ok {
		_spec.SetField(followlist.FieldFeedCategory, field.TypeInt, value)
	}
	if value, ok := fluo.mutation.AddedFeedCategory(); ok {
		_spec.AddField(followlist.FieldFeedCategory, field.TypeInt, value)
	}
	if value, ok := fluo.mutation.IsActive(); ok {
		_spec.SetField(followlist.FieldIsActive, field.TypeBool, value)
	}
	if value, ok := fluo.mutation.IsFavorite(); ok {
		_spec.SetField(followlist.FieldIsFavorite, field.TypeBool, value)
	}
	if value, ok := fluo.mutation.IsRead(); ok {
		_spec.SetField(followlist.FieldIsRead, field.TypeBool, value)
	}
	if value, ok := fluo.mutation.IsUpdated(); ok {
		_spec.SetField(followlist.FieldIsUpdated, field.TypeBool, value)
	}
	_node = &FollowList{config: fluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{followlist.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	fluo.mutation.done = true
	return _node, nil
}

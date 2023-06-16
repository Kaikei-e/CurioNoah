// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"insightstream/ent/predicate"
	"insightstream/ent/users"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// UsersQuery is the builder for querying Users entities.
type UsersQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	inters     []Interceptor
	predicates []predicate.Users
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UsersQuery builder.
func (uq *UsersQuery) Where(ps ...predicate.Users) *UsersQuery {
	uq.predicates = append(uq.predicates, ps...)
	return uq
}

// Limit the number of records to be returned by this query.
func (uq *UsersQuery) Limit(limit int) *UsersQuery {
	uq.limit = &limit
	return uq
}

// Offset to start from.
func (uq *UsersQuery) Offset(offset int) *UsersQuery {
	uq.offset = &offset
	return uq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (uq *UsersQuery) Unique(unique bool) *UsersQuery {
	uq.unique = &unique
	return uq
}

// Order specifies how the records should be ordered.
func (uq *UsersQuery) Order(o ...OrderFunc) *UsersQuery {
	uq.order = append(uq.order, o...)
	return uq
}

// First returns the first Users entity from the query.
// Returns a *NotFoundError when no Users was found.
func (uq *UsersQuery) First(ctx context.Context) (*Users, error) {
	nodes, err := uq.Limit(1).All(newQueryContext(ctx, TypeUsers, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{users.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (uq *UsersQuery) FirstX(ctx context.Context) *Users {
	node, err := uq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Users ID from the query.
// Returns a *NotFoundError when no Users ID was found.
func (uq *UsersQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = uq.Limit(1).IDs(newQueryContext(ctx, TypeUsers, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{users.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (uq *UsersQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := uq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Users entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Users entity is found.
// Returns a *NotFoundError when no Users entities are found.
func (uq *UsersQuery) Only(ctx context.Context) (*Users, error) {
	nodes, err := uq.Limit(2).All(newQueryContext(ctx, TypeUsers, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{users.Label}
	default:
		return nil, &NotSingularError{users.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (uq *UsersQuery) OnlyX(ctx context.Context) *Users {
	node, err := uq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Users ID in the query.
// Returns a *NotSingularError when more than one Users ID is found.
// Returns a *NotFoundError when no entities are found.
func (uq *UsersQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = uq.Limit(2).IDs(newQueryContext(ctx, TypeUsers, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{users.Label}
	default:
		err = &NotSingularError{users.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (uq *UsersQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := uq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UsersSlice.
func (uq *UsersQuery) All(ctx context.Context) ([]*Users, error) {
	ctx = newQueryContext(ctx, TypeUsers, "All")
	if err := uq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Users, *UsersQuery]()
	return withInterceptors[[]*Users](ctx, uq, qr, uq.inters)
}

// AllX is like All, but panics if an error occurs.
func (uq *UsersQuery) AllX(ctx context.Context) []*Users {
	nodes, err := uq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Users IDs.
func (uq *UsersQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	ctx = newQueryContext(ctx, TypeUsers, "IDs")
	if err := uq.Select(users.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (uq *UsersQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := uq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (uq *UsersQuery) Count(ctx context.Context) (int, error) {
	ctx = newQueryContext(ctx, TypeUsers, "Count")
	if err := uq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, uq, querierCount[*UsersQuery](), uq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (uq *UsersQuery) CountX(ctx context.Context) int {
	count, err := uq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (uq *UsersQuery) Exist(ctx context.Context) (bool, error) {
	ctx = newQueryContext(ctx, TypeUsers, "Exist")
	switch _, err := uq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (uq *UsersQuery) ExistX(ctx context.Context) bool {
	exist, err := uq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UsersQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (uq *UsersQuery) Clone() *UsersQuery {
	if uq == nil {
		return nil
	}
	return &UsersQuery{
		config:     uq.config,
		limit:      uq.limit,
		offset:     uq.offset,
		order:      append([]OrderFunc{}, uq.order...),
		inters:     append([]Interceptor{}, uq.inters...),
		predicates: append([]predicate.Users{}, uq.predicates...),
		// clone intermediate query.
		sql:    uq.sql.Clone(),
		path:   uq.path,
		unique: uq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Username string `json:"username,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Users.Query().
//		GroupBy(users.FieldUsername).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (uq *UsersQuery) GroupBy(field string, fields ...string) *UsersGroupBy {
	uq.fields = append([]string{field}, fields...)
	grbuild := &UsersGroupBy{build: uq}
	grbuild.flds = &uq.fields
	grbuild.label = users.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Username string `json:"username,omitempty"`
//	}
//
//	client.Users.Query().
//		Select(users.FieldUsername).
//		Scan(ctx, &v)
func (uq *UsersQuery) Select(fields ...string) *UsersSelect {
	uq.fields = append(uq.fields, fields...)
	sbuild := &UsersSelect{UsersQuery: uq}
	sbuild.label = users.Label
	sbuild.flds, sbuild.scan = &uq.fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UsersSelect configured with the given aggregations.
func (uq *UsersQuery) Aggregate(fns ...AggregateFunc) *UsersSelect {
	return uq.Select().Aggregate(fns...)
}

func (uq *UsersQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range uq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, uq); err != nil {
				return err
			}
		}
	}
	for _, f := range uq.fields {
		if !users.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if uq.path != nil {
		prev, err := uq.path(ctx)
		if err != nil {
			return err
		}
		uq.sql = prev
	}
	return nil
}

func (uq *UsersQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Users, error) {
	var (
		nodes = []*Users{}
		_spec = uq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Users).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Users{config: uq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, uq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (uq *UsersQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := uq.querySpec()
	_spec.Node.Columns = uq.fields
	if len(uq.fields) > 0 {
		_spec.Unique = uq.unique != nil && *uq.unique
	}
	return sqlgraph.CountNodes(ctx, uq.driver, _spec)
}

func (uq *UsersQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   users.Table,
			Columns: users.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: users.FieldID,
			},
		},
		From:   uq.sql,
		Unique: true,
	}
	if unique := uq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := uq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, users.FieldID)
		for i := range fields {
			if fields[i] != users.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := uq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := uq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := uq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := uq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (uq *UsersQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(uq.driver.Dialect())
	t1 := builder.Table(users.Table)
	columns := uq.fields
	if len(columns) == 0 {
		columns = users.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if uq.sql != nil {
		selector = uq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if uq.unique != nil && *uq.unique {
		selector.Distinct()
	}
	for _, p := range uq.predicates {
		p(selector)
	}
	for _, p := range uq.order {
		p(selector)
	}
	if offset := uq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := uq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UsersGroupBy is the group-by builder for Users entities.
type UsersGroupBy struct {
	selector
	build *UsersQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ugb *UsersGroupBy) Aggregate(fns ...AggregateFunc) *UsersGroupBy {
	ugb.fns = append(ugb.fns, fns...)
	return ugb
}

// Scan applies the selector query and scans the result into the given value.
func (ugb *UsersGroupBy) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeUsers, "GroupBy")
	if err := ugb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UsersQuery, *UsersGroupBy](ctx, ugb.build, ugb, ugb.build.inters, v)
}

func (ugb *UsersGroupBy) sqlScan(ctx context.Context, root *UsersQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ugb.fns))
	for _, fn := range ugb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ugb.flds)+len(ugb.fns))
		for _, f := range *ugb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ugb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ugb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UsersSelect is the builder for selecting fields of Users entities.
type UsersSelect struct {
	*UsersQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (us *UsersSelect) Aggregate(fns ...AggregateFunc) *UsersSelect {
	us.fns = append(us.fns, fns...)
	return us
}

// Scan applies the selector query and scans the result into the given value.
func (us *UsersSelect) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeUsers, "Select")
	if err := us.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UsersQuery, *UsersSelect](ctx, us.UsersQuery, us, us.inters, v)
}

func (us *UsersSelect) sqlScan(ctx context.Context, root *UsersQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(us.fns))
	for _, fn := range us.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*us.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := us.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

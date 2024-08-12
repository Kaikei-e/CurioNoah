// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"insightstream/ent/followlists"
	"insightstream/ent/predicate"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FollowListsQuery is the builder for querying FollowLists entities.
type FollowListsQuery struct {
	config
	ctx        *QueryContext
	order      []followlists.OrderOption
	inters     []Interceptor
	predicates []predicate.FollowLists
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FollowListsQuery builder.
func (flq *FollowListsQuery) Where(ps ...predicate.FollowLists) *FollowListsQuery {
	flq.predicates = append(flq.predicates, ps...)
	return flq
}

// Limit the number of records to be returned by this query.
func (flq *FollowListsQuery) Limit(limit int) *FollowListsQuery {
	flq.ctx.Limit = &limit
	return flq
}

// Offset to start from.
func (flq *FollowListsQuery) Offset(offset int) *FollowListsQuery {
	flq.ctx.Offset = &offset
	return flq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (flq *FollowListsQuery) Unique(unique bool) *FollowListsQuery {
	flq.ctx.Unique = &unique
	return flq
}

// Order specifies how the records should be ordered.
func (flq *FollowListsQuery) Order(o ...followlists.OrderOption) *FollowListsQuery {
	flq.order = append(flq.order, o...)
	return flq
}

// First returns the first FollowLists entity from the query.
// Returns a *NotFoundError when no FollowLists was found.
func (flq *FollowListsQuery) First(ctx context.Context) (*FollowLists, error) {
	nodes, err := flq.Limit(1).All(setContextOp(ctx, flq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{followlists.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (flq *FollowListsQuery) FirstX(ctx context.Context) *FollowLists {
	node, err := flq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first FollowLists ID from the query.
// Returns a *NotFoundError when no FollowLists ID was found.
func (flq *FollowListsQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = flq.Limit(1).IDs(setContextOp(ctx, flq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{followlists.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (flq *FollowListsQuery) FirstIDX(ctx context.Context) int {
	id, err := flq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single FollowLists entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one FollowLists entity is found.
// Returns a *NotFoundError when no FollowLists entities are found.
func (flq *FollowListsQuery) Only(ctx context.Context) (*FollowLists, error) {
	nodes, err := flq.Limit(2).All(setContextOp(ctx, flq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{followlists.Label}
	default:
		return nil, &NotSingularError{followlists.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (flq *FollowListsQuery) OnlyX(ctx context.Context) *FollowLists {
	node, err := flq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only FollowLists ID in the query.
// Returns a *NotSingularError when more than one FollowLists ID is found.
// Returns a *NotFoundError when no entities are found.
func (flq *FollowListsQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = flq.Limit(2).IDs(setContextOp(ctx, flq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{followlists.Label}
	default:
		err = &NotSingularError{followlists.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (flq *FollowListsQuery) OnlyIDX(ctx context.Context) int {
	id, err := flq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FollowListsSlice.
func (flq *FollowListsQuery) All(ctx context.Context) ([]*FollowLists, error) {
	ctx = setContextOp(ctx, flq.ctx, ent.OpQueryAll)
	if err := flq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*FollowLists, *FollowListsQuery]()
	return withInterceptors[[]*FollowLists](ctx, flq, qr, flq.inters)
}

// AllX is like All, but panics if an error occurs.
func (flq *FollowListsQuery) AllX(ctx context.Context) []*FollowLists {
	nodes, err := flq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of FollowLists IDs.
func (flq *FollowListsQuery) IDs(ctx context.Context) (ids []int, err error) {
	if flq.ctx.Unique == nil && flq.path != nil {
		flq.Unique(true)
	}
	ctx = setContextOp(ctx, flq.ctx, ent.OpQueryIDs)
	if err = flq.Select(followlists.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (flq *FollowListsQuery) IDsX(ctx context.Context) []int {
	ids, err := flq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (flq *FollowListsQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, flq.ctx, ent.OpQueryCount)
	if err := flq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, flq, querierCount[*FollowListsQuery](), flq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (flq *FollowListsQuery) CountX(ctx context.Context) int {
	count, err := flq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (flq *FollowListsQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, flq.ctx, ent.OpQueryExist)
	switch _, err := flq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (flq *FollowListsQuery) ExistX(ctx context.Context) bool {
	exist, err := flq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FollowListsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (flq *FollowListsQuery) Clone() *FollowListsQuery {
	if flq == nil {
		return nil
	}
	return &FollowListsQuery{
		config:     flq.config,
		ctx:        flq.ctx.Clone(),
		order:      append([]followlists.OrderOption{}, flq.order...),
		inters:     append([]Interceptor{}, flq.inters...),
		predicates: append([]predicate.FollowLists{}, flq.predicates...),
		// clone intermediate query.
		sql:  flq.sql.Clone(),
		path: flq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UUID uuid.UUID `json:"uuid,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.FollowLists.Query().
//		GroupBy(followlists.FieldUUID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (flq *FollowListsQuery) GroupBy(field string, fields ...string) *FollowListsGroupBy {
	flq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &FollowListsGroupBy{build: flq}
	grbuild.flds = &flq.ctx.Fields
	grbuild.label = followlists.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UUID uuid.UUID `json:"uuid,omitempty"`
//	}
//
//	client.FollowLists.Query().
//		Select(followlists.FieldUUID).
//		Scan(ctx, &v)
func (flq *FollowListsQuery) Select(fields ...string) *FollowListsSelect {
	flq.ctx.Fields = append(flq.ctx.Fields, fields...)
	sbuild := &FollowListsSelect{FollowListsQuery: flq}
	sbuild.label = followlists.Label
	sbuild.flds, sbuild.scan = &flq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a FollowListsSelect configured with the given aggregations.
func (flq *FollowListsQuery) Aggregate(fns ...AggregateFunc) *FollowListsSelect {
	return flq.Select().Aggregate(fns...)
}

func (flq *FollowListsQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range flq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, flq); err != nil {
				return err
			}
		}
	}
	for _, f := range flq.ctx.Fields {
		if !followlists.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if flq.path != nil {
		prev, err := flq.path(ctx)
		if err != nil {
			return err
		}
		flq.sql = prev
	}
	return nil
}

func (flq *FollowListsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*FollowLists, error) {
	var (
		nodes = []*FollowLists{}
		_spec = flq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*FollowLists).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &FollowLists{config: flq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, flq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (flq *FollowListsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := flq.querySpec()
	_spec.Node.Columns = flq.ctx.Fields
	if len(flq.ctx.Fields) > 0 {
		_spec.Unique = flq.ctx.Unique != nil && *flq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, flq.driver, _spec)
}

func (flq *FollowListsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(followlists.Table, followlists.Columns, sqlgraph.NewFieldSpec(followlists.FieldID, field.TypeInt))
	_spec.From = flq.sql
	if unique := flq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if flq.path != nil {
		_spec.Unique = true
	}
	if fields := flq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, followlists.FieldID)
		for i := range fields {
			if fields[i] != followlists.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := flq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := flq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := flq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := flq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (flq *FollowListsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(flq.driver.Dialect())
	t1 := builder.Table(followlists.Table)
	columns := flq.ctx.Fields
	if len(columns) == 0 {
		columns = followlists.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if flq.sql != nil {
		selector = flq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if flq.ctx.Unique != nil && *flq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range flq.predicates {
		p(selector)
	}
	for _, p := range flq.order {
		p(selector)
	}
	if offset := flq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := flq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FollowListsGroupBy is the group-by builder for FollowLists entities.
type FollowListsGroupBy struct {
	selector
	build *FollowListsQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (flgb *FollowListsGroupBy) Aggregate(fns ...AggregateFunc) *FollowListsGroupBy {
	flgb.fns = append(flgb.fns, fns...)
	return flgb
}

// Scan applies the selector query and scans the result into the given value.
func (flgb *FollowListsGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, flgb.build.ctx, ent.OpQueryGroupBy)
	if err := flgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FollowListsQuery, *FollowListsGroupBy](ctx, flgb.build, flgb, flgb.build.inters, v)
}

func (flgb *FollowListsGroupBy) sqlScan(ctx context.Context, root *FollowListsQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(flgb.fns))
	for _, fn := range flgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*flgb.flds)+len(flgb.fns))
		for _, f := range *flgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*flgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := flgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// FollowListsSelect is the builder for selecting fields of FollowLists entities.
type FollowListsSelect struct {
	*FollowListsQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (fls *FollowListsSelect) Aggregate(fns ...AggregateFunc) *FollowListsSelect {
	fls.fns = append(fls.fns, fns...)
	return fls
}

// Scan applies the selector query and scans the result into the given value.
func (fls *FollowListsSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fls.ctx, ent.OpQuerySelect)
	if err := fls.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FollowListsQuery, *FollowListsSelect](ctx, fls.FollowListsQuery, fls, fls.inters, v)
}

func (fls *FollowListsSelect) sqlScan(ctx context.Context, root *FollowListsQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(fls.fns))
	for _, fn := range fls.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*fls.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

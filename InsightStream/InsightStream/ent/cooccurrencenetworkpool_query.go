// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"insightstream/ent/cooccurrencenetworkpool"
	"insightstream/ent/predicate"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// CooccurrenceNetworkPoolQuery is the builder for querying CooccurrenceNetworkPool entities.
type CooccurrenceNetworkPoolQuery struct {
	config
	ctx        *QueryContext
	order      []cooccurrencenetworkpool.OrderOption
	inters     []Interceptor
	predicates []predicate.CooccurrenceNetworkPool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CooccurrenceNetworkPoolQuery builder.
func (cnpq *CooccurrenceNetworkPoolQuery) Where(ps ...predicate.CooccurrenceNetworkPool) *CooccurrenceNetworkPoolQuery {
	cnpq.predicates = append(cnpq.predicates, ps...)
	return cnpq
}

// Limit the number of records to be returned by this query.
func (cnpq *CooccurrenceNetworkPoolQuery) Limit(limit int) *CooccurrenceNetworkPoolQuery {
	cnpq.ctx.Limit = &limit
	return cnpq
}

// Offset to start from.
func (cnpq *CooccurrenceNetworkPoolQuery) Offset(offset int) *CooccurrenceNetworkPoolQuery {
	cnpq.ctx.Offset = &offset
	return cnpq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cnpq *CooccurrenceNetworkPoolQuery) Unique(unique bool) *CooccurrenceNetworkPoolQuery {
	cnpq.ctx.Unique = &unique
	return cnpq
}

// Order specifies how the records should be ordered.
func (cnpq *CooccurrenceNetworkPoolQuery) Order(o ...cooccurrencenetworkpool.OrderOption) *CooccurrenceNetworkPoolQuery {
	cnpq.order = append(cnpq.order, o...)
	return cnpq
}

// First returns the first CooccurrenceNetworkPool entity from the query.
// Returns a *NotFoundError when no CooccurrenceNetworkPool was found.
func (cnpq *CooccurrenceNetworkPoolQuery) First(ctx context.Context) (*CooccurrenceNetworkPool, error) {
	nodes, err := cnpq.Limit(1).All(setContextOp(ctx, cnpq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{cooccurrencenetworkpool.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cnpq *CooccurrenceNetworkPoolQuery) FirstX(ctx context.Context) *CooccurrenceNetworkPool {
	node, err := cnpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CooccurrenceNetworkPool ID from the query.
// Returns a *NotFoundError when no CooccurrenceNetworkPool ID was found.
func (cnpq *CooccurrenceNetworkPoolQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = cnpq.Limit(1).IDs(setContextOp(ctx, cnpq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{cooccurrencenetworkpool.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cnpq *CooccurrenceNetworkPoolQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := cnpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CooccurrenceNetworkPool entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CooccurrenceNetworkPool entity is found.
// Returns a *NotFoundError when no CooccurrenceNetworkPool entities are found.
func (cnpq *CooccurrenceNetworkPoolQuery) Only(ctx context.Context) (*CooccurrenceNetworkPool, error) {
	nodes, err := cnpq.Limit(2).All(setContextOp(ctx, cnpq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{cooccurrencenetworkpool.Label}
	default:
		return nil, &NotSingularError{cooccurrencenetworkpool.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cnpq *CooccurrenceNetworkPoolQuery) OnlyX(ctx context.Context) *CooccurrenceNetworkPool {
	node, err := cnpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CooccurrenceNetworkPool ID in the query.
// Returns a *NotSingularError when more than one CooccurrenceNetworkPool ID is found.
// Returns a *NotFoundError when no entities are found.
func (cnpq *CooccurrenceNetworkPoolQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = cnpq.Limit(2).IDs(setContextOp(ctx, cnpq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{cooccurrencenetworkpool.Label}
	default:
		err = &NotSingularError{cooccurrencenetworkpool.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cnpq *CooccurrenceNetworkPoolQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := cnpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CooccurrenceNetworkPools.
func (cnpq *CooccurrenceNetworkPoolQuery) All(ctx context.Context) ([]*CooccurrenceNetworkPool, error) {
	ctx = setContextOp(ctx, cnpq.ctx, ent.OpQueryAll)
	if err := cnpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*CooccurrenceNetworkPool, *CooccurrenceNetworkPoolQuery]()
	return withInterceptors[[]*CooccurrenceNetworkPool](ctx, cnpq, qr, cnpq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cnpq *CooccurrenceNetworkPoolQuery) AllX(ctx context.Context) []*CooccurrenceNetworkPool {
	nodes, err := cnpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CooccurrenceNetworkPool IDs.
func (cnpq *CooccurrenceNetworkPoolQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if cnpq.ctx.Unique == nil && cnpq.path != nil {
		cnpq.Unique(true)
	}
	ctx = setContextOp(ctx, cnpq.ctx, ent.OpQueryIDs)
	if err = cnpq.Select(cooccurrencenetworkpool.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cnpq *CooccurrenceNetworkPoolQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := cnpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cnpq *CooccurrenceNetworkPoolQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cnpq.ctx, ent.OpQueryCount)
	if err := cnpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cnpq, querierCount[*CooccurrenceNetworkPoolQuery](), cnpq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cnpq *CooccurrenceNetworkPoolQuery) CountX(ctx context.Context) int {
	count, err := cnpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cnpq *CooccurrenceNetworkPoolQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cnpq.ctx, ent.OpQueryExist)
	switch _, err := cnpq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cnpq *CooccurrenceNetworkPoolQuery) ExistX(ctx context.Context) bool {
	exist, err := cnpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CooccurrenceNetworkPoolQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cnpq *CooccurrenceNetworkPoolQuery) Clone() *CooccurrenceNetworkPoolQuery {
	if cnpq == nil {
		return nil
	}
	return &CooccurrenceNetworkPoolQuery{
		config:     cnpq.config,
		ctx:        cnpq.ctx.Clone(),
		order:      append([]cooccurrencenetworkpool.OrderOption{}, cnpq.order...),
		inters:     append([]Interceptor{}, cnpq.inters...),
		predicates: append([]predicate.CooccurrenceNetworkPool{}, cnpq.predicates...),
		// clone intermediate query.
		sql:  cnpq.sql.Clone(),
		path: cnpq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		SiteURL string `json:"site_url,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.CooccurrenceNetworkPool.Query().
//		GroupBy(cooccurrencenetworkpool.FieldSiteURL).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cnpq *CooccurrenceNetworkPoolQuery) GroupBy(field string, fields ...string) *CooccurrenceNetworkPoolGroupBy {
	cnpq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CooccurrenceNetworkPoolGroupBy{build: cnpq}
	grbuild.flds = &cnpq.ctx.Fields
	grbuild.label = cooccurrencenetworkpool.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		SiteURL string `json:"site_url,omitempty"`
//	}
//
//	client.CooccurrenceNetworkPool.Query().
//		Select(cooccurrencenetworkpool.FieldSiteURL).
//		Scan(ctx, &v)
func (cnpq *CooccurrenceNetworkPoolQuery) Select(fields ...string) *CooccurrenceNetworkPoolSelect {
	cnpq.ctx.Fields = append(cnpq.ctx.Fields, fields...)
	sbuild := &CooccurrenceNetworkPoolSelect{CooccurrenceNetworkPoolQuery: cnpq}
	sbuild.label = cooccurrencenetworkpool.Label
	sbuild.flds, sbuild.scan = &cnpq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CooccurrenceNetworkPoolSelect configured with the given aggregations.
func (cnpq *CooccurrenceNetworkPoolQuery) Aggregate(fns ...AggregateFunc) *CooccurrenceNetworkPoolSelect {
	return cnpq.Select().Aggregate(fns...)
}

func (cnpq *CooccurrenceNetworkPoolQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cnpq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cnpq); err != nil {
				return err
			}
		}
	}
	for _, f := range cnpq.ctx.Fields {
		if !cooccurrencenetworkpool.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cnpq.path != nil {
		prev, err := cnpq.path(ctx)
		if err != nil {
			return err
		}
		cnpq.sql = prev
	}
	return nil
}

func (cnpq *CooccurrenceNetworkPoolQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CooccurrenceNetworkPool, error) {
	var (
		nodes = []*CooccurrenceNetworkPool{}
		_spec = cnpq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*CooccurrenceNetworkPool).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &CooccurrenceNetworkPool{config: cnpq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cnpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (cnpq *CooccurrenceNetworkPoolQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cnpq.querySpec()
	_spec.Node.Columns = cnpq.ctx.Fields
	if len(cnpq.ctx.Fields) > 0 {
		_spec.Unique = cnpq.ctx.Unique != nil && *cnpq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cnpq.driver, _spec)
}

func (cnpq *CooccurrenceNetworkPoolQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(cooccurrencenetworkpool.Table, cooccurrencenetworkpool.Columns, sqlgraph.NewFieldSpec(cooccurrencenetworkpool.FieldID, field.TypeUUID))
	_spec.From = cnpq.sql
	if unique := cnpq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cnpq.path != nil {
		_spec.Unique = true
	}
	if fields := cnpq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, cooccurrencenetworkpool.FieldID)
		for i := range fields {
			if fields[i] != cooccurrencenetworkpool.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cnpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cnpq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cnpq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cnpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cnpq *CooccurrenceNetworkPoolQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cnpq.driver.Dialect())
	t1 := builder.Table(cooccurrencenetworkpool.Table)
	columns := cnpq.ctx.Fields
	if len(columns) == 0 {
		columns = cooccurrencenetworkpool.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cnpq.sql != nil {
		selector = cnpq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cnpq.ctx.Unique != nil && *cnpq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range cnpq.predicates {
		p(selector)
	}
	for _, p := range cnpq.order {
		p(selector)
	}
	if offset := cnpq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cnpq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CooccurrenceNetworkPoolGroupBy is the group-by builder for CooccurrenceNetworkPool entities.
type CooccurrenceNetworkPoolGroupBy struct {
	selector
	build *CooccurrenceNetworkPoolQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cnpgb *CooccurrenceNetworkPoolGroupBy) Aggregate(fns ...AggregateFunc) *CooccurrenceNetworkPoolGroupBy {
	cnpgb.fns = append(cnpgb.fns, fns...)
	return cnpgb
}

// Scan applies the selector query and scans the result into the given value.
func (cnpgb *CooccurrenceNetworkPoolGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cnpgb.build.ctx, ent.OpQueryGroupBy)
	if err := cnpgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CooccurrenceNetworkPoolQuery, *CooccurrenceNetworkPoolGroupBy](ctx, cnpgb.build, cnpgb, cnpgb.build.inters, v)
}

func (cnpgb *CooccurrenceNetworkPoolGroupBy) sqlScan(ctx context.Context, root *CooccurrenceNetworkPoolQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cnpgb.fns))
	for _, fn := range cnpgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cnpgb.flds)+len(cnpgb.fns))
		for _, f := range *cnpgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cnpgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cnpgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CooccurrenceNetworkPoolSelect is the builder for selecting fields of CooccurrenceNetworkPool entities.
type CooccurrenceNetworkPoolSelect struct {
	*CooccurrenceNetworkPoolQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cnps *CooccurrenceNetworkPoolSelect) Aggregate(fns ...AggregateFunc) *CooccurrenceNetworkPoolSelect {
	cnps.fns = append(cnps.fns, fns...)
	return cnps
}

// Scan applies the selector query and scans the result into the given value.
func (cnps *CooccurrenceNetworkPoolSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cnps.ctx, ent.OpQuerySelect)
	if err := cnps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CooccurrenceNetworkPoolQuery, *CooccurrenceNetworkPoolSelect](ctx, cnps.CooccurrenceNetworkPoolQuery, cnps, cnps.inters, v)
}

func (cnps *CooccurrenceNetworkPoolSelect) sqlScan(ctx context.Context, root *CooccurrenceNetworkPoolQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cnps.fns))
	for _, fn := range cnps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cnps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cnps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

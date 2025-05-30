// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"tmail/ent/attachment"
	"tmail/ent/envelope"
	"tmail/ent/predicate"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EnvelopeQuery is the builder for querying Envelope entities.
type EnvelopeQuery struct {
	config
	ctx             *QueryContext
	order           []envelope.OrderOption
	inters          []Interceptor
	predicates      []predicate.Envelope
	withAttachments *AttachmentQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the EnvelopeQuery builder.
func (eq *EnvelopeQuery) Where(ps ...predicate.Envelope) *EnvelopeQuery {
	eq.predicates = append(eq.predicates, ps...)
	return eq
}

// Limit the number of records to be returned by this query.
func (eq *EnvelopeQuery) Limit(limit int) *EnvelopeQuery {
	eq.ctx.Limit = &limit
	return eq
}

// Offset to start from.
func (eq *EnvelopeQuery) Offset(offset int) *EnvelopeQuery {
	eq.ctx.Offset = &offset
	return eq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (eq *EnvelopeQuery) Unique(unique bool) *EnvelopeQuery {
	eq.ctx.Unique = &unique
	return eq
}

// Order specifies how the records should be ordered.
func (eq *EnvelopeQuery) Order(o ...envelope.OrderOption) *EnvelopeQuery {
	eq.order = append(eq.order, o...)
	return eq
}

// QueryAttachments chains the current query on the "attachments" edge.
func (eq *EnvelopeQuery) QueryAttachments() *AttachmentQuery {
	query := (&AttachmentClient{config: eq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(envelope.Table, envelope.FieldID, selector),
			sqlgraph.To(attachment.Table, attachment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, envelope.AttachmentsTable, envelope.AttachmentsColumn),
		)
		fromU = sqlgraph.SetNeighbors(eq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Envelope entity from the query.
// Returns a *NotFoundError when no Envelope was found.
func (eq *EnvelopeQuery) First(ctx context.Context) (*Envelope, error) {
	nodes, err := eq.Limit(1).All(setContextOp(ctx, eq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{envelope.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (eq *EnvelopeQuery) FirstX(ctx context.Context) *Envelope {
	node, err := eq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Envelope ID from the query.
// Returns a *NotFoundError when no Envelope ID was found.
func (eq *EnvelopeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = eq.Limit(1).IDs(setContextOp(ctx, eq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{envelope.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (eq *EnvelopeQuery) FirstIDX(ctx context.Context) int {
	id, err := eq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Envelope entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Envelope entity is found.
// Returns a *NotFoundError when no Envelope entities are found.
func (eq *EnvelopeQuery) Only(ctx context.Context) (*Envelope, error) {
	nodes, err := eq.Limit(2).All(setContextOp(ctx, eq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{envelope.Label}
	default:
		return nil, &NotSingularError{envelope.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (eq *EnvelopeQuery) OnlyX(ctx context.Context) *Envelope {
	node, err := eq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Envelope ID in the query.
// Returns a *NotSingularError when more than one Envelope ID is found.
// Returns a *NotFoundError when no entities are found.
func (eq *EnvelopeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = eq.Limit(2).IDs(setContextOp(ctx, eq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{envelope.Label}
	default:
		err = &NotSingularError{envelope.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (eq *EnvelopeQuery) OnlyIDX(ctx context.Context) int {
	id, err := eq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Envelopes.
func (eq *EnvelopeQuery) All(ctx context.Context) ([]*Envelope, error) {
	ctx = setContextOp(ctx, eq.ctx, ent.OpQueryAll)
	if err := eq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Envelope, *EnvelopeQuery]()
	return withInterceptors[[]*Envelope](ctx, eq, qr, eq.inters)
}

// AllX is like All, but panics if an error occurs.
func (eq *EnvelopeQuery) AllX(ctx context.Context) []*Envelope {
	nodes, err := eq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Envelope IDs.
func (eq *EnvelopeQuery) IDs(ctx context.Context) (ids []int, err error) {
	if eq.ctx.Unique == nil && eq.path != nil {
		eq.Unique(true)
	}
	ctx = setContextOp(ctx, eq.ctx, ent.OpQueryIDs)
	if err = eq.Select(envelope.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (eq *EnvelopeQuery) IDsX(ctx context.Context) []int {
	ids, err := eq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (eq *EnvelopeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, eq.ctx, ent.OpQueryCount)
	if err := eq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, eq, querierCount[*EnvelopeQuery](), eq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (eq *EnvelopeQuery) CountX(ctx context.Context) int {
	count, err := eq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (eq *EnvelopeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, eq.ctx, ent.OpQueryExist)
	switch _, err := eq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (eq *EnvelopeQuery) ExistX(ctx context.Context) bool {
	exist, err := eq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the EnvelopeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (eq *EnvelopeQuery) Clone() *EnvelopeQuery {
	if eq == nil {
		return nil
	}
	return &EnvelopeQuery{
		config:          eq.config,
		ctx:             eq.ctx.Clone(),
		order:           append([]envelope.OrderOption{}, eq.order...),
		inters:          append([]Interceptor{}, eq.inters...),
		predicates:      append([]predicate.Envelope{}, eq.predicates...),
		withAttachments: eq.withAttachments.Clone(),
		// clone intermediate query.
		sql:  eq.sql.Clone(),
		path: eq.path,
	}
}

// WithAttachments tells the query-builder to eager-load the nodes that are connected to
// the "attachments" edge. The optional arguments are used to configure the query builder of the edge.
func (eq *EnvelopeQuery) WithAttachments(opts ...func(*AttachmentQuery)) *EnvelopeQuery {
	query := (&AttachmentClient{config: eq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eq.withAttachments = query
	return eq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		To string `json:"to,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Envelope.Query().
//		GroupBy(envelope.FieldTo).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (eq *EnvelopeQuery) GroupBy(field string, fields ...string) *EnvelopeGroupBy {
	eq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &EnvelopeGroupBy{build: eq}
	grbuild.flds = &eq.ctx.Fields
	grbuild.label = envelope.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		To string `json:"to,omitempty"`
//	}
//
//	client.Envelope.Query().
//		Select(envelope.FieldTo).
//		Scan(ctx, &v)
func (eq *EnvelopeQuery) Select(fields ...string) *EnvelopeSelect {
	eq.ctx.Fields = append(eq.ctx.Fields, fields...)
	sbuild := &EnvelopeSelect{EnvelopeQuery: eq}
	sbuild.label = envelope.Label
	sbuild.flds, sbuild.scan = &eq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a EnvelopeSelect configured with the given aggregations.
func (eq *EnvelopeQuery) Aggregate(fns ...AggregateFunc) *EnvelopeSelect {
	return eq.Select().Aggregate(fns...)
}

func (eq *EnvelopeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range eq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, eq); err != nil {
				return err
			}
		}
	}
	for _, f := range eq.ctx.Fields {
		if !envelope.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if eq.path != nil {
		prev, err := eq.path(ctx)
		if err != nil {
			return err
		}
		eq.sql = prev
	}
	return nil
}

func (eq *EnvelopeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Envelope, error) {
	var (
		nodes       = []*Envelope{}
		_spec       = eq.querySpec()
		loadedTypes = [1]bool{
			eq.withAttachments != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Envelope).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Envelope{config: eq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, eq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := eq.withAttachments; query != nil {
		if err := eq.loadAttachments(ctx, query, nodes,
			func(n *Envelope) { n.Edges.Attachments = []*Attachment{} },
			func(n *Envelope, e *Attachment) { n.Edges.Attachments = append(n.Edges.Attachments, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (eq *EnvelopeQuery) loadAttachments(ctx context.Context, query *AttachmentQuery, nodes []*Envelope, init func(*Envelope), assign func(*Envelope, *Attachment)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Envelope)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(envelope.AttachmentsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.envelope_attachments
		if fk == nil {
			return fmt.Errorf(`foreign-key "envelope_attachments" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "envelope_attachments" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (eq *EnvelopeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := eq.querySpec()
	_spec.Node.Columns = eq.ctx.Fields
	if len(eq.ctx.Fields) > 0 {
		_spec.Unique = eq.ctx.Unique != nil && *eq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, eq.driver, _spec)
}

func (eq *EnvelopeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(envelope.Table, envelope.Columns, sqlgraph.NewFieldSpec(envelope.FieldID, field.TypeInt))
	_spec.From = eq.sql
	if unique := eq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if eq.path != nil {
		_spec.Unique = true
	}
	if fields := eq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, envelope.FieldID)
		for i := range fields {
			if fields[i] != envelope.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := eq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := eq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := eq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := eq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (eq *EnvelopeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(eq.driver.Dialect())
	t1 := builder.Table(envelope.Table)
	columns := eq.ctx.Fields
	if len(columns) == 0 {
		columns = envelope.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if eq.sql != nil {
		selector = eq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if eq.ctx.Unique != nil && *eq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range eq.predicates {
		p(selector)
	}
	for _, p := range eq.order {
		p(selector)
	}
	if offset := eq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := eq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// EnvelopeGroupBy is the group-by builder for Envelope entities.
type EnvelopeGroupBy struct {
	selector
	build *EnvelopeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (egb *EnvelopeGroupBy) Aggregate(fns ...AggregateFunc) *EnvelopeGroupBy {
	egb.fns = append(egb.fns, fns...)
	return egb
}

// Scan applies the selector query and scans the result into the given value.
func (egb *EnvelopeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, egb.build.ctx, ent.OpQueryGroupBy)
	if err := egb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EnvelopeQuery, *EnvelopeGroupBy](ctx, egb.build, egb, egb.build.inters, v)
}

func (egb *EnvelopeGroupBy) sqlScan(ctx context.Context, root *EnvelopeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(egb.fns))
	for _, fn := range egb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*egb.flds)+len(egb.fns))
		for _, f := range *egb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*egb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := egb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// EnvelopeSelect is the builder for selecting fields of Envelope entities.
type EnvelopeSelect struct {
	*EnvelopeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (es *EnvelopeSelect) Aggregate(fns ...AggregateFunc) *EnvelopeSelect {
	es.fns = append(es.fns, fns...)
	return es
}

// Scan applies the selector query and scans the result into the given value.
func (es *EnvelopeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, es.ctx, ent.OpQuerySelect)
	if err := es.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EnvelopeQuery, *EnvelopeSelect](ctx, es.EnvelopeQuery, es, es.inters, v)
}

func (es *EnvelopeSelect) sqlScan(ctx context.Context, root *EnvelopeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(es.fns))
	for _, fn := range es.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*es.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := es.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

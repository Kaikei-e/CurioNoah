// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"insightstream/ent/migrate"

	entfeeds "insightstream/ent/feeds"
	"insightstream/ent/followlist"
	"insightstream/ent/users"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Feeds is the client for interacting with the Feeds builders.
	Feeds *FeedsClient
	// FollowList is the client for interacting with the FollowList builders.
	FollowList *FollowListClient
	// Users is the client for interacting with the Users builders.
	Users *UsersClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Feeds = NewFeedsClient(c.config)
	c.FollowList = NewFollowListClient(c.config)
	c.Users = NewUsersClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Feeds:      NewFeedsClient(cfg),
		FollowList: NewFollowListClient(cfg),
		Users:      NewUsersClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Feeds:      NewFeedsClient(cfg),
		FollowList: NewFollowListClient(cfg),
		Users:      NewUsersClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Feeds.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Feeds.Use(hooks...)
	c.FollowList.Use(hooks...)
	c.Users.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Feeds.Intercept(interceptors...)
	c.FollowList.Intercept(interceptors...)
	c.Users.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *FeedsMutation:
		return c.Feeds.mutate(ctx, m)
	case *FollowListMutation:
		return c.FollowList.mutate(ctx, m)
	case *UsersMutation:
		return c.Users.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// FeedsClient is a client for the Feeds schema.
type FeedsClient struct {
	config
}

// NewFeedsClient returns a client for the Feeds from the given config.
func NewFeedsClient(c config) *FeedsClient {
	return &FeedsClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `entfeeds.Hooks(f(g(h())))`.
func (c *FeedsClient) Use(hooks ...Hook) {
	c.hooks.Feeds = append(c.hooks.Feeds, hooks...)
}

// Use adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `entfeeds.Intercept(f(g(h())))`.
func (c *FeedsClient) Intercept(interceptors ...Interceptor) {
	c.inters.Feeds = append(c.inters.Feeds, interceptors...)
}

// Create returns a builder for creating a Feeds entity.
func (c *FeedsClient) Create() *FeedsCreate {
	mutation := newFeedsMutation(c.config, OpCreate)
	return &FeedsCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Feeds entities.
func (c *FeedsClient) CreateBulk(builders ...*FeedsCreate) *FeedsCreateBulk {
	return &FeedsCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Feeds.
func (c *FeedsClient) Update() *FeedsUpdate {
	mutation := newFeedsMutation(c.config, OpUpdate)
	return &FeedsUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *FeedsClient) UpdateOne(f *Feeds) *FeedsUpdateOne {
	mutation := newFeedsMutation(c.config, OpUpdateOne, withFeeds(f))
	return &FeedsUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *FeedsClient) UpdateOneID(id uuid.UUID) *FeedsUpdateOne {
	mutation := newFeedsMutation(c.config, OpUpdateOne, withFeedsID(id))
	return &FeedsUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Feeds.
func (c *FeedsClient) Delete() *FeedsDelete {
	mutation := newFeedsMutation(c.config, OpDelete)
	return &FeedsDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *FeedsClient) DeleteOne(f *Feeds) *FeedsDeleteOne {
	return c.DeleteOneID(f.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *FeedsClient) DeleteOneID(id uuid.UUID) *FeedsDeleteOne {
	builder := c.Delete().Where(entfeeds.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &FeedsDeleteOne{builder}
}

// Query returns a query builder for Feeds.
func (c *FeedsClient) Query() *FeedsQuery {
	return &FeedsQuery{
		config: c.config,
		inters: c.Interceptors(),
	}
}

// Get returns a Feeds entity by its id.
func (c *FeedsClient) Get(ctx context.Context, id uuid.UUID) (*Feeds, error) {
	return c.Query().Where(entfeeds.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *FeedsClient) GetX(ctx context.Context, id uuid.UUID) *Feeds {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *FeedsClient) Hooks() []Hook {
	return c.hooks.Feeds
}

// Interceptors returns the client interceptors.
func (c *FeedsClient) Interceptors() []Interceptor {
	return c.inters.Feeds
}

func (c *FeedsClient) mutate(ctx context.Context, m *FeedsMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&FeedsCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&FeedsUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&FeedsUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&FeedsDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Feeds mutation op: %q", m.Op())
	}
}

// FollowListClient is a client for the FollowList schema.
type FollowListClient struct {
	config
}

// NewFollowListClient returns a client for the FollowList from the given config.
func NewFollowListClient(c config) *FollowListClient {
	return &FollowListClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `followlist.Hooks(f(g(h())))`.
func (c *FollowListClient) Use(hooks ...Hook) {
	c.hooks.FollowList = append(c.hooks.FollowList, hooks...)
}

// Use adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `followlist.Intercept(f(g(h())))`.
func (c *FollowListClient) Intercept(interceptors ...Interceptor) {
	c.inters.FollowList = append(c.inters.FollowList, interceptors...)
}

// Create returns a builder for creating a FollowList entity.
func (c *FollowListClient) Create() *FollowListCreate {
	mutation := newFollowListMutation(c.config, OpCreate)
	return &FollowListCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of FollowList entities.
func (c *FollowListClient) CreateBulk(builders ...*FollowListCreate) *FollowListCreateBulk {
	return &FollowListCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for FollowList.
func (c *FollowListClient) Update() *FollowListUpdate {
	mutation := newFollowListMutation(c.config, OpUpdate)
	return &FollowListUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *FollowListClient) UpdateOne(fl *FollowList) *FollowListUpdateOne {
	mutation := newFollowListMutation(c.config, OpUpdateOne, withFollowList(fl))
	return &FollowListUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *FollowListClient) UpdateOneID(id int) *FollowListUpdateOne {
	mutation := newFollowListMutation(c.config, OpUpdateOne, withFollowListID(id))
	return &FollowListUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for FollowList.
func (c *FollowListClient) Delete() *FollowListDelete {
	mutation := newFollowListMutation(c.config, OpDelete)
	return &FollowListDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *FollowListClient) DeleteOne(fl *FollowList) *FollowListDeleteOne {
	return c.DeleteOneID(fl.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *FollowListClient) DeleteOneID(id int) *FollowListDeleteOne {
	builder := c.Delete().Where(followlist.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &FollowListDeleteOne{builder}
}

// Query returns a query builder for FollowList.
func (c *FollowListClient) Query() *FollowListQuery {
	return &FollowListQuery{
		config: c.config,
		inters: c.Interceptors(),
	}
}

// Get returns a FollowList entity by its id.
func (c *FollowListClient) Get(ctx context.Context, id int) (*FollowList, error) {
	return c.Query().Where(followlist.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *FollowListClient) GetX(ctx context.Context, id int) *FollowList {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *FollowListClient) Hooks() []Hook {
	return c.hooks.FollowList
}

// Interceptors returns the client interceptors.
func (c *FollowListClient) Interceptors() []Interceptor {
	return c.inters.FollowList
}

func (c *FollowListClient) mutate(ctx context.Context, m *FollowListMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&FollowListCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&FollowListUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&FollowListUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&FollowListDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown FollowList mutation op: %q", m.Op())
	}
}

// UsersClient is a client for the Users schema.
type UsersClient struct {
	config
}

// NewUsersClient returns a client for the Users from the given config.
func NewUsersClient(c config) *UsersClient {
	return &UsersClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `users.Hooks(f(g(h())))`.
func (c *UsersClient) Use(hooks ...Hook) {
	c.hooks.Users = append(c.hooks.Users, hooks...)
}

// Use adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `users.Intercept(f(g(h())))`.
func (c *UsersClient) Intercept(interceptors ...Interceptor) {
	c.inters.Users = append(c.inters.Users, interceptors...)
}

// Create returns a builder for creating a Users entity.
func (c *UsersClient) Create() *UsersCreate {
	mutation := newUsersMutation(c.config, OpCreate)
	return &UsersCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Users entities.
func (c *UsersClient) CreateBulk(builders ...*UsersCreate) *UsersCreateBulk {
	return &UsersCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Users.
func (c *UsersClient) Update() *UsersUpdate {
	mutation := newUsersMutation(c.config, OpUpdate)
	return &UsersUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UsersClient) UpdateOne(u *Users) *UsersUpdateOne {
	mutation := newUsersMutation(c.config, OpUpdateOne, withUsers(u))
	return &UsersUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UsersClient) UpdateOneID(id int) *UsersUpdateOne {
	mutation := newUsersMutation(c.config, OpUpdateOne, withUsersID(id))
	return &UsersUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Users.
func (c *UsersClient) Delete() *UsersDelete {
	mutation := newUsersMutation(c.config, OpDelete)
	return &UsersDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UsersClient) DeleteOne(u *Users) *UsersDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UsersClient) DeleteOneID(id int) *UsersDeleteOne {
	builder := c.Delete().Where(users.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UsersDeleteOne{builder}
}

// Query returns a query builder for Users.
func (c *UsersClient) Query() *UsersQuery {
	return &UsersQuery{
		config: c.config,
		inters: c.Interceptors(),
	}
}

// Get returns a Users entity by its id.
func (c *UsersClient) Get(ctx context.Context, id int) (*Users, error) {
	return c.Query().Where(users.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UsersClient) GetX(ctx context.Context, id int) *Users {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *UsersClient) Hooks() []Hook {
	return c.hooks.Users
}

// Interceptors returns the client interceptors.
func (c *UsersClient) Interceptors() []Interceptor {
	return c.inters.Users
}

func (c *UsersClient) mutate(ctx context.Context, m *UsersMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UsersCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UsersUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UsersUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UsersDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Users mutation op: %q", m.Op())
	}
}

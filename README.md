q go-todo
## 概要
> 学習用途

golangを用いて、DDDでREST APIを実装。
```plantuml
@startuml
namespace config {
    class Api << (S,Aquamarine) >> {
        + Port string

    }
    class Config << (S,Aquamarine) >> {
    }
    class Database << (S,Aquamarine) >> {
        + Driver string
        + Host string
        + Port string
        + Name string
        + User string
        + Password string

    }
}
"config.Api" *-- "config.Config"
"config.Database" *-- "config.Config"


namespace ent {
    class Client << (S,Aquamarine) >> {
        + Schema *migrate.Schema
        + Todo *TodoClient

        - init()

        + Tx(ctx context.Context) (*Tx, error)
        + BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error)
        + Debug() *Client
        + Close() error
        + Use(hooks ...Hook)
        + Intercept(interceptors ...Interceptor)
        + Mutate(ctx context.Context, m Mutation) (Value, error)

    }
    class CommitFunc << (S,Aquamarine) >> {
        + Commit(ctx context.Context, tx *Tx) error

    }
    interface Committer  {
        + Commit( context.Context,  *Tx) error

    }
    class ConstraintError << (S,Aquamarine) >> {
        - msg string
        - wrap error

        + Error() string
        + Unwrap() error

    }
    class NotFoundError << (S,Aquamarine) >> {
        - label string

        + Error() string

    }
    class NotLoadedError << (S,Aquamarine) >> {
        - edge string

        + Error() string

    }
    class NotSingularError << (S,Aquamarine) >> {
        - label string

        + Error() string

    }
    class RollbackFunc << (S,Aquamarine) >> {
        + Rollback(ctx context.Context, tx *Tx) error

    }
    interface Rollbacker  {
        + Rollback( context.Context,  *Tx) error

    }
    class Todo << (S,Aquamarine) >> {
        - selectValues sql.SelectValues

        + ID int
        + Title string
        + Completed bool
        + CreatedAt time.Time
        + Priority int

        - scanValues(columns []string) ([]any, error)
        - assignValues(columns []string, values []any) error

        + Value(name string) (ent.Value, error)
        + Update() *TodoUpdateOne
        + Unwrap() *Todo
        + String() string

    }
    class TodoClient << (S,Aquamarine) >> {
        - mutate(ctx context.Context, m *TodoMutation) (Value, error)

        + Use(hooks ...Hook)
        + Intercept(interceptors ...Interceptor)
        + Create() *TodoCreate
        + CreateBulk(builders ...*TodoCreate) *TodoCreateBulk
        + Update() *TodoUpdate
        + UpdateOne(t *Todo) *TodoUpdateOne
        + UpdateOneID(id int) *TodoUpdateOne
        + Delete() *TodoDelete
        + DeleteOne(t *Todo) *TodoDeleteOne
        + DeleteOneID(id int) *TodoDeleteOne
        + Query() *TodoQuery
        + Get(ctx context.Context, id int) (*Todo, error)
        + GetX(ctx context.Context, id int) *Todo
        + Hooks() []Hook
        + Interceptors() []Interceptor

    }
    class TodoCreate << (S,Aquamarine) >> {
        - mutation *TodoMutation
        - hooks []Hook

        - defaults()
        - check() error
        - sqlSave(ctx context.Context) (*Todo, error)
        - createSpec() (*Todo, *sqlgraph.CreateSpec)

        + SetTitle(s string) *TodoCreate
        + SetNillableTitle(s *string) *TodoCreate
        + SetCompleted(b bool) *TodoCreate
        + SetNillableCompleted(b *bool) *TodoCreate
        + SetCreatedAt(t time.Time) *TodoCreate
        + SetNillableCreatedAt(t *time.Time) *TodoCreate
        + SetPriority(i int) *TodoCreate
        + Mutation() *TodoMutation
        + Save(ctx context.Context) (*Todo, error)
        + SaveX(ctx context.Context) *Todo
        + Exec(ctx context.Context) error
        + ExecX(ctx context.Context)

    }
    class TodoCreateBulk << (S,Aquamarine) >> {
        - builders []*TodoCreate

        + Save(ctx context.Context) ([]*Todo, error)
        + SaveX(ctx context.Context) []*Todo
        + Exec(ctx context.Context) error
        + ExecX(ctx context.Context)

    }
    class TodoDelete << (S,Aquamarine) >> {
        - hooks []Hook
        - mutation *TodoMutation

        - sqlExec(ctx context.Context) (int, error)

        + Where(ps ...predicate.Todo) *TodoDelete
        + Exec(ctx context.Context) (int, error)
        + ExecX(ctx context.Context) int

    }
    class TodoDeleteOne << (S,Aquamarine) >> {
        - td *TodoDelete

        + Where(ps ...predicate.Todo) *TodoDeleteOne
        + Exec(ctx context.Context) error
        + ExecX(ctx context.Context)

    }
    class TodoGroupBy << (S,Aquamarine) >> {
        - build *TodoQuery

        - sqlScan(ctx context.Context, root *TodoQuery, v any) error

        + Aggregate(fns ...AggregateFunc) *TodoGroupBy
        + Scan(ctx context.Context, v any) error

    }
    class TodoMutation << (S,Aquamarine) >> {
        - op Op
        - typ string
        - id *int
        - title *string
        - completed *bool
        - created_at *time.Time
        - priority *int
        - addpriority *int
        - clearedFields <font color=blue>map</font>[string]<font color=blue>struct</font>{}
        - done bool
        - oldValue <font color=blue>func</font>(context.Context) (*Todo, error)
        - predicates []predicate.Todo

        + Client() *Client
        + Tx() (*Tx, error)
        + ID() (int, bool)
        + IDs(ctx context.Context) ([]int, error)
        + SetTitle(s string)
        + Title() (string, bool)
        + OldTitle(ctx context.Context) (string, error)
        + ResetTitle()
        + SetCompleted(b bool)
        + Completed() (bool, bool)
        + OldCompleted(ctx context.Context) (bool, error)
        + ResetCompleted()
        + SetCreatedAt(t time.Time)
        + CreatedAt() (time.Time, bool)
        + OldCreatedAt(ctx context.Context) (time.Time, error)
        + ResetCreatedAt()
        + SetPriority(i int)
        + Priority() (int, bool)
        + OldPriority(ctx context.Context) (int, error)
        + AddPriority(i int)
        + AddedPriority() (int, bool)
        + ResetPriority()
        + Where(ps ...predicate.Todo)
        + WhereP(ps ...<font color=blue>func</font>(*sql.Selector) )
        + Op() Op
        + SetOp(op Op)
        + Type() string
        + Fields() []string
        + Field(name string) (ent.Value, bool)
        + OldField(ctx context.Context, name string) (ent.Value, error)
        + SetField(name string, value ent.Value) error
        + AddedFields() []string
        + AddedField(name string) (ent.Value, bool)
        + AddField(name string, value ent.Value) error
        + ClearedFields() []string
        + FieldCleared(name string) bool
        + ClearField(name string) error
        + ResetField(name string) error
        + AddedEdges() []string
        + AddedIDs(name string) []ent.Value
        + RemovedEdges() []string
        + RemovedIDs(name string) []ent.Value
        + ClearedEdges() []string
        + EdgeCleared(name string) bool
        + ClearEdge(name string) error
        + ResetEdge(name string) error

    }
    class TodoQuery << (S,Aquamarine) >> {
        - ctx *QueryContext
        - order []todo.OrderOption
        - inters []Interceptor
        - predicates []predicate.Todo
        - sql *sql.Selector
        - path <font color=blue>func</font>(context.Context) (*sql.Selector, error)

        - prepareQuery(ctx context.Context) error
        - sqlAll(ctx context.Context, hooks ...queryHook) ([]*Todo, error)
        - sqlCount(ctx context.Context) (int, error)
        - querySpec() *sqlgraph.QuerySpec
        - sqlQuery(ctx context.Context) *sql.Selector

        + Where(ps ...predicate.Todo) *TodoQuery
        + Limit(limit int) *TodoQuery
        + Offset(offset int) *TodoQuery
        + Unique(unique bool) *TodoQuery
        + Order(o ...todo.OrderOption) *TodoQuery
        + First(ctx context.Context) (*Todo, error)
        + FirstX(ctx context.Context) *Todo
        + FirstID(ctx context.Context) (int, error)
        + FirstIDX(ctx context.Context) int
        + Only(ctx context.Context) (*Todo, error)
        + OnlyX(ctx context.Context) *Todo
        + OnlyID(ctx context.Context) (int, error)
        + OnlyIDX(ctx context.Context) int
        + All(ctx context.Context) ([]*Todo, error)
        + AllX(ctx context.Context) []*Todo
        + IDs(ctx context.Context) ([]int, error)
        + IDsX(ctx context.Context) []int
        + Count(ctx context.Context) (int, error)
        + CountX(ctx context.Context) int
        + Exist(ctx context.Context) (bool, error)
        + ExistX(ctx context.Context) bool
        + Clone() *TodoQuery
        + GroupBy(field string, fields ...string) *TodoGroupBy
        + Select(fields ...string) *TodoSelect
        + Aggregate(fns ...AggregateFunc) *TodoSelect

    }
    class TodoSelect << (S,Aquamarine) >> {
        - sqlScan(ctx context.Context, root *TodoQuery, v any) error

        + Aggregate(fns ...AggregateFunc) *TodoSelect
        + Scan(ctx context.Context, v any) error

    }
    class TodoUpdate << (S,Aquamarine) >> {
        - hooks []Hook
        - mutation *TodoMutation

        - check() error
        - sqlSave(ctx context.Context) (int, error)

        + Where(ps ...predicate.Todo) *TodoUpdate
        + SetTitle(s string) *TodoUpdate
        + SetNillableTitle(s *string) *TodoUpdate
        + SetCompleted(b bool) *TodoUpdate
        + SetNillableCompleted(b *bool) *TodoUpdate
        + SetCreatedAt(t time.Time) *TodoUpdate
        + SetNillableCreatedAt(t *time.Time) *TodoUpdate
        + SetPriority(i int) *TodoUpdate
        + AddPriority(i int) *TodoUpdate
        + Mutation() *TodoMutation
        + Save(ctx context.Context) (int, error)
        + SaveX(ctx context.Context) int
        + Exec(ctx context.Context) error
        + ExecX(ctx context.Context)

    }
    class TodoUpdateOne << (S,Aquamarine) >> {
        - fields []string
        - hooks []Hook
        - mutation *TodoMutation

        - check() error
        - sqlSave(ctx context.Context) (*Todo, error)

        + SetTitle(s string) *TodoUpdateOne
        + SetNillableTitle(s *string) *TodoUpdateOne
        + SetCompleted(b bool) *TodoUpdateOne
        + SetNillableCompleted(b *bool) *TodoUpdateOne
        + SetCreatedAt(t time.Time) *TodoUpdateOne
        + SetNillableCreatedAt(t *time.Time) *TodoUpdateOne
        + SetPriority(i int) *TodoUpdateOne
        + AddPriority(i int) *TodoUpdateOne
        + Mutation() *TodoMutation
        + Where(ps ...predicate.Todo) *TodoUpdateOne
        + Select(field string, fields ...string) *TodoUpdateOne
        + Save(ctx context.Context) (*Todo, error)
        + SaveX(ctx context.Context) *Todo
        + Exec(ctx context.Context) error
        + ExecX(ctx context.Context)

    }
    class Tx << (S,Aquamarine) >> {
        - client *Client
        - clientOnce sync.Once
        - ctx context.Context

        + Todo *TodoClient

        - init()

        + Commit() error
        + OnCommit(f CommitHook)
        + Rollback() error
        + OnRollback(f RollbackHook)
        + Client() *Client

    }
    class ValidationError << (S,Aquamarine) >> {
        - err error

        + Name string

        + Error() string
        + Unwrap() error

    }
    class clientCtxKey << (S,Aquamarine) >> {
    }
    class config << (S,Aquamarine) >> {
        - driver dialect.Driver
        - debug bool
        - log <font color=blue>func</font>(...any)
        - hooks *hooks
        - inters *inters

        - options(opts ...Option)

    }
    class ent.AggregateFunc << (T, #FF7700) >>  {
    }
    class ent.CommitFunc << (T, #FF7700) >>  {
    }
    class ent.CommitHook << (T, #FF7700) >>  {
    }
    class ent.Hook << (T, #FF7700) >>  {
    }
    class ent.InterceptFunc << (T, #FF7700) >>  {
    }
    class ent.Interceptor << (T, #FF7700) >>  {
    }
    class ent.MutateFunc << (T, #FF7700) >>  {
    }
    class ent.Mutation << (T, #FF7700) >>  {
    }
    class ent.Mutator << (T, #FF7700) >>  {
    }
    class ent.Op << (T, #FF7700) >>  {
    }
    class ent.Option << (T, #FF7700) >>  {
    }
    class ent.OrderFunc << (T, #FF7700) >>  {
    }
    class ent.Policy << (T, #FF7700) >>  {
    }
    class ent.Querier << (T, #FF7700) >>  {
    }
    class ent.QuerierFunc << (T, #FF7700) >>  {
    }
    class ent.Query << (T, #FF7700) >>  {
    }
    class ent.QueryContext << (T, #FF7700) >>  {
    }
    class ent.RollbackFunc << (T, #FF7700) >>  {
    }
    class ent.RollbackHook << (T, #FF7700) >>  {
    }
    class ent.Todos << (T, #FF7700) >>  {
    }
    class ent.TraverseFunc << (T, #FF7700) >>  {
    }
    class ent.Traverser << (T, #FF7700) >>  {
    }
    class ent.Value << (T, #FF7700) >>  {
    }
    class ent.queryHook << (T, #FF7700) >>  {
    }
    class ent.todoOption << (T, #FF7700) >>  {
    }
    class hooks << (S,Aquamarine) >> {
        + Todo []ent.Hook

    }
    class inters << (S,Aquamarine) >> {
        + Todo []ent.Interceptor

    }
    class selector << (S,Aquamarine) >> {
        - label string
        - flds *[]string
        - fns []AggregateFunc
        - scan <font color=blue>func</font>(context.Context, any) error

        + ScanX(ctx context.Context, v any)
        + Strings(ctx context.Context) ([]string, error)
        + StringsX(ctx context.Context) []string
        + String(ctx context.Context) (string, error)
        + StringX(ctx context.Context) string
        + Ints(ctx context.Context) ([]int, error)
        + IntsX(ctx context.Context) []int
        + Int(ctx context.Context) (int, error)
        + IntX(ctx context.Context) int
        + Float64s(ctx context.Context) ([]float64, error)
        + Float64sX(ctx context.Context) []float64
        + Float64(ctx context.Context) (float64, error)
        + Float64X(ctx context.Context) float64
        + Bools(ctx context.Context) ([]bool, error)
        + BoolsX(ctx context.Context) []bool
        + Bool(ctx context.Context) (bool, error)
        + BoolX(ctx context.Context) bool

    }
    class txCtxKey << (S,Aquamarine) >> {
    }
    class txDriver << (S,Aquamarine) >> {
        - drv dialect.Driver
        - tx dialect.Tx
        - mu sync.Mutex
        - onCommit []CommitHook
        - onRollback []RollbackHook

        + Tx( context.Context) (dialect.Tx, error)
        + Dialect() string
        + Close() error
        + Commit() error
        + Rollback() error
        + Exec(ctx context.Context, query string, args any, v any) error
        + Query(ctx context.Context, query string, args any, v any) error

    }
    class "ent.Hook" as entHook {
        'This class was created so that we can correctly have an alias pointing to this name. Since it contains dots that can break namespaces
    }
    class "ent.InterceptFunc" as entInterceptFunc {
        'This class was created so that we can correctly have an alias pointing to this name. Since it contains dots that can break namespaces
    }
    class "ent.Interceptor" as entInterceptor {
        'This class was created so that we can correctly have an alias pointing to this name. Since it contains dots that can break namespaces
    }
    class "ent.MutateFunc" as entMutateFunc {
        'This class was created so that we can correctly have an alias pointing to this name. Since it contains dots that can break namespaces
    }
    class "ent.Mutation" as entMutation {
        'This class was created so that we can correctly have an alias pointing to this name. Since it contains dots that can break namespaces
    }
    class "ent.Mutator" as entMutator {
        'This class was created so that we can correctly have an alias pointing to this name. Since it contains dots that can break namespaces
    }
    class "ent.Op" as entOp {
        'This class was created so that we can correctly have an alias pointing to this name. Since it contains dots that can break namespaces
    }
    class "ent.Policy" as entPolicy {
        'This class was created so that we can correctly have an alias pointing to this name. Since it contains dots that can break namespaces
    }
    class "ent.Querier" as entQuerier {
        'This class was created so that we can correctly have an alias pointing to this name. Since it contains dots that can break namespaces
    }
    class "ent.QuerierFunc" as entQuerierFunc {
        'This class was created so that we can correctly have an alias pointing to this name. Since it contains dots that can break namespaces
    }
    class "ent.Query" as entQuery {
        'This class was created so that we can correctly have an alias pointing to this name. Since it contains dots that can break namespaces
    }
    class "ent.QueryContext" as entQueryContext {
        'This class was created so that we can correctly have an alias pointing to this name. Since it contains dots that can break namespaces
    }
    class "ent.TraverseFunc" as entTraverseFunc {
        'This class was created so that we can correctly have an alias pointing to this name. Since it contains dots that can break namespaces
    }
    class "ent.Traverser" as entTraverser {
        'This class was created so that we can correctly have an alias pointing to this name. Since it contains dots that can break namespaces
    }
    class "ent.Value" as entValue {
        'This class was created so that we can correctly have an alias pointing to this name. Since it contains dots that can break namespaces
    }
    class "<font color=blue>func</font>(context.Context, *Tx) error" as fontcolorbluefuncfontcontextContextTxerror {
        'This class was created so that we can correctly have an alias pointing to this name. Since it contains dots that can break namespaces
    }
    class "<font color=blue>func</font>(context.Context, *sqlgraph.QuerySpec) " as fontcolorbluefuncfontcontextContextsqlgraphQuerySpec {
        'This class was created so that we can correctly have an alias pointing to this name. Since it contains dots that can break namespaces
    }
    class "<font color=blue>func</font>(*sql.Selector) " as fontcolorbluefuncfontsqlSelector {
        'This class was created so that we can correctly have an alias pointing to this name. Since it contains dots that can break namespaces
    }
    class "<font color=blue>func</font>(*sql.Selector) string" as fontcolorbluefuncfontsqlSelectorstring {
        'This class was created so that we can correctly have an alias pointing to this name. Since it contains dots that can break namespaces
    }
}
"ent.config" *-- "ent.Client"
"ent.config" *-- "ent.Todo"
"ent.config" *-- "ent.TodoClient"
"ent.config" *-- "ent.TodoCreate"
"ent.config" *-- "ent.TodoCreateBulk"
"ent.config" *-- "ent.TodoDelete"
"ent.selector" *-- "ent.TodoGroupBy"
"ent.config" *-- "ent.TodoMutation"
"ent.config" *-- "ent.TodoQuery"
"ent.TodoQuery" *-- "ent.TodoSelect"
"ent.selector" *-- "ent.TodoSelect"
"ent.config" *-- "ent.TodoUpdate"
"ent.config" *-- "ent.TodoUpdateOne"
"ent.config" *-- "ent.Tx"

"ent.Committer" <|-- "ent.CommitFunc"
"ent.Rollbacker" <|-- "ent.RollbackFunc"

namespace enttest {
    interface TestingT  {
        + FailNow()
        + Error( ...any)

    }
    class enttest.Option << (T, #FF7700) >>  {
    }
    class options << (S,Aquamarine) >> {
        - opts []ent.Option
        - migrateOpts []schema.MigrateOption

    }
}


namespace handler {
    class CreateRequest << (S,Aquamarine) >> {
        + Title string
        + Completed bool
        + Priority int
        + CreatedAt time.Time

    }
    class Handler << (S,Aquamarine) >> {
        - usecase *usecase.TodoUsecase

        + List(w http.ResponseWriter, r *http.Request)
        + Create(w http.ResponseWriter, r *http.Request)

    }
}


namespace hook {
    class Chain << (S,Aquamarine) >> {
        - hooks []ent.Hook

        + Hook() ent.Hook
        + Append(hooks ...ent.Hook) Chain
        + Extend(chain Chain) Chain

    }
    class TodoFunc << (S,Aquamarine) >> {
        + Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error)

    }
    class hook.Condition << (T, #FF7700) >>  {
    }
    class hook.TodoFunc << (T, #FF7700) >>  {
    }
    class "<font color=blue>func</font>(context.Context, ent.Mutation) bool" as fontcolorbluefuncfontcontextContextentMutationbool {
        'This class was created so that we can correctly have an alias pointing to this name. Since it contains dots that can break namespaces
    }
    class "<font color=blue>func</font>(context.Context, *ent.TodoMutation) (ent.Value, error)" as fontcolorbluefuncfontcontextContextentTodoMutationentValueerror {
        'This class was created so that we can correctly have an alias pointing to this name. Since it contains dots that can break namespaces
    }
}


namespace migrate {
    class Schema << (S,Aquamarine) >> {
        - drv dialect.Driver

        + Create(ctx context.Context, opts ...schema.MigrateOption) error
        + WriteTo(ctx context.Context, w io.Writer, opts ...schema.MigrateOption) error

    }
}


namespace predicate {
    class predicate.Todo << (T, #FF7700) >>  {
    }
    class "<font color=blue>func</font>(*sql.Selector) " as fontcolorbluefuncfontsqlSelector {
        'This class was created so that we can correctly have an alias pointing to this name. Since it contains dots that can break namespaces
    }
}


namespace repository {
    class TodoRepository << (S,Aquamarine) >> {
        - ent *ent.Client

        + List(ctx context.Context) ([]*todo.Schema, error)
        + Create(ctx context.Context, t *todo.Schema) (*todo.Schema, error)

    }
}


namespace schema {
    class Todo << (S,Aquamarine) >> {
        + Fields() []ent.Field
        + Edges() []ent.Edge

    }
}
"ent.Schema" *-- "schema.Todo"


namespace server {
    class Server << (S,Aquamarine) >> {
        - cfg *config.Config
        - ent *ent.Client
        - router *chi.Mux

        - initTodo()
        - newRouter()
        - newEnt()

        + InitDomains()
        + Init()
        + Run()

    }
    class server.Options << (T, #FF7700) >>  {
    }
}


namespace todo {
    class Schema << (S,Aquamarine) >> {
        + Title string
        + Completed bool
        + Priority int
        + CretedAt time.Time

    }
    class todo.OrderOption << (T, #FF7700) >>  {
    }
    class "<font color=blue>func</font>(*sql.Selector) " as fontcolorbluefuncfontsqlSelector {
        'This class was created so that we can correctly have an alias pointing to this name. Since it contains dots that can break namespaces
    }
}


namespace usecase {
    class TodoUsecase << (S,Aquamarine) >> {
        - repo *repository.TodoRepository

        + Create(ctx context.Context, r *todo.Schema) (*todo.Schema, error)
        + List(ctx context.Context) ([]*todo.Schema, error)

    }
}


"ent.<font color=blue>func</font>(*TodoMutation) " #.. "ent.todoOption"
"ent.<font color=blue>func</font>(*config) " #.. "ent.Option"
"ent.fontcolorbluefuncfontsqlSelector" #.. "ent.OrderFunc"
"ent.fontcolorbluefuncfontsqlSelectorstring" #.. "ent.AggregateFunc"
"ent.<font color=blue>func</font>(Committer) Committer" #.. "ent.CommitHook"
"ent.<font color=blue>func</font>(Rollbacker) Rollbacker" #.. "ent.RollbackHook"
"ent.fontcolorbluefuncfontcontextContextTxerror" #.. "ent.CommitFunc"
"ent.fontcolorbluefuncfontcontextContextTxerror" #.. "ent.RollbackFunc"
"ent.fontcolorbluefuncfontcontextContextsqlgraphQuerySpec" #.. "ent.queryHook"
"ent.[]*Todo" #.. "ent.Todos"
"ent.entHook" #.. "ent.Hook"
"ent.entInterceptFunc" #.. "ent.InterceptFunc"
"ent.entInterceptor" #.. "ent.Interceptor"
"ent.entMutateFunc" #.. "ent.MutateFunc"
"ent.entMutation" #.. "ent.Mutation"
"ent.entMutator" #.. "ent.Mutator"
"ent.entOp" #.. "ent.Op"
"ent.entPolicy" #.. "ent.Policy"
"ent.entQuerier" #.. "ent.Querier"
"ent.entQuerierFunc" #.. "ent.QuerierFunc"
"ent.entQuery" #.. "ent.Query"
"ent.entQueryContext" #.. "ent.QueryContext"
"ent.entTraverseFunc" #.. "ent.TraverseFunc"
"ent.entTraverser" #.. "ent.Traverser"
"ent.entValue" #.. "ent.Value"
"enttest.<font color=blue>func</font>(*options) " #.. "enttest.Option"
"hook.fontcolorbluefuncfontcontextContextentTodoMutationentValueerror" #.. "hook.TodoFunc"
"hook.fontcolorbluefuncfontcontextContextentMutationbool" #.. "hook.Condition"
"predicate.fontcolorbluefuncfontsqlSelector" #.. "predicate.Todo"
"server.<font color=blue>func</font>(*Server) error" #.. "server.Options"
"todo.fontcolorbluefuncfontsqlSelector" #.. "todo.OrderOption"
@enduml
```

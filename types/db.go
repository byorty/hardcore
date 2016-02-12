package types

type Query interface {
	ToNative() interface{}
	GetArgs() []interface{}
}

type QueryExecuter interface {
	Exec(Query) DBResult
	Query(Query) DBRows
	QueryRow(Query) DBRow
	Custom(Query) DBCustomRow
}

type DBScanner interface {
	Scan(...interface{}) error
}

type DBRow interface {
	DBScanner
	One(Model)
}

type DBCustomRow interface {
	DBScanner
	One(...interface{})
}

type DBRows interface {
	DBScanner
	All(Model)
}

type DBResult interface {
	One(Model)
}

type DBKind int

const (
	SqlDB   DBKind = iota
	NoSqlDB
)

type DB interface {
	QueryExecuter
	GetKind() DBKind
	Close()
	GetQueryWriter() QueryWriter
	SupportLastInsertId() bool
	SupportReturningId() bool
	Prepare(Query) DBStatement
}

type DBStatement interface {
	Exec(ModelDAO, ...interface{})
	Query(ModelDAO, ...interface{})
	QueryRow(ModelDAO, ...interface{})
}

type DBPool interface {
	Add(string, DB) DBPool
	ByDAO(ModelDAO) DB
	ByName(string) DB
}

type QueryWriter interface {
	SetProto(Proto)
	SetTable(string)
	SetLogicChain([]LogicChain)
	SetProjections([]Projection)
	GetArgs() []interface{}
	SetArgs([]interface{})
	AddArg(interface{})
	WriteSelect() interface{}
	WriteInsert() interface{}
	WriteUpdate() interface{}
}

type SqlQueryWriter interface {
	QueryWriter
	AddField(string)
	SetFields([]string)
	GetFields() []string
	WriteField(string, string) string
	WriteTable(string) string
	WriteArg(int, interface{}) interface{}
	GetArgTpl() string
}

type SqlPartWriter interface {
	WriteSqlPart(SqlQueryWriter, Proto, string, int) string
}

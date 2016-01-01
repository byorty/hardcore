package types

type Query interface {
	ToNative() interface{}
	GetArgs() []interface{}
}

type QueryExecuter interface {
	Exec(Query, ModelDAO, Model)
	Query(Query, ModelDAO, Model)
	QueryRow(Query, ModelDAO, Model)
	Custom(Query, ...interface{})
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
}

type Pool interface {
	Add(string, DB) Pool
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

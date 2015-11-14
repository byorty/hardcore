package types

type Query interface {
	ToNative() interface{}
	GetArgs() []interface{}
}

type QueryExecuter interface {
	Exec(Query, DAO, interface{})
	Query(Query, DAO, interface{})
	QueryRow(Query, DAO, interface{})
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
}

type Pool interface {
	Add(string, DB) Pool
	ByDAO(DAO) DB
	ByName(string) DB
}

type QueryWriter interface {
	SetProto(Proto)
	SetTable(string)
	SetLogicChain([]LogicChain)
	SetProjections([]Projection)
	SetArgs([]interface{})
	WriteSelect() interface{}
//	WriteInsert() interface{}
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

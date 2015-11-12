package types

type Query interface {
	ToNative() interface{}
	GetArgs() []interface{}
}

type QueryExecuter interface {
	Exec(Query, DAO, interface{})
	Query(Query, DAO, interface{})
	QueryRow(Query, DAO, interface{})
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
	SetConditions([]Logic)
	WriteSelect() interface{}
}

type SqlQueryWriter interface {
	QueryWriter
	WriteField(string, string) string
	WriteTable(string) string
	WriteArg(int, ProtoProperty) string
}

type SqlPartWriter interface {
	WriteSqlPart(SqlQueryWriter, Proto, string, int) string
}

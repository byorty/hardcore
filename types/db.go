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
}

type Pool interface {
	Add(string, DB) Pool
	ByDAO(DAO) DB
	ByName(string) DB
}

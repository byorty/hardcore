package types

type QueryBuilder interface {
	GetQuery() interface{}
	GetArgs() []interface{}
}

type QueryExecuter interface {
	Exec(QueryBuilder, DAO, interface{})
	Query(QueryBuilder, DAO, interface{})
	QueryRow(QueryBuilder, DAO, interface{})
}

type DB interface {
	QueryExecuter
	Close()
}

type Pool interface {
	Add(string, DB) Pool
	ByDAO(DAO) DB
	ByName(string) DB
}

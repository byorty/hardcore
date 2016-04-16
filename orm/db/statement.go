package db

import (
	"database/sql"
	"github.com/byorty/hardcore/types"
)

type DBStatementImpl struct {
	*sql.Stmt
	sql string
}

func (d *DBStatementImpl) Exec(args ...interface{}) types.DBResult {
	result, err := d.Stmt.Exec(args...)
	return &DBResult{
		Result: result,
		err:    err,
		sql:    d.sql,
		args:   args,
	}
}

func (d *DBStatementImpl) Query(args ...interface{}) types.DBRows {
	rows, err := d.Stmt.Query(args...)
	return &DBRowsImpl{
		Rows: rows,
		err:  err,
		sql:  d.sql,
		args: args,
	}
}

func (d *DBStatementImpl) QueryRow(args ...interface{}) types.DBRow {
	row := d.Stmt.QueryRow(args...)
	return &DBRowImpl{
		Row:  row,
		sql:  d.sql,
		args: args,
	}
}

func (d *DBStatementImpl) Custom(args ...interface{}) types.DBCustomRow {
	row := d.Stmt.QueryRow(args...)
	return &DBCustomRowImpl{
		Row:  row,
		sql:  d.sql,
		args: args,
	}
}

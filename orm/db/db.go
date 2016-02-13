package db

import (
	"database/sql"
	"net/url"
	"github.com/byorty/hardcore/types"
	"github.com/byorty/hardcore/orm/db/writer"
	"github.com/byorty/hardcore/env"
)

var (
	writers = map[string]func() types.QueryWriter{
		"postgres": writer.NewPostgres,
	}
	supportsLastInsertId = map[string]bool{
		"postgres": false,
	}
	supportsReturningId = map[string]bool{
		"postgres": true,
	}
)

type SqlDBImpl struct {
	db                  *sql.DB
	writer              func() types.QueryWriter
	supportLastInsertId bool
	supportReturningId  bool
}

func NewSqlDB(uri string) types.DB {
	configUrl, err := url.Parse(uri)
	if err != nil {
		env.Me().GetLogger().Error("db - can't parse %s, detail - %v", uri, err)
		return nil
	}
	db, err := sql.Open(configUrl.Scheme, configUrl.String())
	if err != nil {
		env.Me().GetLogger().Error("db - can't connect to %s, detail - %v", uri, err)
		return nil
	}
	env.Me().GetLogger().Info("db - connect to %s success", uri)
	return &SqlDBImpl{
		db,
		writers[configUrl.Scheme],
		supportsLastInsertId[configUrl.Scheme],
		supportsReturningId[configUrl.Scheme],
	}
}

func (s *SqlDBImpl) SetMaxIdleConns(conns int) types.DB {
	s.db.SetMaxIdleConns(conns)
	return s
}

func (s *SqlDBImpl) SetMaxOpenConns(conns int) types.DB {
	s.db.SetMaxOpenConns(conns)
	return s
}

func (s SqlDBImpl) Exec(query types.Query) types.DBResult {
	sql := query.ToNative().(string)
	result, err := s.db.Exec(sql, query.GetArgs()...)
	return &DBResult{
		Result: result,
		err: err,
		sql: sql,
		args: query.GetArgs(),
	}
}

func (s SqlDBImpl) Query(query types.Query) types.DBRows {
	sql := query.ToNative().(string)
	rows, err := s.db.Query(sql, query.GetArgs()...)
	return &DBRowsImpl{
		Rows: rows,
		err: err,
		sql: sql,
		args: query.GetArgs(),
	}
}

func (s SqlDBImpl) QueryRow(query types.Query) types.DBRow {
	sql := query.ToNative().(string)
	row := s.db.QueryRow(sql, query.GetArgs()...)
	return &DBRowImpl{
		Row: row,
		sql: sql,
		args: query.GetArgs(),
	}
}

func (s SqlDBImpl) Custom(query types.Query) types.DBCustomRow {
	sql := query.ToNative().(string)
	row := s.db.QueryRow(sql, query.GetArgs()...)
	return &DBCustomRowImpl{
		Row: row,
		sql: sql,
		args: query.GetArgs(),
	}
}

func (s *SqlDBImpl) Close() {
	err := s.db.Close()
	if err != nil {
		env.Me().GetLogger().Error("db - can't close connection, detail - %v", err)
	}
}

func (s SqlDBImpl) GetKind() types.DBKind {
	return types.SqlDB
}

func (s SqlDBImpl) GetQueryWriter() types.QueryWriter {
	return s.writer()
}

func (s SqlDBImpl) SupportLastInsertId() bool {
	return s.supportLastInsertId
}

func (s SqlDBImpl) SupportReturningId() bool {
	return s.supportReturningId
}

func (s SqlDBImpl) Prepare(query types.Query) types.DBStatement {
	sql := query.ToNative().(string)
	stmt, err := s.db.Prepare(sql)
	if err == nil {
		return &DBStatementImpl{
			Stmt: stmt,
			sql: sql,
		}
	} else {
		env.Me().GetLogger().Error("db - can't prepare query %s, detail - %v", sql, err)
		return nil
	}
}

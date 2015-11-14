package db

import (
	"database/sql"
	"net/url"
	"github.com/byorty/hardcore/logger"
	"github.com/byorty/hardcore/types"
	"github.com/byorty/hardcore/db/writer"
)

var (
	writers = map[string]func() types.QueryWriter {
		"postgres": writer.NewPostgres,
	}
)

type sqlDB struct {
	db *sql.DB
	writer func() types.QueryWriter
}

func NewSqlDB(uri string) types.DB {
	configUrl, err := url.Parse(uri)
	if err != nil {
		logger.Err(`db - can't parse "%s", detail - %v`, uri, err)
		return nil
	}
	db, err := sql.Open(configUrl.Scheme, configUrl.String())
	if err != nil {
		logger.Err(`db - can't connect to "%s", detail - %v`, uri, err)
		return nil
	}
	return &sqlDB{db, writers[configUrl.Scheme]}
}

func (s sqlDB) Exec(query types.Query, d types.DAO, model interface{}) {
	_, err := s.db.Exec(query.ToNative().(string), query.GetArgs()...)
	if err != nil {
		logger.Warn(`db - can't exec "%s", detail - %v`, query, err)
	}
}

func (s sqlDB) Query(query types.Query, d types.DAO, models interface{}) {
	rows, err := s.db.Query(query.ToNative().(string), query.GetArgs()...)
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			d.ScanAll(rows, models)
		}
	} else {
		logger.Warn(`db - can't exec query "%s", detail - %v`, query, err)
	}
}

func (s sqlDB) QueryRow(query types.Query, d types.DAO, model interface{}) {
	row := s.db.QueryRow(query.ToNative().(string), query.GetArgs()...)
	if row != nil {
		d.Scan(row, model)
	}
}

func (s sqlDB) Custom(query types.Query, items ...interface{}) {
	row := s.db.QueryRow(query.ToNative().(string), query.GetArgs()...)
	if row != nil {
		row.Scan(items...)
	}
}

func (s *sqlDB) Close() {
	err := s.db.Close()
	if err != nil {
		logger.Err(`db - can't close connection, detail - %v`, err)
	}
}

func (s sqlDB) GetKind() types.DBKind {
	return types.SqlDB
}

func (s sqlDB) GetQueryWriter() types.QueryWriter {
	return s.writer()
}

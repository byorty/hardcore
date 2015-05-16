package db

import (
	"database/sql"
	"net/url"
	"github.com/byorty/hardcore/logger"
	"github.com/byorty/hardcore/types"
)

type sqlDB struct {
	db *sql.DB
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
	return &sqlDB{db}
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
			d.Scan(rows, models)
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

func (s *sqlDB) Close() {
	err := s.db.Close()
	if err != nil {
		logger.Err(`db - can't close connection, detail - %v`, err)
	}
}

func (s sqlDB) GetKind() types.DBKind {
	return types.SqlDB
}

type RawQuery []interface{}

func (r RawQuery) ToNative() interface{} {
	if len(r) > 1 {
		return r[0]
	} else {
		return nil
	}
}

func (r RawQuery) GetArgs() []interface{} {
	if len(r) >= 2 {
		return r[1:]
	} else {
		return nil
	}
}

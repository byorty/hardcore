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

func (s sqlDB) Exec(builder types.QueryBuilder, d types.DAO, model interface{}) {
	query := s.toQuery(builder)
	_, err := s.db.Exec(query, builder.GetArgs()...)
	if err != nil {
		logger.Warn(`db - can't exec "%s", detail - %v`, query, err)
	}
}

func (s sqlDB) Query(builder types.QueryBuilder, d types.DAO, models interface{}) {
	query := s.toQuery(builder)
	rows, err := s.db.Query(query, builder.GetArgs()...)
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			d.Scan(rows, models)
		}
	} else {
		logger.Warn(`db - can't exec query "%s", detail - %v`, query, err)
	}
}

func (s sqlDB) QueryRow(builder types.QueryBuilder, d types.DAO, model interface{}) {
	query := s.toQuery(builder)
	row := s.db.QueryRow(query, builder.GetArgs()...)
	if row != nil {
		d.Scan(row, model)
	}
}

func (s sqlDB) toQuery(builder types.QueryBuilder) string {
	return builder.GetQuery().(string)
}

func (s *sqlDB) Close() {
	err := s.db.Close()
	if err != nil {
		logger.Err(`db - can't close connection, detail - %v`, err)
	}
}



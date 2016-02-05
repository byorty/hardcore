package db

import (
	"database/sql"
	"net/url"
	"github.com/byorty/hardcore/types"
	"github.com/byorty/hardcore/orm/db/writer"
	"github.com/byorty/hardcore/env"
)

var (
	writers = map[string]func() types.QueryWriter {
		"postgres": writer.NewPostgres,
	}
	supportsLastInsertId = map[string]bool {
		"postgres": false,
	}
	supportsReturningId = map[string]bool {
		"postgres": true,
	}
)

type sqlDB struct {
	db *sql.DB
	writer func() types.QueryWriter
	supportLastInsertId bool
	supportReturningId bool
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
	return &sqlDB{
		db,
		writers[configUrl.Scheme],
		supportsLastInsertId[configUrl.Scheme],
		supportsReturningId[configUrl.Scheme],
	}
}

func (s sqlDB) Exec(query types.Query, d types.ModelDAO, model types.Model) {
	sql := query.ToNative().(string)
	result, err := s.db.Exec(sql, query.GetArgs()...)
	if err == nil {
		id, err := result.LastInsertId()
		if err == nil && id > 0 {
			setter := model.Proto().GetByName("id").GetSetter()
			setter(model, id)
		}
		env.Me().GetLogger().Debug("db - exec query %s %v", sql, query.GetArgs())
	} else {
		env.Me().GetLogger().Error("db - can't exec %s %v, detail - %v", sql, query.GetArgs(), err)
	}
}

func (s sqlDB) Query(query types.Query, d types.ModelDAO, models types.Model) {
	sql := query.ToNative().(string)
	rows, err := s.db.Query(sql, query.GetArgs()...)
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			d.ScanAll(rows, models)
		}
		env.Me().GetLogger().Debug("db - exec query %s %v", sql, query.GetArgs())
	} else {
		env.Me().GetLogger().Error("db - can't exec query %s %v, detail - %v", sql, query.GetArgs(), err)
	}
}

func (s sqlDB) QueryRow(query types.Query, d types.ModelDAO, model types.Model) {
	sql := query.ToNative().(string)
	row := s.db.QueryRow(sql, query.GetArgs()...)
	if row == nil {
		env.Me().GetLogger().Debug("db - result hasn't row, query %s %v", sql, query.GetArgs())
	} else {
		d.Scan(row, model)
		env.Me().GetLogger().Debug("db - exec query row %s %v", sql, query.GetArgs())
	}
}

func (s sqlDB) Custom(query types.Query, items ...interface{}) {
	sql := query.ToNative().(string)
	row := s.db.QueryRow(sql, query.GetArgs()...)
	if row == nil {
		env.Me().GetLogger().Debug("db - custom result hasn't row, query %s %v", sql, query.GetArgs())
	} else {
		row.Scan(items...)
		env.Me().GetLogger().Debug("db - exec custom query %s %v", sql, query.GetArgs())
	}
}

func (s *sqlDB) Close() {
	err := s.db.Close()
	if err != nil {
		env.Me().GetLogger().Error("db - can't close connection, detail - %v", err)
	}
}

func (s sqlDB) GetKind() types.DBKind {
	return types.SqlDB
}

func (s sqlDB) GetQueryWriter() types.QueryWriter {
	return s.writer()
}

func (s sqlDB) SupportLastInsertId() bool {
	return s.supportLastInsertId
}

func (s sqlDB) SupportReturningId() bool {
	return s.supportReturningId
}

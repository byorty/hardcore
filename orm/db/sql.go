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

func (s sqlDB) Exec(query types.Query) types.DBResult {
	sql := query.ToNative().(string)
	result, err := s.db.Exec(sql, query.GetArgs()...)
	return &DBResult{
		Result: result,
		err: err,
		sql: sql,
		args: query.GetArgs(),
	}
}

func (s sqlDB) Query(query types.Query) types.DBRows {
	sql := query.ToNative().(string)
	rows, err := s.db.Query(sql, query.GetArgs()...)
	return &DBRowsImpl{
		Rows: rows,
		err: err,
		sql: sql,
		args: query.GetArgs(),
	}
}

func (s sqlDB) QueryRow(query types.Query) types.DBRow {
	sql := query.ToNative().(string)
	row := s.db.QueryRow(sql, query.GetArgs()...)
	return &DBRowImpl{
		Row: row,
		sql: sql,
		args: query.GetArgs(),
	}
}

func (s sqlDB) Custom(query types.Query) types.DBCustomRow {
	sql := query.ToNative().(string)
	row := s.db.QueryRow(sql, query.GetArgs()...)
	return &DBCustomRowImpl{
		Row: row,
		sql: sql,
		args: query.GetArgs(),
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

func (s sqlDB) Prepare(query types.Query) types.DBStatement {
	sql := query.ToNative().(string)
	_, err := s.db.Prepare(sql)
	if err == nil {

	} else {
		env.Me().GetLogger().Error("db - can't prepare query %s, detail - %v", sql, err)
	}
	return nil
}

type DBRowImpl struct {
	*sql.Row
	sql string
	args []interface{}
}

func (d *DBRowImpl) One(model types.Model) {
	err := model.CommonDAO().Scan(d, model)
	if err == nil {
		env.Me().GetLogger().Debug("db - exec row query row %s %v", d.sql, d.args)
	} else {
		env.Me().GetLogger().Debug("db - can`t exec row query %s %v, detail - %v", d.sql, d.args, err)
	}
}

type DBCustomRowImpl struct {
	*sql.Row
	sql string
	args []interface{}
}

func (d *DBCustomRowImpl) One(items ...interface{}) {
	err := d.Scan(items...)
	if err == nil {
		env.Me().GetLogger().Debug("db - exec custom query %s %v", d.sql, d.args)
	} else {
		env.Me().GetLogger().Debug("db -  can`t exec custom query %s %v, detail - %v", d.sql, d.args, err)
	}
}

type DBRowsImpl struct {
	*sql.Rows
	err error
	sql string
	args []interface{}
}

func (d *DBRowsImpl) All(models types.Model) {
	if d.err == nil {
		defer d.Rows.Close()
		for d.Rows.Next() && d.err == nil {
			d.err = models.CommonDAO().ScanAll(d.Rows, models)
		}
		env.Me().GetLogger().Debug("db - exec rows query row %s %v", d.sql, d.args)
	} else {
		env.Me().GetLogger().Debug("db - can`t exec rows query %s %v, detail - %v", d.sql, d.args, d.err)
	}
}

type DBResult struct {
	sql.Result
	err error
	sql string
	args []interface{}
}

func (d *DBResult) One(model types.Model) {
	if d.err == nil {
		id, err := d.Result.LastInsertId()
		if err == nil && id > 0 {
			setter := model.Proto().GetByName("id").GetSetter()
			setter(model, id)
		}
		env.Me().GetLogger().Debug("db - exec query %s %v", d.sql, d.args)
	} else {
		env.Me().GetLogger().Error("db - can't exec %s %v, detail - %v", d.sql, d.args, d.err)
	}
}

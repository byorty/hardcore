package db

import (
	"database/sql"
	"github.com/byorty/hardcore/types"
	"github.com/byorty/hardcore/env"
)

type DBRowImpl struct {
	*sql.Row
	sql  string
	args []interface{}
}

func (d *DBRowImpl) One(model types.Model) {
	err := model.CommonDAO().Scan(d, model)
	if err == nil {
		env.Me().GetLogger().Debug("db - exec row query %s %v", d.sql, d.args)
	} else {
		env.Me().GetLogger().Error("db - can`t exec row query %s %v, detail - %v", d.sql, d.args, err)
	}
}

type DBCustomRowImpl struct {
	*sql.Row
	sql  string
	args []interface{}
}

func (d *DBCustomRowImpl) One(items ...interface{}) {
	err := d.Scan(items...)
	if err == nil {
		env.Me().GetLogger().Debug("db - exec custom query %s %v", d.sql, d.args)
	} else {
		env.Me().GetLogger().Error("db -  can`t exec custom query %s %v, detail - %v", d.sql, d.args, err)
	}
}

type DBRowsImpl struct {
	*sql.Rows
	err  error
	sql  string
	args []interface{}
}

func (d *DBRowsImpl) All(models types.Model) {
	if d.err == nil {
		defer d.Rows.Close()
		for d.Rows.Next() && d.err == nil {
			d.err = models.CommonDAO().ScanAll(d.Rows, models)
		}
		if d.err == nil {
			env.Me().GetLogger().Debug("db - exec rows query %s %v", d.sql, d.args)
		} else {
			env.Me().GetLogger().Error("db - can`t exec rows query %s %v, detail - %v", d.sql, d.args, d.err)
		}
	} else {
		env.Me().GetLogger().Error("db - can`t exec rows query %s %v, detail - %v", d.sql, d.args, d.err)
	}
}

type DBResult struct {
	sql.Result
	err  error
	sql  string
	args []interface{}
}

func (d *DBResult) One(model types.Model) {
	if d.err == nil {
		id, err := d.Result.LastInsertId()
		if err == nil && id > 0 {
			setter := model.Proto().GetByName("id").GetSetter()
			setter(model, id)
		}
		env.Me().GetLogger().Debug("db - exec result query %s %v", d.sql, d.args)
	} else {
		env.Me().GetLogger().Error("db - can't exec %s %v, detail - %v", d.sql, d.args, d.err)
	}
}

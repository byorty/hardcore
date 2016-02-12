package criteria

import (
"database/sql"
//"github.com/byorty/hardcore/types"
)

type PreparedCriteriaImpl struct {
	stmt *sql.Stmt
}

//func (p *PreparedCriteriaImpl) Exec(types.Model, args ...interface{}) {
//
//}
//
//func (p *PreparedCriteriaImpl) Query(types.Model, args ...interface{}) {
//
//}
//
//func (p *PreparedCriteriaImpl) QueryRow(types.Model, args ...interface{}) {
//
//}

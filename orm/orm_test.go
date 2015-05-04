package orm

import (
	"testing"
	"github.com/byorty/hardcore/orm/criteria"
	"github.com/byorty/hardcore/orm/criteria/cond"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/byorty/hardcore/orm/db"
)

type TestStatus struct {
	Id   int
	Name string
}

type Test struct {
	Id     int        `orm:"pk"`
	Name   string     // `orm:"-;name:some_name;default:someValue;required;pk;index;unique;relation:oneToOne;fetch:lazy"`
	Status TestStatus `orm:"relation:oneToOne;fetch:lazy"`
}

type TestDAO struct {

}

func (t TestDAO) GetDB() string {
	return "test"
}

func (t TestDAO) GetTable() string {
	return "test"
}

func (t TestDAO) Scan(row interface{}, model interface{}) {
	testModel := model.(*Test)
	err := row.(*sql.Row).Scan(&testModel.Id, &testModel.Name)
	if err != nil {
		fmt.Println(err)
	}
}

var (
	testDao = new(TestDAO)
//	testProto = dao.ProtoStruct{
//		Dao: testDao,
//		Map: dao.ProtoMap{
//			"Id"    : dao.ProtoProp{"id", dao.None},
//			"Name"  : dao.ProtoProp{"name", dao.None},
//			"Status": dao.ProtoProp{"status_id", dao.OneToOne},
//
//		},
//	}
)

func TestDB(t *testing.T) {
	sqlDb := db.NewSqlDB("postgres://byorty:MK99rc@localhost:5432/kapz?sslmode=disable&client_encoding=utf-8")
	if sqlDb == nil {
		t.Fail()
	}
	db.Pool().
		Add("default", sqlDb)
}

func TestCriteria(t *testing.T) {
	cr := criteria.New(testDao)
	cr.Add(
		cond.Eq("Id", 1),
	)
}

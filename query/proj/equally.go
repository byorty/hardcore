package proj

import (
	"github.com/byorty/hardcore/types"
	"fmt"
)

type EquallyImpl struct {
	BaseImpl
	value interface{}
}

func Eq(name string, value interface{}) *EquallyImpl {
	eq := new(EquallyImpl)
	eq.name = name
	eq.value = value
	return eq
}

func (e *EquallyImpl) UpdateSqlQueryWriter(writer types.SqlQueryWriter, proto types.Proto, table string) {
	property := proto.GetByName(e.name)
	writer.AddField(fmt.Sprintf("%s = %s", writer.WriteTable(property.GetField()), writer.GetArgTpl()))
	writer.AddArg(e.value)
}

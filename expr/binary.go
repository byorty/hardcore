package expr

import (
	"github.com/byorty/hardcore/types"
	"fmt"
)

type Binary struct {
	name string
	arg interface{}
	logic string
}

func (b Binary) GetArg() interface{} {
	return b.arg
}

func (b Binary) WriteSqlPart(writer types.SqlQueryWriter, proto types.Proto, table string, i int) string {
	property := proto.GetByName(b.name)
	return fmt.Sprintf(
		"(%s %s %s)",
		writer.WriteField(table, property.GetField()),
		b.logic,
		writer.WriteArg(i, property),
	)
}

func Eq(name string, arg interface{}) types.Logic {
	return &Binary{name, arg, "="}
}

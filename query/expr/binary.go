package expr

import (
	"fmt"
	"github.com/byorty/hardcore/types"
)

type Binary struct {
	name  string
	arg   interface{}
	logic string
}

func (b Binary) GetArg() interface{} {
	return b.arg
}

func (b Binary) GetName() string {
	return b.name
}

func (b Binary) WriteSqlPart(writer types.SqlQueryWriter, proto types.Proto, table string, i int) string {
	property := proto.GetByName(b.name)
	return fmt.Sprintf(
		"(%s %s %s)",
		writer.WriteField(table, property.GetField()),
		b.logic,
		writer.GetArgTpl(),
	)
}

func (b Binary) AddArg(criteria types.Criteria) {
	criteria.AddArg(b.arg)
}

func Eq(name string, arg interface{}) types.Logic {
	return &Binary{name, arg, "="}
}

func NotEq(name string, arg interface{}) types.Logic {
	return &Binary{name, arg, "!="}
}

func Gt(name string, arg interface{}) types.Logic {
	return &Binary{name, arg, ">"}
}

func GtEq(name string, arg interface{}) types.Logic {
	return &Binary{name, arg, ">="}
}

func Lt(name string, arg interface{}) types.Logic {
	return &Binary{name, arg, "<"}
}

func LtEq(name string, arg interface{}) types.Logic {
	return &Binary{name, arg, "<="}
}

func Like(name string, arg interface{}) types.Logic {
	return &Binary{name, arg, "LIKE"}
}

func NotLike(name string, arg interface{}) types.Logic {
	return &Binary{name, arg, "NOT LIKE"}
}

func ILike(name string, arg interface{}) types.Logic {
	return &Binary{name, arg, "ILIKE"}
}

func NotILike(name string, arg interface{}) types.Logic {
	return &Binary{name, arg, "NOT ILIKE"}
}

func Add(name string, arg interface{}) types.Logic {
	return &Binary{name, arg, "+"}
}

func Sub(name string, arg interface{}) types.Logic {
	return &Binary{name, arg, "-"}
}

func Mul(name string, arg interface{}) types.Logic {
	return &Binary{name, arg, "*"}
}

func Div(name string, arg interface{}) types.Logic {
	return &Binary{name, arg, "/"}
}

func Mod(name string, arg interface{}) types.Logic {
	return &Binary{name, arg, "%"}
}

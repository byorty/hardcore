package expr

import (
	"github.com/byorty/hardcore/types"
	"fmt"
	"strings"
)

type ExprIn struct {
	name string
	args []interface{}
}

func In(name string, args []int) types.Logic {
	in := newExprIn(name, len(args))
	for i, arg := range args {
		in.args[i] = arg
	}
	return in
}

func newExprIn(name string, len int) *ExprIn {
	in := new(ExprIn)
	in.name = name
	in.args = make([]interface{}, len)
	return in
}

func (e *ExprIn) convertArgs(args ...interface{}) {
	fmt.Println(args)
}

func (e ExprIn) GetArg() interface{} {
	return e.args
}

func (e ExprIn) WriteSqlPart(writer types.SqlQueryWriter, proto types.Proto, table string, i int) string {
	argTpls := make([]string, len(e.args))
	for i, _ := range e.args {
		argTpls[i] = writer.GetArgTpl()
	}
	property := proto.GetByName(e.name)
	return fmt.Sprintf(
		"(%s IN (%s))",
		writer.WriteField(table, property.GetField()),
		strings.Join(argTpls, ", "),
	)
}

func (e ExprIn) AddArg(criteria types.SelectCriteria) {
	for _, arg := range e.args {
		criteria.AddArg(arg)
	}
}


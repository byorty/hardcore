package expr

import (
	"github.com/byorty/hardcore/types"
	"fmt"
	"strings"
)

type ExprIn struct {
	name string
	args types.Slice
}

func In(name string, args types.Slice) types.Logic {
	return &ExprIn{name, args}
}

func (e ExprIn) GetName() string {
	return e.name
}

func (e ExprIn) GetArg() interface{} {
	return e.args
}

func (e ExprIn) WriteSqlPart(writer types.SqlQueryWriter, proto types.Proto, table string, i int) string {
	var argTpls []string
	if e.args == nil {
		argTpls = append(argTpls, writer.GetArgTpl())
	} else {
		argTpls = make([]string, e.args.Len())
		for i := 0;i < e.args.Len();i++ {
			argTpls = append(argTpls, writer.GetArgTpl())
		}
	}
	property := proto.GetByName(e.name)
	return fmt.Sprintf(
		"(%s IN (%s))",
		writer.WriteField(table, property.GetField()),
		strings.Join(argTpls, ", "),
	)
}

func (e ExprIn) AddArg(criteria types.Criteria) {
	if e.args == nil {
		criteria.AddArg(nil)
	} else {
		for i := 0;i < e.args.Len();i++ {
			criteria.AddArg(e.args.GetRaw(i))
		}
	}
}


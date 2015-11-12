package expr

import (
	"github.com/byorty/hardcore/types"
	"fmt"
)

type Binary struct {
	left  string
	right interface{}
	logic string
}

func (b Binary) GetLeft() string {
	return b.left
}

func (b Binary) GetRight() interface{} {
	return b.right
}

func (b Binary) GetLogic() string {
	return b.logic
}

func (b Binary) WriteSqlPart(writer types.SqlQueryWriter, proto types.Proto, table string, i int) string {
	property := proto.GetByName(b.GetLeft())
	return fmt.Sprintf(
		"(%s %s %s)",
		writer.WriteField(table, property.GetField()),
		b.GetLogic(),
		writer.WriteArg(i, property),
	)
}

func Eq(left string, right interface{}) types.Logic {
	return &Binary{left, right, "="}
}

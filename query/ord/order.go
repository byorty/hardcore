package ord

import (
	"fmt"
	"github.com/byorty/hardcore/types"
)

type OrderDirect int

const (
	AscDirect OrderDirect = iota
	DescDirect
)

var (
	sqlDirects = map[OrderDirect]string{
		AscDirect:  "ASC",
		DescDirect: "DESC",
	}
)

type OrderImpl struct {
	name   string
	direct OrderDirect
}

func (o OrderImpl) WriteSqlPart(writer types.SqlQueryWriter, proto types.Proto, table string, i int) string {
	property := proto.GetByName(o.name)
	return fmt.Sprintf(
		"%s %s",
		writer.WriteField(table, property.GetField()),
		sqlDirects[o.direct],
	)
}

func Asc(name string) types.Order {
	return &OrderImpl{name, AscDirect}
}

func Desc(name string) types.Order {
	return &OrderImpl{name, DescDirect}
}

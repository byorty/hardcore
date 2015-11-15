package proj

import (
	"github.com/byorty/hardcore/types"
	"fmt"
)

type CountImpl struct {
	BaseImpl
}

func Count(name string) *CountImpl {
	count := new(CountImpl)
	count.name = name
	return count
}

func (c *CountImpl) UpdateSqlQueryWriter(writer types.SqlQueryWriter, proto types.Proto, table string) {
	property := proto.GetByName(c.name)
	writer.AddField(fmt.Sprintf("COUNT(%s)", writer.WriteField(table, property.GetField())))
}

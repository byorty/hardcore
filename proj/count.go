package proj

import (
	"github.com/byorty/hardcore/types"
	"fmt"
)

type ProjCount struct {
	BaseImpl
}

func Count(name string) *ProjCount {
	count := new(ProjCount)
	count.name = name
	return count
}

func (p *ProjCount) WriteSqlPart(writer types.SqlQueryWriter, proto types.Proto, table string, i int) string {
	property := proto.GetByName(p.name)
	return fmt.Sprintf("COUNT(%s)", writer.WriteField(table, property.GetField()))
}

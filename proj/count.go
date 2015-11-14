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

func (p *ProjCount) UpdateSqlQueryWriter(writer types.SqlQueryWriter, proto types.Proto, table string) {
	property := proto.GetByName(p.name)
	writer.AddField(fmt.Sprintf("COUNT(%s)", writer.WriteField(table, property.GetField())))
}

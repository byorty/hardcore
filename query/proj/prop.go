package proj
import (
	"github.com/byorty/hardcore/types"
	"fmt"
)

type PropImpl struct {
	BaseImpl
}

func Prop(name string) *PropImpl {
	prop := new(PropImpl)
	prop.name = name
	return prop
}

func (p *PropImpl) UpdateSqlQueryWriter(writer types.SqlQueryWriter, proto types.Proto, table string) {
	property := proto.GetByName(p.name)
	writer.AddField(fmt.Sprintf("%s", writer.WriteField(table, property.GetField())))
}

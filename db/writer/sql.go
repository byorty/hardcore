package writer

import (
	"bytes"
	"strings"
	"fmt"
	"github.com/byorty/hardcore/types"
)

type SqlImpl struct {
	BaseImpl
}

func (s SqlImpl) writeSelect(writer types.SqlQueryWriter) interface{} {
	buf := new(bytes.Buffer)
	buf.WriteString("SELECT ")

	properties := make([]string, 0)
	if len(s.projections) > 0 {
		for i, projection := range s.projections {
			if projection.IsWriteSqlPart() {
				properties = append(properties, projection.WriteSqlPart(writer, s.proto, s.table, i))
			}
		}
	} else {
		for _, property := range s.proto.GetSlice() {
			properties = append(properties, writer.WriteField(s.table, property.GetField()))
		}
	}
	buf.WriteString(strings.Join(properties, ", "))
	buf.WriteString(fmt.Sprintf(" FROM %s", writer.WriteTable(s.table)))

	chainsLen := len(s.chains)
	if chainsLen > 0 {
		parts := make([]string, chainsLen)
		for i, chain := range s.chains {
			parts[i] = chain.WriteSqlPart(writer, s.proto, s.table, i)
		}
		buf.WriteString(" WHERE ")
		buf.WriteString(strings.Join(parts, " "))
	}

	return buf.String()
}

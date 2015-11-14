package writer

import (
	"bytes"
	"strings"
	"fmt"
	"github.com/byorty/hardcore/types"
)

type SqlImpl struct {
	BaseImpl
	fields []string
}

func (s *SqlImpl) AddField(field string) {
	s.fields = append(s.fields, field)
}

func (s *SqlImpl) SetFields(fields []string) {
	s.fields = fields
}

func (s *SqlImpl) GetFields() []string {
	return s.fields
}

func (s SqlImpl) writeSelect(writer types.SqlQueryWriter) interface{} {
	buf := new(bytes.Buffer)
	buf.WriteString("SELECT ")

	writer.SetFields(make([]string, 0))
	if len(s.projections) > 0 {
		for _, projection := range s.projections {
			projection.UpdateSqlQueryWriter(writer, s.proto, s.table)
		}
	} else {
		for _, property := range s.proto.GetSlice() {
			writer.AddField(writer.WriteField(s.table, property.GetField()))
		}
	}
	buf.WriteString(strings.Join(writer.GetFields(), ", "))
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

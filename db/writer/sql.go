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

	properties := make([]string, len(s.proto.GetSlice()))
	for i, property := range s.proto.GetSlice() {
		properties[i] = writer.WriteField(s.table, property.GetField())
	}
	buf.WriteString(strings.Join(properties, ", "))
	buf.WriteString(fmt.Sprintf(" FROM %s", writer.WriteTable(s.table)))

	conditionsLen := len(s.conditions)
	if conditionsLen > 0 {
		conditions := make([]string, conditionsLen)
		for i, condition := range s.conditions {
			conditions[i] = condition.WriteSqlPart(writer, s.proto, s.table, i)
		}

		buf.WriteString(" WHERE ")
		buf.WriteString(strings.Join(conditions, " "))
	}

	return buf.String()
}

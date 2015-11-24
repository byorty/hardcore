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

func (s SqlImpl) writeSelect(writer types.SqlQueryWriter) string {
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

	args := make([]interface{}, len(s.args))
	for i, arg := range s.args {
		args[i] = writer.WriteArg(i, arg)
	}

	return fmt.Sprintf(buf.String(), args...)
}

func (s SqlImpl) writeInsert(writer types.SqlQueryWriter) string {
	buf := new(bytes.Buffer)
	buf.WriteString("INSERT INTO ")
	buf.WriteString(writer.WriteTable(s.table))
	buf.WriteString(" (")

	writer.SetFields(make([]string, 0))
	for _, property := range s.proto.GetSlice() {
		if property.GetField() != "id" {
			writer.AddField(writer.WriteTable(property.GetField()))
		}
	}
	buf.WriteString(strings.Join(writer.GetFields(), ", "))
	buf.WriteString(") VALUES (")

	args := make([]string, len(s.args))
	for i, arg := range s.args {
		args[i] = fmt.Sprintf(writer.GetArgTpl(), writer.WriteArg(i, arg))
	}

	buf.WriteString(strings.Join(args, ", "))
	buf.WriteString(")")
	return buf.String()
}

func (s SqlImpl) writeUpdate(writer types.SqlQueryWriter) interface{} {
	oldArgs := writer.GetArgs()

	buf := new(bytes.Buffer)
	buf.WriteString("UPDATE ")
	buf.WriteString(writer.WriteTable(s.table))
	buf.WriteString(" SET ")

	writer.SetFields(make([]string, 0))
	writer.SetArgs(make([]interface{}, 0))

	for _, projection := range s.projections {
		projection.UpdateSqlQueryWriter(writer, s.proto, s.table)
	}
	buf.WriteString(strings.Join(writer.GetFields(), ", "))

	chainsLen := len(s.chains)
	if chainsLen > 0 {
		parts := make([]string, chainsLen)
		for i, chain := range s.chains {
			parts[i] = chain.WriteSqlPart(writer, s.proto, s.table, i)
		}
		buf.WriteString(" WHERE ")
		buf.WriteString(strings.Join(parts, " "))
	}

	newArgs := make([]interface{}, 0)
	newArgs = append(newArgs, writer.GetArgs()...)
	newArgs = append(newArgs, oldArgs...)
	writer.SetArgs(newArgs)

	placeholders := make([]interface{}, len(newArgs))
	for i, arg := range newArgs {
		placeholders[i] = writer.WriteArg(i, arg)
	}

	return fmt.Sprintf(buf.String(), placeholders...)
}

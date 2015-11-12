package writer

import (
	"github.com/byorty/hardcore/types"
	"fmt"
)

type PostgresImpl struct {
	SqlImpl
}

func NewPostgres() types.QueryWriter {
	return new(PostgresImpl)
}

func (p *PostgresImpl) WriteSelect() interface{} {
	return p.writeSelect(p)
}

func (p *PostgresImpl) WriteField(table string, field string) string {
	return fmt.Sprintf("%s.\"%s\"", p.WriteTable(table), field)
}

func (p *PostgresImpl) WriteTable(table string) string {
	return fmt.Sprintf("\"%s\"", table)
}

func (p *PostgresImpl) WriteArg(i int, property types.ProtoProperty) string {
	return fmt.Sprintf("$%d", i + 1)
}





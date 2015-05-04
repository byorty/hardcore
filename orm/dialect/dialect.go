package dialect

type PostgresqlDialect struct {

}

type PostgresqlDialectWriter interface {
	ToPostgresqlDialect(PostgresqlDialect)
}

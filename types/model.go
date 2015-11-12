package types

type StraightMappingModel interface {
	DAOConnected
	Prototyped
}

type SqlModelScanner interface {
	Scan(...interface{}) error
}

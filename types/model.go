package types

type Model interface {
	DAOConnected
	Prototyped
}

type SqlModelScanner interface {
	Scan(...interface{}) error
}

type DAOConnected interface {
	CommonDAO() ModelDAO
}

type ModelDAO interface {
	Proto() Proto
	GetDB() string
	GetTable() string
	ScanAll(interface{}, interface{})
	Scan(interface{}, interface{})
	All(Query, Model)
	One(Query, Model)
	Custom(ModelDAO, Query, ...interface{})
	Save(Model)
	Add(Model)
	Insert(Query, Model)
	Update(Query, Model)
}

type IntModelDAO interface {
	ModelDAO
	ById(int) ModelScanner
	ByIds([]int) ModelScanner
}

type Int8ModelDAO interface {
	ModelDAO
	ById(int8) ModelScanner
	ByIds([]int8) ModelScanner
}

type Int16ModelDAO interface {
	ModelDAO
	ById(int16) ModelScanner
	ByIds([]int16) ModelScanner
}

type Int32ModelDAO interface {
	ModelDAO
	ById(int32) ModelScanner
	ByIds([]int32) ModelScanner
}

type Int64ModelDAO interface {
	ModelDAO
	ById(int64) ModelScanner
	ByIds([]int64) ModelScanner
}

type UintModelDAO interface {
	ModelDAO
	ById(uint) ModelScanner
	ByIds([]uint) ModelScanner
}

type Uint8ModelDAO interface {
	ModelDAO
	ById(uint8) ModelScanner
	ByIds([]uint8) ModelScanner
}

type Uint16ModelDAO interface {
	ModelDAO
	ById(uint16) ModelScanner
	ByIds([]uint16) ModelScanner
}

type Uint32ModelDAO interface {
	ModelDAO
	ById(uint32) ModelScanner
	ByIds([]uint32) ModelScanner
}

type Uint64ModelDAO interface {
	ModelDAO
	ById(uint64) ModelScanner
	ByIds([]uint64) ModelScanner
}
package types

type Model interface {
	DAOConnected
	Prototyped
	IsScanned() bool
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
	Take(Model)
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

type IntModel interface {
	Model
	IntIdentifiable
	KindDAO() IntModelDAO
}

type Int8Model interface {
	Model
	Int8Identifiable
	KindDAO() Int8ModelDAO
}

type Int16Model interface {
	Model
	Int16Identifiable
	KindDAO() Int16ModelDAO
}

type Int32Model interface {
	Model
	Int32Identifiable
	KindDAO() Int32ModelDAO
}

type Int64Model interface {
	Model
	Int64Identifiable
	KindDAO() Int64ModelDAO
}

type UintModel interface {
	Model
	UintIdentifiable
	KindDAO() UintModelDAO
}

type Uint8Model interface {
	Model
	Uint8Identifiable
	KindDAO() Uint8ModelDAO
}

type Uint16Model interface {
	Model
	Uint16Identifiable
	KindDAO() Uint16ModelDAO
}

type Uint32Model interface {
	Model
	Uint32Identifiable
	KindDAO() Uint32ModelDAO
}

type Uint64Model interface {
	Model
	Uint64Identifiable
	KindDAO() Uint64ModelDAO
}

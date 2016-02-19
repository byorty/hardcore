package types

type Model interface {
	DAOConnected
	Prototyped
	IsScanned() bool
}

type DAOConnected interface {
	CommonDAO() ModelDAO
}

type ModelDAO interface {
	Proto() Proto
	GetDB() string
	GetTable() string
	ScanAll(interface{}, interface{}) error
	Scan(interface{}, interface{}) error
	SelectAll(Query, Model)
	SelectOne(Query, Model)
	Custom(ModelDAO, Query, ...interface{})
//	Save(Model)
//	Add(Model)
//	Take(Model)
	Insert(Query, Model)
	Update(Query, Model)
	AutoInit(DB)
	Init(DB)
}

type IntModelDAO interface {
	ModelDAO
	ById(int) DBRow
	ByIds([]int) ModelScanner
}

type Int8ModelDAO interface {
	ModelDAO
	ById(int8) DBRow
	ByIds([]int8) ModelScanner
}

type Int16ModelDAO interface {
	ModelDAO
	ById(int16) DBRow
	ByIds([]int16) ModelScanner
}

type Int32ModelDAO interface {
	ModelDAO
	ById(int32) DBRow
	ByIds([]int32) ModelScanner
}

type Int64ModelDAO interface {
	ModelDAO
	ById(int64) DBRow
	ByIds([]int64) ModelScanner
}

type UintModelDAO interface {
	ModelDAO
	ById(uint) DBRow
	ByIds([]uint) ModelScanner
}

type Uint8ModelDAO interface {
	ModelDAO
	ById(uint8) DBRow
	ByIds([]uint8) ModelScanner
}

type Uint16ModelDAO interface {
	ModelDAO
	ById(uint16) DBRow
	ByIds([]uint16) ModelScanner
}

type Uint32ModelDAO interface {
	ModelDAO
	ById(uint32) DBRow
	ByIds([]uint32) ModelScanner
}

type Uint64ModelDAO interface {
	ModelDAO
	ById(uint64) DBRow
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

package types

type DAOConnected interface {
	DAO() DAO
}

type DAO interface {
	Proto() Proto
	GetDB() string
	GetTable() string
	ScanAll(interface{}, interface{})
	Scan(interface{}, interface{})
	All(Query, StraightMappingModel)
	One(Query, StraightMappingModel)
	Custom(DAO, Query, ...interface{})
	Save(StraightMappingModel)
	Add(StraightMappingModel)
	Insert(Query, StraightMappingModel)
	Update(Query, StraightMappingModel)
	ById(int) StraightMappingModelScanner
	ByIds([]int) StraightMappingModelScanner
}

type EnumDAO interface {
    EnumScanner
    GetList() []Named
    Eq(Named, interface{}) bool
    Scan(src, dest Named)
}

type IntEnumDAO interface {
    ById(int) EnumScanner
//    ByIds([]int) EnumScanner
}



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
//	Customs(DAO, Query, ...[]interface{})
//	Save(StraightMappingModel)
//	Add(StraightMappingModel)
	Insert(StraightMappingModel, ...interface{})
	ById(int) SelectCriteria
	ByIds([]int) SelectCriteria
//	ByIds([]interface{}) SelectCriteria
//	ByIds(...int) SelectCriteria
}

type IdentityMap interface {
	Get(string, interface{}) bool
}

package types

type LogicChainKind int

const (
	AndLogicChainKind = iota
	OrLogicChainKind
)

type Projection interface {
	SqlPartWriter
	IsWriteSqlPart() bool
//	SetTable(string)
//	GetTable() string
}

type LogicChain interface {
	SqlPartWriter
	Add(Logic)
}

type Logic interface {
	SqlPartWriter
	GetArg() interface{}
}

type Criteria interface {
	Query
	One(StraightMappingModel)
	All(StraightMappingModel)
}

type CustomCriteria interface {
	Custom(...interface{})
//	Customs(...[]interface{})
}

type SelectCriteria interface {
	Criteria
	CustomCriteria
	Add(Projection) SelectCriteria
	And(Logic) SelectCriteria
	Or(Logic) SelectCriteria
}

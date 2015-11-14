package types

type LogicChainKind int

const (
	AndLogicChainKind = iota
	OrLogicChainKind
)

type Projection interface {
	Alias(string)
	UpdateSqlQueryWriter(SqlQueryWriter, Proto, string)
}

type LogicChain interface {
	SqlPartWriter
	Add(Logic)
}

type Logic interface {
	SqlPartWriter
	AddArg(SelectCriteria)
	GetArg() interface{}
}

type Criteria interface {
	Query
	AddArg(interface{})
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

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
	AddArg(Criteria)
	GetArg() interface{}
	GetName() string
}

type Order interface {
	SqlPartWriter
}

type Criteria interface {
	Query
	ModelScanner
	AddArg(interface{})
}

type CustomCriteria interface {
	Custom(...interface{})
}

type SelectCriteria interface {
	Criteria
	CustomCriteria
	Add(Projection) SelectCriteria
	And(Logic) SelectCriteria
	Or(Logic) SelectCriteria
	Order(Order) SelectCriteria
	Limit(int) SelectCriteria
	Offset(int) SelectCriteria
}

type UpdateCriteria interface {
	Criteria
	Add(Projection) UpdateCriteria
	And(Logic) UpdateCriteria
	Or(Logic) UpdateCriteria
}

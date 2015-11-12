package types

type LogicChainType int

const (
	AndLogicChainType = iota
	OrLogicChainType
)

type LogicChain interface {
	SqlPartWriter
	GetType() LogicChainType
}

type Logic interface {
	SqlPartWriter
	GetLeft() string
	GetRight() interface{}
	GetLogic() string
}

type Criteria interface {
	Query
	Add(Logic) Criteria
	One(StraightMappingModel)
	All(StraightMappingModel)
}

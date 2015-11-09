package types

type Logic interface {
	GetLeft() string
	GetRight() interface{}
	GetLogic() string
}

type Criteria interface {
	Query
	Add(Logic) Criteria
	One(StraightMappingModel)
	All(interface{})
}

package types

type Logic interface {
	GetLeft() string
	GetRight() interface{}
	GetLogic() string
}

type Criteria interface {
	Add(...Logic) Criteria
	One(interface{})
	All(interface{})
}

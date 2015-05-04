package cond

import "github.com/byorty/hardcore/types"

type binaryCondition struct {
	left  string
	right interface{}
	logic string
}

func (b binaryCondition) GetLeft() string {
	return b.left
}

func (b binaryCondition) GetRight() interface{} {
	return b.right
}

func (b binaryCondition) GetLogic() string {
	return b.logic
}

func Eq(left string, right interface{}) types.Logic {
	return &binaryCondition{left, right, "="}
}

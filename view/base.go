package view

import "github.com/byorty/hardcore/types"

type BaseImpl struct {
	scope types.RequestScope
}

func (b *BaseImpl) SetScope(scope types.RequestScope) {
	b.scope = scope
}

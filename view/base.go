package view

import "github.com/byorty/hardcore/types"

type BaseImpl struct {
	scope types.RequestScope
	ctrl  types.ActionController
}

func (b *BaseImpl) SetScope(scope types.RequestScope) {
	b.scope = scope
}

func (b *BaseImpl) SetController(ctrl types.ActionController) {
	b.ctrl = ctrl
}

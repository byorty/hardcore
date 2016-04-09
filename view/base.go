package view

import (
	"github.com/byorty/hardcore/types"
	"net/http"
)

type BaseImpl struct {
	scope  types.RequestScope
	ctrl   types.ActionController
	status int
}

func (b *BaseImpl) SetScope(scope types.RequestScope) types.View {
	b.scope = scope
	return b
}

func (b *BaseImpl) SetController(ctrl types.ActionController) types.View {
	b.ctrl = ctrl
	return b
}

func (b *BaseImpl) Render() {
	if b.status == 0 {
		b.status = http.StatusOK
	}
	b.scope.GetWriter().WriteHeader(b.status)
}

package ws

import (
    "github.com/byorty/hardcore/types"
)

func (h *Handler) CallAction(action interface{}, scope types.RequestScope) {
	if callable, ok := action.(types.CallableAction); ok {
		callable.Call(h, scope)
	} else {
		v := action.(func(*Handler, types.RequestScope) types.View)(h, scope)
		v.SetController(h)
		v.SetScope(scope)
		v.Render()
	}
}

var (
)

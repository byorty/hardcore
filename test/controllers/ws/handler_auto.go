package ws

import (
	"github.com/byorty/hardcore/types"
)

func (h *Handler) CallAction(action interface{}, rs types.RequestScope) {
	action.(func(*Handler, types.WebsocketScope))(h, rs.(types.WebsocketScope))
}

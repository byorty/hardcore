package scope

import (
	"github.com/byorty/hardcore/types"
	"github.com/gorilla/websocket"
)

type WebsocketImpl struct {
	types.RequestScope
	conn *websocket.Conn
}

func NewWebsocket(rs types.RequestScope, conn *websocket.Conn) types.WebsocketScope {
	return &WebsocketImpl{
		RequestScope: rs,
		conn:         conn,
	}
}

func (w WebsocketImpl) GetConnection() *websocket.Conn {
	return w.conn
}

package scope

import (
	"golang.org/x/net/websocket"
	"github.com/byorty/hardcore/types"
)

type WebsocketImpl struct {
	types.RequestScope
	conn *websocket.Conn
}

func NewWebsocket(rs types.RequestScope, conn *websocket.Conn) types.WebsocketScope {
	return &WebsocketImpl{
		RequestScope: rs,
		conn: conn,
	}
}

func (w WebsocketImpl) GetConnection() *websocket.Conn {
	return w.conn
}

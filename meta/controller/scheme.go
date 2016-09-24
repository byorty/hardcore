package controller

const (
	HttpScheme  Scheme = "http"
	HttpsScheme Scheme = "https"
	WsScheme    Scheme = "ws"
	WssScheme   Scheme = "wss"
)

type Scheme string

func (s Scheme) IsWebsocket() bool {
	return  s == WsScheme || s == WssScheme
}

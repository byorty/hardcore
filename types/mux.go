package types

import "net/http"

type MiddlewareFunc func(RequestScope)

type ActionController interface {
	CallAction(interface{}, RequestScope)
}

type CallableAction interface {
	Call(interface{}, RequestScope)
}

type ControllerConstructor func() ActionController

type Router interface {
	http.Handler
}

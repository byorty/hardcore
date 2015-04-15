package mux

import (
	"net/url"
	"net/http"
)

type MiddlewareFunc func(*RequestScope)

type ActionController interface {
	CallAction(interface{}, *RequestScope)
}

type ControllerFunc func() ActionController

type PathParams map[string]string

type RequestScope struct {
	UrlStr     string
	Url        *url.URL
	Writer     http.ResponseWriter
	Request    *http.Request
	PathParams map[string]string
}

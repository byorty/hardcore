package mux

import (
	"regexp"
	"github.com/byorty/hardcore"
)

type Handler struct {
	method    string
	scheme    *regexp.Regexp
	host      *regexp.Regexp
	path      *regexp.Regexp
	headers   map[string]*regexp.Regexp

	beforeMiddlewares []hardcore.MiddlewareFunc
	construct         hardcore.ControllerFunc
	handler           interface{}
	afterMiddlewares  []hardcore.MiddlewareFunc
}

package mux

import (
	"regexp"
	"github.com/byorty/hardcore"
)

type Handler struct {
	tpl       string
	urlRegexp regexp.Regexp
	construct hardcore.ControllerFunc
	handler   interface{}
}

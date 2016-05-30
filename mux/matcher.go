package mux

import (
	"github.com/byorty/hardcore/scope"
	"github.com/byorty/hardcore/types"
	"net/http"
	"regexp"
)

type ParamMatcherKind int

const (
	RequireParamMatcherKind ParamMatcherKind = iota
	OptionalParamMatcherKind
	EndPathParamMatcherKind
)

var (
	paramMatcherKindByByte = map[byte]ParamMatcherKind{
		'?': OptionalParamMatcherKind,
		'>': EndPathParamMatcherKind,
	}
)

type ParamMatcher struct {
	name      string
	len       int
	lastIndex int
	kind      ParamMatcherKind
}

type HeaderMatcher struct {
	key    string
	regexp *regexp.Regexp
	vars   []string
}

func newHeaderMatcher(key, value string) *HeaderMatcher {
	//tpl, vars := parseTpl(value)
	return &HeaderMatcher{
		key: key,
		//regexp: regexp.MustCompile(tpl),
		//vars:   vars,
	}
}

func (h *HeaderMatcher) Match(requestScope types.RequestScope) bool {
	matches := h.regexp.FindStringSubmatch(requestScope.GetRequest().Header.Get(h.key))
	match := len(matches) > 0
	if match {
		if requestScope.GetHeaderParams() == nil {
			requestScope.SetHeaderParams(scope.NewRequestScopeParams())
		}
		for i, match := range matches {
			if i > 0 {
				requestScope.SetHeaderParam(h.vars[i-1], match)
			}
		}
	}
	return match
}

type Matcher struct {
	path      string
	pathLen   int
	lastIndex int
	params    []*ParamMatcher
	paramsLen int
	headers   []*HeaderMatcher

	scopeConstruct    types.RequestScopeConstructor
	beforeMiddlewares []types.MiddlewareFunc
	construct         types.ControllerConstructor
	handler           interface{}
	afterMiddlewares  []types.MiddlewareFunc
}

func (m *Matcher) Match(path string, req *http.Request, rw http.ResponseWriter) (bool, types.RequestScope) {
	pathLen := len(path)
	pathLastIndex := pathLen - 1
	var params types.RequestScopeParams
	var i, j, x int
	for j < pathLen && i < m.pathLen {
		switch {
		case m.path[i] == ':', j == pathLastIndex && i < m.lastIndex && m.path[i+1] == ':':
			i++
			param := m.params[x]
			if param.kind == OptionalParamMatcherKind {
				i++
				i += param.len
			} else {
				i += param.len
			}
			start := j
			if param.kind == EndPathParamMatcherKind {
				i++
				j = pathLen
			} else {
				for j < pathLen && path[j] != '/' {
					j++
				}
			}
			if params == nil {
				params = scope.NewRequestScopeParams()
			}
			params.Set(param.name, path[start:j])
			if param.kind == OptionalParamMatcherKind {
				if start == j {
					i++
					j++
				}
			}
			x++

		case m.path[i] == path[j]:
			i++
			j++

		default:
			return false, nil
		}
	}
	if i != m.pathLen || j != pathLen {
		return false, nil
	}
	rs := m.scopeConstruct()
	rs.SetRequest(req)
	rs.SetWriter(rw)
	rs.SetPathParams(params)
	return true, rs
}

type Matchers []*Matcher

func (m *Matchers) Add(matcher *Matcher) {
	*m = append(*m, matcher)
}

func (m Matchers) Len() int {
	return len(m)
}

func (m Matchers) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m Matchers) Less(i, j int) bool {
	return m[i].path < m[j].path
}

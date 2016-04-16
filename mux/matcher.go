package mux

import (
	"bytes"
	"fmt"
	"github.com/byorty/hardcore/scope"
	"github.com/byorty/hardcore/types"
	"net/http"
	"regexp"
	"strings"
)

type ParamMatcher struct {
	name string
	len  int
}

func parseTpl(tpl string) (string, []string) {
	buf := new(bytes.Buffer)
	var level, start int
	vars := make([]string, 0)
	for i := 0; i < len(tpl); i++ {
		switch tpl[i] {
		case '{':
			if level++; level == 1 {
				start = i
			}
			break
		case '}':
			if level--; level == 0 {
				parts := strings.SplitN(tpl[start+1:i], ":", 2)
				if len(parts) == 2 {
					vars = append(vars, parts[0])
					subTpl, subVars := parseTpl(parts[1])
					if len(subVars) > 0 {
						vars = append(vars, subVars...)
						buf.WriteString(subTpl)
					} else {
						buf.WriteString(parts[1])
					}
				} else {
					panic(fmt.Sprintf("mux: missing name or pattern in %q", tpl[start:i]))
				}
			} else if level < 0 {
				panic(fmt.Sprintf("mux: unbalanced braces in %q", tpl))
			}
			break
		default:
			if level == 0 {
				buf.WriteByte(tpl[i])
			}
		}
	}
	if level != 0 {
		panic(fmt.Sprintf("mux: unbalanced braces in %q", tpl))
	}
	return buf.String(), vars
}

type HeaderMatcher struct {
	key    string
	regexp *regexp.Regexp
	vars   []string
}

func newHeaderMatcher(key, value string) *HeaderMatcher {
	tpl, vars := parseTpl(value)
	return &HeaderMatcher{
		key:    key,
		regexp: regexp.MustCompile(tpl),
		vars:   vars,
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
	path    string
	pathLen int
	//urlParams *ParamMatcher
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
	var params types.RequestScopeParams
	var i, j, x int
	for j < pathLen {
		switch {
		case i >= m.pathLen:
			return false, nil

		case m.path[i] == ':':
			param := m.params[x]
			i += param.len + 1
			start := j
			for j < pathLen && path[j] != '/' {
				j++
			}
			if params == nil {
				params = scope.NewRequestScopeParams()
			}
			params.Set(param.name, path[start:j])
			x++

		case m.path[i] == path[j]:
			i++
			j++

		default:
			return false, nil
		}
	}
	if i != m.pathLen {
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

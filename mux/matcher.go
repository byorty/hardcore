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
	regexp *regexp.Regexp
	vars   []string
}

func newParamMatcher(str string) *ParamMatcher {
	tpl, vars := parseTpl(str)
	return &ParamMatcher{
		regexp: regexp.MustCompile(tpl),
		vars:   vars,
	}
}

func (p *ParamMatcher) Match(url string) (bool, types.RequestScopeParams) {
	var params types.RequestScopeParams
	matches := p.regexp.FindStringSubmatch(url)
	match := len(matches) > 0
	if match {
		params = scope.NewRequestScopeParams()
		for i, match := range matches {
			if i > 0 {
				params.Set(p.vars[i-1], match)
			}
		}
	}
	return match, params
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
	urlParams *ParamMatcher
	headers   []*HeaderMatcher

	scopeConstruct    types.RequestScopeConstructor
	beforeMiddlewares []types.MiddlewareFunc
	construct         types.ControllerConstructor
	handler           interface{}
	afterMiddlewares  []types.MiddlewareFunc
}

func (m *Matcher) Match(url string, req *http.Request, rw http.ResponseWriter) (bool, types.RequestScope) {
	var scope types.RequestScope
	matchHeaders := true
	matchUrl, pathParams := m.urlParams.Match(url)
	if matchUrl {
		scope = m.scopeConstruct()
		scope.SetRequest(req)
		scope.SetWriter(rw)
		scope.SetPathParams(pathParams)

		for _, header := range m.headers {
			if !header.Match(scope) {
				matchHeaders = false
			}
		}
	}
	return matchUrl && matchHeaders, scope
}

type Matchers []*Matcher

func (m *Matchers) Add(matcher *Matcher) {
	*m = append(*m, matcher)
}

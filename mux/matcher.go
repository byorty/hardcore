package mux

import (
	"regexp"
	"fmt"
	"strings"
	"bytes"
)

type ParamMatcher struct {
	regexp *regexp.Regexp
	vars   []string
}

func newParamMatcher(str string) *ParamMatcher {
	tpl, vars := parseTpl(str)
	return &ParamMatcher{
		regexp: regexp.MustCompile(tpl),
		vars  : vars,
	}
}

func (p *ParamMatcher) Match(scope *RequestScope) bool {
	matches := p.regexp.FindStringSubmatch(scope.urlStr)
	match := len(matches) > 0
	if match {
		scope.PathParams = make(RequestScopeParams)
		for i, match := range matches {
			if i > 0 {
				scope.PathParams[p.vars[i - 1]] = match
			}
		}
	}
	return match
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
				parts := strings.SplitN(tpl[start + 1:i], ":", 2)
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
		key   : key,
		regexp: regexp.MustCompile(tpl),
		vars  : vars,
	}
}

func (h *HeaderMatcher) Match(scope *RequestScope) bool {
	matches := h.regexp.FindStringSubmatch(scope.Request.Header.Get(h.key))
	match := len(matches) > 0
	if match {
		if scope.HeaderParams == nil {
			scope.HeaderParams = make(RequestScopeParams)
		}
		for i, match := range matches {
			if i > 0 {
				scope.HeaderParams[h.vars[i - 1]] = match
			}
		}
	}
	return match
}

type Matcher struct {
	urlParams *ParamMatcher
	headers   []*HeaderMatcher

	beforeMiddlewares []MiddlewareFunc
	construct         ControllerFunc
	handler           interface{}
	afterMiddlewares  []MiddlewareFunc
}

func (m *Matcher) Match(scope *RequestScope) bool {
	matchUrl := m.urlParams.Match(scope)
	matchHeaders := true
	for _, header := range m.headers {
		if !header.Match(scope) {
			matchHeaders = false
		}
	}
	return matchUrl && matchHeaders
}

type Matchers []*Matcher

func (m *Matchers) Add(matcher *Matcher) {
	*m = append(*m, matcher)
}

package mux

import (
	"regexp"
	"fmt"
	"strings"
	"bytes"
)

type ParamsMatcher struct {
	regexp *regexp.Regexp
	vars   []string
}

func newParamsMatcher(str string) *ParamsMatcher {
	tpl, vars := parseTpl(str)
	return &ParamsMatcher{
		regexp: regexp.MustCompile(tpl),
		vars  : vars,
	}
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

type Matcher struct {
	urlParams *ParamsMatcher
	headers map[string]*ParamsMatcher

	beforeMiddlewares []MiddlewareFunc
	construct         ControllerFunc
	handler           interface{}
	afterMiddlewares  []MiddlewareFunc
}

func (m *Matcher) Match(scope *RequestScope) bool {
	matches := m.urlParams.regexp.FindStringSubmatch(scope.UrlStr)
	matchUrl := len(matches) > 0
	if matchUrl {
		scope.PathParams = make(PathParams)
		for i, match := range matches {
			if i > 0 {
				scope.PathParams[m.urlParams.vars[i - 1]] = match
			}
		}
	}
	return matchUrl
}

type Matchers []*Matcher

func (m *Matchers) Add(matcher *Matcher) {
	*m = append(*m, matcher)
}

package controller

import (
	"github.com/byorty/hardcore/meta/types"
	"strings"
	"fmt"
	"github.com/byorty/hardcore/utils"
)

type Action struct {
	Name string `xml:"name,attr"`
	Route string `xml:"path,attr"`
	Method string `xml:"method,attr"`
	Return string `xml:"return"`
	Params []*Param `xml:"params>param"`
	params []types.ActionParam
}

func (a Action) GetName() string {
	return a.Name
}

func (a Action) GetRoute() string {
	if len(a.Route) == 0 {
		a.Route = "/"
	}
	return a.Route
}

func (a Action) GetMethod() string {
	if len(a.Method) == 0 {
		a.Method = "GET"
	}
	return utils.UpperFirst(strings.ToLower(a.Method))
}

func (a Action) GetParams() []types.ActionParam {
	return a.params
}

func (a *Action) SetParams(params []types.ActionParam) {
	a.params = params
}

func (a Action) HasForm() bool {
	if len(a.params) == 1 && a.params[0].GetKind() == RequestScopeKind {
		return false
	} else {
		return true
	}
}

func (a Action) GetDefineKinds() string {
	kinds := make([]string, len(a.params))

	for i, param := range a.params {
		kinds[i] = param.GetDefineKind()
	}

	return strings.Join(kinds, ", ")
}

func (a Action) GetDefineParams() string {
	params := make([]string, len(a.params))

	for i, param := range a.params {
		params[i] = fmt.Sprintf("%s %s", param.GetName(), param.GetDefineKind())
	}

	return strings.Join(params, ", ")
}

func (a Action) GetDefineVars() string {
	vars := make([]string, len(a.params))

	for i, param := range a.params {
		vars[i] = param.GetDefineVarName()
	}

	return strings.Join(vars, ", ")
}

func (a Action) GetReturn() string {
	if a.HasForm() {
		if len(a.Return) == 0 {
			a.Return = "types.EncodeView"
		}
	} else {
		a.Return = "types.View"
	}
	return a.Return
}
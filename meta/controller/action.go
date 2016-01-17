package controller

import (
	"github.com/byorty/hardcore/meta/types"
	"strings"
	"fmt"
)

type Action struct {
	Name string `xml:"name,attr"`
	Route string `xml:"route,attr"`
	Method string `xml:"method,attr"`
	Params []*Param `xml:"params>param"`
	params []types.ActionParam
}

func (a Action) GetName() string {
	return a.Name
}

func (a Action) GetRoute() string {
	return a.Route
}

func (a Action) GetMethod() string {
	return a.Method
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
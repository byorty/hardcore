package view

import (
	"github.com/byorty/hardcore/scope"
	"github.com/byorty/hardcore/types"
)

type TmplImpl struct {
	BaseImpl
	filename string
	vars     map[string]interface{}
}

func Tmpl(filename string) types.TmplView {
	tmpl := new(TmplImpl)
	tmpl.filename = filename
	tmpl.vars = make(map[string]interface{})
	return tmpl
}

func (t *TmplImpl) SetTemplate(template string) types.TmplView {
	t.filename = template
	return t
}

func (t *TmplImpl) Set(key string, value interface{}) types.TmplView {
	t.vars[key] = value
	return t
}

func (t *TmplImpl) Render() {
	tmplCache := scope.App().GetTmplCache()
	if tmpl, ok := tmplCache[t.filename]; ok {
		clonedTmpl, err := tmpl.Clone()
		if err == nil {
			err = clonedTmpl.Execute(t.scope.GetWriter(), t.vars)
			if err != nil {
				scope.App().GetLogger().Error("tmpl - can't exec template %s, details - %v", t.filename, err)
			}
		} else {
			scope.App().GetLogger().Error("tmpl - can't clone template %s, details - %v", t.filename, err)
		}
	} else {
		scope.App().GetLogger().Error("tmpl - template %s not found", t.filename)
	}
}

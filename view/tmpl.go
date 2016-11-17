package view

import (
	"github.com/byorty/hardcore/is"
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

func (t *TmplImpl) SetStatus(status int) types.TmplView {
	t.status = status
	return t
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
	t.BaseImpl.Render()
	if viewer, ok := t.ctrl.(types.ExtendedViewer); ok {
		for _, tmpl := range viewer.GetHeaderTmpls() {
			t.render(tmpl)
		}
		t.render(t.filename)
		for _, tmpl := range viewer.GetFooterTmpls() {
			t.render(tmpl)
		}
	} else {
		t.render(t.filename)
	}
}

func (t *TmplImpl) render(tmplName string) {
	tmplCache := scope.App().GetTmplCache()
	if tmpl, ok := tmplCache[tmplName]; ok {
		clonedTmpl, err := tmpl.Clone()
		if is.Nil(err) {
			err = clonedTmpl.Execute(t.scope.GetWriter(), t.vars)
			if err != nil {
				scope.App().GetLogger().Error("tmpl - can't exec template %s, details - %v", tmplName, err)
			}
		} else {
			scope.App().GetLogger().Error("tmpl - can't clone template %s, details - %v", tmplName, err)
		}
	} else {
		scope.App().GetLogger().Error("tmpl - template %s not found", tmplName)
	}
}

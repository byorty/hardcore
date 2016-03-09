package view

import (
	"fmt"
	"github.com/byorty/hardcore/scope"
	"github.com/byorty/hardcore/types"
	"html/template"
	"path"
	"sync"
)

var (
	tmplCache = make(map[string]*template.Template)
	tmplMutex = new(sync.Mutex)
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
	var tmpl *template.Template
	var ok bool
	var err error
	filename := path.Join(scope.App().GetRootPath(), scope.App().GetTmplPath(), t.filename)

	tmplMutex.Lock()
	if tmpl, ok = tmplCache[filename]; ok {
		tmpl, err = tmpl.Clone()
		if err != nil {
			scope.App().GetLogger().Error("tmpl - can't clone template %s.%s, details - %v", filename, scope.App().GetTmplExt(), err)
		}
	} else {
		tmpl, err = template.ParseFiles(fmt.Sprintf("%s.%s", filename, scope.App().GetTmplExt()))
		if err == nil {
			tmplCache[filename] = tmpl
		} else {
			scope.App().GetLogger().Error("tmpl - can't parse file %s.%s, details - %v", filename, scope.App().GetTmplExt(), err)
		}
	}
	tmplMutex.Unlock()

	if tmpl != nil {
		err = tmpl.Execute(t.scope.GetWriter(), t.vars)
		if err != nil {
			scope.App().GetLogger().Error("tmpl - can't exec template %s.%s, details - %v", filename, scope.App().GetTmplExt(), err)
		}
	}
}

package plugin

import (
	"github.com/byorty/hardcore/scope"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"github.com/byorty/hardcore/types"
)

var (
	tmplFiles = make([]string, 0)
)

func NewTmpl() types.ApplicationPlugin {
	return NewTmplByName(scope.DefaultName)
}

func NewTmplByName(name string) types.ApplicationPlugin {
	return newByName(new(TmplImpl), name)
}

type TmplImpl struct{
	BaseImpl
}

func (t *TmplImpl) Run() {
	logger := scope.AppByName(t.name).GetLogger()
	dirname := filepath.Join(scope.AppByName(t.name).GetRootPath(), scope.AppByName(t.name).GetTmplPath())
	filepath.Walk(dirname, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			tmplFiles = append(tmplFiles, path)
		}
		return nil
	})

	tmplCache := make(map[string]*template.Template)
	for _, tmplFile := range tmplFiles {
		rel, err := filepath.Rel(dirname, tmplFile)
		if err == nil {
			relParts := strings.Split(rel, ".")

			if len(relParts) == 2 {
				tmplName := relParts[0]
				tmpl := template.New(tmplName)
				tmpl.Delims(scope.AppByName(t.name).GetTmplDelims())
				var buf []byte
				buf, err = ioutil.ReadFile(tmplFile)
				if err == nil {
					_, err = tmpl.Parse(string(buf))
					if err == nil {
						tmplCache[tmplName] = tmpl
					} else {
						logger.Error("tmpl - can't parse %s, details - %v", tmplFile, err)
					}
				} else {
					logger.Error("tmpl - can't read filename %s", rel)
				}
			} else {
				logger.Error("tmpl - can't give filename from %s", rel)
			}
		} else {
			logger.Error("tmpl - %v", err)
		}
	}
	scope.AppByName(t.name).SetTmplCache(tmplCache)
}

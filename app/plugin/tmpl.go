package plugin

import (
	"github.com/byorty/hardcore/scope"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var (
	tmplFiles = make([]string, 0)
)

type TmplImpl struct{}

func (t *TmplImpl) Run() {
	logger := scope.App().GetLogger()
	dirname := filepath.Join(scope.App().GetRootPath(), scope.App().GetTmplPath())
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
				tmpl.Delims(scope.App().GetTmplDelims())
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
	scope.App().SetTmplCache(tmplCache)
}

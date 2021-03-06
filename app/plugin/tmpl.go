package plugin

import (
	"github.com/byorty/hardcore/is"
	"github.com/byorty/hardcore/scope"
	"github.com/byorty/hardcore/types"
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

func NewTmpl() types.ApplicationPlugin {
	return new(TmplImpl)
}

func (t *TmplImpl) Run() {
	logger := scope.App().GetLogger()
	dirname := filepath.Join(scope.App().GetRootPath(), scope.App().GetTmplPath())
	filepath.Walk(dirname, func(path string, info os.FileInfo, err error) error {
		if err == nil {
			if !info.IsDir() {
				tmplFiles = append(tmplFiles, path)
			}
			return nil
		} else {
			scope.App().GetLogger().Error(err)
			return err
		}
	})

	tmplCache := make(map[string]*template.Template)
	for _, tmplFile := range tmplFiles {
		rel, err := filepath.Rel(dirname, tmplFile)
		if is.Nil(err) {
			relParts := strings.Split(rel, ".")

			if len(relParts) == 2 {
				tmplName := relParts[0]
				tmpl := template.New(tmplName)
				tmpl.Delims(scope.App().GetTmplDelims())
				var buf []byte
				buf, err = ioutil.ReadFile(tmplFile)
				if is.Nil(err) {
					_, err = tmpl.Parse(string(buf))
					if is.Nil(err) {
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

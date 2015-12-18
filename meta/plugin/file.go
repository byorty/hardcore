package plugin

import (
    "github.com/byorty/hardcore/meta"
    "bytes"
    "text/template"
)

type File struct {}

func (f *File) Do(env *meta.Environment) {
    for _, file := range env.Configuration.Files {
        buf := new(bytes.Buffer)
        tmpl := template.New(file.Name  + "_tpl")
        tmpl.Parse(file.Tpl)
        tmpl.Execute(buf, file.Params)
        env.Logger.Info(buf.String())
    }
}


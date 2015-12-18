package plugin

import (
    "github.com/byorty/hardcore/meta"
    "bytes"
    "text/template"
    "github.com/byorty/hardcore/meta/common"
    "github.com/astaxie/beego/utils"
    "os"
    "bufio"
)

type File struct {}

func (f *File) Do(env *meta.Environment) {
    for _, file := range env.Configuration.Files {
        buf := new(bytes.Buffer)
        tmpl := template.New(file.Name  + "_tpl")
        tmpl.Parse(file.Tpl)
        tmpl.Execute(buf, file.Params)

        if file.Kind == common.AutoFileKind {
            if utils.FileExists(file.Name) {
                os.Remove(file.Name)
            }
            f.create(env, file)
        } else {
            os.Remove(file.Name)
            if !utils.FileExists(file.Name) {
                f.create(env, file)
            }
        }
    }
}

func (f *File) create(env *meta.Environment, file common.File) {
    newFile, err := os.Create(file.Name)
    if err == nil {
        writer := bufio.NewWriter(newFile)
        tmpl := template.New(file.Name + "_tpl")
        tmpl.Parse(file.Tpl)
        tmpl.Execute(writer, file.Params)
        writer.Flush()
    } else {
        env.Logger.Error(err)
    }
}


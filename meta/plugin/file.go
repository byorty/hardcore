package plugin

import (
    "bytes"
    "text/template"
    "github.com/astaxie/beego/utils"
    "os"
    "bufio"
	"github.com/byorty/hardcore/meta/types"
	"github.com/byorty/hardcore/log"
)

type File struct {
	logger log.Logger
}

func (f *File) Do(env types.Environment) {
    f.logger = env.GetLogger()
    for _, file := range env.GetConfiguration().GetFiles() {
        buf := new(bytes.Buffer)
        tmpl := template.New(file.GetName()  + "_tpl")
        _, err := tmpl.Parse(file.GetTpl())
        if err != nil {
			f.logger.Error(err)
        }
        err = tmpl.Execute(buf, file.GetParams())
        if err != nil {
			f.logger.Error(err)
        }

        if file.GetKind() == types.AutoFileKind {
            if utils.FileExists(file.GetName()) {
                os.Remove(file.GetName())
            }
            f.create(file)
        } else {
            os.Remove(file.GetName())
            if !utils.FileExists(file.GetName()) {
                f.create(file)
            }
        }
    }
}

func (f *File) create(file types.File) {
    newFile, err := os.Create(file.GetName())
    if err == nil {
        writer := bufio.NewWriter(newFile)
        tmpl := template.New(file.GetName() + "_tpl")
        tmpl.Parse(file.GetTpl())
        tmpl.Execute(writer, file.GetParams())
        writer.Flush()
    } else {
		f.logger.Error(err)
    }
}


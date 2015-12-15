package plugin
import (
    "github.com/byorty/hardcore/meta"
    "github.com/byorty/hardcore/utils"
    "os"
)

type Directory struct {}

func (d *Directory) Do(env *meta.Environment) {
    conf := env.Configuration
    for _, container := range conf.ControllerContainers {
        d.createDirectory(env, container.Path)
    }
}

func (d *Directory) createDirectory(env *meta.Environment, path string) {
    if !utils.FileExists(path) {
        err := os.MkdirAll(path, os.ModePerm)
        if err != nil {
            env.Logger.Error(err)
            os.Exit(1)
        }
    }
}
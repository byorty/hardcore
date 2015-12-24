package plugin

import (
    "github.com/byorty/hardcore/utils"
    "os"
    "github.com/byorty/hardcore/meta/types"
)

type Directory struct {}

func (d *Directory) Do(env types.Environment) {
    for _, container := range env.GetConfiguration().GetContainers() {
        if !utils.FileExists(container.GetPath()) {
            err := os.MkdirAll(container.GetPath(), os.ModePerm)
            if err != nil {
                env.GetLogger().Error(err)
                os.Exit(1)
            }
        }
    }
}
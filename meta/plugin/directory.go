package plugin

import (
    "github.com/byorty/hardcore/utils"
    "os"
    "github.com/byorty/hardcore/meta/types"
    "path/filepath"
    "github.com/byorty/hardcore/log"
)

type Directory struct {
    logger log.Logger
}

func (d *Directory) Do(env types.Environment) {
    d.logger = env.GetLogger()
    for _, container := range env.GetConfiguration().GetContainers() {
        d.mkdir(container.GetPath())
    }
    d.mkdir(filepath.Join(env.GetAbsPath(), "configs"))
}

func (d *Directory) mkdir(dir string) {
    if !utils.FileExists(dir) {
        err := os.MkdirAll(dir, os.ModePerm)
        if err != nil {
            d.logger.Error(err)
        }
    }
}
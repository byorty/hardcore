package plugin

import (
    "github.com/byorty/hardcore/meta"
    "fmt"
    "strings"
    "path/filepath"
)

type Init struct {}

func (i *Init) Do(env *meta.Environment) {
    i.initControlles(env)
}

func (i *Init) initControlles(env *meta.Environment) {
    conf := env.Configuration
    for _, container := range conf.ControllerContainers {
        path := strings.Replace(container.Package, ".", string(filepath.Separator), -1)
        container.Import = fmt.Sprintf(
            "%s/%s",
            env.ImportPart,
            path,
        )
        container.Path = filepath.Join(env.AbsPath, path)
        pkgParts := strings.Split(container.Package, ".")
        container.ShortPackage = pkgParts[len(pkgParts) - 1]
    }

    for _, container := range conf.ControllerContainers {
        for _, controller := range container.Controllers {
            if len(controller.Extends) > 0 {
                for _, extend := range controller.Extends {
                    hasController, ctrlImport := i.hasController(conf, extend.Name)
                    if hasController {
                        extend.Import = ctrlImport
                    } else {
                        //                        env.Logger.Error("controller %s not found", extend.Name)
                        //                        os.Exit(1)
                    }
                }
            }
        }
    }
}

func (i *Init) hasController(conf *meta.Configuration, ctrl string) (bool, string)  {
    ctrl = i.clearName(ctrl)
    var ctrlImport string
    hasController := false
    for _, container := range conf.ControllerContainers {
        for _, controller := range container.Controllers {
            if ctrl == fmt.Sprintf("%s.%s", container.ShortPackage, controller.Name) {
                hasController = true
                break
            }
        }
        if hasController {
            ctrlImport = container.Import
            break
        }
    }
    return hasController, ctrlImport
}

func (i *Init) clearName(str string) string  {
    if len(str) > 0 && str[0] == uint8('*') {
        str = str[1:]
    }
    return str
}
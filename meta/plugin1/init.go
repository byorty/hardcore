package plugin

import (
    "github.com/byorty/hardcore/meta"
    "fmt"
    "strings"
    "path/filepath"
    "os"
    "regexp"
)

var (
    renameRegex = regexp.MustCompile(`([A-Z0-9])`)
)

type Init struct {}

func (i *Init) Do(env *meta.Environment) {
    i.initModels(env)
    i.initControlles(env)
}

func (i *Init) initModels(env *meta.Environment) {
    for _, container := range env.Configuration.ModelContainers {
        container.Path, container.Import, container.ShortPackage = i.createPathAndImportAndShortPackage(env, container.Package)
    }
    for _, container := range env.Configuration.ModelContainers {
        for _, enum := range container.Enums {
            enum.Filename, enum.AutoFilename = i.createFilenameAndAutoFilename(container.Path, enum.Name)
        }
        for _, model := range container.Models {
            model.Filename, model.AutoFilename = i.createFilenameAndAutoFilename(container.Path, model.Name)
        }
    }
    for _, container := range env.Configuration.ModelContainers {
        for _, model := range container.Models {
            model.Imports = make([]string, 0)
            if len(model.Extends) > 0 {
                for _, extend := range model.Extends {
                    hasModel, modelImport := i.hasModel(env.Configuration, extend.Name)
                    if hasModel {
                        extend.Import = modelImport
                    } else {
                        env.Logger.Error("parent model %s not found", extend.Name)
                        os.Exit(1)
                    }
                }
            }
            for _, prop := range model.Properties {
                if prop.HasRelation() {
                    var hasImport bool
                    var newImport string
                    hasImport, newImport = i.hasModel(env.Configuration, prop.Kind)
                    if !hasImport {
                        hasImport, newImport = i.hasEnum(env.Configuration, prop.Kind)
                    }
                    if hasImport {
                        prop.Kind = i.clearName(prop.Kind)
                        if container.Import == newImport {
                            parts := strings.Split(prop.Kind, ".")
                            prop.Kind = parts[len(parts) - 1]
                        } else {
                            model.AddImport(newImport)
                        }
                    }
                }
            }
        }
    }
}

func (i *Init) initControlles(env *meta.Environment) {
    for _, container := range env.Configuration.ControllerContainers {
        container.Path, container.Import, container.ShortPackage = i.createPathAndImportAndShortPackage(env, container.Package)
    }

    for _, container := range env.Configuration.ControllerContainers {
        for _, controller := range container.Controllers {
            controller.Filename, controller.AutoFilename = i.createFilenameAndAutoFilename(container.Path, controller.Name)
            if len(controller.Extends) > 0 {
                for _, extend := range controller.Extends {
                    hasController, ctrlImport := i.hasController(env.Configuration, extend.Name)
                    if hasController {
                        extend.Import = ctrlImport
                    } else {
                        env.Logger.Error("parent controller %s not found", extend.Name)
                        os.Exit(1)
                    }
                }
            }
        }
    }
}

func (i *Init) hasModel(conf *meta.Configuration, mdl string) (bool, string)  {
    mdl = i.clearName(mdl)
    var modelImport string
    hasModel := false
    for _, container := range conf.ModelContainers {
        for _, model := range container.Models {
            if mdl == fmt.Sprintf("%s.%s", container.ShortPackage, model.Name) {
                hasModel = true
                break
            }
        }
        if hasModel {
            modelImport = container.Import
            break
        }
    }
    return hasModel, modelImport
}

func (i *Init) hasEnum(conf *meta.Configuration, e string) (bool, string)  {
    e = i.clearName(e)
    var enumImport string
    hasEnum := false
    for _, container := range conf.ModelContainers {
        for _, enum := range container.Enums {
            if e == fmt.Sprintf("%s.%s", container.ShortPackage, enum.Name) {
                hasEnum = true
                break
            }
        }
        if hasEnum {
            enumImport = container.Import
            break
        }
    }
    return hasEnum, enumImport
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

func (i *Init) createPathAndImportAndShortPackage(env *meta.Environment, pkg string) (string, string, string) {
    path := strings.Replace(pkg, ".", string(filepath.Separator), -1)
    imp := fmt.Sprintf(
        "%s/%s",
        env.ImportPart,
        path,
    )
    pkgParts := strings.Split(pkg, ".")
    return filepath.Join(env.AbsPath, path), imp, pkgParts[len(pkgParts) - 1]
}

func (i *Init) createFilename(name string) string {
    name = strings.ToLower(renameRegex.ReplaceAllString(name, "_$1"))
    if name[0] == '_' {
        name = name[1:]
    }
    return name
}

func (i *Init) createFilenameAndAutoFilename(path, name string) (string, string) {
    filename := filepath.Join(path, i.createFilename(name))
    autoFilename := fmt.Sprintf("%s_auto", filename)
    return filename, autoFilename
}
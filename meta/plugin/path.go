package plugin

import (
    "github.com/byorty/hardcore/log"
    "github.com/byorty/hardcore/meta/types"
    "strings"
    "path/filepath"
    "fmt"
    "regexp"
)

var (
    filenameRegex = regexp.MustCompile(`([A-Z0-9])`)
)

type Path struct {
    logger log.Logger
    containers []types.Container
}

func (p *Path) Do(env types.Environment) {
    p.logger = env.GetLogger()
    p.containers = env.GetConfiguration().GetContainers()
    for _, container := range p.containers {
        container.SetPath(strings.Replace(container.GetPackage(), ".", string(filepath.Separator), -1))
        container.SetImport(fmt.Sprintf(
            "%s/%s",
            env.GetImportPart(),
            container.GetPath(),
        ))
        pkgParts := strings.Split(container.GetPackage(), ".")
        container.SetShortPackage(pkgParts[len(pkgParts) - 1])

        for _, entity := range container.GetEntities() {
            name := strings.ToLower(filenameRegex.ReplaceAllString(entity.GetName(), "_$1"))
            if name[0] == '_' {
                name = name[1:]
            }
            entity.SetFilename(filepath.Join(container.GetPath(), name))
            entity.SetAutoFilename(fmt.Sprintf("%s_auto", entity.GetFilename()))
        }
    }

//    for _, container := range p.containers {
//        for _, entity := range container.GetEntities() {
//            if entity.GetRawExtends() != nil {
//                p.logger.Debug("check %s parent entities", entity.GetFullname())
//                parentEntities := make([]types.Entity, 0)
//                for _, extend := range entity.GetRawExtends() {
//                    parentEntity := p.getEntity(extend)
//                    if parentEntity != nil && parentEntity.GetKind() == entity.GetKind() {
//                        parentEntities = append(parentEntities, parentEntity)
//                        p.logger.Debug("found %s parent entity", parentEntity.GetFullname())
//                    } else {
//                        p.logger.Error("parent entity %s not found", extend)
//                    }
//                }
//                entity.SetExtends(parentEntities)
//            }
//            switch entity.GetKind() {
//            case types.ControllerEntityKind:
//                return
//            case types.ModelEntityKind:
//
//                return
//            case types.EnumEntityKind:
//
//                return
//            }
//        }
//    }
}

func (i *Path) getEntity(name string) types.Entity {
    var needle types.Entity = nil
    for _, container := range i.containers {
        for _, entity := range container.GetEntities() {
            if name == entity.GetFullname() {
                needle = entity
                break
            }
        }
    }
    return needle
}
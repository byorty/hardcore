package plugin

import (
    "github.com/byorty/hardcore/log"
    "github.com/byorty/hardcore/meta/types"
    "strings"
    "path/filepath"
    "fmt"
)

type Init struct {
    logger log.Logger
    containers []types.Container
}

func (i *Init) Do(env types.Environment) {
    i.logger = env.GetLogger()
    i.containers = env.GetConfiguration().GetContainers()
    for _, container := range i.containers {
        container.SetPath(strings.Replace(container.GetPackage(), ".", string(filepath.Separator), -1))
        container.SetImport(fmt.Sprintf(
            "%s/%s",
            env.GetImportPart(),
            container.GetPath(),
        ))
        pkgParts := strings.Split(container.GetPackage(), ".")
        container.SetShortPackage(pkgParts[len(pkgParts) - 1])

        for _, entity := range container.GetEntities() {
            if entity.GetRawExtends() != nil {
                for _, extend := range entity.GetRawExtends() {

                }
            }
        }
    }
}

func (i *Init) getEntity(name string) types.Entity {
    var needle types.Entity
    for _, container := range i.containers {
        for _, entity := range container.GetEntities() {

        }
    }
}
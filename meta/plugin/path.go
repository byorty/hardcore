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
    for _, container := range env.GetConfiguration().GetContainers() {
        path := strings.Replace(container.GetPackage(), ".", string(filepath.Separator), -1)
        container.SetPath(filepath.Join(env.GetAbsPath(), path))
        container.SetImport(fmt.Sprintf(
            "%s/%s",
            env.GetImportPart(),
            path,
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
}
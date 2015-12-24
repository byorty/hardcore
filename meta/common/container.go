package common
import "github.com/byorty/hardcore/meta/types"

type Container struct {
    Package string `xml:"package,attr"`
    shortPackage string
    importName string
    path string
    entities []types.Entity
}

func (c Container) GetPackage() string {
    return c.Package
}

func (c *Container) SetShortPackage(shortPackage string) {
    c.shortPackage = shortPackage
}

func (c Container) GetShortPackage() string {
    return c.shortPackage
}

func (c *Container) SetImport(importName string) {
    c.importName = importName
}

func (c Container) GetImport() string {
    return c.importName
}

func (c *Container) SetPath(path string) {
    c.path = path
}

func (c Container) GetPath() string {
    return c.path
}

func (c *Container) SetEntities(entities []types.Entity) {
    c.entities = entities
}

func (c Container) GetEntities() []types.Entity {
    return c.entities
}

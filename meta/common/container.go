package common

import (
	"fmt"
	"github.com/byorty/hardcore/meta/types"
	"path/filepath"
	"strings"
)

type Container struct {
	Package      string `xml:"package,attr"`
	shortPackage string
	importName   string
	path         string
	entities     []types.Entity
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

func (c *Container) Init(env types.Environment) {
	path := strings.Replace(c.GetPackage(), ".", string(filepath.Separator), -1)
	c.SetPath(filepath.Join(env.GetAbsPath(), path))
	if env.GetImportPart() == "" {
		c.SetImport(path)
	} else {
		c.SetImport(fmt.Sprintf(
			"%s/%s",
			env.GetImportPart(),
			path,
		))
	}
	pkgParts := strings.Split(c.GetPackage(), ".")
	c.SetShortPackage(pkgParts[len(pkgParts)-1])
}

func (c *Container) Merge(slice types.EntitySlice) {
	entities := c.GetEntities()
	for i := 0; i < slice.Len(); i++ {
		entities = append(entities, slice.Get(i))
	}
	c.SetEntities(entities)
}

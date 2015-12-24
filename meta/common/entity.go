package common

import "github.com/byorty/hardcore/meta/types"

type Entity struct {
    Name string `xml:"name,attr"`
    Extends []string `xml:"extends>extend"`
    filename string
    autoFilename string
    imports []string
    entities []types.Entity
}

func (e Entity) GetName() string {
    return e.Name
}

func (e *Entity) SetExtends(entities []types.Entity) {
    e.entities = entities
}

func (e Entity) GetExtends() []types.Entity {
    return e.entities
}

func (e *Entity) SetImports(imports []string) {
    e.imports = imports
}

func (e Entity) GetImports() []string {
    return e.imports
}

func (e *Entity) SetFilename(filename string) {
    e.filename = filename
}

func (e Entity) GetFilename()string {
    return e.filename
}

func (e *Entity) SetAutoFilename(autoFilename string) {
    e.autoFilename = autoFilename
}

func (e Entity) GetAutoFilename()string {
    return e.autoFilename
}

func (e Entity) GetRawExtends() []string {
    return e.Extends
}
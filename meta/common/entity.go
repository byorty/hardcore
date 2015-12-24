package common

import (
    "github.com/byorty/hardcore/meta/types"
    "fmt"
)

type Entity struct {
    Name string `xml:"name,attr"`
    Extends []string `xml:"extends>extend"`
    filename string
    autoFilename string
    imports []string
    entities []types.Entity
    container types.Container
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

func (e *Entity) AddImport(newImport string) {
    hasImport := false
    for _, existsImport := range e.imports {
        if existsImport == newImport {
            hasImport = true
            break
        }
    }
    if !hasImport {
        e.imports = append(e.imports, newImport)
    }
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

func (e *Entity) SetContainer(container types.Container) {
    e.container = container
}

func (e Entity) GetContainer() types.Container {
    return e.container
}

func (e Entity) GetFullname() string {
    return fmt.Sprintf("%s.%s", e.container.GetShortPackage(), e.GetName())
}

func (e Entity) GetPointerName() string {
    return e.writePointer(e.GetName())
}

func (e Entity) writePointer(name string) string {
    return fmt.Sprintf("*%s", name)
}

func (e Entity) GetPointerFullName() string {
    return e.writePointer(e.GetFullname())
}

func (e *Entity) ClearName() {
    if e.Name[0] == uint8('*') {
        e.Name = e.Name[1:]
    }
}
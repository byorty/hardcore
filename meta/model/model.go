package model

import (
    "github.com/byorty/hardcore/meta/common"
    "github.com/byorty/hardcore/meta/types"
)

type Model struct {
    common.Entity
    Source string `xml:"source,attr"`
    Identifier *Property `xml:"properties>identifier"`
    Properties []*Property `xml:"properties>property"`
    Pattern Pattern `xml:"pattern"`
    Imports []string
}

func (m Model) GetKind() types.EntityKind {
    return types.ModelEntityKind
}

func (m *Model) AddImport(newImport string) {
    hasImport := false
    for _, existsImport := range m.Imports {
        if existsImport == newImport {
            hasImport = true
            break
        }
    }
    if !hasImport {
        m.Imports = append(m.Imports, newImport)
    }
}

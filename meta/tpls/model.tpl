package {{.Package}}

import ({{range .Imports}}
    "{{.}}"{{end}}
)

type {{.Name}} struct {
    {{.AutoName}}
}

type {{.MultipleName}} []*{{.Name}}

type {{.DaoName}} struct {
    {{.AutoDaoName}}
}

func ({{.ShortName}} *{{.DaoName}}) Init(db types.DB) {}
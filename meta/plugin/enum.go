package plugin

import (
    "github.com/byorty/hardcore/meta"
    "strings"
    "fmt"
    "github.com/byorty/hardcore/meta/model"
)

var (
    enumTpl = `package {{.Package}}

type {{.Name}} {{.Kind}}
`

    autoEnumTpl = `{{$name := .Name}}` +
`package {{.Package}}

import ({{range .AutoImports}}
    "{{.}}"{{end}}
)

func ({{.ShortName}} {{.Name}}) GetRawId() interface{} {
    return {{.ShortName}}
}

func ({{.ShortName}} {{.Name}}) GetId() {{.Kind}} {
    return {{.Kind}}({{.ShortName}})
}

func ({{.ShortName}} {{.Name}}) GetName() string {
    return {{.MapName}}[{{.ShortName}}]
}

func ({{.ShortName}} {{.Name}}) DAO() *{{.Name}}DAO {
    return new({{.Name}}DAO)
}

type {{.Name}}DAO struct {
    dao.Enum
}

func ({{.ShortName}} {{.Name}}DAO) GetList() []types.Named {
    return {{.SliceName}}
}

func ({{.ShortName}} {{.Name}}DAO) ById(id {{.Kind}}) types.EnumScanner {
    u.SetId(id)
    return {{.ShortName}}
}

func ({{.ShortName}} {{.Name}}DAO) One(enum types.Named) {
    {{.ShortName}}.FindOne({{.ShortName}}, enum)
}

func ({{.ShortName}} {{.Name}}DAO) Eq(named types.Named, id interface{}) bool {
    return named.({{.Name}}).GetId() == id
}

func ({{.ShortName}} {{.Name}}DAO) Scan(src, dest types.Named) {
    enum := dest.(*{{.Name}})
    (*enum) = src.({{.Name}})
}

const ({{range $k, $v := .Constants}}
    {{$v.Name}}{{if $v.IsWriteKind $k}} {{$name}}{{end}}{{if $v.HasValue}} = {{$v.Value}}{{end}}{{end}}
)

var (
    {{.SliceName}} = []types.Named{ {{range .Constants}}
        {{.Name}},{{end}}
    }
    {{.MapName}} = map[{{.Name}}]string{ {{range .Constants}}
        {{.Name}}: "{{.Label}}",{{end}}
    }
)
`
)

type Enum struct {}

func (e *Enum) Do(env *meta.Environment) {
    for _, container := range env.Configuration.ModelContainers {
        for _, enum := range container.Enums {
            if len(enum.Kind) == 0 {
                enum.Kind = model.IntEnumKind
            } else {
                if !model.HasEnumKind(enum.Kind) {
                    env.Logger.Error("unknown enum type %v", enum.Kind)
                }
            }

            hasIota := false
            for _, constant := range enum.Constants {
                if strings.Contains(constant.Value, "iota") && !hasIota {
                    hasIota = true
                }
                constant.Value = enum.GetValue(constant)
            }
            if !hasIota {
                for _, constant := range enum.Constants {
                    if !constant.HasValue() {
                        constant.Value = "iota"
                        break
                    }
                }
            }

            autoImports := make([]string, 0)
            autoImports = append(autoImports, "github.com/byorty/hardcore/types")
            autoImports = append(autoImports, "github.com/byorty/hardcore/orm/dao")

            varName := fmt.Sprintf("%s%s", strings.ToLower(enum.Name[0:1]), enum.Name[1:])
            tmplParams := map[string]interface{}{
                "ShortName": strings.ToLower(enum.Name[0:1]),
                "Name": enum.Name,
                "Package": container.ShortPackage,
                "SliceName": fmt.Sprintf("%sList", varName),
                "MapName": fmt.Sprintf("%sMap", varName),
                "Constants": enum.Constants,
                "Kind": enum.Kind,
                "AutoImports": autoImports,
            }

            env.Configuration.AddFile(enum.Filename, enumTpl, tmplParams)
            env.Configuration.AddFile(enum.AutoFilename, autoEnumTpl, tmplParams)
        }
    }
}
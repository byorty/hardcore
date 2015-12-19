package plugin

import (
    "github.com/byorty/hardcore/meta"
    "strings"
    "fmt"
    "github.com/byorty/hardcore/meta/model"
    "github.com/byorty/hardcore/utils"
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

func ({{.ShortName}} {{.Name}}) DAO() types.{{.Kind.DAO}} {
    return {{.VarDaoName}}
}

type {{.DaoName}} struct {}

func ({{.ShortName}} {{.DaoName}}) GetList() []types.Named {
    return {{.SliceName}}
}

func ({{.ShortName}} {{.DaoName}}) ById(id {{.Kind}}) types.EnumScanner {
    return dao.NewEnumScanner(id, {{.ShortName}})
}

func ({{.ShortName}} {{.DaoName}}) Eq(named types.Named, id interface{}) bool {
    return named.({{.Name}}).GetId() == id
}

func ({{.ShortName}} {{.DaoName}}) Scan(src, dest types.Named) {
    enum := dest.(*{{.Name}})
    (*enum) = src.({{.Name}})
}

const ({{range $k, $v := .Constants}}
    {{$v.Name}}{{if $v.IsWriteKind $k}} {{$name}}{{end}}{{if $v.HasValue}} = {{$v.Value}}{{end}}{{end}}
)

var (
    {{.VarDaoName}} {{.DaoName}}
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
                if !enum.Kind.IsValid() {
                    env.Logger.Error("unknown enum type %v for %s", enum.Kind, enum.Name)
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

            varName := utils.LowerFirst(enum.Name)
            tmplParams := map[string]interface{}{
                "ShortName": strings.ToLower(enum.Name[0:1]),
                "Name": enum.Name,
                "DaoName": fmt.Sprintf("%sDao", enum.Name),
                "Package": container.ShortPackage,
                "SliceName": fmt.Sprintf("%sList", varName),
                "MapName": fmt.Sprintf("%sMap", varName),
                "Constants": enum.Constants,
                "Kind": enum.Kind,
                "AutoImports": autoImports,
                "VarDaoName": fmt.Sprintf("%sDao", varName),
            }

            env.Configuration.AddAutoFile(enum.AutoFilename, autoEnumTpl, tmplParams)
            env.Configuration.AddFile(enum.Filename, enumTpl, tmplParams)
        }
    }
}
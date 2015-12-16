package plugin

import (
    "github.com/byorty/hardcore/meta"
    "strings"
    "fmt"
    "bytes"
    "text/template"
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

func ({{.ShortName}} {{.Name}}) GetId() {{.Kind}} {
    return {{.Kind}}({{.ShortName}})
}

func ({{.ShortName}} {{.Name}}) GetName() string {
    return {{.MapName}}[{{.ShortName}}];
}

const ({{range $k, $v := .Constants}}
    {{$v.Name}}{{if $v.IsWriteKind $k}} {{$name}}{{end}}{{if $v.HasValue}} = {{$v.Value}}{{end}}{{end}}
)

var (
    {{.SliceName}} = []types.Enum{ {{range .Constants}}
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
    env.Logger.Info(env.Configuration.ModelContainers)
    for _, container := range env.Configuration.ModelContainers {
        for _, enum := range container.Enums {
            if len(enum.Kind) == 0 {
                enum.Kind = "int"
            }

            hasIota := false
            for _, constant := range enum.Constants {
                if strings.Contains(constant.Value, "iota") {
                    hasIota = true
                }
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

            varName := fmt.Sprintf("%s%s", strings.ToLower(enum.Name[0:1]), enum.Name[1:])
            tmplParams := map[string]interface{}{
                "ShortName": strings.ToLower(enum.Name[0:1]),
                "Name": enum.Name,
                "Package": container.Package,
                "SliceName": fmt.Sprintf("%sList", varName),
                "MapName": fmt.Sprintf("%sMap", varName),
                "Constants": enum.Constants,
                "Kind": enum.Kind,
                "AutoImports": autoImports,
            }

            buf := new(bytes.Buffer)
            tmpl := template.New(enum.Filename  + "_tpl")
            tmpl.Parse(enumTpl)
            tmpl.Execute(buf, tmplParams)
            env.Logger.Info(buf.String())

            buf = new(bytes.Buffer)
            tmpl = template.New(enum.AutoFilename  + "_tpl")
            tmpl.Parse(autoEnumTpl)
            tmpl.Execute(buf, tmplParams)
            env.Logger.Info(buf.String())
        }
    }
}
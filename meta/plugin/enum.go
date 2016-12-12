package plugin

import (
	"fmt"
	"github.com/byorty/hardcore/meta/model"
	"github.com/byorty/hardcore/meta/types"
	"github.com/byorty/hardcore/utils"
	"strings"
)

var (
	enumTpl = `package {{.Package}}

type {{.Name}} {{.Kind}}

type {{.DaoName}} struct {}
`

	autoEnumTpl = `{{$name := .Name}}` +
		`package {{.Package}}

import ({{range .AutoImports}}
	"{{.}}"{{end}}
)

func ({{.ShortName}} {{.Name}}) GetRawId() interface{} {
	return {{.ShortName}}.GetId()
}

func ({{.ShortName}} {{.Name}}) GetId() {{.Kind}} {
	return {{.Kind}}({{.ShortName}})
}

func ({{.ShortName}} {{.Name}}) GetName() string {
	return {{.MapName}}[{{.ShortName}}]
}

func ({{.ShortName}} {{.Name}}) DAO() {{.Kind.DAO}} {
	return {{.VarDaoName}}
}

func ({{.ShortName}} {{.Name}}) CastDAO() {{.DaoName}} {
	return {{.VarDaoName}}
}

func ({{.ShortName}} {{.Name}}) GetProtoKind() types.ProtoKind {
	return types.{{.ProtoKind}}
}

type {{.MultipleName}} []{{.Name}}

func ({{.ShortName}} {{.MultipleName}}) Len() int {
	return len({{.ShortName}})
}

func ({{.ShortName}} {{.MultipleName}}) Less(x, y int) bool {
	return {{.ShortName}}[x].GetId() < {{.ShortName}}[y].GetId()
}

func ({{.ShortName}} {{.MultipleName}}) Swap(x, y int) {
	{{.ShortName}}[x], {{.ShortName}}[y] = {{.ShortName}}[y], {{.ShortName}}[x]
}

func ({{.ShortName}} {{.MultipleName}}) GetRaw(x int) interface{} {
	return {{.ShortName}}.Get(x)
}

func ({{.ShortName}} {{.MultipleName}}) Get(x int) {{.Name}} {
	return {{.ShortName}}[x]
}

func ({{.ShortName}} {{.DaoName}}) GetSlice() types.Slice {
	return {{.SliceName}}
}

func ({{.ShortName}} {{.DaoName}}) ById(id {{.Kind}}) types.EnumScanner {
	scanner := new(_{{.Name}}Scanner)
	scanner.id = id
	return scanner
}

type _{{.Name}}Scanner struct {
	dao.EnumScannerImpl
	dest *{{.Name}}
	id {{.Kind}}
}

func ({{.ShortName}} *_{{.Name}}Scanner) Scan(i int) {
	(*{{.ShortName}}.dest) = {{.SliceName}}.Get(i)
}

func ({{.ShortName}} *_{{.Name}}Scanner) One(named types.Named) {
	{{.ShortName}}.dest = named.(*{{.Name}})
	{{.ShortName}}.Find({{.ShortName}}, {{.SliceName}})
}

func ({{.ShortName}} *_{{.Name}}Scanner) Eq(i int) bool {
	return {{.ShortName}}.id == {{.SliceName}}.Get(i).GetId()
}

const ({{range $k, $v := .Constants}}
	{{$v.Name}}{{if $v.IsWriteKind $k}} {{$name}}{{end}}{{if $v.HasValue}} = {{$v.Value}}{{end}}{{end}}
)

var (
	{{.VarDaoName}} {{.DaoName}}
	{{.SliceName}} = {{.MultipleName}}{ {{range .Constants}}
		{{.Name}},{{end}}
	}
	{{.MapName}} = map[{{.Name}}]string{ {{range .Constants}}
		{{.Name}}: "{{.Label}}",{{end}}
	}
)`
)

type Enum struct{}

func (e Enum) Do(env types.Environment) {
	for _, container := range env.GetConfiguration().GetContainers() {
		if container.GetContainerKind() == types.EntityContainerKind {
			for _, entity := range container.GetEntities() {
				if entity.GetEntityKind() == types.EnumEntityKind {
					enum := entity.(*model.Enum)

					hasIota := false
					for _, constant := range enum.Constants {
						if strings.Contains(constant.Value, "iota") && !hasIota {
							hasIota = true
						}
						constant.Value = enum.GetValue(constant)
					}
					if enum.GetKind() != model.StringEnumKind {
						if !hasIota {
							for _, constant := range enum.Constants {
								if !constant.HasValue() {
									constant.Value = "iota + 1"
									break
								}
							}
						}
					}

					varName := utils.LowerFirst(enum.Name)
					tmplParams := map[string]interface{}{
						"ShortName": strings.ToLower(enum.Name[0:1]),
						"Name":      enum.Name,
						"DaoName":   fmt.Sprintf("%sDao", enum.Name),
						"Package":   container.GetShortPackage(),
						"SliceName": fmt.Sprintf("%sList", varName),
						"MapName":   fmt.Sprintf("%sNames", varName),
						"Constants": enum.Constants,
						"Kind":      enum.GetKind(),
						"ProtoKind": fmt.Sprintf("Proto%sEnumKind", utils.UpperFirst(string(enum.GetKind()))),
						"AutoImports": []string{
							types.DefaultImport,
							types.DaoImport,
						},
						"MultipleName": entity.GetMultipleName(),
						"VarDaoName":   fmt.Sprintf("%sDao", varName),
					}

					env.GetConfiguration().AddAutoFile(enum.GetAutoFilename(), autoEnumTpl, tmplParams)
					env.GetConfiguration().AddFile(enum.GetFilename(), enumTpl, tmplParams)
				}
			}
		}
	}
}

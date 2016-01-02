package plugin

import (
	"github.com/byorty/hardcore/log"
	"github.com/byorty/hardcore/meta/types"
	"github.com/byorty/hardcore/meta/model"
	"strings"
	"github.com/byorty/hardcore/utils"
	"fmt"
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

func ({{.ShortName}} {{.Name}}) DAO() {{.Kind.DAO}} {
	return {{.VarDaoName}}
}

type {{.MultipleName}} []*{{.Name}}

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

func ({{.ShortName}} {{.MultipleName}}) Get(x int) *{{.Name}} {
	return {{.ShortName}}[x]
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

type Enum struct {
	logger log.Logger
}

func (e Enum) Do(env types.Environment) {
	e.logger = env.GetLogger()

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
					if !hasIota {
						for _, constant := range enum.Constants {
							if !constant.HasValue() {
								constant.Value = "iota"
								break
							}
						}
					}

					varName := utils.LowerFirst(enum.Name)
					tmplParams := map[string]interface{}{
						"ShortName": strings.ToLower(enum.Name[0:1]),
						"Name": enum.Name,
						"DaoName": fmt.Sprintf("%sDao", enum.Name),
						"Package": container.GetShortPackage(),
						"SliceName": fmt.Sprintf("%sList", varName),
						"MapName": fmt.Sprintf("%sMap", varName),
						"Constants": enum.Constants,
						"Kind": enum.GetKind(),
						"AutoImports": []string{
							types.DefaultImport,
							types.DaoImport,
						},
						"MultipleName": entity.GetMultipleName(),
						"VarDaoName": fmt.Sprintf("%sDao", varName),
					}

					env.GetConfiguration().AddAutoFile(enum.GetAutoFilename(), autoEnumTpl, tmplParams)
					env.GetConfiguration().AddFile(enum.GetFilename(), enumTpl, tmplParams)
				}
			}
		}
	}
}
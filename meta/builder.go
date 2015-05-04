package meta

import (
	"os"
	"fmt"
	"path/filepath"
	"regexp"
	"reflect"
	"text/template"
	"bufio"
)

type Builder interface {
	Build(...interface{}) []*BuildResult
}

type BuildResult struct {
	Pkg     string
	PkgPath string
	Bytes   []byte
	Imports []string
}

type autoFile struct {
	Pkg     string
	Parts   []string
	Imports []string
}

type builderKind int

const (
	enumBuilderKind       builderKind = iota
	controllerBuilderKind
)

var (
	builders = map[builderKind]Builder{
		enumBuilderKind      : new(enumBuilder),
		controllerBuilderKind: new(controllerBuilder),
	}
	itemsByKind = make(map[builderKind][]interface{})
	endYRegex = regexp.MustCompile(`y$`)
	endSRegex = regexp.MustCompile(`s$`)
	IdentifiableByType = map[string]string{
		reflect.Int.String()       : "types.IntIdentifiable",
		reflect.Int8.String()      : "types.Int8Identifiable",
		reflect.Int16.String()     : "types.Int16Identifiable",
		reflect.Int32.String()     : "types.Int32Identifiable",
		reflect.Int64.String()     : "types.Int64Identifiable",
		reflect.Uint.String()      : "types.UintIdentifiable",
		reflect.Uint8.String()     : "types.Uint8Identifiable",
		reflect.Uint16.String()    : "types.Uint16Identifiable",
		reflect.Uint32.String()    : "types.Uint32Identifiable",
		reflect.Uint64.String()    : "types.Uint64Identifiable",
		reflect.Float32.String()   : "types.Float32Identifiable",
		reflect.Float64.String()   : "types.Float64Identifiable",
		reflect.Complex64.String() : "types.Complex64Identifiable",
		reflect.Complex128.String(): "types.Complex128Identifiable",
		reflect.String.String()    : "types.StringIdentifiable",
		"rune"                     : "types.RuneIdentifiable",
	}
	AutoFileTpl = `package {{.Package}}

{{if .HasImports}}import ({{range .Imports}}
    "{{.}}"{{end}}
){{end}}

{{range .Parts}}{{.}}{{end}}`
)

func RegisterEnums(enums ...interface{}) {
	itemsByKind[enumBuilderKind] = enums
}

func RegisterControllers(controllers ...interface{}) {
	itemsByKind[controllerBuilderKind] = controllers
}

func Build() {
	results := make([]*BuildResult, 0)
	for kind, items := range itemsByKind {
		results = append(results, builders[kind].Build(items...)...)
	}
	autoFiles := make(map[string]*autoFile)
	for _, result := range results {
		if aFile, ok := autoFiles[result.PkgPath]; ok {
			for _, resultImport := range result.Imports {
				found := false
				for _, autoImport := range aFile.Imports {
					if resultImport == autoImport {
						found = true
						break
					}
				}
				if !found {
					aFile.Imports = append(aFile.Imports, resultImport)
				}
			}
		} else {
			autoFiles[result.PkgPath] = new(autoFile)
			autoFiles[result.PkgPath].Pkg = result.Pkg
			autoFiles[result.PkgPath].Parts = make([]string, 0)
			autoFiles[result.PkgPath].Imports = result.Imports
		}
		autoFiles[result.PkgPath].Parts = append(autoFiles[result.PkgPath].Parts, string(result.Bytes))
	}

	projectPath := os.Getenv("GOPATH")
	if len(projectPath) == 0 {
		panic("GOPATH should be defined")
	}

	for path, autoFile := range autoFiles {
		filename := fmt.Sprintf("%ssrc/%s/auto.go", projectPath, path)
		dir := filepath.Dir(filename)
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			panic(fmt.Sprintf("dir \"%s\" not exists", dir))
		}
		var file *os.File
		_, err := os.Stat(filename)
		if os.IsExist(err) {
			os.Remove(filename)
		}
		file, err = os.Create(filename)
		if err == nil {
			writer := bufio.NewWriter(file)
			tmplParams := map[string]interface{}{
				"Package": autoFile.Pkg,
				"HasImports": len(autoFile.Imports) > 0,
				"Imports": autoFile.Imports,
				"Parts"  : autoFile.Parts,
			}
			tmpl := template.New(filename + "Template")
			tmpl.Parse(AutoFileTpl)
			tmpl.Execute(writer, tmplParams)
			writer.Flush()
		} else {
			panic(err)
		}
	}
}

package meta

import (
	"os"
	"fmt"
	"path/filepath"
	"bufio"
	"github.com/byorty/hardcore"
	"regexp"
	"reflect"
)

type Builder interface {
	Build(...interface{}) []*BuildResult
}

type BuildResult struct {
	Pkg     string
	PkgPath string
	Bytes   []byte
}

type autoFile struct {
	Pkg   string
	Parts [][]byte
}

type builderKind int

const (
	enumBuilderKind builderKind = iota
)

var (
	builders = map[builderKind]Builder{
		enumBuilderKind: new(enumBuilder),
	}
	itemsByKind = make(map[builderKind][]interface{})
	imports = make([]string, 0)
	endYRegex = regexp.MustCompile(`y$`)
	endSRegex = regexp.MustCompile(`s$`)
	IdentifiableByType = map[string]string{
		reflect.Int.String()       : "hardcore.IntIdentifiable",
		reflect.Int8.String()      : "hardcore.Int8Identifiable",
		reflect.Int16.String()     : "hardcore.Int16Identifiable",
		reflect.Int32.String()     : "hardcore.Int32Identifiable",
		reflect.Int64.String()     : "hardcore.Int64Identifiable",
		reflect.Uint.String()      : "hardcore.UintIdentifiable",
		reflect.Uint8.String()     : "hardcore.Uint8Identifiable",
		reflect.Uint16.String()    : "hardcore.Uint16Identifiable",
		reflect.Uint32.String()    : "hardcore.Uint32Identifiable",
		reflect.Uint64.String()    : "hardcore.Uint64Identifiable",
		reflect.Float32.String()   : "hardcore.Float32Identifiable",
		reflect.Float64.String()   : "hardcore.Float64Identifiable",
		reflect.Complex64.String() : "hardcore.Complex64Identifiable",
		reflect.Complex128.String(): "hardcore.Complex128Identifiable",
		reflect.String.String()    : "hardcore.StringIdentifiable",
		"rune"                     : "hardcore.RuneIdentifiable",
	}
)

func AddImport(imp string) {
	imports = append(imports, imp)
}

func RegisterEnums(enums ...interface{}) {
	itemsByKind[enumBuilderKind] = enums
}

func Build() {
	AddImport("github.com/byorty/hardcore")
	results := make([]*BuildResult, 0)
	for kind, items := range itemsByKind {
		results = append(results, builders[kind].Build(items...)...)
	}
	autoFiles := make(map[string]*autoFile)
	for _, result := range results {
		if _, ok := autoFiles[result.PkgPath]; !ok {
			autoFiles[result.PkgPath] = new(autoFile)
			autoFiles[result.PkgPath].Pkg = result.Pkg
			autoFiles[result.PkgPath].Parts = make([][]byte, 0)
		}
		autoFiles[result.PkgPath].Parts = append(autoFiles[result.PkgPath].Parts, result.Bytes)
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
		if os.IsNotExist(err) {
			file, err = os.Create(filename)
		} else {
			file, err = os.OpenFile(filename, os.O_WRONLY, os.ModePerm)
		}
		if err == nil {
			writer := bufio.NewWriter(file)
			writer.WriteString("package ")
			writer.WriteString(autoFile.Pkg)
			writer.WriteRune(hardcore.EOL)
			writer.WriteRune(hardcore.EOL)
			if len(imports) > 0 {
				writer.WriteString("import (")
				writer.WriteRune(hardcore.EOL)
				for _, imp := range imports {
					writer.WriteString(fmt.Sprintf("    \"%s\"", imp))
				}
				writer.WriteRune(hardcore.EOL)
				writer.WriteString(")")
				writer.WriteRune(hardcore.EOL)
				writer.WriteRune(hardcore.EOL)
			}
			for _, part := range autoFile.Parts {
				writer.Write(part)
			}
			writer.Flush()
		} else {
			panic(err)
		}
	}
}

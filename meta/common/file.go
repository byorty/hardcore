package common

const FileExt string = ".go"

type FileKind int

const (
    AutoFileKind FileKind = iota
    NormalFileKind
)

type File struct {
    Name string
    Tpl string
    Params map[string]interface{}
    Kind FileKind
}
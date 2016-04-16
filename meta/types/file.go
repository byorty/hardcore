package types

const FileExt string = ".go"

type FileKind int

const (
	AutoFileKind FileKind = iota
	NormalFileKind
)

type File interface {
	GetName() string
	GetTpl() string
	GetParams() map[string]interface{}
	GetKind() FileKind
}

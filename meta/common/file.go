package common

import "github.com/byorty/hardcore/meta/types"

type File struct {
    name string
    tpl string
    params map[string]interface{}
    kind types.FileKind
}

func NewFile(name, tpl string, params map[string]interface{}, kind types.FileKind) types.File {
	return &File {
		name: name,
		tpl: tpl,
		params: params,
		kind: kind,
	}
}

func (f File) GetName() string {
	return f.name
}

func (f File) GetTpl() string {
	return f.tpl
}

func (f File) GetParams() map[string]interface{} {
	return f.params
}

func (f File) GetKind() types.FileKind {
	return f.kind
}
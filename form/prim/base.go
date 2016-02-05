package prim

import "github.com/byorty/hardcore/types"

type BaseImpl struct {
	name string
	error string
	missing string
	wrong string
	required bool
	customs map[int]string
	dest interface{}
	source types.PrimitiveSource
}

func (b *BaseImpl) init(name string) {
	b.name = name
	b.missing = "missing argument"
	b.wrong = "wrong argument"
}

func (b BaseImpl) GetName() string {
	return b.name
}

func (b BaseImpl) GetError() string {
	return b.error
}

func (b *BaseImpl) Required() {
	b.required = true
}

func (b *BaseImpl) Missing(missing string) {
	b.missing = missing
}

func (b *BaseImpl) Wrong(wrong string) {
	b.wrong = wrong
}

func (b *BaseImpl) Custom(i int, error string) {
	if b.customs == nil {
		b.customs = make(map[int]string)
	}
	b.customs[i] = error
}

func (b *BaseImpl) Export(dest interface{}) {
	b.dest = dest
}

func (b *BaseImpl) SetSource(source types.PrimitiveSource) {
	b.source = source
}

func (b BaseImpl) GetSource() types.PrimitiveSource {
	return b.source
}

func (b BaseImpl) IsRequired() bool {
	return b.required
}

func (b *BaseImpl) SetError(error string) {
	b.error = error
}

func (b BaseImpl) GetWrong() string {
	return b.wrong
}

func (b BaseImpl) GetMissing() string {
	return b.missing
}

func (b BaseImpl) do(primitive types.Primitive, rawValue interface{}) bool {
	strValue, ok := rawValue.(string)
	if ok {
		if primitive.IsRequired() {
			if len(strValue) == 0 {
				primitive.SetError(primitive.GetMissing())
				return false
			} else {
				return primitive.ImportFromString(strValue)
			}
		} else {
			return primitive.ImportFromString(strValue)
		}

	} else {
		primitive.SetError(primitive.GetWrong())
		return ok
	}
}

package prim

import "github.com/byorty/hardcore/types"

type Base struct {
	name string
	error string
	missing string
	wrong string
	required bool
	customs map[int]string
	dest interface{}
	source types.PrimitiveSource
}

func (b Base) GetName() string {
	return b.name
}

func (b Base) GetError() string {
	return b.error
}

func (b *Base) Required() {
	b.required = true
}

func (b *Base) Missing(missing string) {
	b.missing = missing
}

func (b *Base) Wrong(wrong string) {
	b.wrong = wrong
}

func (b *Base) Custom(i int, error string) {
	if b.customs == nil {
		b.customs = make(map[int]string)
	}
	b.customs[i] = error
}

func (b *Base) Export(dest interface{}) {
	b.dest = dest
}

func (b *Base) SetSource(source types.PrimitiveSource) {
	b.source = source
}

func (b Base) GetSource() types.PrimitiveSource {
	return b.source
}

func (b Base) IsRequired() bool {
	return b.required
}

func (b *Base) SetError(error string) {
	b.error = error
}

func (b Base) GetWrong() string {
	return b.wrong
}

func (b Base) GetMissing() string {
	return b.missing
}

func (b Base) do(primitive types.Primitive, rawValue interface{}) bool {
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

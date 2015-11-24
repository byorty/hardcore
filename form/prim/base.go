package prim

type BasePrimitive struct {
	name string
	error string
	missing string
	wrong string
	required bool
	customs map[int]string
	dest interface{}
}

func (b *BasePrimitive) GetName() string {
	return b.name
}

func (b *BasePrimitive) GetError() string {
	return b.error
}

func (b *BasePrimitive) Required() {
	b.required = true
}

func (b *BasePrimitive) Missing(missing string) {
	b.missing = missing
}

func (b *BasePrimitive) Wrong(wrong string) {
	b.wrong = wrong
}

func (b *BasePrimitive) Custom(i int, error string) {
	if b.customs == nil {
		b.customs = make(map[int]string)
	}
	b.customs[i] = error
}

func (b *BasePrimitive) Export(dest interface{}) {
	b.dest = dest
}

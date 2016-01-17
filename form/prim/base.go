package prim

type Base struct {
	name string
	error string
	missing string
	wrong string
	required bool
	customs map[int]string
	dest interface{}
}

func (b *Base) GetName() string {
	return b.name
}

func (b *Base) GetError() string {
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

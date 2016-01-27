package exporter

type PropertyImpl struct {
	name       string
}

func NewProperty(name string) PropertyImpl {
	return PropertyImpl{name: name}
}

func (p PropertyImpl) GetName() string {
	return p.name
}
package exporter

import "github.com/byorty/hardcore/types"

type LimitPropertyImpl struct {
	PropertyImpl
	closure func(types.Named) interface{}
}

func newLimitProperty(name string, closure func(named types.Named) interface{}) types.ExportableProperty {
	return &LimitPropertyImpl{
		NewProperty(name),
		closure,
	}
}

func (l LimitPropertyImpl) GetValue(model interface{}) interface{} {
	return l.closure(model.(types.Named))
}

func NewNamed(named types.Named) types.Exporter {
	exp := new(BaseImpl)
	exp.SetProperties(namedProperties)
	exp.SetExportable(named)
	return exp
}

var (
	namedProperties = []types.ExportableProperty{
		newLimitProperty("id", func(limit types.Named) interface{} {
			return limit.GetRawId()
		}),
		newLimitProperty("name", func(limit types.Named) interface{} {
			return limit.GetName()
		}),
	}
)

package form

import (
	"github.com/byorty/hardcore/exporter"
	"github.com/byorty/hardcore/types"
)

type FormErrorPropertyImpl struct {
	exporter.PropertyImpl
	closure func(types.FormError) interface{}
}

func NewFormErrorProperty(name string, closure func(error types.FormError) interface{}) types.ExportableProperty {
	return &FormErrorPropertyImpl{
		exporter.NewProperty(name),
		closure,
	}
}

func (u FormErrorPropertyImpl) GetValue(model interface{}) interface{} {
	return u.closure(model.(types.FormError))
}

func ExportError(error types.FormError) types.Exporter {
	exp := new(exporter.BaseImpl)
	exp.SetProperties(errorProperties)
	exp.SetExportable(error)
	return exp
}

func ExportErrorMessage(error types.FormError) types.Exporter {
	exp := new(exporter.BaseImpl)
	exp.SetProperties(errorMessageProperties)
	exp.SetExportable(error)
	return exp
}

func ExportErrors(errors types.FormErrors) types.Exporter {
	exp := new(exporter.BaseImpl)
	exp.SetProperties(errorProperties)
	exp.SetExportable(errors)
	return exp
}

var (
	errorMessageProperties = []types.ExportableProperty{
		NewFormErrorProperty("message", func(error types.FormError) interface{} {
			return error.GetMessage()
		}),
	}
	errorProperties = []types.ExportableProperty{
		NewFormErrorProperty("name", func(error types.FormError) interface{} {
			return error.GetName()
		}),
		NewFormErrorProperty("message", func(error types.FormError) interface{} {
			return error.GetMessage()
		}),
	}
)
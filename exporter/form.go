package exporter

import (
	"github.com/byorty/hardcore/types"
)

type FormErrorPropertyImpl struct {
	PropertyImpl
	closure func(types.FormError) interface{}
}

func newFormErrorProperty(name string, closure func(error types.FormError) interface{}) types.ExportableProperty {
	return &FormErrorPropertyImpl{
		NewProperty(name),
		closure,
	}
}

func (u FormErrorPropertyImpl) GetValue(model interface{}) interface{} {
	return u.closure(model.(types.FormError))
}

func NewFormError(error types.FormError) types.Exporter {
	exp := new(BaseImpl)
	exp.SetProperties(errorProperties)
	exp.SetExportable(error)
	return exp
}

func NewFormErrorMessage(error types.FormError) types.Exporter {
	exp := new(BaseImpl)
	exp.SetProperties(errorMessageProperties)
	exp.SetExportable(error)
	return exp
}

func NewFormErrors(errors types.FormErrors) types.Exporter {
	exp := new(BaseImpl)
	exp.SetProperties(errorProperties)
	exp.SetExportable(errors)
	return exp
}

var (
	errorMessageProperties = []types.ExportableProperty{
		newFormErrorProperty("message", func(error types.FormError) interface{} {
			return error.GetMessage()
		}),
	}
	errorProperties = []types.ExportableProperty{
		newFormErrorProperty("name", func(error types.FormError) interface{} {
			return error.GetName()
		}),
		newFormErrorProperty("message", func(error types.FormError) interface{} {
			return error.GetMessage()
		}),
	}
)

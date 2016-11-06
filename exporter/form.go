package exporter

import (
	"github.com/byorty/hardcore/types"
)

type FormErrorImpl struct {
	error types.FormError
	props []FormPropertyImpl
	kind types.ProtoKind
}

func NewFormError(error types.FormError) types.Exporter {
	return newFormError(error, formMessageProperties)
}

func newFormError(error types.FormError, props []FormPropertyImpl) types.Exporter {
	exp := new(FormErrorImpl)
	exp.props = props
	exp.error = error
	exp.kind = types.ProtoModelKind
	return exp
}

func (f *FormErrorImpl) Len() int {
	return len(f.props)
}

func (f *FormErrorImpl) Get(i int) types.ExportableProperty {
	return f.props[i]
}

func (f *FormErrorImpl) GetProtoKind() types.ProtoKind {
	return f.kind
}

func (f *FormErrorImpl) Export(i int, encoder types.Encoder) {
	f.props[i].closure(f.error, encoder)
}

type FormErrorsImpl struct {
	errors types.FormErrors
}

func NewFormErrors(errors types.FormErrors) types.SliceExporter {
	exp := new(FormErrorsImpl)
	exp.errors = errors
	return exp
}

func (f *FormErrorsImpl) Len() int {
	return f.errors.Len()
}

func (f *FormErrorsImpl) ExportItem(i int, encoder types.Encoder) {
	encoder.One(newFormError(f.errors.Get(i), formMessagesProperties))
}

type FormPropertyImpl struct {
	PropertyImpl
	closure func(types.FormError, types.Encoder)
}

func newFormErrorProperty(name string, closure func(types.FormError, types.Encoder)) FormPropertyImpl {
	return FormPropertyImpl{
		NewPropertyWithKind(name, types.ProtoModelKind),
		closure,
	}
}

var (
	formMessageProperties = []FormPropertyImpl{
		newFormErrorProperty("message", func(error types.FormError, encoder types.Encoder) {
			encoder.EncodeString(error.GetMessage())
		}),
		newFormErrorProperty("code", func(error types.FormError, encoder types.Encoder) {
			encoder.EncodeInt(error.GetCode())
		}),
	}
	formMessagesProperties = []FormPropertyImpl{
		newFormErrorProperty("name", func(error types.FormError, encoder types.Encoder) {
			encoder.EncodeString(error.GetName())
		}),
		newFormErrorProperty("message", func(error types.FormError, encoder types.Encoder) {
			encoder.EncodeString(error.GetMessage())
		}),
		newFormErrorProperty("code", func(error types.FormError, encoder types.Encoder) {
			encoder.EncodeInt(error.GetCode())
		}),
	}
)

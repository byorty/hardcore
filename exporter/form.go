package exporter

import (
	"github.com/byorty/hardcore/types"
	"github.com/byorty/hardcore/is"
)

type FormErrorImpl struct {
	error  types.FormError
	errors types.FormErrors
	props  []FormPropertyImpl
	kind   types.ProtoKind
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
	if is.Eq(f.kind, types.ProtoModelKind) {
		return len(f.props)
	} else {
		return f.errors.Len()
	}
}

func (f *FormErrorImpl) Get(i int) types.ExportableProperty {
	return f.props[i]
}

func (f *FormErrorImpl) GetProtoKind() types.ProtoKind {
	return f.kind
}

func (f *FormErrorImpl) Export(i int, encoder types.Encoder) {
	if is.Eq(f.kind, types.ProtoModelKind) {
		f.props[i].closure(f.error, encoder)
	} else {
		encoder.Encode(newFormError(f.errors.Get(i), formMessagesProperties))
	}
}

func NewFormErrors(errors types.FormErrors) types.Exporter {
	exp := new(FormErrorImpl)
	exp.props = formMessageProperties
	exp.errors = errors
	exp.kind = types.ProtoModelSliceKind
	return exp
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

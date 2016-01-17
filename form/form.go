package form

import "github.com/byorty/hardcore/types"

type FormImpl struct {
	errors FormErrorsImpl
	primitives []types.Primitive
}

func New() types.Form {
	return &FormImpl{
		errors: NewFormErrors(),
		primitives: make([]types.Primitive, 0),
	}
}

func (f *FormImpl) Add(primitive types.Primitive) types.Form {
	f.primitives = append(f.primitives, primitive)
	return f
}

func (f *FormImpl) AddErrorMessage(name, message string) types.Form {
	f.AddError(NewError(name, message))
	return f
}

func (f *FormImpl) AddErrorMessageWithCode(name, message string, code int) types.Form {
	f.AddError(NewErrorWithCode(name, message, code))
	return f
}

func (f *FormImpl) AddError(err types.FormError) types.Form {
	f.errors = append(f.errors, err)
	return f
}

func (f *FormImpl) Check(verifiable types.FormVerifiable) bool {
	for _, primitive := range f.primitives {
		verifiable.Verify(f, primitive)
	}

	return f.errors.Len() == 0
}

func (f FormImpl) GetErrors() types.FormErrors {
	return f.errors
}

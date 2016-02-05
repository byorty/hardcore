package form

import "github.com/byorty/hardcore/types"

type FormErrorImpl struct {
	name string
	message string
	code int
}

func (f FormErrorImpl) GetName() string {
	return f.name
}

func (f FormErrorImpl) GetMessage() string {
	return f.message
}

func (f FormErrorImpl) GetCode() int {
	return f.code
}

func NewError(name, message string) types.FormError {
	return newError(name, message)
}

func newError(name, message string) *FormErrorImpl {
	return &FormErrorImpl{
		name: name,
		message: message,
	}
}

func NewErrorWithCode(name, message string, code int) types.FormError {
	err := newError(name, message)
	err.code = code
	return err
}

type FormErrorsImpl []types.FormError

func NewFormErrors() types.FormErrors {
	errors := make(FormErrorsImpl, 0)
	return &errors
}

func (f *FormErrorsImpl) Add(error types.FormError) {
	(*f) = append((*f), error)
}

func (f FormErrorsImpl) Len() int {
	return len(f)
}

func (f FormErrorsImpl) Less(x, y int) bool {
	return f[x].GetMessage() < f[y].GetMessage()
}

func (f FormErrorsImpl) Swap(x, y int) {
	f[x], f[y] = f[y], f[x]
}

func (f FormErrorsImpl) GetRaw(x int) interface{} {
	return f.Get(x)
}

func (f FormErrorsImpl) Get(x int) types.FormError {
	return f[x]
}

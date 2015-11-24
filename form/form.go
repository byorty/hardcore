package form

import "github.com/byorty/hardcore/types"

type FormImpl map[string]types.Primitive

func New() types.Form {
	f := make(FormImpl)
	return f
}

func (f FormImpl) Add(primitive types.Primitive) types.Form {
	f[primitive.GetName()] = primitive
	return f
}

func (f FormImpl) Check(scope types.RequestScope) (bool, map[string]string) {
	errors := make(map[string]string, 0)
	request := scope.GetRequest()

	for _, primitive := range f {
		value := scope.GetPathParams().GetString(primitive.GetName())
		if len(value) == 0 {
			value = request.FormValue(primitive.GetName())
		}
		if len(value) == 0 {
			value = request.PostFormValue(primitive.GetName())
		}
		if !primitive.Import(value) {
			errors[primitive.GetName()] = primitive.GetError()
		}
	}

	return len(errors) > 0, errors
}

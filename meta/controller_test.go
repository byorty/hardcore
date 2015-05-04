package meta

import (
	"testing"
)

type SomeController struct {}

func TestController(t *testing.T) {
	code := `func NewSomeController() mux.ActionController {
    return new(SomeController)
}

func (t *SomeController) CallAction(action interface{}, scope mux.RequestScope) {
    action.(func(*SomeController, mux.RequestScope))(t, scope)
}
`
	builder := new(controllerBuilder)
	results := builder.Build(
		new(SomeController),
	)

	t.Log(code)
	t.Log(string(results[0].Bytes))
	if string(results[0].Bytes) != code {
		t.Fail()
	}
}

package api

import (
	"github.com/byorty/hardcore/form"
	"github.com/byorty/hardcore/form/prim"
	"github.com/byorty/hardcore/types"
	"github.com/byorty/hardcore/view"
)

func (t *Test) CallAction(action interface{}, scope types.RequestScope) {
	if callable, ok := action.(types.CallableAction); ok {
		callable.Call(t, scope)
	} else {
		action.(func(*Test, types.RequestScope))(t, scope)
	}
}

type TestView func(*Test, string) types.View

func (t TestView) Call(rawCtrl interface{}, scope types.RequestScope) {
	form := form.New()
	var text string
	textPrim := prim.String("text")
	textPrim.Required()
	textPrim.SetSource(types.GetPrimitiveSource)
	textPrim.Export(&text)
	form.Add(textPrim)

	var v types.View
	if form.Check(scope) {
		ctrl := rawCtrl.(*Test)
		v = t(ctrl, text)
	} else {
		handler, ok := rawCtrl.(types.FormErrorsHandler)
		if ok {
			v = handler.HandleFormErrors(form.GetErrors())
		} else {
			v = view.BadRequest()
		}
	}
	v.SetScope(scope)
	v.Render()
}

var (
	TestViewAction TestView = (*Test).View
)

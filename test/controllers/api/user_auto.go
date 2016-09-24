package api

import (
    "github.com/byorty/hardcore/types"
    "github.com/byorty/hardcore/form"
    "github.com/byorty/hardcore/form/prim"
    "github.com/byorty/hardcore/view"
    "github.com/byorty/hardcore/test/models"
)

func (u *User) CallAction(action interface{}, scope types.RequestScope) {
	if callable, ok := action.(types.CallableAction); ok {
		callable.Call(u, scope)
	} else {
		v := action.(func(*User, types.RequestScope) types.View)(u, scope)
		v.SetController(u)
		v.SetScope(scope)
		v.Render()
	}
}

type UserList func(*User, int) types.EncodeView

func (u UserList) Call(rawCtrl interface{}, scope types.RequestScope) {
	form := form.New()
	var page int
	pagePrim := prim.Int("page")
	pagePrim.SetSource(types.GetPrimitiveSource)
	pagePrim.Export(&page)
	form.Add(pagePrim)

	var v types.View
	if form.Check(scope) {
		ctrl := rawCtrl.(*User)
		v = u(ctrl, page)
		v.SetController(ctrl)
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

type UserView func(*User, *models.User) types.EncodeView

func (u UserView) Call(rawCtrl interface{}, scope types.RequestScope) {
	form := form.New()
	var user models.User
	userPrim := prim.Int64Model("user")
	userPrim.Required()
	userPrim.SetSource(types.PathPrimitiveSource)
	userPrim.Export(&user)
	form.Add(userPrim)

	var v types.View
	if form.Check(scope) {
		ctrl := rawCtrl.(*User)
		v = u(ctrl, &user)
		v.SetController(ctrl)
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
	UserListAction UserList = (*User).List
	UserViewAction UserView = (*User).View
)

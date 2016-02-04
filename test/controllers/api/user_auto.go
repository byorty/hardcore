package api

import (
    "github.com/byorty/hardcore/types"
    "github.com/byorty/hardcore/form"
    "github.com/byorty/hardcore/form/prim"
    "github.com/byorty/hardcore/test/models"
)

func (u *User) CallAction(action interface{}, scope types.RequestScope) {
	if callable, ok := action.(types.CallableAction); ok {
		callable.Call(u, scope)
	} else {
		action.(func(*User, types.RequestScope))(u, scope)
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

	var view types.View
	if form.Check(scope) {
		ctrl := rawCtrl.(*User)
		view = u(ctrl, page)
	} else {
//		handler, ok := u.(types.FormErrorsHandler)
//		if ok {
//			view = handler.HandleFormErrors(form.GetErrors())
//		} else {
			handler, ok := rawCtrl.(types.FormErrorsHandler)
			if ok {
				view = handler.HandleFormErrors(form.GetErrors())
			}
//		}
	}
	view.SetScope(scope)
	view.Render()
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

	var view types.View
	if form.Check(scope) {
		ctrl := rawCtrl.(*User)
		view = u(ctrl, &user)
	} else {
//		handler, ok := u.(types.FormErrorsHandler)
//		if ok {
//			view = handler.HandleFormErrors(form.GetErrors())
//		} else {
			handler, ok := rawCtrl.(types.FormErrorsHandler)
			if ok {
				view = handler.HandleFormErrors(form.GetErrors())
			}
//		}
	}
	view.SetScope(scope)
	view.Render()
}


var (
	UserListAction UserList = (*User).List
	UserViewAction UserView = (*User).View
)

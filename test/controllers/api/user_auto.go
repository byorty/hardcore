package api

import (
	"github.com/byorty/hardcore/form"
	"github.com/byorty/hardcore/form/prim"
	"github.com/byorty/hardcore/test/models"
	"github.com/byorty/hardcore/types"
	"github.com/byorty/hardcore/view"
)

func (u *User) CallAction(action interface{}, rs types.RequestScope) {
	if callable, ok := action.(types.CallableAction); ok {
		callable.Call(u, rs)
	} else {
		v := action.(func(*User, types.RequestScope) types.View)(u, rs)
		v.SetController(u)
		v.SetScope(rs)
		v.Render()
	}
}

type UserList func(*User, int) types.EncodeView

func (u UserList) Call(rawCtrl interface{}, rs types.RequestScope) {
	form := form.New()
	var page int
	pagePrim := prim.Int("page")
	pagePrim.SetSource(types.GetPrimitiveSource)
	pagePrim.Export(&page)
	form.Add(pagePrim)

	var v types.View
	if form.Check(rs) {
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
	v.SetScope(rs)
	v.Render()
}

type UserView func(*User, *models.User) types.EncodeView

func (u UserView) Call(rawCtrl interface{}, rs types.RequestScope) {
	form := form.New()
	var user models.User
	userPrim := prim.Int64Model("user")
	userPrim.Required()
	userPrim.SetSource(types.PathPrimitiveSource)
	userPrim.Export(&user)
	form.Add(userPrim)

	var v types.View
	if form.Check(rs) {
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
	v.SetScope(rs)
	v.Render()
}

var (
	UserListAction UserList = (*User).List
	UserViewAction UserView = (*User).View
)

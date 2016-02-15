package api

import (
    "github.com/byorty/hardcore/types"
    "github.com/byorty/hardcore/form"
    "github.com/byorty/hardcore/form/prim"
    "github.com/byorty/hardcore/view"
    "github.com/byorty/hardcore/test/models"
)

func (p *Post) CallAction(action interface{}, scope types.RequestScope) {
	if callable, ok := action.(types.CallableAction); ok {
		callable.Call(p, scope)
	} else {
		action.(func(*Post, types.RequestScope))(p, scope)
	}
}

type PostList func(*Post, int, string) types.EncodeView

func (p PostList) Call(rawCtrl interface{}, scope types.RequestScope) {
	form := form.New()
	var page int
	pagePrim := prim.Int("page")
	pagePrim.SetSource(types.GetPrimitiveSource)
	pagePrim.Export(&page)
	form.Add(pagePrim)
	var search string
	searchPrim := prim.String("search")
	searchPrim.SetSource(types.GetPrimitiveSource)
	searchPrim.Export(&search)
	form.Add(searchPrim)

	var v types.View
	if form.Check(scope) {
		ctrl := rawCtrl.(*Post)
		v = p(ctrl, page, search)
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

type PostView func(*Post, *models.Post) types.EncodeView

func (p PostView) Call(rawCtrl interface{}, scope types.RequestScope) {
	form := form.New()
	var post models.Post
	postPrim := prim.Int64Model("post")
	postPrim.Required()
	postPrim.SetSource(types.PathPrimitiveSource)
	postPrim.Export(&post)
	form.Add(postPrim)

	var v types.View
	if form.Check(scope) {
		ctrl := rawCtrl.(*Post)
		v = p(ctrl, &post)
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

type PostEdit func(*Post, types.Form, *models.Post, string, string) types.EncodeView

func (p PostEdit) Call(rawCtrl interface{}, scope types.RequestScope) {
	form := form.New()
	var post models.Post
	postPrim := prim.Int64Model("post")
	postPrim.Required()
	postPrim.SetSource(types.PathPrimitiveSource)
	postPrim.Export(&post)
	form.Add(postPrim)
	var name string
	namePrim := prim.String("name")
	namePrim.Required()
	namePrim.SetSource(types.PostPrimitiveSource)
	namePrim.Export(&name)
	form.Add(namePrim)
	var description string
	descriptionPrim := prim.String("description")
	descriptionPrim.Required()
	descriptionPrim.SetSource(types.PostPrimitiveSource)
	descriptionPrim.Export(&description)
	form.Add(descriptionPrim)

	var v types.View
	if form.Check(scope) {
		ctrl := rawCtrl.(*Post)
		v = p(ctrl, form, &post, name, description)
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
	PostListAction PostList = (*Post).List
	PostViewAction PostView = (*Post).View
	PostEditAction PostEdit = (*Post).Edit
)

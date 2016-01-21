package api

import (
    "github.com/byorty/hardcore/types"
    "github.com/byorty/hardcore/form"
    "github.com/byorty/hardcore/form/prim"
    "github.com/byorty/hardcore/test/models"
)

func (p *Post) CallAction(action interface{}, scope types.RequestScope) {
	if callable, ok := action.(types.CallableAction); ok {
		callable.Call(p, scope)
	} else {
		action.(func(*Post, types.RequestScope))(p, scope)
	}
}


type PostList func(*Post, int, string)

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

	form.Check(scope)

	ctrl := rawCtrl.(*Post)
	p(ctrl, page, search)
//	p(ctrl, NewPostListForm(ctrl.(*Post)))
}

type PostView func(*Post, *models.Post)

func (p PostView) Call(rawCtrl interface{}, scope types.RequestScope) {
	form := form.New()
	var post models.Post
	postPrim := prim.Int64Model("post")
	postPrim.Required()
	postPrim.SetSource(types.PathPrimitiveSource)
	postPrim.Export(&post)
	form.Add(postPrim)

	form.Check(scope)

	ctrl := rawCtrl.(*Post)
	p(ctrl, &post)
//	p(ctrl, NewPostViewForm(ctrl.(*Post)))
}

type PostEdit func(*Post, types.Form, *models.Post, string, string)

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

	form.Check(scope)

	ctrl := rawCtrl.(*Post)
	p(ctrl, form, &post, name, description)
//	p(ctrl, NewPostEditForm(ctrl.(*Post)))
}


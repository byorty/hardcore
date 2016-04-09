package view

import (
	"github.com/byorty/hardcore/types"
	"net/http"
)

type RedirectImpl struct {
	BaseImpl
	url string
}

func Redirect(url string) types.RedirectView {
	redirect := new(RedirectImpl)
	redirect.url = url
	redirect.status = http.StatusMovedPermanently
	return redirect
}

func (r *RedirectImpl) SetUrl(url string) types.RedirectView {
	r.url = url
	return r
}

func (r RedirectImpl) Render() {
	http.Redirect(r.scope.GetWriter(), r.scope.GetRequest(), r.url, r.status)
}

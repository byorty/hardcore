package view

import (
	"github.com/byorty/hardcore/types"
	"net/http"
)

type RedirectImpl struct {
	BaseImpl
	url    string
	status int
}

func Redirect(url string) types.RedirectView {
	redirect := new(RedirectImpl)
	redirect.url = url
	return redirect
}

func (r *RedirectImpl) SetUrl(url string) types.RedirectView {
	r.url = url
	return r
}

func (r *RedirectImpl) SetStatus(status int) types.RedirectView {
	r.status = status
	return r
}

func (r RedirectImpl) Render() {
	if r.status == 0 {
		r.status = http.StatusMovedPermanently
	}
	http.Redirect(r.scope.GetWriter(), r.scope.GetRequest(), r.url, r.status)
}

package view

import (
	"github.com/byorty/hardcore/types"
	"net/http"
)

var (
	badRequest = []byte("Bad Request")
)

type RawImpl struct {
	BaseImpl
	bytes []byte
}

func Raw(bytes []byte) types.RawView {
	return &RawImpl{
		bytes: bytes,
	}
}

func (r *RawImpl) SetStatus(status int) types.RawView {
	r.status = status
	return r
}

func (r *RawImpl) Render() {
	rw := r.scope.GetWriter()
	r.BaseImpl.Render()
	rw.Write(r.bytes)
}

func BadRequest() types.RawView {
	r := Raw(badRequest)
	r.SetStatus(http.StatusBadRequest)
	return r
}

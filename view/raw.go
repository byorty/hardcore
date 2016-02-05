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
	status int
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
	if r.status == 0 {
		r.status = http.StatusOK
	}
	rw.WriteHeader(r.status)
	rw.Write(r.bytes)
}

func BadRequest() types.RawView {
	return Raw(badRequest).SetStatus(http.StatusBadRequest)
}


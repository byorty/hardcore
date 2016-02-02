package view

import (
	"github.com/byorty/hardcore/types"
	"net/http"
	"github.com/byorty/hardcore/encoder"
)

type EncodeImpl struct {
	BaseImpl
	status int
	encoder types.Encoder
	exporter types.Exporter
}

func Encode(exporter types.Exporter) types.EncodeView {
	return &EncodeImpl{
		exporter: exporter,
	}
}

func (e *EncodeImpl) SetStatus(status int) types.EncodeView {
	e.status = status
	return e
}

func (e *EncodeImpl) SetEncoder(encoder types.Encoder) types.EncodeView {
	e.encoder = encoder
	return e
}

func (e EncodeImpl) Render() {
	req := e.scope.GetRequest()
	rw := e.scope.GetWriter()

	if e.encoder == nil {
		switch req.Header.Get("Content-Type") {
		case "application/xml", "text/xml":
		case "application/json": e.encoder = encoder.NewJson()
		default: e.encoder = encoder.NewJson()
		}
	}
	if e.status == 0 {
		e.status = http.StatusOK
	}
	rw.WriteHeader(e.status)
	rw.Write(e.encoder.Encode(e.exporter))
}
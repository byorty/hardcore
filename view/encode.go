package view

import (
	"github.com/byorty/hardcore/encoder"
	"github.com/byorty/hardcore/is"
	"github.com/byorty/hardcore/types"
)

type EncodeImpl struct {
	BaseImpl
	encoder  types.Encoder
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

	if is.Nil(e.encoder) {
		switch req.Header.Get("Content-Type") {
		case "application/xml", "text/xml":
		case "application/json":
			e.encoder = encoder.NewJson()
		default:
			e.encoder = encoder.NewJson()
		}
	}
	//	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	e.BaseImpl.Render()
	if e.exporter.GetProtoKind().IsSlice() {
		rw.Write(e.encoder.All(e.exporter))
	} else {
		rw.Write(e.encoder.One(e.exporter))
	}
}

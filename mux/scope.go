package mux

import (
	"net/http"
	"github.com/byorty/hardcore/types"
)

type RequestScopeImpl struct {
	writer       http.ResponseWriter
	request      *http.Request
	pathParams   types.RequestScopeParams
	headerParams types.RequestScopeParams
}

func NewRequestScope() types.RequestScope {
	return new(RequestScopeImpl)
}

func (r RequestScopeImpl) GetWriter() http.ResponseWriter {
	return r.writer
}

func (r *RequestScopeImpl) SetWriter(writer http.ResponseWriter) {
	r.writer = writer
}

func (r RequestScopeImpl) GetRequest() *http.Request {
	return r.request
}

func (r *RequestScopeImpl) SetRequest(request *http.Request) {
	r.request = request
}

func (r RequestScopeImpl) GetPathParams() types.RequestScopeParams {
	return r.pathParams
}

func (r *RequestScopeImpl) SetPathParams(params types.RequestScopeParams) {
	r.pathParams = params
}

func (r RequestScopeImpl) GetHeaderParams() types.RequestScopeParams {
	return r.headerParams
}

func (r *RequestScopeImpl) SetHeaderParams(params types.RequestScopeParams) {
	r.headerParams = params
}

func (r *RequestScopeImpl) SetHeaderParam(key, value string) {
	r.headerParams.Set(key, value)
}
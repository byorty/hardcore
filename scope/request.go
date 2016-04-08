package scope

import (
	"github.com/byorty/hardcore/types"
	"net/http"
)

type RequestImpl struct {
	writer       http.ResponseWriter
	request      *http.Request
	pathParams   types.RequestScopeParams
	headerParams types.RequestScopeParams
	session      types.SessionScope
	isPrevent    bool
}

func NewRequest() types.RequestScope {
	return new(RequestImpl)
}

func (r RequestImpl) GetWriter() http.ResponseWriter {
	return r.writer
}

func (r *RequestImpl) SetWriter(writer http.ResponseWriter) {
	r.writer = writer
}

func (r RequestImpl) GetRequest() *http.Request {
	return r.request
}

func (r *RequestImpl) SetRequest(request *http.Request) {
	r.request = request
}

func (r RequestImpl) GetPathParams() types.RequestScopeParams {
	return r.pathParams
}

func (r *RequestImpl) SetPathParams(params types.RequestScopeParams) {
	r.pathParams = params
}

func (r RequestImpl) GetHeaderParams() types.RequestScopeParams {
	return r.headerParams
}

func (r *RequestImpl) SetHeaderParams(params types.RequestScopeParams) {
	r.headerParams = params
}

func (r *RequestImpl) SetHeaderParam(key, value string) {
	r.headerParams.Set(key, value)
}

func (r RequestImpl) GetSession() types.SessionScope {
	return r.session
}

func (r *RequestImpl) SetSession(session types.SessionScope) {
	r.session = session
}

func (r RequestImpl) Verify(form types.Form, primitive types.Primitive) {
	var value string

	switch primitive.GetSource() {
	case types.PathPrimitiveSource:
		value = r.GetPathParams().GetString(primitive.GetName())
	case types.PostPrimitiveSource:
		value = r.request.PostFormValue(primitive.GetName())
	case types.GetPrimitiveSource:
		value = r.request.FormValue(primitive.GetName())
	}

	if !primitive.Import(value) {
		form.AddErrorMessage(primitive.GetName(), primitive.GetError())
	}
}

func (r *RequestImpl) Prevent() {
	r.isPrevent = true
}

func (r RequestImpl) IsPrevented() bool {
	return r.isPrevent
}

func (r RequestImpl) NotPrevented() bool {
	return !r.isPrevent
}

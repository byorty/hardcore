package controller

import "github.com/byorty/hardcore/meta/types"

type BaseInjection struct {
	param types.ActionParam
}

func (b *BaseInjection) SetParam(param types.ActionParam) {
	b.param = param
}

func (b BaseInjection) GetImports() []string {
	return []string{}
}

func (b BaseInjection) GetAutoImports() []string {
	return []string{}
}

func (b BaseInjection) GetBody() string {
	return ""
}

type RequestScopeInjection struct {
	BaseInjection
}

func NewRequestScopeInjection() types.Injection {
	var inject RequestScopeInjection
	return &inject
}

func (r *RequestScopeInjection) IsMustWrite() bool {
	return false
}

func (r *RequestScopeInjection) GetKind() string {
	return "types.RequestScope"
}

type FormInjection struct {
	BaseInjection
}

func NewFormInjection() types.Injection {
	var inject FormInjection
	return &inject
}

func (f *FormInjection) IsMustWrite() bool {
	return false
}

func (f *FormInjection) GetKind() string {
	return "types.Form"
}

type PaginatorInjection struct {
	BaseInjection
}

func NewPaginatorInjection() types.Injection {
	var inject PaginatorInjection
	return &inject
}

func (p *PaginatorInjection) IsMustWrite() bool {
	return true
}

func (p *PaginatorInjection) GetKind() string {
	return "types.Paginator"
}

func (p *PaginatorInjection) GetBody() string {
	return `paginator := helper.NewPaginator()
	paginatorPrim := prim.Paginator("page")
	paginatorPrim.SetSource(types.GetPrimitiveSource)
	paginatorPrim.Export(paginator)
	form.Add(paginatorPrim)`
}

func (p PaginatorInjection) GetAutoImports() []string {
	return []string{
		types.HelperImport,
	}
}

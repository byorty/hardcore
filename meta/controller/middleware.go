package controller

import (
	"github.com/byorty/hardcore/meta/common"
	"github.com/byorty/hardcore/meta/types"
)

type Middleware struct {
	common.Entity
	entity types.Entity
}

func (m Middleware) GetEntity() types.Entity {
	return m.entity
}

func (m *Middleware) SetEntity(entity types.Entity) {
	m.entity = entity
}

type BeforeMiddleware struct {
	Middleware
}

func (b BeforeMiddleware) GetEntityKind() types.EntityKind {
	return types.BeforeMiddlewareEntityKind
}

type BeforeMiddlewares []*BeforeMiddleware

func (b BeforeMiddlewares) Len() int {
	return len(b)
}

func (b BeforeMiddlewares) Get(i int) types.Entity {
	return b[i]
}

type AfterMiddleware struct {
	Middleware
}

func (a AfterMiddleware) GetEntityKind() types.EntityKind {
	return types.AfterMiddlewareEntityKind
}

type AfterMiddlewares []*AfterMiddleware

func (a AfterMiddlewares) Len() int {
	return len(a)
}

func (a AfterMiddlewares) Get(i int) types.Entity {
	return a[i]
}

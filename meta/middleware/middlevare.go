package middleware

import (
	"github.com/byorty/hardcore/meta/common"
	"github.com/byorty/hardcore/meta/types"
)

type Middleware struct {
	common.Entity
}

func (m Middleware) GetEntityKind() types.EntityKind {
	return types.MiddlewareEntityKind
}

type Middlewares []*Middleware

func (m Middlewares) Len() int {
	return len(m)
}

func (m Middlewares) Get(i int) types.Entity {
	return m[i]
}

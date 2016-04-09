package controller

import "github.com/byorty/hardcore/meta/types"

type Initializer struct {
	Befores     BeforeMiddlewares `xml:"befores>before"`
	Afters      AfterMiddlewares  `xml:"afters>afters"`
	middlewares []types.Middleware
}

func (i *Initializer) PostInit(env types.Environment) {
	i.initMiddlewares(env, i.Befores)
	i.initMiddlewares(env, i.Afters)
}

func (i *Initializer) initMiddlewares(env types.Environment, slice types.EntitySlice) {
	for i := 0; i < slice.Len(); i++ {
		middleware := slice.Get(i).(types.Middleware)
		entity := env.GetConfiguration().GetEntity(middleware.GetName())
		if entity == nil {
			env.GetLogger().Error(
				"can't find middleware %s",
				middleware.GetName(),
			)
		} else {
			middleware.SetEntity(entity)
		}
	}
}

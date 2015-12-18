package plugin

import "github.com/byorty/hardcore/meta"

type Model struct {}

func (e *Model) Do(env *meta.Environment) {
    for _, container := range env.Configuration.ModelContainers {
        for _, model := range container.Models {

            env.Logger.Info(model)
        }
    }
}

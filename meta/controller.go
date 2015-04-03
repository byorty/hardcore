package meta

var ControllerTmpl = `

`

type controllerBuilder struct {}

func (e *controllerBuilder) Build(controllers ...interface{}) []*BuildResult {
	results := make([]*BuildResult, 0)

	return results
}

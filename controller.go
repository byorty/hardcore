package hardcore

type Controller interface {
	CallAction(interface{})
}

type ControllerConstructor func() Controller

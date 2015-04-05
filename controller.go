package hardcore

type Controller interface {
	CallAction(interface{})
}

type ControllerFunc func() Controller

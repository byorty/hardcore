package types

type Controller interface {
	GetRoute() string
	GetActions() []Action
}

type Action interface {
	GetName() string
	GetRoute() string
	GetMethod() string
	GetParameters() []ActionParameter
}

type ActionParameter interface {

}

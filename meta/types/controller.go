package types

type ControllerEntity interface {
	Entity
	GetRoute() string
	GetActions() []Action
	SetActions([]Action)
}

type Action interface {
	GetName() string
	GetRoute() string
	GetRouteMethod() string
	GetParams() []ActionParam
	SetParams([]ActionParam)
	HasForm() bool
	GetDefineKinds() string
	GetDefineParams() string
	GetDefineVars() string
	GetReturn() string
}

type ActionParam interface {
	GetName() string
	IsRequired() bool
	GetSource() string
	GetKind() string
	GetEntity() Entity
	SetEntity(Entity)
	GetDefineKind() string
	GetDefineVarKind() string
	GetDefineVarName() string
	GetPrimitive() string
	IsInjection() bool
	GetInjection() Injection
}

type Injection interface {
	SetParam(ActionParam)
	IsMustWrite() bool
	GetKind() string
	GetBody() string
	GetImports() []string
	GetAutoImports() []string
}

type Middleware interface {
	Entity
	GetEntity() Entity
	SetEntity(Entity)
}

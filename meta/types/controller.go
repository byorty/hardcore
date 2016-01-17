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
	GetMethod() string
	GetParams() []ActionParam
	SetParams([]ActionParam)
	HasForm() bool
	GetDefineKinds() string
	GetDefineParams() string
	GetDefineVars() string
}

type ActionParam interface {
	GetName() string
	IsRequired() bool
	GetSource() string
	GetKind() string
	GetEntity() Entity
	SetEntity(entity Entity)
	GetDefineKind() string
	GetDefineVarKind() string
	GetDefineVarName() string
	GetPrimitive() string
	IsReserved() bool
}

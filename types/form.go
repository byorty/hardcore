package types

type Primitive interface {
	GetName() string
	GetError() string
	SetSource(PrimitiveSource)
	GetSource() PrimitiveSource
	Required()
	Missing(string)
	Wrong(string)
	Custom(int, string)
	Import(interface{}) bool
	Export(interface{})
}

type PrimitiveSource int

const (
	RoutePrimitiveSource PrimitiveSource = iota
	GetPrimitiveSource
	PostPrimitiveSource
)

type Form interface {
	AddError(FormError) Form
	AddErrorMessage(string, string) Form
	AddErrorMessageWithCode(string, string, int) Form
	Add(Primitive) Form
	Check(FormVerifiable) bool
	GetErrors() FormErrors
}

type FormError interface {
	GetName() string
	GetMessage() string
	GetCode() int
}

type FormErrors interface {
	Slice
	Get(int) FormError
}

type FormVerifiable interface {
	Verify(Form, Primitive)
}
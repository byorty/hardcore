package types

type View interface {
	SetController(ActionController)
	SetScope(RequestScope)
	Render()
}

type EncodeView interface {
	View
	SetStatus(int) EncodeView
	SetEncoder(Encoder) EncodeView
}

type TmplView interface {
	View
	Set(string, interface{}) TmplView
	SetTemplate(string) TmplView
}

type RawView interface {
	View
	SetStatus(int) RawView
}

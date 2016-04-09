package types

type View interface {
	SetController(ActionController) View
	SetScope(RequestScope) View
	Render()
}

type EncodeView interface {
	View
	SetStatus(int) EncodeView
	SetEncoder(Encoder) EncodeView
}

type TmplView interface {
	View
	SetStatus(int) TmplView
	Set(string, interface{}) TmplView
	SetTemplate(string) TmplView
}

type RawView interface {
	View
	SetStatus(int) RawView
}

type StringView interface {
	View
	SetStatus(int) StringView
}

type RedirectView interface {
	View
}

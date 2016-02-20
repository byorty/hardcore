package types

type View interface {
	SetScope(RequestScope)
	Render()
}

type EncodeView interface {
	View
	SetStatus(int) EncodeView
	SetEncoder(Encoder) EncodeView
}

type HtmlView interface {
	View
	SetTemplate(string)
}

type RawView interface {
	View
	SetStatus(int) RawView
}
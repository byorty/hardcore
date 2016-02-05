package types

type Server interface {
	SetEnv(Environment)
	Serve()
}

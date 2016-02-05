package types

type Logger interface {
	Close()
	Finest(interface{}, ...interface{})
	Fine(interface{}, ...interface{})
	Debug(interface{}, ...interface{})
	Trace(interface{}, ...interface{})
	Info(interface{}, ...interface{})
	Warn(interface{}, ...interface{}) error
	Error(interface{}, ...interface{}) error
	Critical(interface{}, ...interface{}) error
}

package types

import "time"

type Environment interface {
	GetProjectName() string
	SetProjectName(string) Environment
	GetRouter() Router
	SetRouter(Router) Environment
	GetHostname() string
	SetHostname(string) Environment
	GetPort() int
	SetPort(int) Environment
	GetReadTimeout() time.Duration
	SetReadTimeout(time.Duration) Environment
	GetWriteTimeout() time.Duration
	SetWriteTimeout(time.Duration) Environment
	GetRootPath() string
	SetRootPath(string) Environment
	GetLogger() Logger
	SetLogger(Logger) Environment
}

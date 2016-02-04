package server

import (
	"github.com/byorty/hardcore/types"
	"time"
	"github.com/byorty/hardcore/utils"
)

var environment types.Environment

type EnvironmentImpl struct {
	projectName string
	router types.Router
	hostname string
	port int
	readTimeout time.Duration
	writeTimeout time.Duration
	rootPath string
}

func Environment() types.Environment {
	if environment == nil {
		environment = NewEnvironment()
	}
	return environment
}

func NewEnvironment() types.Environment {
	return &EnvironmentImpl{
		hostname: "localhost",
		port: 8080,
		readTimeout: 10 * time.Second,
		writeTimeout: 10 * time.Second,
		rootPath: utils.Pwd(),
	}
}

func (e EnvironmentImpl) GetProjectName() string {
	return e.projectName
}

func (e *EnvironmentImpl) SetProjectName(projectName string) types.Environment {
	e.projectName = projectName
	return e
}

func (e EnvironmentImpl) GetRouter() types.Router {
	return e.router
}

func (e *EnvironmentImpl) SetRouter(router types.Router) types.Environment {
	e.router = router
	return e
}

func (e EnvironmentImpl) GetHostname() string {
	return e.hostname
}

func (e *EnvironmentImpl) SetHostname(hostname string) types.Environment {
	e.hostname = hostname
	return e
}

func (e EnvironmentImpl) GetPort() int {
	return e.port
}

func (e *EnvironmentImpl) SetPort(port int) types.Environment {
	e.port = port
	return e
}

func (e EnvironmentImpl) GetReadTimeout() time.Duration {
	return e.readTimeout
}

func (e *EnvironmentImpl) SetReadTimeout(readTimeout time.Duration) types.Environment {
	e.readTimeout = readTimeout
	return e
}

func (e EnvironmentImpl) GetWriteTimeout() time.Duration {
	return e.writeTimeout
}

func (e *EnvironmentImpl) SetWriteTimeout(writeTimeout time.Duration) types.Environment {
	e.writeTimeout = writeTimeout
	return e
}

func (e EnvironmentImpl) GetRootPath() string {
	return e.rootPath
}

func (e *EnvironmentImpl) SetRootPath(rootPath string) types.Environment {
	e.rootPath = rootPath
	return e
}
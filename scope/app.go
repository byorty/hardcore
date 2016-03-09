package scope

import (
	"github.com/byorty/hardcore/log"
	"github.com/byorty/hardcore/types"
	"github.com/byorty/hardcore/utils"
	"time"
)

var app types.ApplicationScope

type AppImpl struct {
	projectName  string
	router       types.Router
	hostname     string
	port         int
	readTimeout  time.Duration
	writeTimeout time.Duration
	rootPath     string
	logger       types.Logger
	daos         []types.ModelDAO
}

func New() types.ApplicationScope {
	return &AppImpl{
		hostname:     "localhost",
		port:         8080,
		readTimeout:  10 * time.Second,
		writeTimeout: 10 * time.Second,
		rootPath:     utils.Pwd(),
		logger:       log.NewDefaultLogger(log.FINEST),
	}
}

func (a AppImpl) GetProjectName() string {
	return a.projectName
}

func (a *AppImpl) SetProjectName(projectName string) types.ApplicationScope {
	a.projectName = projectName
	return a
}

func (a AppImpl) GetRouter() types.Router {
	return a.router
}

func (a *AppImpl) SetRouter(router types.Router) types.ApplicationScope {
	a.router = router
	return a
}

func (a AppImpl) GetHostname() string {
	return a.hostname
}

func (a *AppImpl) SetHostname(hostname string) types.ApplicationScope {
	a.hostname = hostname
	return a
}

func (a AppImpl) GetPort() int {
	return a.port
}

func (a *AppImpl) SetPort(port int) types.ApplicationScope {
	a.port = port
	return a
}

func (a AppImpl) GetReadTimeout() time.Duration {
	return a.readTimeout
}

func (a *AppImpl) SetReadTimeout(readTimeout time.Duration) types.ApplicationScope {
	a.readTimeout = readTimeout
	return a
}

func (a AppImpl) GetWriteTimeout() time.Duration {
	return a.writeTimeout
}

func (a *AppImpl) SetWriteTimeout(writeTimeout time.Duration) types.ApplicationScope {
	a.writeTimeout = writeTimeout
	return a
}

func (a AppImpl) GetRootPath() string {
	return a.rootPath
}

func (a *AppImpl) SetRootPath(rootPath string) types.ApplicationScope {
	a.rootPath = rootPath
	return a
}

func (a AppImpl) GetLogger() types.Logger {
	return a.logger
}

func (a *AppImpl) SetLogger(logger types.Logger) types.ApplicationScope {
	a.logger = logger
	return a
}

func (a AppImpl) GetDAOs() []types.ModelDAO {
	return a.daos
}

func (a *AppImpl) SetDAOs(daos []types.ModelDAO) types.ApplicationScope {
	a.daos = daos
	return a
}

func App() types.ApplicationScope {
	if app == nil {
		app = New()
	}
	return app
}

func Set(a types.ApplicationScope) {
	app = a
}

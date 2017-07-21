package scope

import (
	"crypto/tls"
	"github.com/byorty/hardcore/is"
	"github.com/byorty/hardcore/log"
	"github.com/byorty/hardcore/types"
	"github.com/byorty/hardcore/utils"
	"github.com/gorilla/websocket"
	"html/template"
	"time"
	"context"
)

var app types.ApplicationScope

type AppImpl struct {
	projectName        string
	router             types.Router
	hostname           string
	port               int
	securityPort       int
	readTimeout        time.Duration
	writeTimeout       time.Duration
	rootPath           string
	tmplPath           string
	tmplCache          map[string]*template.Template
	logger             types.Logger
	daos               []types.ModelDAO
	cookieName         string
	enableSession      bool
	sessionProvider    types.Cache
	routes             []types.Route
	staticDir          string
	staticPath         string
	startDelim         string
	endDelim           string
	enableWebsocket    bool
	exit               chan bool
	certFilename       string
	privateKeyFilename string
	tlsConfig          *tls.Config
	upgrader           *websocket.Upgrader
	ctx                context.Context
}

func New() types.ApplicationScope {
	return &AppImpl{
		hostname:     "localhost",
		port:         8080,
		securityPort: 4443,
		readTimeout:  10 * time.Second,
		writeTimeout: 10 * time.Second,
		rootPath:     utils.Pwd(),
		logger:       log.NewDefaultLogger(log.FINEST),
		startDelim:   "{{",
		endDelim:     "}}",
		exit:         make(chan bool),
		ctx:          context.Background(),
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

func (a AppImpl) GetTmplPath() string {
	return a.tmplPath
}

func (a *AppImpl) SetTmplPath(tmplPath string) types.ApplicationScope {
	a.tmplPath = tmplPath
	return a
}

func (a AppImpl) GetTmplCache() map[string]*template.Template {
	return a.tmplCache
}

func (a *AppImpl) SetTmplCache(tmplCache map[string]*template.Template) types.ApplicationScope {
	a.tmplCache = tmplCache
	return a
}

func (a AppImpl) GetCookieName() string {
	return a.cookieName
}

func (a *AppImpl) SetCookieName(cookieName string) types.ApplicationScope {
	a.cookieName = cookieName
	return a
}

func (a AppImpl) GetEnableSession() bool {
	return a.enableSession
}

func (a *AppImpl) SetEnableSession(enableSession bool) types.ApplicationScope {
	a.enableSession = enableSession
	return a
}

func (a AppImpl) GetSessionProvider() types.Cache {
	return a.sessionProvider
}

func (a *AppImpl) SetSessionProvider(sessionProvider types.Cache) types.ApplicationScope {
	a.sessionProvider = sessionProvider
	return a
}

func (a AppImpl) GetRoutes() []types.Route {
	return a.routes
}

func (a *AppImpl) SetRoutes(routes []types.Route) types.ApplicationScope {
	a.routes = routes
	return a
}

func (a AppImpl) GetStaticDir() string {
	return a.staticDir
}

func (a *AppImpl) SetStaticDir(staticDir string) types.ApplicationScope {
	a.staticDir = staticDir
	return a
}

func (a AppImpl) GetStaticPath() string {
	return a.staticPath
}

func (a *AppImpl) SetStaticPath(staticPath string) types.ApplicationScope {
	a.staticPath = staticPath
	return a
}

func (a AppImpl) GetTmplDelims() (string, string) {
	return a.startDelim, a.endDelim
}

func (a *AppImpl) SetTmplDelims(startDelim, endDelim string) types.ApplicationScope {
	a.startDelim = startDelim
	a.endDelim = endDelim
	return a
}

func (a AppImpl) GetEnableWebsocket() bool {
	return a.enableWebsocket
}

func (a *AppImpl) SetEnableWebsocket(enableWebsocket bool) types.ApplicationScope {
	a.enableWebsocket = true
	return a
}

func (a AppImpl) IsExit() chan bool {
	return a.exit
}

func (a *AppImpl) Exit() {
	time.Sleep(2 * time.Second)
	go func() {
		a.exit <- true
	}()
}

func (a AppImpl) GetCertFilename() string {
	return a.certFilename
}

func (a *AppImpl) SetCertFilename(certFilename string) types.ApplicationScope {
	a.certFilename = certFilename
	return a
}

func (a AppImpl) GetPrivateKeyFilename() string {
	return a.privateKeyFilename
}

func (a *AppImpl) SetPrivateKeyFilename(privateKeyFilename string) types.ApplicationScope {
	a.privateKeyFilename = privateKeyFilename
	return a
}

func (a AppImpl) GetTlsConfig() *tls.Config {
	return a.tlsConfig
}

func (a *AppImpl) SetTlsConfig(tlsConfig *tls.Config) types.ApplicationScope {
	a.tlsConfig = tlsConfig
	return a
}

func (a AppImpl) GetSecurePort() int {
	return a.securityPort
}

func (a *AppImpl) SetSecurityPort(securityPort int) types.ApplicationScope {
	a.securityPort = securityPort
	return a
}

func (a AppImpl) GetUpgrader() *websocket.Upgrader {
	return a.upgrader
}

func (a *AppImpl) SetUpgrader(upgrader *websocket.Upgrader) types.ApplicationScope {
	a.upgrader = upgrader
	return a
}

func (a AppImpl) GetContext() context.Context {
	return a.ctx
}

func App() types.ApplicationScope {
	if is.Nil(app) {
		app = New()
	}
	return app
}

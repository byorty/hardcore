package helper

import (
	"crypto/rand"
	"github.com/byorty/hardcore/scope"
	"github.com/byorty/hardcore/types"
	"github.com/byorty/hardcore/utils"
	"io"
	"net/http"
	"sync"
	"time"
)

var sessionManager types.SessionManager

func SessionManager() types.SessionManager {
	if sessionManager == nil {
		sessionManager = &SessionManagerImpl{
			mtx:    new(sync.Mutex),
			maxAge: 1209600,
		}
	}
	return sessionManager
}

type SessionManagerImpl struct {
	mtx      *sync.Mutex
	provider types.Cache
	maxAge   int
}

func (s *SessionManagerImpl) Create(rs types.RequestScope) types.SessionScope {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	buf := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, buf); err != nil {
		scope.App().GetLogger().Error("session - can't make session id")
	}
	session := scope.NewSession(utils.Md5(string(buf) + time.Now().String()))
	s.provider.Set(session.GetId(), session)
	cookie := http.Cookie{
		Name:   scope.App().GetCookieName(),
		Value:  session.GetId(),
		Domain: scope.App().GetHostname(),
		Path:   "/",
		MaxAge: s.maxAge,
	}
	http.SetCookie(rs.GetWriter(), &cookie)
	return session
}

func (s *SessionManagerImpl) Get(rs types.RequestScope) types.SessionScope {
	id, err := rs.GetRequest().Cookie(scope.App().GetCookieName())
	if err == nil && id.Value != "" {
		session := s.provider.Get(id.Value)
		if session == nil {
			return s.Create(rs)
		} else {
			return session.(types.SessionScope)
		}
	} else {
		return s.Create(rs)
	}
}

func (s *SessionManagerImpl) Remove(rs types.RequestScope) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	id, err := rs.GetRequest().Cookie(scope.App().GetCookieName())
	if err == nil && id.Value != "" {
		s.provider.Remove(id.Value)
	}
	cookie := http.Cookie{
		Name:     scope.App().GetCookieName(),
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now(),
		MaxAge:   -1,
	}
	http.SetCookie(rs.GetWriter(), &cookie)
}

func (s *SessionManagerImpl) SetProvider(cache types.Cache) types.SessionManager {
	s.provider = cache
	return s
}

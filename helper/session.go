package helper

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/byorty/hardcore/scope"
	"github.com/byorty/hardcore/types"
	"io"
	"net/http"
	"sync"
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
	mtx    *sync.Mutex
	cache  types.Cache
	maxAge int
}

func (s *SessionManagerImpl) Create(rs types.RequestScope) types.SessionScope {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	buf := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, buf); err != nil {
		scope.App().GetLogger().Error("session - can't make session id")
	}
	id := base64.URLEncoding.EncodeToString(buf)

	s.cache.Set(id)
	cookie := &http.Cookie{
		Name:   scope.App().GetCookieName(),
		Value:  id,
		Domain: scope.App().GetHostname(),
		Path:   "/",
		MaxAge: s.maxAge,
	}
	http.SetCookie(rs.GetWriter(), cookie)
}

func (s *SessionManagerImpl) Get(rs types.RequestScope) types.SessionScope {

}

func (s *SessionManagerImpl) Remove(rs types.RequestScope) {

}

func (s *SessionManagerImpl) SetStorage(cache types.Cache) types.SessionManager {
	s.cache = cache
	return s
}

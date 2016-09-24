package plugin

import (
	"github.com/byorty/hardcore/helper"
	"github.com/byorty/hardcore/scope"
	"github.com/byorty/hardcore/types"
)

type SessionImpl struct{}

func NewSession() types.ApplicationPlugin {
	return new(SessionImpl)
}

func (s *SessionImpl) Run() {
	if scope.App().GetSessionProvider() != nil {
		helper.SessionManager().SetProvider(scope.App().GetSessionProvider())
	}
}

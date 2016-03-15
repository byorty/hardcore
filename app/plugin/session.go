package plugin

import (
	"github.com/byorty/hardcore/helper"
	"github.com/byorty/hardcore/scope"
)

type SessionImpl struct{}

func (s *SessionImpl) Run() {
	if scope.App().GetSessionProvider() != nil {
		helper.SessionManager().SetProvider(scope.App().GetSessionProvider())
	}
}

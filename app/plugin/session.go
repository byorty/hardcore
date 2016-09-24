package plugin

import (
	"github.com/byorty/hardcore/helper"
	"github.com/byorty/hardcore/scope"
	"github.com/byorty/hardcore/types"
)

func NewSession() types.ApplicationPlugin {
	return NewSessionByName(scope.DefaultName)
}

func NewSessionByName(name string) types.ApplicationPlugin {
	return newByName(new(SessionImpl), name)
}

type SessionImpl struct{
	BaseImpl
}

func (s *SessionImpl) Run() {
	if scope.AppByName(s.name).GetSessionProvider() != nil {
		helper.SessionManager().SetProvider(scope.AppByName(s.name).GetSessionProvider())
	}
}

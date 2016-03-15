package scope

import "github.com/byorty/hardcore/types"

type SessionImpl struct {
	id   string
	data map[string]interface{}
}

func NewSession(id string) types.SessionScope {
	return &SessionImpl{
		id:   id,
		data: make(map[string]interface{}),
	}
}

func (s SessionImpl) Get(key string) interface{} {
	if value, ok := s.get(key); ok {
		return value
	} else {
		return nil
	}
}

func (s SessionImpl) get(key string) (interface{}, bool) {
	value, ok := s.data[key]
	return value, ok
}

func (s SessionImpl) GetId() string {
	return s.id
}

func (s SessionImpl) Has(key string) bool {
	_, ok := s.get(key)
	return ok
}

func (s *SessionImpl) Remove(key string) types.SessionScope {
	delete(s.data, key)
	return s
}

func (s *SessionImpl) Set(key string, value interface{}) types.SessionScope {
	s.data[key] = value
	return s
}

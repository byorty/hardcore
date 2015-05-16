package proto

import "github.com/byorty/hardcore/types"

type protoMap map[string]types.ProtoProperty

func (m protoMap) HasKey(key string) bool {
	_, ok := m[key]
	return ok
}

func (m protoMap) Get(key string) types.ProtoProperty {
	if prop, ok := m[key]; ok {
		return prop
	} else {
		return nil
	}
}

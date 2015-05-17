package proto

import "github.com/byorty/hardcore/types"

type protoStruct struct {
	propMap protoMap
	props   protoSlice
}

func New() types.Proto {
	return &protoStruct{
		propMap: make(protoMap),
		props  : make(protoSlice, 0),
	}
}

func (p protoStruct) HasKey(key string) bool {
	return p.propMap.HasKey(key)
}

func (p protoStruct) HasNum(i int) bool {
	return p.props.HasNum(i)
}

func (p protoStruct) GetByName(key string) types.ProtoProperty {
	return p.propMap.Get(key)
}

func (p protoStruct) GetByNum(i int) types.ProtoProperty {
	return p.props.Get(i)
}

func (p *protoStruct) Set(name string, property types.ProtoProperty) types.Proto {
	p.propMap[name] = property
	p.props = append(p.props, property)
	return p
}

func (p *protoStruct) GetMap() map[string]types.ProtoProperty {
	return p.propMap
}

func (p *protoStruct) GetSlice() []types.ProtoProperty {
	return p.props
}

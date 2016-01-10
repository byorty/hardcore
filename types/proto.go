package types

type Prototyped interface {
	Proto() Proto
}

type Proto interface {
	HasKey(string) bool
	HasNum(int) bool
	GetByName(string) ProtoProperty
	GetByNum(int) ProtoProperty
	Set(string, ProtoProperty) Proto
	GetMap() map[string]ProtoProperty
	GetSlice() []ProtoProperty
}

type ProtoMap interface {
	HasKey(string) bool
	Get(string) ProtoProperty
}

type ProtoSlice interface {
	HasNum(int) bool
	Get(int) ProtoProperty
}

type ProtoKind int

const (
	ProtoBasicKind ProtoKind = iota
	ProtoIntKind
	ProtoInt8Kind
	ProtoInt16Kind
	ProtoInt32Kind
	ProtoInt64Kind
	ProtoUintKind
	ProtoUint8Kind
	ProtoUint16Kind
	ProtoUint32Kind
	ProtoUint64Kind
	ProtoFloat32Kind
	ProtoFloat64Kind
	ProtoStringKind
	ProtoTimeKind
	ProtoEnumKind
	ProtoModelKind
	ProtoSliceKind
)

type ProtoRelation int

const (
	ProtoNoneRelation       ProtoRelation = iota
	ProtoOneToOneRelation
	ProtoOneToManyRelation
	ProtoManyToManyRelation
)

type ProtoProperty interface {
	GetField() string
	GetKind() ProtoKind
	GetRelation() ProtoRelation
	GetProto() Proto
	IsRequired() bool
	GetSetter() Setter
	GetGetter() Getter
}

type Setter func(interface{}, interface{})
type Getter func(interface{}) interface{}

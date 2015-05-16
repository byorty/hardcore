package types

type Proto interface {
	HasKey(string) bool
	HasNum(int) bool
	GetByName(string) ProtoProperty
	GetByNum(int) ProtoProperty
	Set(string, ProtoProperty) Proto
	ToMap() ProtoMap
	ToSlice() ProtoSlice
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
	ProtoEnumKind
	ProtoModelKind
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
}

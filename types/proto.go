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

type ProtoEntity interface {
	GetProtoKind() ProtoKind
}

type ProtoKind int

func (p ProtoKind) IsModel() bool {
	return p == ProtoModelKind
}

func (p ProtoKind) IsScalar() bool {
	return p <= ProtoBoolKind
}

func (p ProtoKind) IsSlice() bool {
	return p >= ProtoIntSliceKind && p <= ProtoModelSliceKind
}

func (p ProtoKind) IsNumber() bool {
	return p >= ProtoIntKind && p <= ProtoFloat64Kind
}

func (p ProtoKind) IsBool() bool {
	return p == ProtoBoolKind
}

func (p ProtoKind) IsString() bool {
	return p == ProtoStringKind
}

const (
	ProtoIntKind ProtoKind = iota
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
	ProtoBoolKind
	ProtoTimeKind
	ProtoEnumKind
	ProtoModelKind
	ProtoIntSliceKind
	ProtoInt8SliceKind
	ProtoInt16SliceKind
	ProtoInt32SliceKind
	ProtoInt64SliceKind
	ProtoUintSliceKind
	ProtoUint8SliceKind
	ProtoUint16SliceKind
	ProtoUint32SliceKind
	ProtoUint64SliceKind
	ProtoFloat32SliceKind
	ProtoFloat64SliceKind
	ProtoStringSliceKind
	ProtoBoolSliceKind
	ProtoTimeSliceKind
	ProtoEnumSliceKind
	ProtoModelSliceKind
	ProtoUnkownKind
)

type ProtoRelation int

const (
	ProtoNoneRelation ProtoRelation = iota
	ProtoOneToOneRelation
	ProtoOneToManyRelation
	ProtoManyToManyRelation
)

type ProtoProperty interface {
	GetName() string
	SetName(string)
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

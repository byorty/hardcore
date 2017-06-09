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

func (p ProtoKind) IsBase() bool {
	return p <= ProtoStringKind
}

func (p ProtoKind) IsSlice() bool {
	return p >= ProtoIntSliceKind && p <= ProtoModelSliceKind
}

func (p ProtoKind) IsModelSlice() bool {
	return p == ProtoModelSliceKind
}

func (p ProtoKind) IsEnum() bool {
	return p >= ProtoIntEnumKind && p <= ProtoStringEnumKind
}

func (p ProtoKind) IsNotStringEnum() bool {
	return p >= ProtoIntEnumKind && p <= ProtoBoolEnumKind
}

func (p ProtoKind) IsStringEnum() bool {
	return p == ProtoStringEnumKind
}

func (p ProtoKind) IsNumber() bool {
	return p >= ProtoIntKind && p <= ProtoFloat64Kind
}

func (p ProtoKind) IsNumberSlice() bool {
	return p >= ProtoIntSliceKind && p <= ProtoFloat64SliceKind
}

func (p ProtoKind) IsBool() bool {
	return p == ProtoBoolKind
}

func (p ProtoKind) IsBoolSlice() bool {
	return p == ProtoBoolSliceKind
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
	ProtoBoolKind
	ProtoByteKind
	ProtoStringKind
	ProtoTimeKind
	ProtoIntEnumKind
	ProtoInt8EnumKind
	ProtoInt16EnumKind
	ProtoInt32EnumKind
	ProtoInt64EnumKind
	ProtoUintEnumKind
	ProtoUint8EnumKind
	ProtoUint16EnumKind
	ProtoUint32EnumKind
	ProtoUint64EnumKind
	ProtoFloat32EnumKind
	ProtoFloat64EnumKind
	ProtoBoolEnumKind
	ProtoByteEnumKind
	ProtoStringEnumKind
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
	ProtoByteSliceKind
	ProtoStringSliceKind
	ProtoBoolSliceKind
	ProtoTimeSliceKind
	ProtoEnumSliceKind
	ProtoModelSliceKind
	ProtoUnknownKind
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

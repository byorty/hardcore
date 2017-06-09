package types

type ModelEntity interface {
	Entity
	GetPattern() Pattern
	GetSource() string
	SetProperties([]Property)
	GetProperties() []Property
	GetTable() string
}

type Property interface {
	GetName() string
	GetKind() string
	GetField() string
	IsRequired() bool
	GetRelation() Relation
	SetUpperName(string)
	GetUpperName() string
	SetEntity(Entity)
	GetEntity() Entity
	HasRelation() bool
	SetSelfPackage(bool)
	IsSelfPackage() bool
	GetDefineKind() string
	GetMethodDefineKind() string
	GetVariableKind() string
	GetProtoKind() string
	SetRelationProperty(Property)
	GetRelationProperty() Property
	IsIdentifier() bool
	NotIdentifier() bool
}

type Relation string

const (
	NoneRelation       Relation = ""
	OneToOneRelation            = "OneToOne"
	OneToManyRelation           = "OneToMany"
	ManyToManyRelation          = "ManyToMany"
)

func (r Relation) IsNone() bool {
	return r == NoneRelation
}

func (r Relation) IsOneToOne() bool {
	return r == OneToOneRelation
}

func (r Relation) IsOneToMany() bool {
	return r == OneToManyRelation
}

func (r Relation) IsManyToMany() bool {
	return r == ManyToManyRelation
}

func (r Relation) HasMany() bool {
	return r.IsOneToMany() || r.IsManyToMany()
}

func (r Relation) AsProtoRelation() string {
	return protoRelations[r]
}

var (
	protoRelations = map[Relation]string{
		NoneRelation:       "ProtoNoneRelation",
		OneToOneRelation:   "ProtoOneToOneRelation",
		OneToManyRelation:  "ProtoOneToManyRelation",
		ManyToManyRelation: "ProtoManyToManyRelation",
	}
	SliceTypes = map[string]string{
		"types.IntSlice":     "int",
		"types.Int8Slice":    "int8",
		"types.Int16Slice":   "int16",
		"types.Int32Slice":   "int32",
		"types.Int64Slice":   "int64",
		"types.UintSlice":    "uint",
		"types.Uint8Slice":   "uint8",
		"types.Uint16Slice":  "uint16",
		"types.Uint32Slice":  "uint32",
		"types.Uint64Slice":  "uint64",
		"types.Float32Slice": "float32",
		"types.Float64Slice": "float64",
		"types.StringSlice":  "string",
	}
	SliceProtoKinds = map[string]string{
		"types.IntSlice":     "ProtoIntSliceKind",
		"types.Int8Slice":    "ProtoInt8SliceKind",
		"types.Int16Slice":   "ProtoInt16SliceKind",
		"types.Int32Slice":   "ProtoInt32SliceKind",
		"types.Int64Slice":   "ProtoInt64SliceKind",
		"types.UintSlice":    "ProtoUintSliceKind",
		"types.Uint8Slice":   "ProtoUint8SliceKind",
		"types.Uint16Slice":  "ProtoUint16SliceKind",
		"types.Uint32Slice":  "ProtoUint32SliceKind",
		"types.Uint64Slice":  "ProtoUint64SliceKind",
		"types.Float32Slice": "ProtoFloat32SliceKind",
		"types.Float64Slice": "ProtoFloat64SliceKind",
		"types.StringSlice":  "ProtoStringSliceKind",
	}
)

type Pattern string

const (
	ValueObjectPattern     Pattern = "ValueObject"
	StraightMappingPattern Pattern = "StraightMapping"
)

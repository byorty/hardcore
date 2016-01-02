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
	GetVariableKind() string
	GetProtoKind() string
	SetRelationProperty(Property)
	GetRelationProperty() Property
}

type Relation string

const (
    NoneRelation Relation = ""
    OneToOneRelation = "OneToOne"
    OneToManyRelation = "OneToMany"
    ManyToManyRelation = "ManyToMany"
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
	protoRelations = map[Relation]string {
		NoneRelation: "ProtoNoneRelation",
		OneToOneRelation: "ProtoOneToOneRelation",
		OneToManyRelation: "ProtoOneToManyRelation",
		ManyToManyRelation: "ProtoManyToManyRelation",
	}
)

type Pattern string

const (
    ValueObjectPattern = "ValueObject"
    StraightMappingPattern = "StraightMapping"
)

package types

type ModelEntity interface {
    Entity
	GetPattern() Pattern
    GetSource() string
    SetProperties([]Property)
    GetProperties() []Property
}

type Property interface {
    GetName() string
    GetKind() string
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

type Pattern string

const (
    NilPattern Pattern = ""
    ValueObjectPattern = "ValueObject"
    StraightMappingPattern = "StraightMapping"
)

package model

type Relation string

const (
    NoneRelation Relation = ""
    OneToOneRelation = "oneToOne"
    OneToManyRelation = "oneToMany"
    ManyToManyRelation = "manyToMany"
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

func (r Relation) NeedMany() bool {
    return r.IsOneToMany() || r.IsManyToMany()
}
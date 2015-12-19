package model

type Relation string

const (
    NoneRelation Relation = ""
    OneToOneRelation = "oneToOne"
    OneToManyRelation = "oneToMany"
    ManyManyRelation = "manyToMany"
)
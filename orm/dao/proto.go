package dao

type RelationKind int

const (
	None       RelationKind = iota
	OneToOne
	OneToMany
	ManyToMany
)

type ProtoStruct struct {
	Dao *DAO
	Map ProtoMap
}

type ProtoMap map[string]ProtoProp

type ProtoProp struct {
	Name         string
	RelationKind RelationKind
	RelationProp *ProtoProp
}

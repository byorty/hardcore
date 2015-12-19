package model

type Pattern string

const (
    NilPattern Pattern = ""
    ValueObjectPattern = "ValueObject"
    StraightMappingPattern = "StraightMapping"
)
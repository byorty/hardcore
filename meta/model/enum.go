package model

import (
    "fmt"
    "github.com/byorty/hardcore/meta/common"
"github.com/byorty/hardcore/meta/types"
)

type EnumKind string

func (e EnumKind) IsValid() bool {
    isValid := false
    for _, kind := range enumKinds {
        if kind == e {
            isValid = true
            break
        }
    }
    return isValid
}

func (e EnumKind) DAO() string {
    return daoByKinds[e]
}

const (
    IntEnumKind EnumKind = "int"
    Int8EnumKind = "in8"
    Int16EnumKind = "int16"
    Int32EnumKind = "int32"
    Int64EnumKind = "int64"
    UintEnumKind = "uint"
    Uint8EnumKind = "uin8"
    Uint16EnumKind = "uint16"
    Uint32EnumKind = "uint32"
    Uint64EnumKind = "uint64"
    Float32EnumKind = "float32"
    Float64EnumKind = "float64"
    Complex64EnumKind = "complex64"
    Complex128EnumKind = "complex128"
    StringEnumKind = "string"
    RuneEnumKind = "rune"
)

var (
    enumKinds = []EnumKind{
        IntEnumKind,
        Int8EnumKind,
        Int16EnumKind,
        Int32EnumKind,
        Int64EnumKind,
        UintEnumKind,
        Uint8EnumKind,
        Uint16EnumKind,
        Uint32EnumKind,
        Uint64EnumKind,
        Float32EnumKind,
        Float64EnumKind,
        Complex64EnumKind,
        Complex128EnumKind,
        StringEnumKind,
        RuneEnumKind,
    }
    daoByKinds = map[EnumKind]string {
        IntEnumKind: "IntEnumDAO",
        Int8EnumKind: "Int8EnumDAO",
        Int16EnumKind: "Int16EnumDAO",
        Int32EnumKind: "Int32EnumDAO",
        Int64EnumKind: "Int64EnumDAO",
        UintEnumKind: "UintEnumDAO",
        Uint8EnumKind: "Uint8EnumDAO",
        Uint16EnumKind: "Uint16EnumDAO",
        Uint32EnumKind: "Uint32EnumDAO",
        Uint64EnumKind: "Uint64EnumDAO",
        Float32EnumKind: "Float32EnumDAO",
        Float64EnumKind: "Float64EnumDAO",
        Complex64EnumKind: "Complex64EnumDAO",
        Complex128EnumKind: "Complex128EnumDAO",
        StringEnumKind: "StringEnumDAO",
        RuneEnumKind: "RuneEnumDAO",
    }
)

func HasEnumKind(kind EnumKind) bool {
    hasKind := false
    for _, enumKind := range enumKinds {
        if enumKind == kind {
            hasKind = true
            break
        }
    }
    return hasKind
}

type Enum struct {
    common.Entity
    Kind EnumKind `xml:"type,attr"`
    Constants []*Constant `xml:"constants>constant"`
    Filename string
    AutoFilename string
}

func (e Enum) GetKind() types.EntityKind {
    return types.EnumEntityKind
}

func (e Enum) GetValue(constant *Constant) string {
    if e.Kind == StringEnumKind && constant.HasValue() {
        return fmt.Sprintf(`"%s"`, constant.Value)
    } else {
        return constant.Value
    }
}
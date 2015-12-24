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
        IntEnumKind: "types.IntEnumDAO",
        Int8EnumKind: "types.Int8EnumDAO",
        Int16EnumKind: "types.Int16EnumDAO",
        Int32EnumKind: "types.Int32EnumDAO",
        Int64EnumKind: "types.Int64EnumDAO",
        UintEnumKind: "types.UintEnumDAO",
        Uint8EnumKind: "types.Uint8EnumDAO",
        Uint16EnumKind: "types.Uint16EnumDAO",
        Uint32EnumKind: "types.Uint32EnumDAO",
        Uint64EnumKind: "types.Uint64EnumDAO",
        Float32EnumKind: "types.Float32EnumDAO",
        Float64EnumKind: "types.Float64EnumDAO",
        Complex64EnumKind: "types.Complex64EnumDAO",
        Complex128EnumKind: "types.Complex128EnumDAO",
        StringEnumKind: "types.StringEnumDAO",
        RuneEnumKind: "types.RuneEnumDAO",
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
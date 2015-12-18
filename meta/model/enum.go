package model

import (
    "fmt"
)

type EnumKind string

const (
    IntEnumKind EnumKind = "int"
    Int8EnumKind = "in8"
    Int16EnumKind = "int16"
    Int32EnumKind = "int32"
    Int64EnumKind = "int64"
    UintntEnumKind = "uint"
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
        UintntEnumKind,
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
    Name string `xml:"name,attr"`
    Kind EnumKind `xml:"type,attr"`
    Constants []*Constant `xml:"constants>constant"`
    Filename string
    AutoFilename string
}

func (e Enum) GetValue(constant *Constant) string {
    if e.Kind == StringEnumKind && constant.HasValue() {
        return fmt.Sprintf(`"%s"`, constant.Value)
    } else {
        return constant.Value
    }
}
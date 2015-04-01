package meta

import (
    "github.com/byorty/hardcore"
)
var (
    __autoEnumSomeStatusInst0 = &SomeStatus{Id: 1, Name: "one"}
    __autoEnumSomeStatusInst1 = &SomeStatus{Id: 2, Name: "two"}
    __autoEnumSomeStatusInst2 = &SomeStatus{Id: 3, Name: "three"}
    __autoSomeStatusMap = map[int]*SomeStatus{
        1: __autoEnumSomeStatusInst0,
        2: __autoEnumSomeStatusInst1,
        3: __autoEnumSomeStatusInst2,
    }
    __autoSomeStatusAsEnumMap = map[int]hardcore.Enum{
        1: __autoEnumSomeStatusInst0,
        2: __autoEnumSomeStatusInst1,
        3: __autoEnumSomeStatusInst2,
    }
    __autoSomeStatusSlice = []*SomeStatus{
        __autoEnumSomeStatusInst0,
        __autoEnumSomeStatusInst1,
        __autoEnumSomeStatusInst2,
    }
    __autoSomeStatusAsEnumSlice = []hardcore.Enum{
        __autoEnumSomeStatusInst0,
        __autoEnumSomeStatusInst1,
        __autoEnumSomeStatusInst2,
    }
)

func GetSomeStatusMap() map[int]*SomeStatus {
    return __autoSomeStatusMap
}

func GetSomeStatusAsEnumMap() map[int]hardcore.Enum {
    return __autoSomeStatusAsEnumMap
}

func GetSomeStatusSlice() []*SomeStatus {
    return __autoSomeStatusSlice
}

func GetSomeStatusAsEnumSlice() []hardcore.Enum {
    return __autoSomeStatusAsEnumSlice
}

func GetSomeStatusById(id int) *SomeStatus {
    switch id {
    case 1: return __autoEnumSomeStatusInst0
    case 2: return __autoEnumSomeStatusInst1
    case 3: return __autoEnumSomeStatusInst2
    default: return nil
    }
}

func GetSomeStatusAsEnumById(id int) hardcore.Enum {
    switch id {
    case 1: return __autoEnumSomeStatusInst0
    case 2: return __autoEnumSomeStatusInst1
    case 3: return __autoEnumSomeStatusInst2
    default: return nil
    }
}

var (
    __autoEnumSomeTypeInst0 = &SomeType{Id: FirstType, Name: "one"}
    __autoEnumSomeTypeInst1 = &SomeType{Id: SecondType, Name: "two"}
    __autoEnumSomeTypeInst2 = &SomeType{Id: ThreeType, Name: "three"}
    __autoSomeTypeMap = map[int]*SomeType{
        FirstType: __autoEnumSomeTypeInst0,
        SecondType: __autoEnumSomeTypeInst1,
        ThreeType: __autoEnumSomeTypeInst2,
    }
    __autoSomeTypeAsEnumMap = map[int]hardcore.Enum{
        FirstType: __autoEnumSomeTypeInst0,
        SecondType: __autoEnumSomeTypeInst1,
        ThreeType: __autoEnumSomeTypeInst2,
    }
    __autoSomeTypeSlice = []*SomeType{
        __autoEnumSomeTypeInst0,
        __autoEnumSomeTypeInst1,
        __autoEnumSomeTypeInst2,
    }
    __autoSomeTypeAsEnumSlice = []hardcore.Enum{
        __autoEnumSomeTypeInst0,
        __autoEnumSomeTypeInst1,
        __autoEnumSomeTypeInst2,
    }
)

func GetSomeTypeMap() map[int]*SomeType {
    return __autoSomeTypeMap
}

func GetSomeTypeAsEnumMap() map[int]hardcore.Enum {
    return __autoSomeTypeAsEnumMap
}

func GetSomeTypeSlice() []*SomeType {
    return __autoSomeTypeSlice
}

func GetSomeTypeAsEnumSlice() []hardcore.Enum {
    return __autoSomeTypeAsEnumSlice
}

func GetSomeTypeById(id int) *SomeType {
    switch id {
    case FirstType: return __autoEnumSomeTypeInst0
    case SecondType: return __autoEnumSomeTypeInst1
    case ThreeType: return __autoEnumSomeTypeInst2
    default: return nil
    }
}

func GetSomeTypeAsEnumById(id int) hardcore.Enum {
    switch id {
    case FirstType: return __autoEnumSomeTypeInst0
    case SecondType: return __autoEnumSomeTypeInst1
    case ThreeType: return __autoEnumSomeTypeInst2
    default: return nil
    }
}


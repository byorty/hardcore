package meta

import (
	"testing"
)

type SomeStatus struct {
	Id   int    `enum:"1,2,3"`
	Name string `enum:"one,two,three"`
}

func (s *SomeStatus) GetId() int {
	return s.Id
}

func (s *SomeStatus) GetName() string {
	return s.Name
}

const (
	FirstType int = 1 << iota
	SecondType
	ThreeType
)

type SomeType struct {
	Id   int    `enum:"FirstType, SecondType, ThreeType"`
	Name string `enum:"one, two, three"`
}

func (s *SomeType) GetId() int {
	return s.Id
}

func (s *SomeType) GetName() string {
	return s.Name
}

func TestEnum1(t *testing.T) {
	builder := new(enumBuilder)
	results := builder.Build(
		new(SomeStatus),
	)
	code := `var (
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

`
	t.Log(code)
	t.Log(string(results[0].Bytes))
	if string(results[0].Bytes) != code {
		t.Fail()
	}
}

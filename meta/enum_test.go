package meta

import (
	"testing"
)

type SomeStatus struct {
	Id   int    `enum:"1,2,3"`
	Name string `enum:"one,two,three"`
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

	code := `func NewSomeStatusById(id int) *SomeStatus {
    s := new(SomeStatus).ById(id)
    if s == nil {
        return nil
    } else {
        return s.(*SomeStatus)
    }
}

func (s *SomeStatus) GetId() int {
    return s.Id
}

func (s *SomeStatus) GetName() string {
    return s.Name
}

func (s *SomeStatus) ById(id int) hardcore.IntIdentifiable {
    switch id {
    case 1:
        s.Id = 1
        s.Name = "one"
        break
    case 2:
        s.Id = 2
        s.Name = "two"
        break
    case 3:
        s.Id = 3
        s.Name = "three"
        break
    default: return nil
    }
    return s
}

type SomeStatuses []*SomeStatus

func (s *SomeStatuses) All() {
    (*s) = append((*s), NewSomeStatusById(1))
    (*s) = append((*s), NewSomeStatusById(2))
    (*s) = append((*s), NewSomeStatusById(3))
}

func (s *SomeStatuses) ByIds(ids []int) {
    enumMap := NewSomeStatusMap()
    enumMap.All()
    for _, id := range ids {
        if enum, ok := enumMap[id]; ok {
            (*s) = append((*s), enum)
        }
    }
}

type SomeStatusMap map[int]*SomeStatus

func NewSomeStatusMap() SomeStatusMap {
    return make(SomeStatusMap)
}

func (s SomeStatusMap) All() {
    s[1] = NewSomeStatusById(1)
    s[2] = NewSomeStatusById(2)
    s[3] = NewSomeStatusById(3)
}

func (s SomeStatusMap) ByIds(ids []int) {
    enumMap := NewSomeStatusMap()
    enumMap.All()
    for _, id := range ids {
        if enum, ok := enumMap[id]; ok {
            s[id] = enum
        }
    }
}

`
	t.Log(code)
	t.Log(string(results[0].Bytes))
	if string(results[0].Bytes) != code {
		t.Fail()
	}
}

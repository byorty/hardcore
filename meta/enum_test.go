package meta

import (
	"testing"
	"fmt"
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

func (s *SomeStatus) ById(id int) types.IntIdentifiable {
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

func NewSomeStatuses() SomeStatuses {
    return make(SomeStatuses, 0)
}

func (s *SomeStatuses) All() {
    (*s) = append((*s), NewSomeStatusById(1))
    (*s) = append((*s), NewSomeStatusById(2))
    (*s) = append((*s), NewSomeStatusById(3))
}

func NewAllSomeStatuses() SomeStatuses {
    var s SomeStatuses
    s.All()
    return s
}

func (s *SomeStatuses) ByIds(ids []int) {
    for _, id := range ids {
        enum := NewSomeStatusById(id)
        if enum != nil {
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
    for _, id := range ids {
        enum := NewSomeStatusById(id)
        if enum != nil {
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

	fmt.Println(getValue(SomeEnum1).Kind())
	var se SomeEnum
	fmt.Println(getValue(&se).Kind())
}

type SomeEnum int

const (
	SomeEnum1 SomeEnum = iota
	SomeEnum2
	SomeEnum3
)

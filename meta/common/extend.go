package common

type Extend struct {
	Name   string `xml:",chardata"`
	Import string
}

func (e Extend) IsPointer() bool {
	return e.Name[:1] == "*"
}

func (e Extend) GetClearName() string {
	if e.IsPointer() {
		return e.Name[1:]
	} else {
		return e.Name
	}
}

package exporter

import (
	"fmt"
	"github.com/byorty/hardcore/utils"
)

type Property struct {
	Name      string `xml:"name,attr"`
	hasGetter bool
}

func (p Property) HasGetter() bool {
	return p.hasGetter
}

func (p *Property) SetHasGetter(hasGetter bool) {
	p.hasGetter = hasGetter
}

func (p Property) GetName() string {
	return p.Name
}

func (p Property) GetGetterName() string {
	return fmt.Sprintf("Get%s", utils.UpperFirst(p.Name))
}

package exporter

import (
	"fmt"
	"github.com/byorty/hardcore/utils"
	"github.com/byorty/hardcore/meta/types"
)

type Property struct {
	Name      string `xml:"name,attr"`
	Alias     string `xml:"as,attr"`
	hasGetter bool
	property  types.Property
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

func (p Property) GetAliasName() string {
	if p.Alias != "" {
		return p.Alias
	}
	return p.Name
}

func (p Property) GetGetterName() string {
	return fmt.Sprintf("Get%s", utils.UpperFirst(p.Name))
}

func (p *Property) SetProperty(property types.Property) {
	p.property = property
}

func (p Property) GetMethod() string {
	kind := p.property.GetKind()
	if kind == "time.Time" {
		kind = "time"
	}
	return fmt.Sprintf("Encode%s", utils.UpperFirst(kind))
}

func (p Property) GetProtoKind() string {
	return p.property.GetProtoKind()
}

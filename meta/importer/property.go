package importer

import (
	"fmt"
	"github.com/byorty/hardcore/meta/types"
	"github.com/byorty/hardcore/utils"
)

type Property struct {
	Name  string `xml:"name,attr"`
	Kind  string `xml:"type,attr"`
	Alias string `xml:"as,attr"`
	prop  types.Property
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

func (p Property) GetSetterName() string {
	return fmt.Sprintf("Set%s", utils.UpperFirst(p.Name))
}

func (p Property) HasModelProperty() bool {
	return p.prop != nil
}

func (p *Property) SetModelProperty(prop types.Property) {
	p.prop = prop
}

func (p Property) GetModelProperty() types.Property {
	return p.prop
}

func (p Property) GetKind() string {
	return p.Kind
}

func (p Property) GetProtoKind() string {
	return fmt.Sprintf("Proto%sKind", utils.UpperFirst(p.GetKind()))
}

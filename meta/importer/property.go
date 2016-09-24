package importer

import (
	"fmt"
	"github.com/byorty/hardcore/meta/types"
	"github.com/byorty/hardcore/utils"
)

type Property struct {
	Name string                       `xml:"name,attr"`
	Kind types.ImportablePropertyKind `xml:"type,attr"`
	prop types.Property
}

func (p Property) GetName() string {
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

func (p Property) GetKind() types.ImportablePropertyKind {
	return p.Kind
}

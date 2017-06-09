package importer

import (
	"fmt"
	"github.com/byorty/hardcore/is"
	"github.com/byorty/hardcore/meta/model"
	"github.com/byorty/hardcore/meta/types"
	"github.com/byorty/hardcore/utils"
	"strings"
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

func (p Property) GetGetterName() string {
	return fmt.Sprintf("Get%s", utils.UpperFirst(p.Name))
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
	if is.NotNil(p.prop) {
		return p.prop.GetProtoKind()
	} else {
		return fmt.Sprintf("Proto%sKind", utils.UpperFirst(p.GetKind()))
	}
}

func (p Property) GetMethod() string {
	kind := p.prop.GetKind()
	if p.prop.GetRelation().IsNone() {
		if kind == "time.Time" {
			kind = "time"
		}
		if sliceKind, ok := types.SliceTypes[kind]; ok {
			kind = sliceKind
		}
	} else {
		if p.prop.GetEntity().GetEntityKind().IsEnum() {
			kind = string(p.prop.GetEntity().(*model.Enum).GetKind())
		}
	}
	return fmt.Sprintf("Decode%s", utils.UpperFirst(kind))
}

func (p Property) IsSlice() bool {
	return strings.Contains(p.GetProtoKind(), "Slice")
}

func (p Property) IsNotModelSlice() bool {
	return !p.IsModelSlice()
}

func (p Property) IsModelSlice() bool {
	return p.GetProtoKind() == "ProtoModelSliceKind"
}

func (p Property) GetSliceConstruct() string {
	var kind string
	if is.NotNil(p.prop) {
		kind = p.prop.GetKind()
	} else {
		kind = p.Kind
	}
	if sliceKind, ok := types.SliceTypes[kind]; ok {
		kind = sliceKind
	}
	return fmt.Sprintf("slice.New%s()", utils.UpperFirst(kind))
}
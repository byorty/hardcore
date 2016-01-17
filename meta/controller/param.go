package controller

import (
	"github.com/byorty/hardcore/meta/types"
	"fmt"
	"github.com/byorty/hardcore/utils"
	"github.com/byorty/hardcore/meta/model"
)

const RequestScopeKind = "types.RequestScope"
const FormScopeKind = "types.Form"

var reservedParams = map[string]string{
	"scope": RequestScopeKind,
	"form": FormScopeKind,
}

type Param struct {
	Name string `xml:"name,attr"`
	Required bool `xml:"required,attr"`
	Source string `xml:"from,attr"`
	Kind string `xml:"type,attr"`
	entity types.Entity
}

func (p Param) GetName() string {
	return p.Name
}

func (p Param) IsRequired() bool {
	return p.Required
}

func (p Param) GetSource() string {
	switch p.Source {
	case "path": return "types.PathPrimitiveSource"
	case "form": return "types.PostPrimitiveSource"
	default: return "types.GetPrimitiveSource"
	}
}

func (p Param) GetKind() string {
	return p.Kind
}

func (p *Param) SetEntity(entity types.Entity) {
	p.entity = entity
}

func (p Param) GetEntity() types.Entity {
	return p.entity
}

func (p Param) GetDefineKind() string {
	return p.getDefineKind(true)
}

func (p Param) getDefineKind(needPointer bool) string {
	if p.entity == nil {
		if p.IsReserved() {
			return reservedParams[p.Name]
		} else {
			return p.Kind
		}
	} else {
		if p.isMultiple() {
			return p.entity.GetFullMultipleName()
		} else {
			if needPointer {
				return p.entity.GetPointerFullName()
			} else {
				return p.entity.GetFullName()
			}
		}
	}
}

func (p Param) isMultiple() bool {
	return p.Kind == p.entity.GetFullMultipleName()
}

func (p Param) GetDefineVarKind() string {
	return p.getDefineKind(false)
}

func (p Param) GetDefineVarName() string {
	if p.entity == nil {
		return p.Name
	} else {
		if p.isMultiple() {
			return p.Name
		} else {
			return fmt.Sprintf("&%s", p.Name)
		}
	}
}

func (p Param) IsReserved() bool {
	_, ok := reservedParams[p.Name];
	return ok
}

func (p Param) GetPrimitive() string {
	if p.entity == nil {
		return utils.UpperFirst(p.Kind)
	} else {
		var kind string
		if p.entity.GetEntityKind() == types.ModelEntityKind {
			kind = p.entity.(types.ModelEntity).GetProperties()[0].GetKind()
		} else {
			kind = string(p.entity.(*model.Enum).Kind)
		}
		if p.isMultiple() {
			return fmt.Sprintf("%sIdentifiers", utils.UpperFirst(kind))
		} else {
			return fmt.Sprintf("%sIdentifier", utils.UpperFirst(kind))
		}
	}
}

package controller

import (
	"fmt"
	"github.com/byorty/hardcore/meta/model"
	"github.com/byorty/hardcore/meta/types"
	"github.com/byorty/hardcore/utils"
)

const RequestScopeKind = "types.RequestScope"
const FormScopeKind = "types.Form"

var injections = map[string]func() types.Injection{
	"scope":     NewRequestScopeInjection,
	"form":      NewFormInjection,
	"paginator": NewPaginatorInjection,
}

type Param struct {
	Name      string `xml:"name,attr"`
	Required  bool   `xml:"required,attr"`
	Source    string `xml:"from,attr"`
	Kind      string `xml:"type,attr"`
	entity    types.Entity
	injection types.Injection
}

func (p Param) GetName() string {
	return p.Name
}

func (p Param) IsRequired() bool {
	return p.Required
}

func (p Param) GetSource() string {
	switch p.Source {
	case "path":
		return "types.PathPrimitiveSource"
	case "form":
		return "types.PostPrimitiveSource"
	default:
		return "types.GetPrimitiveSource"
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
		if p.IsInjection() {
			return p.GetInjection().GetKind()
		} else {
			return p.Kind
		}
	} else {
		if p.isMultiple() {
			return p.entity.GetFullMultipleName()
		} else {
			if p.entity.GetEntityKind() == types.ModelEntityKind {
				if needPointer {
					return p.entity.GetPointerFullName()
				} else {
					return p.entity.GetFullName()
				}
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
			if p.entity.GetEntityKind() == types.ModelEntityKind {
				return fmt.Sprintf("&%s", p.Name)
			} else {
				return p.Name
			}
		}
	}
}

func (p Param) IsInjection() bool {
	_, ok := injections[p.Name]
	return ok
}

func (p Param) GetPrimitive() string {
	if p.entity == nil {
		return utils.UpperFirst(p.Kind)
	} else {
		var kind string
		var entityKind string
		if p.entity.GetEntityKind() == types.ModelEntityKind {
			kind = p.entity.(types.ModelEntity).GetProperties()[0].GetKind()
			entityKind = "Model"
		} else {
			kind = string(p.entity.(*model.Enum).Kind)
			entityKind = "Enum"
		}
		if p.isMultiple() {
			return fmt.Sprintf("%s%ss", utils.UpperFirst(kind), entityKind)
		} else {
			return fmt.Sprintf("%s%s", utils.UpperFirst(kind), entityKind)
		}
	}
}

func (p Param) GetInjection() types.Injection {
	if p.IsInjection() && p.injection == nil {
		p.injection = injections[p.Name]()
		p.injection.SetParam(&p)
	}
	return p.injection
}

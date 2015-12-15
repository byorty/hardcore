package plugin

import "github.com/byorty/hardcore/meta"

var (
    enumTpl = `package {{.Package}}

type {{.Name}} {{.Kind}}
`

    autoEnumTpl = `package {{.Package}}

func ({{.ShortName}} *{{.Name}}) ById(id int) {
	existsId := UserRole(id)
	if _, ok := userRoleNames[existsId]; ok {
		(*u) = existsId
	}
}

func ({{.ShortName}} {{.Name}}) GetId() {{.Kind}} {
	return {{.Kind}}(u)
}

func ({{.ShortName}} {{.Name}}) GetName() string {
	return userRoleNames[u];
}

const (
	GuestUserRole  UserRole = iota
	LoggedUserRole
)

var (
	userRoles = []types.Enum{
		GuestUserRole,
		LoggedUserRole,
	}
	userRoleNames = map[UserRole]string{
		GuestUserRole: "Guest",
		LoggedUserRole: "User",
	}
)
`
)

type Enum struct {}

func (e *Enum) Do(env *meta.Environment) {
    for _, container := range env.Configuration.ModelContainers {
        for _, enum := range container.Enums {

        }
    }
}
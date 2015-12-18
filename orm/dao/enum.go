package dao
import "github.com/byorty/hardcore/types"

type Enum struct {
    id interface{}
}

func (e *Enum) SetId(id interface{})  {
    e.id = id
}

func (e Enum) FindOne(dao types.EnumDAO, enum types.Named) {
    for _, named := range dao.GetList()  {
        if dao.Eq(named, e.id) {
            dao.Scan(named, enum)
        }
    }
}
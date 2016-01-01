package types

const (
	DefaultImport = "github.com/byorty/hardcore/types"
	DaoImport = "github.com/byorty/hardcore/orm/dao"
	ProtoImport = "github.com/byorty/hardcore/proto"
)

type Configuration interface {
    SetContainers([]Container)
    GetContainers() []Container
	AddAutoFile(string, string, map[string]interface{})
	AddFile(string, string, map[string]interface{})
	GetFiles() []File
}

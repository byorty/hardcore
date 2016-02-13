package types

const (
	DefaultImport = "github.com/byorty/hardcore/types"
	DaoImport = "github.com/byorty/hardcore/orm/dao"
	ProtoImport = "github.com/byorty/hardcore/proto"
	ExporterImport = "github.com/byorty/hardcore/exporter"
	MuxImport = "github.com/byorty/hardcore/mux"
	HelperImport = "github.com/byorty/hardcore/helper"
	PoolImport = "github.com/byorty/hardcore/pool"
)

type Configuration interface {
    SetContainers([]Container)
    GetContainers() []Container
	AddAutoFile(string, string, map[string]interface{})
	AddFile(string, string, map[string]interface{})
	GetFiles() []File
}

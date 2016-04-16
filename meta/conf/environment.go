package conf

import (
	"github.com/byorty/hardcore/log"
	"github.com/byorty/hardcore/meta/types"
)

type Environment struct {
	absPath       string
	metaPath      string
	importPart    string
	logger        log.Logger
	configuration types.Configuration
}

func NewEnvironment(absPath, metaPath, importPart string, logger log.Logger, configuration types.Configuration) types.Environment {
	return &Environment{
		absPath,
		metaPath,
		importPart,
		logger,
		configuration,
	}
}

func (e Environment) GetAbsPath() string {
	return e.absPath
}

func (e Environment) GetMetaPath() string {
	return e.metaPath
}

func (e Environment) GetImportPart() string {
	return e.importPart
}

func (e Environment) GetLogger() log.Logger {
	return e.logger
}

func (e Environment) GetConfiguration() types.Configuration {
	return e.configuration
}

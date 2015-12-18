package meta

import (
	"github.com/byorty/hardcore/log"
)

type Environment struct {
	AbsPath string
	MetaPath string
    ImportPart string
	Logger log.Logger
	Configuration *Configuration
}

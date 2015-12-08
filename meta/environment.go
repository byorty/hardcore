package meta

import "github.com/byorty/log4go"

type Environment struct {
	AbsPath string
	MetaPath string
	Logger log4go.Logger
	Configuration *Configuration
}

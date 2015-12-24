package types

import "github.com/byorty/hardcore/log"

type Environment interface {
    GetAbsPath() string
    GetMetaPath() string
    GetImportPart() string
    GetLogger() log.Logger
    GetConfiguration() Configuration
}
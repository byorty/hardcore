package model
import "github.com/byorty/hardcore/meta/common"

type Container struct {
    common.Container
    Enums []*Enum `xml:"enum"`
}
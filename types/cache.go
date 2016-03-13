package types

type Cache interface {
	Get(string) interface{}
	Has(string) bool
	Set(string, interface{}) Cache
	Remove(string) Cache
	GetKeys() []string
}

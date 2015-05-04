package types

type DAO interface {
	GetDB() string
	GetTable() string
	Scan(data interface{}, model interface{})
}

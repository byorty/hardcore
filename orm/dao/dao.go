package dao

type DAO interface {
	GetDB() string
	GetTable() string
	Scan(data interface{}, model interface{})
}

type SqlDAO struct {

}


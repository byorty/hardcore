package types

type DAOConnected interface {
	CommonDAO() DAO
}

type DAO interface {
	GetProto() Proto
	GetDB() string
	GetTable() string
	Scan(data interface{}, model interface{})
}

type IdentityMap interface {
	Get(string, interface{}) bool

}


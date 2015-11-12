package types

type DAOConnected interface {
	CommonDAO() DAO
}

type DAO interface {
	GetProto() Proto
	GetDB() string
	GetTable() string
	ScanAll(interface{}, interface{})
	Scan(interface{}, interface{})
}

type IdentityMap interface {
	Get(string, interface{}) bool
}

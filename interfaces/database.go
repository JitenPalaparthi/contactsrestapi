package interfaces

type DbConnecter interface {
	GetConnection(dns string) (interface{}, error)
}

package seller

type IRepository interface {
	FindByUUID(uuid string) (*Seller, error)
	list() ([]*Seller, error)
	getTop(size int) ([]*Seller, error)
}

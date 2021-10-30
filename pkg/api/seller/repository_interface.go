package seller

//go:generate mockgen -source=$GOFILE -destination=./repository_mock.go -package=$GOPACKAGE
type IRepository interface {
	FindByUUID(uuid string) (*Seller, error)
	list() ([]*Seller, error)
	getTop(size int) ([]*Seller, error)
}

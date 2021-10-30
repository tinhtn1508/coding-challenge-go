package v1

//go:generate mockgen -source=$GOFILE -destination=./repository_mock.go -package=$GOPACKAGE
type IRepository interface {
	delete(product *product) error
	insert(product *product) error
	update(product *product) error
	list(offset int, limit int) ([]*product, error)
	findByUUID(uuid string) (*product, error)
}

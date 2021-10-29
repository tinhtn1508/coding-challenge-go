package v2

type IRepository interface {
	delete(product *product) error
	insert(product *product) error
	update(product *product) error
	list(offset int, limit int) ([]*product, error)
	findByUUID(uuid string) (*product, error)
}

package database

type Product struct {
	ID   int32
	Name string
}

type ProductRepository interface {
	FindById(id int32) (*Product, error)
	Create(name string) (*Product, error)
}

type DBProductRepository struct {
	cursor int32
	DB     map[int32]*Product
}

func (r *DBProductRepository) FindById(id int32) (*Product, error) {
	return r.DB[id], nil
}

func (r *DBProductRepository) Create(name string) (*Product, error) {
	r.cursor++
	product := &Product{
		r.cursor,
		name,
	}
	r.DB[product.ID] = product
	return product, nil
}

func NewDBProductRepository() *DBProductRepository {
	return &DBProductRepository{
		0,
		make(map[int32]*Product),
	}
}

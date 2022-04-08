package products

import (
	"fmt"
	"github/tiagov07/c1-tm/GoWeb/practice3enviromentVar/pkg/store"
)

type Product struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Type  string  `json:"type"`
	Count int     `json:"count"`
	Price float64 `json:"price"`
}

var products []Product
var lastID int

type Repository interface {
	GetAll() ([]Product, error)
	Store(id int, name, productType string, count int, price float64) (Product, error)
	LastID() (int, error)
	Update(id int, name, productType string, count int, price float64) (Product, error)
	UpdateName(id int, name string) (Product, error)
	Delete(id int) error
}

type repository struct {
	db store.Store
}

func (r *repository) GetAll() ([]Product, error) {
	var products []Product
	r.db.Read(&products)
	return products, nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}

func (r *repository) Store(id int, name, productType string, count int, price float64) (Product, error) {
	var products []Product
	r.db.Read(&products)
	lastID = len(products)
	newProduct := Product{lastID, name, productType, count, price}
	products = append(products, newProduct)

	if err := r.db.Write(products); err != nil {
		return Product{}, err
	}
	return newProduct, nil
}

func (r *repository) Update(id int, name, productType string, count int, price float64) (Product, error) {
	product := Product{Name: name, Type: productType, Count: count, Price: price}
	var products []Product
	r.db.Read(&products)
	if len(products) > 0 {
		updated := false
		for i := range products {
			if products[i].Id == id {
				product.Id = id
				products[i] = product
				updated = true
			}
		}
		if !updated {
			return Product{}, fmt.Errorf("Product %d not found", id)
		} else {
			if err := r.db.Write(products); err != nil {
				return Product{}, fmt.Errorf(" file can't be writed")
			}
		}
	} else {
		return Product{}, fmt.Errorf(" there aren't elements to update")
	}
	return product, nil
}

func (r *repository) UpdateName(id int, name string) (Product, error) {
	var product Product
	var products []Product
	r.db.Read(&products)
	if len(products) > 0 {
		updated := false
		for i := range products {
			if products[i].Id == id {
				product = products[i]
				product.Id = id
				product.Name = name
				products[i] = product
				updated = true
			}
		}
		if !updated {
			return Product{}, fmt.Errorf("Product %d not found", id)
		} else {
			if err := r.db.Write(products); err != nil {
				return Product{}, fmt.Errorf(" file can't be writed")
			}
		}
	} else {
		return Product{}, fmt.Errorf(" there aren't elements to update")
	}
	return product, nil
}

func (r *repository) Delete(id int) error {
	var index int
	var products []Product
	r.db.Read(&products)
	if len(products) > 0 {
		deleted := false
		for i := range products {
			if products[i].Id == id {
				index = i
				deleted = true
			}
		}

		if !deleted {
			return fmt.Errorf("Product %d not found", id)
		}
	} else {
		return fmt.Errorf(" there aren't elements to delete")
	}
	products = append(products[:index], products[index+1:]...)
	if err := r.db.Write(products); err != nil {
		return fmt.Errorf(" file can't be updated")
	}
	return nil
}

func NewRepository(db store.Store) Repository {
	return &repository{db: db}
}

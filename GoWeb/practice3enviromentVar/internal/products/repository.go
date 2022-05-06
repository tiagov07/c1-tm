package products

import (
	"database/sql"
	"github/tiagov07/c1-tm/GoWeb/practice3enviromentVar/pkg/db"
	"log"
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
	Store(product Product) (Product, error)
	GetAll() ([]Product, error)
	GetByName(name string) (Product, error)
	LastID() (int, error)
	Update(id int, name, productType string, count int, price float64) (Product, error)
	UpdateName(id int, name string) (Product, error)
	Delete(id int) error
}

type repository struct {
	db *sql.DB
}

func (r *repository) GetByName(name string) (Product, error) {
	var product []Product
	var p Product
	db := db.StorageDB
	rows, err := db.Query("select name, type, count, price from products where name LIKE '%?%'", name)
	if err != nil {
		log.Println(err)
		return Product{}, nil
	}
	for rows.Next() {
		if err := rows.Scan(&p.Id, &p.Name, &p.Type, &p.Count, &p.Price); err != nil {
			log.Println(err.Error())
			return Product{}, nil
		}
		product = append(product, p)
	}
	return p, nil
}

func (r *repository) GetAll() ([]Product, error) {
	var products []Product
	return products, nil
}

func (r *repository) LastID() (int, error) {
	var ps []Product
	if len(ps) == 0 {
		return 0, nil
	}
	return ps[len(ps)-1].Id, nil
}

func (r *repository) Store(product Product) (Product, error) {
	db := db.StorageDB
	stmt, err := db.Prepare("INSERT INTO products(name, type, count, price) VALUES( ?, ?, ?, ?)")

	if err != nil {
		return Product{}, err
	}
	defer stmt.Close()
	var result sql.Result
	result, err = stmt.Exec(product.Name, product.Type, product.Count, product.Price)

	if err != nil {
		return Product{}, err
	}
	insertedId, _ := result.LastInsertId()
	product.Id = int(insertedId)

	return product, nil
}

func (r *repository) Update(id int, name, productType string, count int, price float64) (Product, error) {
	stmt, err := r.db.Prepare("UPDATE products SET name = ?, productType = ?, count = ?, price = ? WHERE id = ?")
	if err != nil {
		return Product{}, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(name, productType, count, price)
	if err != nil {
		return Product{}, err
	}
	product := Product{Id: id, Name: name, Type: productType, Count: count, Price: price}
	return product, nil
}

func (r *repository) UpdateName(id int, name string) (Product, error) {
	var product Product
	return product, nil
}

func (r *repository) Delete(id int) error {
	return nil
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

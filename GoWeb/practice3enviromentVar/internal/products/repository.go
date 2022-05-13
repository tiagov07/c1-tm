package products

import (
	"context"
	"database/sql"
	"github/tiagov07/c1-tm/GoWeb/practice3enviromentVar/pkg/db"
	"log"
)

type Product struct {
	Id              int     `json:"id"`
	Name            string  `json:"name"`
	Type            string  `json:"type"`
	Count           int     `json:"count"`
	Price           float64 `json:"price"`
	Id_warehouse    string  `json:"id_warehouse,omitempty"`
	WarehouseAdress string  `json:"warehouse_address,omitempty"`
}

var products []Product
var lastID int

type Repository interface {
	GetOne(id int) (Product, error)
	GetOneWithContext(ctx context.Context, id int) (Product, error)
	Store(name, productType string, count int, price float64) (Product, error)
	GetAll() ([]Product, error)
	GetFullData(id int) (Product, error)
	GetByName(name string) (Product, error)
	LastID() (int, error)
	Update(id int, name, productType string, count int, price float64) (Product, error)
	UpdateName(id int, name string) (Product, error)
	Delete(id int) error
}

const (
	GetAllProducts   = "SELECT Id, Name, Type, Count, Price FROM products"
	CreateNewProduct = "INSERT INTO products(name, type, count, price) VALUES( ?, ?, ?, ?)"
	UpdateProduct    = "UPDATE products SET name = ?, type = ?, count = ?, price = ? WHERE id = ?"
	DeleteProduct    = "DELETE FROM products WHERE id = ?"
	GetOneProduct    = "SELECT Id, Name, Type, Count, Price FROM products WHERE id = ?"
)

type repository struct {
	db *sql.DB
}

func (r *repository) GetByName(name string) (Product, error) {
	var products []Product
	db := db.StorageDB
	rows, err := db.Query("SELECT * FROM products WHERE name = ?", name)
	if err != nil {
		log.Println(err)
		return Product{}, nil
	}
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.Id, &p.Name, &p.Type, &p.Count, &p.Price); err != nil {
			log.Println(err.Error())
			return Product{}, nil
		}
		products = append(products, p)
	}
	return products[0], nil
}

func (r *repository) GetOne(id int) (Product, error) {
	var products []Product
	db := db.StorageDB
	rows, err := db.Query(GetOneProduct, id)
	if err != nil {
		log.Println(err)
		return Product{}, err
	}

	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.Id, &p.Name, &p.Type, &p.Count, &p.Price); err != nil {
			log.Println(err.Error())
			return Product{}, nil
		}
		products = append(products, p)
	}
	return products[0], nil

}

func (r *repository) GetOneWithContext(ctx context.Context, id int) (Product, error) {
	var product Product
	db := db.StorageDB
	getQuery := GetOneProduct

	rows, err := db.QueryContext(ctx, getQuery, id)
	if err != nil {
		log.Println(err)
		return Product{}, err
	}

	for rows.Next() {
		if err := rows.Scan(&product.Id, &product.Name, &product.Type, &product.Count, &product.Price); err != nil {
			log.Fatal(err)
			return Product{}, err
		}
	}
	return product, nil
}

func (r *repository) GetFullData(id int) (Product, error) {
	var product Product
	db := db.StorageDB
	innerJoin := "SELECT products.id, products.name, products.type, products.count, products.price, warehouses.name, warehouses.adress " +
		"FROM products INNER JOIN warehouses ON products.id_warehouse = warehouses.id " +
		"WHERE products.id = ?"
	rows, err := db.Query(innerJoin, id)
	if err != nil {
		log.Println(err)
		return product, err
	}
	for rows.Next() {
		if err := rows.Scan(&product.Id, &product.Name, &product.Type, &product.Count, &product.Price, &product.Id_warehouse,
			&product.WarehouseAdress); err != nil {
			log.Fatal(err)
			return product, err
		}
	}
	return product, nil
}

func (r *repository) GetAll() ([]Product, error) {
	var products []Product
	db := db.StorageDB
	rows, err := db.Query(GetAllProducts)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for rows.Next() {
		var product Product
		if err := rows.Scan(&product.Id, &product.Name, &product.Type, &product.Count, &product.Price); err != nil {
			log.Fatal(err)
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (r *repository) LastID() (int, error) {
	var ps []Product
	if len(ps) == 0 {
		return 0, nil
	}
	return ps[len(ps)-1].Id, nil
}

func (r *repository) Store(name, productType string, count int, price float64) (Product, error) {
	db := db.StorageDB
	stmt, err := db.Prepare(CreateNewProduct)

	if err != nil {
		return Product{}, err
	}
	defer stmt.Close()
	p := Product{Name: name, Type: productType, Count: count, Price: price}

	result, err := stmt.Exec(name, productType, count, price)

	if err != nil {
		return Product{}, err
	}
	insertedId, _ := result.LastInsertId()
	p.Id = int(insertedId)

	return p, nil
}

func (r *repository) Update(id int, name, productType string, count int, price float64) (Product, error) {
	stmt, err := r.db.Prepare(UpdateProduct)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	product := Product{Id: id, Name: name, Type: productType, Count: count, Price: price}
	_, err = stmt.Exec(name, productType, count, price, id)
	if err != nil {
		return Product{}, err
	}
	return product, nil
}

func (r *repository) UpdateName(id int, name string) (Product, error) {
	var product Product
	return product, nil
}

func (r *repository) Delete(id int) error {
	db := db.StorageDB
	stmt, err := db.Prepare(DeleteProduct)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)

	if err != nil {
		return err
	}
	return nil
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

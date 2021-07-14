package data

import (
	"fmt"

	"database/sql"
)

// swagger:model
type Product struct {
	//The product id
	//
	// required: true
	// min: 1
	ID int `json:"id"`

	//The name for this product
	// required: true
	Name string `json:"name" validate:"required"`

	//The description for this product
	Description string `json:"description"`

	//The price for this product
	//
	// required: true
	// min: 0
	Price float32 `json:"price" validate:"gt=0,required"`

	//The SKU for this product
	SKU string `json:"sku"`
}

type Products []*Product

var ErrProductNotFound = fmt.Errorf("Product not found")

func GetProductList(db *sql.DB) (Products, error) {
	rows, err := db.Query("select * from Products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	productDB := []*Product{}

	for rows.Next() {
		p := &Product{}
		err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.SKU)
		if err != nil {
			return nil, err
		}
		productDB = append(productDB, p)
	}
	return productDB, nil
}

func AddProduct(p *Product, db *sql.DB) error {
	_, err := db.Exec("insert into Products (name,description,price,sku) values ($1,$2,$3,$4)",
		p.Name,
		p.Description,
		p.Price,
		p.SKU)
	return err
}

func findProduct(id int, db *sql.DB) (bool, error) {
	rows, err := db.Query("SELECT exists (SELECT * FROM Products WHERE id=$1)", id)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	b := false
	for rows.Next() {
		err = rows.Scan(&b)
		if err != nil {
			return false, err
		}
	}

	return b, nil
}

func UpdateProduct(id int, p *Product, db *sql.DB) error {
	emp, err := findProduct(id, db)
	if err != nil {
		return err
	}
	if emp {
		_, err := db.Exec("UPDATE Products SET name = $1, description = $2, price = $3, sku = $4 WHERE id=$5", p.Name, p.Description, p.Price, p.SKU, id)
		if err != nil {
			return err
		}
		return nil
	} else {
		return ErrProductNotFound
	}
}

func DeleteProduct(id int, db *sql.DB) error {
	emp, err := findProduct(id, db)
	if err != nil {
		return err
	}
	if emp {
		_, err := db.Exec("delete from Products where id=$1", id)
		if err != nil {
			return err
		}
		return nil
	} else {
		return ErrProductNotFound
	}
}

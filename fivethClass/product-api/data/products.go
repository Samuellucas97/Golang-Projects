package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Product struct {
	ID				int `json:"id"`
	Name			string `json:"name"`
	Description		string `json:"description"`
	Price 			float32 `json:"price"`
	SKU				string `json:"sku"`
	CreatedOn		string `json:"-"`
	UpdatedOn		string `json:"-"`
	DeletedOn		string `json:"-"`
}

func (p *Product) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(p)
}

type Products []*Product


func (p *Products) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(p)
}

func GetProducts() Products {
	return productList
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func UpdateProduct(id int, p *Product) error {
	_, productPosition, err := findProduct(id)
	if err != nil {
		 return err
	}

	p.ID = id
	productList[productPosition] = p

	return nil
}

var ErrorProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}

	return nil, -1, ErrorProductNotFound
}

func getNextID() int {
	product := productList[len(productList)-1]
	return product.ID+1
}

// Defining products statically
var productList = []*Product {
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "fds2123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       6.89,
		SKU:         "bnda123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
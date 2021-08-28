package handlers

import (
	"awesomeProject/thirdFourthClass/product-api/data"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetProducts(res http.ResponseWriter, _ *http.Request) {
	p.l.Println("Handle GET Products")

	// Fetch the products
	productsList := data.GetProducts()

	// Converting Golang element to JSON
	err := productsList.ToJSON(res)
	if err != nil {
		http.Error(res, "Could not to Marshal JSON", http.StatusInternalServerError)
	}
}

func (p *Products) AddProduct(res http.ResponseWriter, req *http.Request) {
	p.l.Println("Handle POST Product")

	prod := processingJSON(res, req)
	p.l.Printf("Prod: %#v", prod)

	data.AddProduct(prod)
}

func (p *Products) UpdateProduct(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	idString := vars["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(res, "Could not convert id", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle PUT", idString)

	prod := processingJSON(res, req)
	p.l.Printf("Prod: %#v", prod)

	err = data.UpdateProduct(id, prod)
	if err == data.ErrorProductNotFound {
		http.Error(res, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(res, "Internal server error", http.StatusInternalServerError)
		return
	}

}

func processingJSON(res http.ResponseWriter, req *http.Request) *data.Product {
	prod := &data.Product{}
	err := prod.FromJSON(req.Body)
	if err != nil {
		http.Error(res, "Could not to Marshal JSON", http.StatusBadRequest)
	}
	return prod
}

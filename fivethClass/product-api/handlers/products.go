package handlers

import (
	"awesomeProject/thirdFourthClass/product-api/data"
	"context"
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

	prod := req.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)
}

func (p *Products) UpdateProduct(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	idString := vars["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(res, "Could not convert id", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle PUT Product", idString)
	prod := req.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(id, &prod)
	if err == data.ErrorProductNotFound {
		http.Error(res, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(res, "Internal server error", http.StatusInternalServerError)
		return
	}

}

type KeyProduct struct {}

func (p Products) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing product", err)
			http.Error(rw, "Error reading product", http.StatusBadRequest)
			return
		}

		// add the product to the context
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}


//func processingJSON(res http.ResponseWriter, req *http.Request) *data.Product {
//	prod := &data.Product{}
//	err := prod.FromJSON(req.Body)
//	if err != nil {
//		http.Error(res, "Could not to Marshal JSON", http.StatusBadRequest)
//	}
//	return prod
//}

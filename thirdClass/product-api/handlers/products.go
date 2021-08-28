package handlers

import (
	"awesomeProject/thirdClass/product-api/data"
	"log"
	"net/http"
)

type Products struct {
	l *log.Logger
}

// NewProducts is a function (like in Java/C++)
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}
// This is a method (like in Java/C++)
func (p *Products) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		p.getProducts(res, req)
		return
	}

	res.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(res http.ResponseWriter, req *http.Request) {
	productsList := data.GetProducts()

	// Converting Golang element to JSON
	err := productsList.ToJSON(res)
	if err != nil {
		http.Error(res, "Could not to use Marshal JSON", http.StatusInternalServerError)
	}
}
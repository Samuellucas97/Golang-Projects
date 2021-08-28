package handlers

import (
	"awesomeProject/thirdFourthClass/product-api/data"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		p.getProducts(res, req)
		return
	}

	if req.Method == http.MethodPost {
		p.addProduct(res, req)
		return
	}

	if req.Method == http.MethodPut {
		p.l.Println("Handle PUT Product", req.URL.Path)
		r := regexp.MustCompile(`/([0-9]+)`)
		pathUrl := req.URL.Path
		g := r.FindAllStringSubmatch(pathUrl, -1)


		if  len(g) != 1 {
			p.l.Println("Invalid URI more than one id")
			http.Error(res, "Invalid URI", http.StatusBadRequest)
			return
		}

		if  len(g[0]) != 2 {
			p.l.Println("Invalid URI more than one capture group")
			http.Error(res, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(res, "Problem converting id from string to int format", http.StatusInternalServerError)
			return
		}
				
		p.updateProduct(id, res, req)
		return
	}

	res.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(res http.ResponseWriter, req *http.Request) {
	p.l.Println("Handle GET Products")

	// Fetch the products
	productsList := data.GetProducts()

	// Converting Golang element to JSON
	err := productsList.ToJSON(res)
	if err != nil {
		http.Error(res, "Could not to Marshal JSON", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(res http.ResponseWriter, req *http.Request) {
	p.l.Println("Handle POST Product")

	prod := ProcessingJSON(res, req)
	p.l.Printf("Prod: %#v", prod)

	data.AddProduct(prod)
}

func (p *Products) updateProduct(id int, res http.ResponseWriter, req *http.Request) {
	prod := ProcessingJSON(res, req)
	p.l.Printf("Prod: %#v", prod)

	err := data.UpdateProduct(id, prod)
	if err == data.ErrorProductNotFound {
		http.Error(res, "Product not found", http.StatusNotFound)
	}

	if err != nil {
		http.Error(res, "Internal server error", http.StatusInternalServerError)
	}

}

func ProcessingJSON(res http.ResponseWriter, req *http.Request) *data.Product {
	prod := &data.Product{}
	err := prod.FromJSON(req.Body)
	if err != nil {
		http.Error(res, "Could not to Marshal JSON", http.StatusBadRequest)
	}
	return prod
}

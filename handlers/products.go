package handlers

import (
	"log"
	"net/http"
	"path"
	"strconv"

	"github.com/zzibert/building-microservices-with-go/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		p.getProducts(rw, r)
		return
	case http.MethodPost:
		p.addProduct(rw, r)
		return
	case http.MethodPut:
		p.updateProduct(rw, r)
		return
	case http.MethodDelete:
		p.deleteProduct(rw, r)
		return
	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Products")

	product := &data.Product{}
	err := product.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}
	p.l.Printf("Prod %#v", product)
	data.AddProduct(product)
}

func (p *Products) updateProduct(rw http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		http.Error(rw, "Unable to parse the id", http.StatusBadRequest)
	}

	product := data.GetProduct(id)
	if product == nil {
		http.Error(rw, "Unable to find the product", http.StatusNotFound)
		return
	}
	err = product.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}
}

func (p *Products) deleteProduct(rw http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		http.Error(rw, "Unable to parse the id", http.StatusBadRequest)
		return
	}

	product := data.DeleteProduct(id)
	if product == nil {
		http.Error(rw, "Unable to find product", http.StatusNotFound)
		return
	}
	products := data.Products{product}

	err = products.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusBadRequest)
	}

}

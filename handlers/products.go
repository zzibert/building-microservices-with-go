package handlers

import (
	"log"
	"net/http"

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
	case "GET":
		p.getProducts(rw, r)
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

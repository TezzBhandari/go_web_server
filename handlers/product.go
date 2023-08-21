package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/TezzBhandari/go_web_server/data"
)

type Product struct {
	log *log.Logger
}

func NewProduct(log *log.Logger) *Product {
	return &Product{
		log: log,
	}
}

func (p *Product) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	products := data.GetProducts()

	responseJson, err := json.Marshal(products)
	if err != nil {
		http.Error(rw, "Unable to marshall json", http.StatusInternalServerError)
	}
	rw.Header().Add("Content-Type", "application/json")
	rw.Write(responseJson)
}

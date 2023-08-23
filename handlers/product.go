package handlers

import (
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
	if r.Method == http.MethodGet {
		getProducts(rw, r)
		return
	}

	// catches all other method
	rw.WriteHeader(http.StatusMethodNotAllowed)

}

func getProducts(rw http.ResponseWriter, r *http.Request) {
	products := data.GetProducts()

	// responseJson, err := json.Marshal(products)
	rw.Header().Add("Content-Type", "application/json")
	err := products.ToJson(rw)
	if err != nil {
		http.Error(rw, "Unable to marshall json", http.StatusInternalServerError)
	}
	// 	rw.Write(responseJson)
}

func PostProduct(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("not implemented yet"))
}

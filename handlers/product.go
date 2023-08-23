package handlers

import (
	"log"
	"net/http"

	"github.com/TezzBhandari/go_web_server/data"
)

type Products struct {
	log *log.Logger
}

func NewProduct(log *log.Logger) *Products {
	return &Products{
		log: log,
	}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}

	// catches all other method
	rw.WriteHeader(http.StatusMethodNotAllowed)

}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.log.Println("Handle GET Products")
	products := data.GetProducts()

	// responseJson, err := json.Marshal(products)
	rw.Header().Add("Content-Type", "application/json")
	err := products.ToJson(rw)
	if err != nil {
		http.Error(rw, "Unable to marshall json", http.StatusInternalServerError)
	}
	// 	rw.Write(responseJson)
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.log.Println("Handle POST Products")

	product := &data.Product{}
	err := product.FromJson(r.Body)
	if err != nil {
		http.Error(rw, "Unable to Marshal Json", http.StatusBadRequest)
	}

	p.log.Printf("product: %#v", product)

	data.AddProduct(product)

	rw.WriteHeader(http.StatusCreated)
	rw.Write([]byte("successfull"))

	// rw.Write([]byte("not implemented yet"))
}

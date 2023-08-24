package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/TezzBhandari/go_web_server/data"
	"github.com/gorilla/mux"
)

// Defines Products Handler
type Products struct {
	log *log.Logger
}

type ProductKey struct{}

// creates new product handler
func NewProduct(log *log.Logger) *Products {
	return &Products{
		log: log,
	}
}

// func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodGet {
// 		p.getProducts(rw, r)
// 		return
// 	}

// 	if r.Method == http.MethodPost {
// 		p.addProduct(rw, r)
// 		return
// 	}

// 	if r.Method == http.MethodPut {
// 		p.log.Println("PUT", r.URL.Path)
// 		// expect the id in the URI
// 		reg := regexp.MustCompile(`/([0-9]+)`)
// 		g := reg.FindAllStringSubmatch(r.URL.Path, -1)

// 		if len(g) != 1 {
// 			p.log.Println("Invalid URI more than one id")
// 			http.Error(rw, "Invalid URI", http.StatusBadRequest)
// 			return
// 		}

// 		if len(g[0]) != 2 {
// 			p.log.Println("Invalid URI more than one capture group")
// 			http.Error(rw, "Invalid URI", http.StatusBadRequest)
// 			return
// 		}

// 		idString := g[0][1]
// 		id, err := strconv.Atoi(idString)
// 		if err != nil {
// 			p.log.Println("Invalid URI unable to convert to numer", idString)
// 			http.Error(rw, "Invalid URI", http.StatusBadRequest)
// 			return
// 		}

// 		p.updateProducts(id, rw, r)
// 		return
// 	}

// 	// catches all other method
// 	rw.WriteHeader(http.StatusMethodNotAllowed)

// }

// returns all the products
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
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

// adds product
func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {

	p.log.Println("Handle POST Products")

	product := r.Context().Value(ProductKey{}).(*data.Product)

	p.log.Printf("product: %#v", product)

	data.AddProduct(product)

	rw.WriteHeader(http.StatusCreated)
	rw.Write([]byte("successfull"))

	// rw.Write([]byte("not implemented yet"))
}

// updates product
func (p Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, conversionError := strconv.Atoi(vars["id"])

	if conversionError != nil {
		http.Error(rw, "Invalid ID", http.StatusBadRequest)
		return
	}

	p.log.Println("Handle PUT Product ", id)
	prod := r.Context().Value(ProductKey{}).(*data.Product)

	err := data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}

func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}

		err := prod.FromJson(r.Body)
		if err != nil {
			http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
			return
		}

		// need to pass the json objecct to next handler using request context
		ctx := context.WithValue(r.Context(), ProductKey{}, prod)

		// create a new request and copies the request context to it and returns request object
		r = r.WithContext(ctx)
		next.ServeHTTP(rw, r)
	})
}

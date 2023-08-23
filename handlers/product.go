package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

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

	if r.Method == http.MethodPut {
		p.log.Println("PUT", r.URL.Path)
		// expect the id in the URI
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)

		if len(g) != 1 {
			p.log.Println("Invalid URI more than one id")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			p.log.Println("Invalid URI more than one capture group")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			p.log.Println("Invalid URI unable to convert to numer", idString)
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		p.updateProducts(id, rw, r)
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

func (p *Products) updateProducts(id int, rw http.ResponseWriter, r *http.Request) {
	p.log.Println("Handle PUT Product")

	prod := &data.Product{}

	err := prod.FromJson(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}

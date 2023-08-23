package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"product_name"`
	Description string  `json:"product_description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

func (p *Product) FromJson(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(p)
}

func AddProduct(product *Product) {
	product.ID = getNextID()

	productList = append(productList, product)

}

func getNextID() int {
	lastItem := productList[len(productList)-1]

	return lastItem.ID + 1
}

type Products []*Product

func (products *Products) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(products)
}

var productList = Products{
	{
		ID:          1,
		Name:        "Latte",
		Description: "forthy milky coffee",
		Price:       2.45,
		SKU:         "abc123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Esspresso",
		Description: "short and strong coffee without milk ",
		Price:       1.95,
		SKU:         "bcd234",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}

func GetProducts() Products {
	return productList
}

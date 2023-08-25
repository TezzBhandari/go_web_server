package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"product_name" validate:"required"`
	Description string  `json:"product_description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", ValidateSKU)
	return validate.Struct(p)
}

func ValidateSKU(fl validator.FieldLevel) bool {
	regx := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := regx.FindAllString(fl.Field().String(), -1)
	return len(matches) == 1
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

// On the Data

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

func UpdateProduct(id int, p *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}

	p.ID = id
	productList[pos] = p

	return nil
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}

	return nil, -1, ErrProductNotFound
}

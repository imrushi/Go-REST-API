package handler

import (
	"encoding/json"
	"fmt"
	"io"
)

type Product struct {
	ID          int     `json:"id,omitempty"`
	Name        string  `json:"name"`
	Description string  `json:"description,omitempty"`
	Price       float32 `json:"price"`
}

// Products is a collection of Product
type Products []*Product

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Bottle",
		Description: "A bottle for drink.",
		Price:       30.0,
	},
	&Product{
		ID:          2,
		Name:        "Milk",
		Description: "Drink milk.",
		Price:       60.0,
	},
}

func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

// GetProducts returns a list of products
func GetProductsList() Products {
	return productList
}

func AddProductList(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
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

func getNextID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

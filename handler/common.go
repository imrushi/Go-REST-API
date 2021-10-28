package handler

import (
	"encoding/json"
	"io"
)

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float32
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

func getNextID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

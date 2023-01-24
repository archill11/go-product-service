package data

import (
	"encoding/json"
	"io"
	"time"
)

// Entity
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product

// метод переводит список сущностей в JSON
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

// метод получает сущность из JSON
func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

// метод возвращает все сущности из БД
func GetProducts() Products {
	return productList
}

// метод добавляет сужность в БД
func AddProduct(prod *Product) {
	prod.ID = getNextID()
	productList = append(productList, prod)
}

// метод получает id для новой сущности
func getNextID() int {
	lp := productList[len(productList)-1] // получаем последни id из БД
	return lp.ID + 1
}

// метод полнестью заменяет сущность
func PutProduct(prod *Product) {
	for i, v := range productList { // ищем в БД сущность с нужным id
		if v.ID == prod.ID {
			productList[i] = prod
		}
	}
}

// БД
var productList = Products{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frotty milky coffe",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffe without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}

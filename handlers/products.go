package handlers

import (
	"log"
	"myapp/data"
	"net/http"
	"regexp"
	"strconv"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	if req.Method == http.MethodGet { // GET
		p.GetProducts(rw, req)
		return
	}
	
	if req.Method == http.MethodPost { // POST
		p.addProducts(rw, req)
		return
	}

	if req.Method == http.MethodPut { // PUT
		path := req.URL.Path
		reg := regexp.MustCompile(`/product/([0-9]+)`) // получаем параметр(id сущности) из адреса
		g := reg.FindAllStringSubmatch(path, -1)

		if len(g) != 1 {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		if len(g[0]) != 2 {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := g[0][1]               // берем id из регулярного выражения
		id, err := strconv.Atoi(idString) // кастим string в int
		if err != nil {
			return
		}

		p.putProducts(id, rw, req)
		return
	}

	// catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) GetProducts(rw http.ResponseWriter, req *http.Request) {
	lp := data.GetProducts() // получаем данные из БД
	err := lp.ToJSON(rw)     // возвращаем на клиент в виде JSON
	if err != nil {
		http.Error(rw, "Ooops", http.StatusInternalServerError)
		return
	}
}

func (p *Products) addProducts(rw http.ResponseWriter, req *http.Request) {
	prod := &data.Product{}        // создаем новую пустую сущность
	err := prod.FromJSON(req.Body) // добавляем поля сущности из тела запроса
	if err != nil {
		http.Error(rw, "Ooops", http.StatusInternalServerError)
		return
	}
	data.AddProduct(prod) // добавляем сущность в БД
}

func (p *Products) putProducts(id int, rw http.ResponseWriter, req *http.Request) {
	prod := &data.Product{} // создаем новую пустую сущность
	prod.ID = id            // устанавливаем id сущности
	err := prod.FromJSON(req.Body)
	if err != nil {
		http.Error(rw, "Ooops", http.StatusInternalServerError)
		return
	}
	data.PutProduct(prod) // заменяем сущность в БД
}

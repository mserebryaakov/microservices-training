package handler

import (
	"net/http"

	"github.com/mserebryaakov/microservices-training/data"
)

// swagger:route GET /products products listProducts
// Return a list of products from the database
// responses:
//	200: productsResponse

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET")
	lp, err := data.GetProductList(p.db)
	if err != nil {
		p.l.Println("Error get product from db", err)
		http.Error(rw, "Error get request", http.StatusInternalServerError)
	}
	err = lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

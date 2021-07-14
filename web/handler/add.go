package handler

import (
	"net/http"

	"github.com/mserebryaakov/microservices-training/data"
)

// swagger:route POST /products products createProduct
// Create product with the enter id from the database
// responses:
//	201: noContentResponse
//	422: errorValidation

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST")

	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err := data.AddProduct(&prod, p.db)
	if err != nil {
		p.l.Println("Error add product from db", err)
		http.Error(rw, "Error post request", http.StatusInternalServerError)
		return
	}
}

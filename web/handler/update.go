package handler

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/mserebryaakov/microservices-training/data"
)

// swagger:route PUT /products/{id} products updateProduct
// Update product with the enter id from the database
// responses:
//	201: noContentResponse
//	400: errorResponse
//	422: errorValidation

func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle PUT", id)

	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(id, &prod, p.db)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		p.l.Println("Error update product from db", err)
		http.Error(rw, "Error put request", http.StatusInternalServerError)
		return
	}
}

package handler

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/mserebryaakov/microservices-training/data"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Delete product with the enter id from the database
// responses:
//	201: noContent

func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle DELETE", id)

	err = data.DeleteProduct(id)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Id not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Error", http.StatusInternalServerError)
		return
	}
}

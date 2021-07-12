package handler

import "github.com/mserebryaakov/microservices-training/data"

// A list of products returns in the response
// swagger:response productsResponse
type productsResponseWrapper struct {
	// All products in the system
	// in:Body
	Body []data.Product
}

// swagger:parameters createProduct updateProduct
type productParamWrapper struct {
	// Product in the system
	// in: path
	Body data.Product
}

// swagger:response NoContent
type productNoContent struct {
}

// swagger:parameters deleteProduct
type productIDParameterWrapper struct {
	// The id of the product to delete from the database
	// in: path
	// required: true
	ID int `json:"id"`
}

// Validation ERROR
// swagger:response errorValidation
type errorValidation struct {
}

// Response ERROR
// swagger:response errorResponse
type errorResponse struct {
}

// Package classification of Product API.
//
// Documentation for Product API
//
//     Schemes: http
//     Host: localhost
//     Version: 1.0.0
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
// swagger:meta
package handler

import "github.com/mserebryaakov/microservices-training/data"

// A list of products returns in the response
// swagger:response productsResponse
type productsResponseWrapper struct {
	// All products in the system
	// in: body
	Body []data.Product
}

// swagger:parameters createProduct updateProduct
type productParamWrapper struct {
	// Product in the system
	// Note: the id must be ignored when creating a product
	// in: body
	Body data.Product
}

// No content is returned by this API endpoint
// swagger:response noContentResponse
type noContentResponseWrapper struct {
}

// swagger:parameters deleteProduct updateProduct
type productIDParameterWrapper struct {
	// The id of the product to delete from the database
	// in: path
	// required: true
	ID int `json:"id"`
}

// Validation ERROR
// swagger:response errorValidation
type errorValidationWrapper struct {
}

// Response ERROR
// swagger:response errorResponse
type errorResponseWrapper struct {
}

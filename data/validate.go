package data

import "github.com/go-playground/validator"

func (p *Product) Validate() error {
	validate := validator.New()

	return validate.Struct(p)
}

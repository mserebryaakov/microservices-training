package data

import "testing"

func TestCheckValidate(t *testing.T) {
	p := &Product{
		Name:  "Tea",
		Price: 0.8,
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}

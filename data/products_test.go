package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := Product{
		Name:  "nic",
		Price: 1.0,
		SKU:   "abc-cdfdfd-dfdf",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}

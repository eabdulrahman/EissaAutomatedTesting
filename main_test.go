package main

import "testing"

var tests = []struct {
	name  string
	state string
	input []Product
	want  string
}{
	{"Washinton", "WA", cart1, "Success"},
	{"Idaho", "ID", cart1, "Success"},
	{"Montana", "MT", cart2, "Success"},
	{"Massachusetts", "MA", cart2, "Failure"},
}

var cart1 = []Product{
	{"Orange", "Fruit", 15},
	{"Infant Milk", "Wic Eligible Food", 20},
	{"MS Office", "Software", 100},
	{"Soap", "Detergent", 50},
}

var cart2 = []Product{
	{"Potato", "Fruit", 15},
	{"Banana", "Wic Eligible Food", 20},
	{"Windows OS", "Software", 100},
}

func TestCheckout(t *testing.T) {

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := Checkout(tt.state, tt.input)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}

package main

import "fmt"

type Product struct {
	Name  string
	Type  string
	Price float32
}

func Checkout(state string, productsList []Product) string {
	var taxPercentage float32
	var unitTax float32
	var totalTax float32
	var productsPrices float32

	//calculate the tax percentage per State
	switch state {
	case "MT":
		taxPercentage = 0
	case "ID":
		taxPercentage = 0.06
	case "WA":
		taxPercentage = 0.0938
	default:
		//in case the sate in not one of the above
		fmt.Printf("UNFORTUNATELY, WE DO NOT HAVE BRANCHES IN %v.\n", state)
	}

	//if the state is on of the mentioned above, then calculate the total price + tax
	if (state == "MT") || (state == "ID") || (state == "WA") {
		for i, _ := range productsList {
			if productsList[i].Price > 0 {

				//No tax for WIC products in all states and Software in WA
				if (productsList[i].Type == "Wic Eligible Food") || (state == "WA" && productsList[i].Type == "Software") {
					unitTax = 0.0
				} else {
					//calculate the tax per product
					unitTax = productsList[i].Price * taxPercentage
				}
				fmt.Printf("%v is : $%v + Tax $%v \n", productsList[i].Name, productsList[i].Price, unitTax)
				totalTax += unitTax                     //increment the total tax in every iteration by a product tax
				productsPrices += productsList[i].Price //increase the total price in every iteration by a product price
			}
		}
		totalPrice := totalTax + productsPrices //calculate the total Price
		fmt.Printf("Sales Tax : $%v\nTotal Price : $%v\n", totalTax, totalPrice)

		return "Success" //the purchase operation was successful
	} else {
		return "Failure" //the purchase operation failed
	}
}

func main() {

}

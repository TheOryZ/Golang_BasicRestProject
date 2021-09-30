package main

import (
	"fmt"
	"goModul/api"
)

func main() {
	//POST
	api.AddProduct()
	//GET
	products, _ := api.GetAllProducts()

	for i := 0; i < len(products); i++ {
		fmt.Println(products[i].ProductName)
	}

	//POST
	api.AddCategory()
	//GET
	categories, _ := api.GetAllCategories()

	for i := 0; i < len(categories); i++ {
		fmt.Println(categories[i].CategoryName)
	}
}

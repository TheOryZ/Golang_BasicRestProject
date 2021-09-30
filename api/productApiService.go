package api

import (
	"bytes"
	"encoding/json"
	"goModul/model"
	"io/ioutil"
	"net/http"
)

func GetAllProducts() ([]model.Product, error) {
	resp, err := http.Get("http://localhost:3000/products")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	var products []model.Product
	json.Unmarshal(bodyBytes, &products)
	return products, nil
}

func AddProduct() (model.Product, error) {
	// TODO : get product information
	product := model.Product{Id: 4, ProductName: "Dummy Product 4", CategoryId: 2, UnitPrice: 6000.00}
	jsonProduct, _ := json.Marshal(product)
	resp, err := http.Post("http://localhost:3000/products", "application/json;charset=utf-8", bytes.NewBuffer(jsonProduct))
	if err != nil {
		return model.Product{}, err // TODO : cannot use nil as type model.Product in return argument ?? **Closed : 'Cause struct is not array type and we cannot return nil
	}
	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	var productResp model.Product
	json.Unmarshal(bodyBytes, &productResp)
	return productResp, nil
}

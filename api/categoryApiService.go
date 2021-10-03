package api

import (
	"bytes"
	"encoding/json"
	"goModul/model"
	"io/ioutil"
	"net/http"
)

func GetAllCategories() ([]model.Category, error) {
	resp, err := http.Get("http://localhost:3000/categories")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	var categories []model.Category
	json.Unmarshal(bodyBytes, &categories)
	return categories, nil
}

func AddCategory() (*model.Category, error) {
	// TODO : get category information
	category := model.Category{Id: 3, CategoryName: "Dummy Category 3"}
	jsonCategory, _ := json.Marshal(category)
	resp, err := http.Post("http://localhost:3000/categories", "application/json;charset=utf-8", bytes.NewBuffer(jsonCategory))
	if err != nil {
		return nil, err // TODO : cannot use nil as type model.Product in return argument ?? **Closed : 'Cause struct is not array type and we cannot return nil
	}
	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	var categoryResp model.Category
	json.Unmarshal(bodyBytes, &categoryResp)
	return &categoryResp, nil
}

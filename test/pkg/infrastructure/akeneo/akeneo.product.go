package akeneo

import (
	"encoding/json"
	"fmt"
	"github/Abraxas-365/akeneo-connector/pkg/core/models"
	"net/http"
	"strconv"
)

func (a *akeneo) UploadProducts(products models.ProductsAkeneo) error {
	token, err := a.Login()
	if err != nil {
		fmt.Println("ERR", err)
		return err
	}
	if err := uploadData(a.url+"/api/rest/v1/products", token, products); err != nil {
		fmt.Println("ERR", err)
		return err
	}

	return nil
}

func (a *akeneo) GetProduct(code int) (models.ProductAkeneo, error) {
	token, err := a.Login()
	if err != nil {
		fmt.Println("ERR", err)
		return models.ProductAkeneo{}, err
	}

	var bearer = "Bearer " + token
	fmt.Println("akeneo", bearer)
	var client http.Client

	req, err := http.NewRequest("GET", a.url+"/api/rest/v1/products/"+strconv.Itoa(code), nil)
	if err != nil {
		return models.ProductAkeneo{}, err
	}
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-Type", "application/vnd.akeneo.collection+json")
	resp, err := client.Do(req)
	if err != nil {
		return models.ProductAkeneo{}, err
	}
	var product models.ProductAkeneo
	if err := json.NewDecoder(resp.Body).Decode(&product); err != nil {
		return models.ProductAkeneo{}, err
	}

	return product, nil
}

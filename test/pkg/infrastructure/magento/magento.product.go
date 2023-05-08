package magento

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github/Abraxas-365/akeneo-connector/pkg/core/models"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (m *magnetoRepo) GetProductById(id int) (models.ProductsMagento, error) {
	// var resp models.ProductFromMagento

	resp := struct {
		Iteams []models.ProductMagento `json:"items"`
	}{}
	token, err := m.Login()
	if err != nil {
		return models.ProductsMagento{}, err
	}

	var bearer = "Bearer " + token
	fmt.Println(bearer)
	var client http.Client
	req, err := http.NewRequest("GET",
		m.url+"/rest/default/V1/products?searchCriteria[filter_groups][0][filters][0][field]=sku&searchCriteria[filter_groups][0][filters][0][value]="+
			strconv.Itoa(id),
		nil)
	if err != nil {
		return models.ProductsMagento{}, err
	}
	req.Header.Add("Authorization", bearer)
	res, err := client.Do(req)
	if err != nil {
		return models.ProductsMagento{}, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err = json.Unmarshal(body, &resp); err != nil {
		return models.ProductsMagento{}, err
	}

	return resp.Iteams, nil
}

func (m *magnetoRepo) GetProductsByPage(page int, pageSize int) (models.ProductsMagento, error) {

	resp := struct {
		Iteams []models.ProductMagento `json:"items"`
	}{}
	token, err := m.Login()
	if err != nil {
		return models.ProductsMagento{}, err
	}

	var bearer = "Bearer " + token
	var client http.Client
	req, err := http.NewRequest("GET",
		"https://local.abcdin245.cl/rest/default/V1/products?searchCriteria[pageSize]="+
			strconv.Itoa(pageSize)+
			"&searchCriteria[currentPage]="+
			strconv.Itoa(page),
		nil)
	if err != nil {
		return models.ProductsMagento{}, err
	}
	req.Header.Add("Authorization", bearer)
	res, err := client.Do(req)
	if err != nil {
		return models.ProductsMagento{}, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err = json.Unmarshal(body, &resp); err != nil {
		return models.ProductsMagento{}, err
	}
	return resp.Iteams, nil
}

func (m *magnetoRepo) UpdateProduct(code int, product models.ProductMagento) error {
	token, err := m.Login()
	if err != nil {
		return err
	}
	var bearer = "Bearer " + token
	fmt.Println("magento", bearer)

	var client http.Client
	var buf bytes.Buffer
	bodyInput := struct {
		Product models.ProductMagento `json:"product"`
	}{}
	bodyInput.Product = product
	if err := json.NewEncoder(&buf).Encode(bodyInput); err != nil {
		return err
	}
	req, err := http.NewRequest("PUT", m.url+"/rest/V1/products/"+strconv.Itoa(code), &buf)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	// body, err := ioutil.ReadAll(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(body, &resp)

	fmt.Println(string(body))
	buf.Reset()

	return nil
}

package magento

import (
	"encoding/json"
	"fmt"
	"github/Abraxas-365/akeneo-connector/pkg/core/models"
	"io/ioutil"
	"net/http"
)

func (m *magnetoRepo) GetOptions() (models.OptionsMagentoResponse, error) {
	resp := struct {
		Iteams models.OptionsMagentoResponse `json:"items"`
	}{}
	token, err := m.Login()
	if err != nil {
		return models.OptionsMagentoResponse{}, err
	}

	var bearer = "Bearer " + token
	fmt.Println(bearer)
	var client http.Client
	req, err := http.NewRequest("GET",
		m.url+`/rest/V1/products/attributes?searchCriteria[filterGroups][0][filters][0][field]=frontend_input&searchCriteria[filterGroups][0][filters][0][condition_type]=in&searchCriteria[filterGroups][0][filters][0][value]=select,multiselect&searchCriteria[filterGroups][1][filters][0][field]=is_visible&searchCriteria[filterGroups][1][filters][0][condition_type]=eq&searchCriteria[filterGroups][1][filters][0][value]=1`,
		nil)
	if err != nil {
		return models.OptionsMagentoResponse{}, err
	}
	req.Header.Add("Authorization", bearer)
	res, err := client.Do(req)
	if err != nil {
		return models.OptionsMagentoResponse{}, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err = json.Unmarshal(body, &resp); err != nil {
		return models.OptionsMagentoResponse{}, err
	}

	return resp.Iteams, nil
}

func (m *magnetoRepo) OptionsInicialLoad() error {
	if _, err := m.db.Exec(`
INSERT into akeneo_connector_entities (import,code,entity_id)
Select Distinct
	'option' as import,
	CONCAT(ea.attribute_code,'_',eao.option_id) as code,
	eao.option_id as 'entity_id'
From eav_attribute ea 
Inner Join eav_attribute_option eao 
	On eao.attribute_id = ea.attribute_id 
Left Join eav_attribute_option_value eaov1 
	ON eaov1.option_id = eao.option_id 
	And eaov1.store_id = 0
Left Join eav_attribute_option_value eaov2 
	ON eaov2.option_id = eao.option_id 
	And eaov1.store_id = 1
Where ea.entity_type_id = 4
And ea.frontend_input in ('select', 'multiselect') 
Order By ea.attribute_id, eao.sort_order asc;
`); err != nil {
		return err
	}

	return nil
}

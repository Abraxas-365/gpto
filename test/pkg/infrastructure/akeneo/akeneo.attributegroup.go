package akeneo

import (
	"encoding/json"
	"fmt"
	"github/Abraxas-365/akeneo-connector/pkg/core/models"
	"net/http"
)

func (a *akeneo) UploadAttributeGroups(attributesGroups models.AttributeGroups) error {
	token, err := a.Login()
	if err != nil {
		return err
	}
	if err := uploadData(a.url+"/api/rest/v1/attribute-groups", token, attributesGroups); err != nil {
		return err
	}

	return nil
}
func (a *akeneo) GetAttributeGroup(code string) ([]string, error) {
	token, err := a.Login()
	if err != nil {
		return []string{}, err
	}
	var bearer = "Bearer " + token
	fmt.Println("akeneo", bearer)
	var client http.Client

	req, err := http.NewRequest("GET", a.url+"/api/rest/v1/attribute-groups/"+code, nil)
	if err != nil {
		return []string{}, err
	}
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return []string{}, err
	}
	attributeGroup := struct {
		Attributes []string `json:"attributes"`
	}{}
	if err := json.NewDecoder(resp.Body).Decode(&attributeGroup); err != nil {
		return []string{}, err
	}
	return attributeGroup.Attributes, nil
}

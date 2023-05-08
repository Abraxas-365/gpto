package akeneo

import (
	"github/Abraxas-365/akeneo-connector/pkg/core/models"
	"log"
)

func (a *akeneo) UploadOptions(options models.OptionsAkeneo, attributeCode string) error {
	token, err := a.Login()
	if err != nil {
		log.Println(err)
		return err
	}
	if err := uploadData(a.url+"/api/rest/v1/attributes/"+attributeCode+"/options", token, options); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (a *akeneo) GenerateAttributeOptionsValueMap() (models.AttributeOptionsValueMapAkeneo, error) {
	rows, err := a.db.Query(`
SELECT pim_catalog_attribute.code, pim_catalog_attribute_option.code, pim_catalog_attribute_option_value.value
FROM pim_catalog_attribute
JOIN pim_catalog_attribute_option ON pim_catalog_attribute_option.attribute_id = pim_catalog_attribute.id 
JOIN pim_catalog_attribute_option_value ON pim_catalog_attribute_option_value.option_id = pim_catalog_attribute_option.id;
`)
	if err != nil {
		return models.AttributeOptionsValueMapAkeneo{}, err
	}
	defer rows.Close()

	result := make(models.AttributeOptionsValueMapAkeneo)

	for rows.Next() {
		var attributeCode string
		var optionCode string
		var optionValue string

		err = rows.Scan(&attributeCode, &optionCode, &optionValue)
		if err != nil {
			return models.AttributeOptionsValueMapAkeneo{}, err
		}

		if _, ok := result[attributeCode]; !ok {
			result[attributeCode] = make(map[string]string)
		}

		result[attributeCode][optionCode] = optionValue
	}

	return result, nil
}

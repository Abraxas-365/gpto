package application

import "github/Abraxas-365/akeneo-connector/pkg/core/models"

func (a *app) SyncAttributeGroups() (models.AttributeGroups, error) {
	attributeGroups, err := a.magento.GetAttributeGroups()
	if err != nil {
		return models.AttributeGroups{}, err
	}

	if err := a.akeneo.UploadAttributeGroups(attributeGroups); err != nil {
		return nil, err
	}

	return attributeGroups, nil

}

package application

import (
	"github/Abraxas-365/akeneo-connector/pkg/core/models"
)

func (a *app) SyncAttributes() (models.Attributes, error) {
	attributes, err := a.magento.GetAttributes()
	if err != nil {
		return models.Attributes{}, err
	}

	if err := a.akeneo.UploadAttributes(attributes); err != nil {
		return models.Attributes{}, err
	}

	return attributes, nil

}

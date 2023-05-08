package akeneo

import (
	"github/Abraxas-365/akeneo-connector/pkg/core/models"
)

func (a *akeneo) UploadCategories(categories models.Categories) error {
	token, err := a.Login()
	if err != nil {
		return err
	}

	if err := uploadData(a.url+"/api/rest/v1/categories", token, categories); err != nil {
		return err
	}

	return nil
}

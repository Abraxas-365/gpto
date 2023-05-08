package akeneo

import (
	"github/Abraxas-365/akeneo-connector/pkg/core/models"
)

func (a *akeneo) UploadFamilies(families models.FamiliesAkeneo) error {
	token, err := a.Login()
	if err != nil {
		return err
	}
	if err := uploadData(a.url+"/api/rest/v1/families", token, families); err != nil {
		return err
	}

	return nil
}

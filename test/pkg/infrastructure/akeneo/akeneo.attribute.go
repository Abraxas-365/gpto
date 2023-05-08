package akeneo

import (
	"github/Abraxas-365/akeneo-connector/pkg/core/models"
)

func (a *akeneo) UploadAttributes(attributes models.Attributes) error {
	token, err := a.Login()
	if err != nil {
		return err
	}
	data := attributes.ToAkeneoStruct()
	if err := uploadData(a.url+"/api/rest/v1/attributes", token, data); err != nil {
		return err
	}

	return nil
}

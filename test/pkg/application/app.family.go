package application

import "github/Abraxas-365/akeneo-connector/pkg/core/models"

func (a *app) SyncFamilies() (models.Families, error) {
	families, err := a.magento.GetFamily()
	if err != nil {
		return models.Families{}, err
	}
	var akeneoFamilies models.FamiliesAkeneo
	for _, elem := range families {
		family := models.NewAkeneoFamilyBuilder()
		akeneoFamilies = append(akeneoFamilies, family.AddCode(elem.Code).
			AddLabels(elem.Labels).
			AddAttributeSetAsLabel(elem.AttributeSetAsLabel).
			AddAttributes(elem.Attributes).
			Build(),
		)
	}

	if err := a.akeneo.UploadFamilies(akeneoFamilies); err != nil {
		return nil, err
	}

	return families, nil

}

func (a *app) GetFamilyNameWithId(id int) (string, error) {
	return a.magento.GetFamilyNameWithId(id)
}

package ports

import "github/Abraxas-365/akeneo-connector/pkg/core/models"

type Akeneo interface {
	Login() (string, error)
	UploadCategories(categories models.Categories) error
	UploadAttributes(attributes models.Attributes) error
	UploadFamilies(families models.FamiliesAkeneo) error
	UploadAttributeGroups(attributeGroups models.AttributeGroups) error
	UploadProducts(products models.ProductsAkeneo) error
	UploadOptions(options models.OptionsAkeneo, attribute string) error
	GetProduct(code int) (models.ProductAkeneo, error)
	GenerateAttributeOptionsValueMap() (models.AttributeOptionsValueMapAkeneo, error)
	GetAttributeGroup(code string) ([]string, error)
}

package ports

import models "github/Abraxas-365/akeneo-connector/pkg/core/models"

type Magento interface {
	GetCategories() (models.Categories, error)
	GetAttributes() (models.Attributes, error)
	GetAttributeType(code string) (string, error)
	GetFamily() (models.Families, error)
	GetOptions() (models.OptionsMagentoResponse, error)
	OptionsInicialLoad() error
	CategoryInicialLoad() error
	GetAttributeGroups() (models.AttributeGroups, error)
	GetFamilyNameWithId(id int) (string, error)
	GetAttributeSetIdWithFamily(family string) (int, error)
	GetProductById(id int) (models.ProductsMagento, error)
	GetProductsByPage(page int, pageSize int) (models.ProductsMagento, error)
	GetCategoriesById(id int) ([]string, error)
	UpdateProduct(code int, product models.ProductMagento) error
	Login() (string, error)
}

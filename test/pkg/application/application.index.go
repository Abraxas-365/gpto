package application

import (
	"github/Abraxas-365/akeneo-connector/pkg/core/models"
	"github/Abraxas-365/akeneo-connector/pkg/core/ports"
	"github/Abraxas-365/akeneo-connector/pkg/core/service"
)

type Application interface {
	SyncCategories(isNew bool) error
	SyncAttributes() (models.Attributes, error)
	SyncFamilies() (models.Families, error)
	SyncOptions() (models.OptionsAkeneo, error)
	SyncAttributeGroups() (models.AttributeGroups, error)
	SyncProductById(akeneoProduct models.ProductAkeneoBuilder, id int) (models.ProductAkeneo, error)
	SyncProducts(akeneoProductBuilder models.ProductAkeneoBuilder, page int) (models.ProductsAkeneo, error)
	GetFamilyNameWithId(id int) (string, error)
	ConectorProduct(code int) (models.ProductMagento, error)
}

type app struct {
	magento ports.Magento
	akeneo  ports.Akeneo
	service service.Service
}

func ApplicationFactory(magento ports.Magento, akeneo ports.Akeneo, service service.Service) Application {
	return &app{
		magento,
		akeneo,
		service,
	}
}

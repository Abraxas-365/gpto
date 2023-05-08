package service

import (
	"github/Abraxas-365/akeneo-connector/pkg/core/models"
	"github/Abraxas-365/akeneo-connector/pkg/core/ports"
)

type Service interface {
	AddAkeneoValues(attributes models.ProductAttributesMagento, builder models.ProductAkeneoBuilder) error
	AddMagentoAttributes(values map[string]models.Values, builder models.ProductMagentoBuilder, additianalGrop []string) error
}

type service struct {
	magento ports.Magento
	akeneo  ports.Akeneo
}

func ServiceFactory(magento ports.Magento, akeneo ports.Akeneo) Service {
	return &service{
		magento: magento,
		akeneo:  akeneo,
	}
}

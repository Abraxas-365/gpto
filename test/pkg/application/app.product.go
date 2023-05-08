package application

import (
	"errors"
	"fmt"
	"github/Abraxas-365/akeneo-connector/pkg/core/models"
	"log"
	"strconv"
)

func (a *app) SyncProductById(akeneoProductBuilder models.ProductAkeneoBuilder, productId int) (models.ProductAkeneo, error) {
	productMagento, err := a.magento.GetProductById(productId)
	if err != nil {
		return models.ProductAkeneo{}, err
	} else if len(productMagento) == 0 {
		return models.ProductAkeneo{}, nil
	}

	familyName, err := a.magento.GetFamilyNameWithId(productMagento[0].FamilyId)
	if err != nil {
		return models.ProductAkeneo{}, nil
	}
	categories, err := a.magento.GetCategoriesById(*productMagento[0].Id)
	if err != nil {
		return models.ProductAkeneo{}, err
	}

	err = a.service.AddAkeneoValues(productMagento[0].CustomAttributes, akeneoProductBuilder)
	if err != nil {
		return models.ProductAkeneo{}, err
	}
	product := akeneoProductBuilder.AddIdentifier(productMagento[0].Sku).
		AddName(productMagento[0].Name).
		AddSku(productMagento[0].Sku).
		AddIsInStock(productMagento[0].ExtensionAttributes.StockItem.IsInStock.ToStringNumber()).
		AddVisibility(strconv.Itoa(productMagento[0].Visibility)).
		AddPrice(productMagento[0].Price).AddStatus(productMagento[0].Status.ToBool()).
		AddCategories(categories).
		AddFamily(familyName).
		Build()

	if err := a.akeneo.UploadProducts(models.ProductsAkeneo{product}); err != nil {
		return models.ProductAkeneo{}, err
	}

	return product, nil
}

func (a *app) SyncProducts(akeneoProductBuilder models.ProductAkeneoBuilder, page int) (models.ProductsAkeneo, error) {
	var sycquedProducts models.ProductsAkeneo

	products, err := a.magento.GetProductsByPage(page, 99)
	if err != nil {
		return models.ProductsAkeneo{}, err
	} else if len(products) == 0 {
		return models.ProductsAkeneo{}, errors.New("No more products")
	}

	for _, elem := range products {

		familyName, err := a.magento.GetFamilyNameWithId(elem.FamilyId)
		if err != nil {
			return models.ProductsAkeneo{}, nil
		}

		categories, err := a.magento.GetCategoriesById(*elem.Id)
		if err != nil {
			return models.ProductsAkeneo{}, err
		}

		akeneoProduct := akeneoProductBuilder.AddIdentifier("sku").
			AddName(elem.Name).
			AddVisibility(strconv.Itoa(elem.Visibility)).
			AddSku(elem.Sku).
			AddPrice(elem.Price).AddStatus(elem.Status.ToBool()).
			AddCategories(categories).
			AddName(familyName)
		err = a.service.AddAkeneoValues(elem.CustomAttributes, akeneoProduct)
		if err != nil {
			return models.ProductsAkeneo{}, err
		}

		sycquedProducts = append(sycquedProducts, akeneoProduct.Build())
	}

	if err := a.akeneo.UploadProducts(sycquedProducts); err != nil {
		return models.ProductsAkeneo{}, err
	}
	return sycquedProducts, nil
}

func (a *app) ConectorProduct(code int) (models.ProductMagento, error) {
	akeneoProduct, err := a.akeneo.GetProduct(code)
	if err != nil {
		return models.ProductMagento{}, err
	}

	magentoProductBuilder := models.NewProductMagentoBuilder().
		AddStatus(akeneoProduct.Enabled.ToInt()).
		AddSku(akeneoProduct.Identifier)

	magetoFamily, err := a.magento.GetAttributeSetIdWithFamily(akeneoProduct.Family)
	if err != nil {
		return models.ProductMagento{}, err
	}
	magentoProductBuilder.AddAttributeSet(magetoFamily)

	//Get attributeGroup
	attributeGroup, err := a.akeneo.GetAttributeGroup("test")
	if err != nil {
		log.Println("attributeGroup")
		return models.ProductMagento{}, err
	}

	if err := a.service.AddMagentoAttributes(akeneoProduct.Values, magentoProductBuilder, attributeGroup); err != nil {
		return models.ProductMagento{}, err
	}
	magentoProduct := magentoProductBuilder.Build()
	fmt.Println(magentoProduct)

	if err := a.magento.UpdateProduct(code, magentoProduct); err != nil {
		log.Println(err)
		return models.ProductMagento{}, err
	}

	return magentoProduct, nil
}

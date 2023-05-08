package application

import (
	"github/Abraxas-365/akeneo-connector/pkg/core/models"
)

func (a *app) SyncOptions() (models.OptionsAkeneo, error) {

	var akeneoOptions models.OptionsAkeneo

	if err := a.magento.OptionsInicialLoad(); err != nil {
		return models.OptionsAkeneo{}, err
	}

	// magentoOptions, err := a.magento.GetOptions()
	// if err != nil {
	// 	return models.OptionsAkeneo{}, err
	// }
	// var count int
	// for _, element := range magentoOptions {
	// 	if element.IsVisible == true {
	// 		for _, option := range element.Options {
	// 			count++
	// 			akeneoOptionBuilder := models.NewOptionAkeneoBuilder()
	// 			akeneoOption := akeneoOptionBuilder.AddCode(option.Value).
	// 				AddLabels(option.Label).
	// 				AddAttribute(element.AttributeCode).
	// 				AddSortOrder(count).Build()
	// 			akeneoOptions.AddOption(akeneoOption)
	// 		}
	//
	// 	}
	//
	// }
	//
	// attributes, err := a.magento.GetAttributes()
	// if err != nil {
	// 	return models.OptionsAkeneo{}, err
	// }
	//
	// for _, attribute := range attributes {
	// 	//aqui puedo hacer una optimisacion para que no recorra todo el array
	// 	var options models.OptionsAkeneo
	// 	for _, option := range akeneoOptions {
	// 		if option.Attribute == attribute.Code {
	// 			options.AddOption(option)
	// 		}
	// 	}
	// 	if len(options) > 0 {
	// 		fmt.Println("attribute", attribute.Code)
	// 		if err := a.akeneo.UploadOptions(options, attribute.Code); err != nil {
	// 			return models.OptionsAkeneo{}, err
	// 		}
	// 		time.Sleep(15000 * time.Millisecond)
	// 	}
	// }
	//
	return akeneoOptions, nil
}

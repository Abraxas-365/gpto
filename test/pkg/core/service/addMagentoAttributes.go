package service

import (
	"encoding/json"
	"github/Abraxas-365/akeneo-connector/pkg/core/models"
	"log"
	"strconv"
)

func (s *service) AddMagentoAttributes(values map[string]models.Values, builder models.ProductMagentoBuilder, additionalGrop []string) error {
	/* json con los atributos de extras */
	attributesJson := models.ProductAttributesMagento{}
	/*mapa con el mapeo de las opciones que traje de la base de datos*/
	optionMap, err := s.akeneo.GenerateAttributeOptionsValueMap()
	if err != nil {
		return err
	}
	/*preparar el json con attributos adicionales*/
	for _, element := range additionalGrop {
		attributeType, err := s.magento.GetAttributeType(element)
		if err != nil {
			return err
		}
		if value, ok := values[element]; ok {
			data := s.getDataAcordingToType(attributeType, element, optionMap, value[0].Data)
			attributesJson = append(attributesJson, models.NewProductAttributeMagentoBuilder().AddAttributeCode(element).AddValue(data).Build())
		}
		//eliminar key del map para que no se repita en los demas attributos
		delete(values, element)
	}

	/*Recorrer el set attributeValues de akeneo*/
	for k, v := range values {

		switch k {
		/*Casos fijos que van en el header del json que pide magento*/
		case "name":
			builder.AddName(v[0].Data.(string))

		case "price":
			priceInt, err := strconv.ParseFloat(v[0].Data.([]any)[0].(map[string]any)["amount"].(string), 32)
			if err != nil {
				log.Println("Precio no es string", err)
			}
			builder.AddPrice(int(priceInt))

		case "visibility":
			visibilityStr := v[0].Data.(string)
			visibilityInt, err := strconv.Atoi(visibilityStr)
			if err != nil {
				log.Println("visibility no es string")
			}
			builder.AddVisibility(visibilityInt)

		default:
			/*Casos de los attributos que van al conjunto de attributos que pide magento*/
			attributeType, err := s.magento.GetAttributeType(k)
			if err != nil {
				return err
			}
			data := s.getDataAcordingToType(attributeType, k, optionMap, v[0].Data)

			/*meter el product*/
			builder.AddAttributes(
				models.NewProductAttributeMagentoBuilder().
					AddAttributeCode(k).
					AddValue(data).
					Build(),
			)

		}
	}

	//enviar todos los attribute como uno solo
	jsonString, err := json.Marshal(attributesJson)
	builder.AddAttributes(models.NewProductAttributeMagentoBuilder().
		AddAttributeCode("bottom_text_pdp"). //nombre del attributo que tendra el html
		AddValue(string(jsonString)).
		Build())

	return nil
}

func (s *service) getDataAcordingToType(attributeType string,
	attCode string,
	optionMap models.AttributeOptionsValueMapAkeneo,
	data any) any {
	if attributeType == "pim_catalog_price" || attributeType == "pim_catalog_price_collection" {
		return data.([]any)[0].(map[string]any)["amount"].(string)

	} else if attributeType == "pim_catalog_boolean" {
		return models.StatusBool(data.(bool)).ToStringNumber()

	} else if attCode == "long_description" && (attributeType == "pim_catalog_simpleselect" ||
		attributeType == "pim_catalog_multiselect") {
		return optionMap[attCode][data.(string)]
	} else {
		return data
	}
}

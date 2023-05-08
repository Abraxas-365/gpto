package service

import "github/Abraxas-365/akeneo-connector/pkg/core/models"

func (s *service) AddAkeneoValues(attributes models.ProductAttributesMagento, builder models.ProductAkeneoBuilder) error {
	for _, element := range attributes {

		valueType, err := s.magento.GetAttributeType(element.AttributeCode)
		if err != nil {
			return err
		}
		switch {
		case element.AttributeCode == "category_ids" ||
			element.AttributeCode == "gallery" ||
			element.AttributeCode == "news_from_date" ||
			element.AttributeCode == "news_to_date" ||
			element.AttributeCode == "status" ||
			element.AttributeCode == "price_view" ||
			element.AttributeCode == "special_from_date" ||
			element.AttributeCode == "special_to_date" ||
			element.AttributeCode == "card_price_from" ||
			element.AttributeCode == "card_price_to" ||
			element.AttributeCode == "msrp_display_actual_price_type" ||
			element.AttributeCode == "cost" ||
			element.AttributeCode == "tier_price" ||
			element.AttributeCode == "msrp" ||
			element.AttributeCode == "custom_design_from" ||
			element.AttributeCode == "custom_design_to" ||
			element.AttributeCode == "custom_layout" ||
			element.AttributeCode == "shipment_type" ||
			element.AttributeCode == "email_template" ||
			element.AttributeCode == "is_redeemable" ||
			element.AttributeCode == "use_config_is_redeemable" ||
			element.AttributeCode == "lifetime" ||
			element.AttributeCode == "use_config_lifetime" ||
			element.AttributeCode == "use_config_email_template" ||
			element.AttributeCode == "allow_message" ||
			element.AttributeCode == "use_config_allow_message" ||
			element.AttributeCode == "minimal_price" ||
			element.AttributeCode == "image_label" ||
			element.AttributeCode == "required_options" ||
			element.AttributeCode == "small_image_label" ||
			element.AttributeCode == "url_path" ||
			element.AttributeCode == "samples_title" ||
			element.AttributeCode == "links_title" ||
			element.AttributeCode == "computacion" ||
			element.AttributeCode == "thumbnail_label" ||
			element.AttributeCode == "has_options":
			break

		default:
			builder.AddValues(element.AttributeCode, element.Value, valueType)

		}
	}
	return nil
}

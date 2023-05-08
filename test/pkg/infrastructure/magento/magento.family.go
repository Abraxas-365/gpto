package magento

import (
	"github/Abraxas-365/akeneo-connector/pkg/core/models"
	"log"
	"strconv"
)

func (m *magnetoRepo) GetFamily() (models.Families, error) {

	rows, err := m.db.Query(`
Select Distinct 
    snakeCase(REPLACE(TRIM(SUBSTRING_INDEX(eas.attribute_set_name, '(', 1)), 'Pim', '')) as 'code', 
    REPLACE(TRIM(SUBSTRING_INDEX(eas.attribute_set_name, '(', 1)), 'Pim', '') as 'label-es_CL',
    'sku' as 'attribute_as_label',
    GROUP_CONCAT(ea.attribute_code SEPARATOR ',') as 'attributes'
From eav_attribute_set eas 
Inner Join eav_entity_attribute eea 
    On eea.attribute_set_id = eas.attribute_set_id 
Inner Join eav_attribute ea 
    ON ea.attribute_id = eea.attribute_id 
    And ea.entity_type_id = (Select eet.entity_type_id From eav_entity_type eet Where eet.entity_type_code = 'catalog_product')
    And (ea.backend_type != 'static' || ea.attribute_code = 'sku')
    And ea.attribute_code not in (
        'gallery', 'news_from_date', 'news_to_date', 'status', 'price_view', 'special_from_date', 
        'special_to_date', 'card_price_from', 'card_price_to', 'msrp_display_actual_price_type',
        'cost', 'tier_price', 'msrp', 'custom_design_from', 'custom_design_to', 'custom_layout',
        'shipment_type','email_template','is_redeemable','use_config_is_redeemable','lifetime',
        'use_config_lifetime','use_config_email_template','allow_message','use_config_allow_message',
        'minimal_price'
    )
Inner Join catalog_eav_attribute cea 
    On cea.attribute_id = ea.attribute_id 
    And cea.is_visible = 1
Where eas.entity_type_id = (Select eet.entity_type_id From eav_entity_type eet Where eet.entity_type_code = 'catalog_product')
Group By eas.attribute_set_name
Order By eas.attribute_set_id;
`)

	if err != nil {
		return models.Families{}, err
	}

	families := models.Families{}
	for rows.Next() {
		var family models.Family
		err = rows.Scan(&family.Code, &family.Labels.EsCl, &family.AttributeSetAsLabel, &family.Attributes)
		if err != nil {
			return models.Families{}, err
		}
		families = append(families, family)
	}

	return families, nil
}

func (m *magnetoRepo) GetFamilyNameWithId(id int) (string, error) {
	var name *string

	rows, err := m.db.Query(`Select snakeCase(REPLACE(TRIM(SUBSTRING_INDEX(eas.attribute_set_name, '(', 1)), 'Pim', '')) as 'family' 
From eav_attribute_set eas 
Where eas.entity_type_id = (Select eet.entity_type_id From eav_entity_type eet Where eet.entity_type_code = 'catalog_product')
        And eas.attribute_set_id = ` + strconv.Itoa(id))
	for rows.Next() {
		err = rows.Scan(&name)
		if err != nil {
			log.Println(err)
			return "", err
		}
	}
	return *name, nil
}

func (m *magnetoRepo) GetAttributeSetIdWithFamily(family string) (int, error) {
	var id int

	rows, err := m.db.Query(`
Select eas.attribute_set_id 
From eav_attribute_set eas 
Where eas.entity_type_id = (Select eet.entity_type_id 
From eav_entity_type eet 
Where eet.entity_type_code = 'catalog_product') 
And snakeCase(REPLACE(TRIM(SUBSTRING_INDEX(eas.attribute_set_name, '(', 1)), 'Pim', '')) = ` + "'" + family + "'")
	log.Println(family)
	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			log.Println(err)
			return 0, err
		}
	}

	return id, nil
}

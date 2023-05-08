package magento

import (
	"fmt"
	"github/Abraxas-365/akeneo-connector/pkg/core/models"
)

func (m *magnetoRepo) GetAttributes() (models.Attributes, error) {

	attributes := models.Attributes{}
	// INPUT TEXT
	rows, err := m.db.Query(`
Select *
From (
	Select Distinct 
		ea.attribute_code as 'code',    
		ifnull(ea.frontend_label, properCase(ea.attribute_code)) as 'label-es_CL',    
		replace(ifnull(min(eag.attribute_group_code), 'other'), '-', '_') as 'group', 
		0 as 'localizable',    
		255 as 'max_characters', 
		case when cea.is_global = 1 then 0 else 1 end as 'scopable',
		cea.position as 'sort_order',    
		'pim_catalog_text' as 'type',   
		ea.is_unique as 'unique',    
		cea.is_filterable_in_grid as 'useable_as_grid_filter'
	From eav_attribute ea 
	Left Join catalog_eav_attribute cea 
	    On cea.attribute_id = ea.attribute_id 
	Left Join eav_entity_attribute eea  
	    On eea.attribute_id = ea.attribute_id 
	Left Join eav_attribute_group eag 
	    On eag.attribute_group_id = eea.attribute_group_id 
	Where ea.entity_type_id = (
		Select eet.entity_type_id From eav_entity_type eet Where eet.entity_type_code = 'catalog_product'
	)
	And (
		(ea.backend_type = 'varchar' and ea.frontend_input in ('text', 'textarea')) Or    
		(ea.backend_type = 'varchar' and ea.frontend_input is null) Or 
		(ea.backend_type = 'text' and  ea.frontend_input = 'text')
	) 
	And ea.attribute_code != 'sku' # El sku viene por defecto en akeneo y es de tipo "Identifier"
	And cea.is_visible = 1
	Group by ea.attribute_code
	Order By ea.attribute_id
) as foo
Where foo.group not in ('advanced_pricing', 'bundle_items', 'schedule_design_update')
`)
	if err != nil {
		return models.Attributes{}, err
	}

	inputText := models.Attributes{}
	for rows.Next() {
		var attribute models.Attribute
		err = rows.Scan(
			&attribute.Code,
			&attribute.Labels.EsCl,
			&attribute.Group,
			&attribute.Localizable,
			&attribute.MaxCharacters,
			&attribute.Scopable,
			&attribute.SortOrder,
			&attribute.Type,
			&attribute.Unique,
			&attribute.UseableAsGridFilter)
		if err != nil {
			fmt.Println("inputText")
			return models.Attributes{}, err
		}
		inputText = append(inputText, attribute)
	}

	//TEXT AREA
	rows, err = m.db.Query(`
Select *
From (
	Select Distinct 
		ea.attribute_code as 'code',    
		ifnull(ea.frontend_label, properCase(ea.attribute_code)) as 'label-es_CL',    
		replace(ifnull(min(eag.attribute_group_code), 'other'), '-', '_') as 'group', 
		0 as 'localizable',
		case when cea.is_global = 1 then 0 else 1 end as 'scopable',
		cea.position as 'sort_order',    
		'pim_catalog_textarea' as 'type',   
		ea.is_unique as 'unique',    
		cea.is_filterable_in_grid as 'useable_as_grid_filter',
		1 as 'wysiwyg_enabled'
	From eav_attribute ea 
	Left Join catalog_eav_attribute cea 
	    On cea.attribute_id = ea.attribute_id 
	Left Join eav_entity_attribute eea  
	    On eea.attribute_id = ea.attribute_id 
	Left Join eav_attribute_group eag 
	    On eag.attribute_group_id = eea.attribute_group_id 
	Where ea.entity_type_id = (
		Select eet.entity_type_id From eav_entity_type eet Where eet.entity_type_code = 'catalog_product'
	)
	And ea.backend_type = 'text'
	And ea.frontend_input = 'textarea'
	And cea.is_visible = 1
	Group by ea.attribute_code
	Order By ea.attribute_id
) as foo
Where foo.group not in ('advanced_pricing', 'bundle_items', 'schedule_design_update')
`)

	if err != nil {
		fmt.Println("textArea")
		return models.Attributes{}, err
	}

	textArea := models.Attributes{}
	for rows.Next() {
		var attribute models.Attribute
		err = rows.Scan(
			&attribute.Code,
			&attribute.Labels.EsCl,
			&attribute.Group,
			&attribute.Localizable,
			&attribute.Scopable,
			&attribute.SortOrder,
			&attribute.Type,
			&attribute.Unique,
			&attribute.UseableAsGridFilter,
			&attribute.WisingEnable,
		)
		if err != nil {
			fmt.Println("textArea")
			return models.Attributes{}, err
		}
		textArea = append(textArea, attribute)
	}

	//Numbers
	rows, err = m.db.Query(`
Select *
From (
	Select Distinct 
		ea.attribute_code as 'code',    
		ifnull(ea.frontend_label, properCase(ea.attribute_code)) as 'label-es_CL',    
		case when ea.backend_type = 'decimal' then 1 else 0 end as 'decimals_allowed',
		replace(ifnull(min(eag.attribute_group_code), 'other'), '-', '_') as 'group', 
		0 as 'localizable',
		case when cea.is_global = 1 then 0 else 1 end as 'scopable',
		0 as 'negative_allowed',
		cea.position as 'sort_order',    
		'pim_catalog_number' as 'type',   
		ea.is_unique as 'unique',    
		cea.is_filterable_in_grid as 'useable_as_grid_filter'
	From eav_attribute ea 
	Left Join catalog_eav_attribute cea 
	    On cea.attribute_id = ea.attribute_id 
	Left Join eav_entity_attribute eea  
	    On eea.attribute_id = ea.attribute_id 
	Left Join eav_attribute_group eag 
	    On eag.attribute_group_id = eea.attribute_group_id 
	Where ea.entity_type_id = (
		Select eet.entity_type_id From eav_entity_type eet Where eet.entity_type_code = 'catalog_product'
	)
	And (
		(ea.backend_type = 'int' and  ea.frontend_input = 'text') Or 
		(ea.backend_type = 'int' and  ea.frontend_input is null) Or
	 	(ea.backend_type = 'decimal' And ea.frontend_input = 'weight')
	)
	And cea.is_visible = 1
	Group by ea.attribute_code
	Order By ea.attribute_id
) as foo
Where foo.group not in ('advanced_pricing', 'bundle_items', 'schedule_design_update')
`)
	if err != nil {
		fmt.Println("Numbers")
		return models.Attributes{}, err
	}

	numbers := models.Attributes{}
	for rows.Next() {
		var attribute models.Attribute
		err = rows.Scan(
			&attribute.Code,
			&attribute.Labels.EsCl,
			&attribute.DecimalAllowed,
			&attribute.Group,
			&attribute.Localizable,
			&attribute.Scopable,
			&attribute.NegativeAllowed,
			&attribute.SortOrder,
			&attribute.Type,
			&attribute.Unique,
			&attribute.UseableAsGridFilter,
		)
		if err != nil {
			fmt.Println("Numbers")
			return models.Attributes{}, err
		}
		numbers = append(numbers, attribute)
	}

	//PRICES
	rows, err = m.db.Query(`
Select *
From (
	Select Distinct 
		ea.attribute_code as 'code',    
		ifnull(ea.frontend_label, properCase(ea.attribute_code)) as 'label-es_CL',    
		1 as 'decimals_allowed',
		case 
			when ea.attribute_code = 'special_price' then 'product_details'
			else replace(ifnull(min(eag.attribute_group_code), 'other'), '-', '_') 
		end as 'group', 
		0 as 'localizable',     
		case when cea.is_global = 1 then 0 else 1 end as 'scopable',
		cea.position as 'sort_order',    
		'pim_catalog_price_collection' as 'type',   
		ea.is_unique as 'unique', 
		cea.is_filterable_in_grid as 'useable_as_grid_filter'
	From eav_attribute ea 
	Left Join catalog_eav_attribute cea 
	    On cea.attribute_id = ea.attribute_id 
	Left Join eav_entity_attribute eea  
	    On eea.attribute_id = ea.attribute_id 
	Left Join eav_attribute_group eag 
	    On eag.attribute_group_id = eea.attribute_group_id 
	Where ea.entity_type_id = (
		Select eet.entity_type_id From eav_entity_type eet Where eet.entity_type_code = 'catalog_product'
	)
	And ea.backend_type = 'decimal' 
	And ea.frontend_input in ('price', 'text')
	And cea.is_visible = 1
	Group by ea.attribute_code
	Order By ea.attribute_id
) as foo
Where foo.group not in ('advanced_pricing', 'bundle_items', 'schedule_design_update')
`)
	if err != nil {
		fmt.Println("Prices")
		return models.Attributes{}, err
	}

	prices := models.Attributes{}
	for rows.Next() {
		var attribute models.Attribute
		err = rows.Scan(
			&attribute.Code,
			&attribute.Labels.EsCl,
			&attribute.DecimalAllowed,
			&attribute.Group,
			&attribute.Localizable,
			&attribute.Scopable,
			&attribute.SortOrder,
			&attribute.Type,
			&attribute.Unique,
			&attribute.UseableAsGridFilter,
		)
		if err != nil {
			fmt.Println("Prices")
			return models.Attributes{}, err
		}
		prices = append(prices, attribute)
	}

	//OTHERS
	rows, err = m.db.Query(`
Select *
From (
	Select Distinct 
		ea.attribute_code as 'code',
		ifnull(ea.frontend_label, properCase(ea.attribute_code)) as 'label-es_CL',
		case 
			when ea.frontend_input in ('gallery', 'media_image', 'image') then 'image_management'
			when ea.attribute_code = 'custom_design' then 'design' 
			else replace(ifnull(min(eag.attribute_group_code), 'other'), '-', '_') 
		end as 'group', 
		0 as 'localizable',
		case when cea.is_global = 1 then 0 else 1 end as 'scopable',
		cea.position as 'sort_order',
		ea.is_unique as 'unique',
		cea.is_filterable_in_grid as 'useable_as_grid_filter',
		case 
			when ea.backend_type = 'datetime' or ea.frontend_input = 'date' then 'pim_catalog_date'
			when ea.frontend_input = 'boolean' then 'pim_catalog_boolean'
			when ea.frontend_input = 'select' then 'pim_catalog_simpleselect'
			when ea.frontend_input = 'multiselect' then 'pim_catalog_multiselect'
			when ea.frontend_input in ('gallery', 'media_image', 'image') then 'pim_catalog_image'
		end as 'type'
	From eav_attribute ea 
	Left Join catalog_eav_attribute cea 
	    On cea.attribute_id = ea.attribute_id 
	Left Join eav_entity_attribute eea  
	    On eea.attribute_id = ea.attribute_id 
	Left Join eav_attribute_group eag 
	    On eag.attribute_group_id = eea.attribute_group_id 
	Where ea.entity_type_id = (
		Select eet.entity_type_id From eav_entity_type eet Where eet.entity_type_code = 'catalog_product'
	)
	And (
		ea.backend_type in ('datetime') OR 
		ea.frontend_input in ('date', 'boolean', 'select', 'multiselect', 'gallery', 'media_image', 'image')
	)
	And ea.attribute_code not in (
		'gallery', 'news_from_date', 'news_to_date', 'status', 'price_view', 'special_from_date', 'special_to_date', 
		'card_price_from', 'card_price_to', 'msrp_display_actual_price_type', 'custom_design_from', 'custom_design_to', 
		'custom_layout', 'shipment_type'
	)
	And cea.is_visible = 1
	Group by ea.attribute_code
	Order By ea.attribute_id
) as foo
Where foo.group not in ('advanced_pricing', 'bundle_items', 'schedule_design_update')
`)
	if err != nil {
		fmt.Println("other")
		return models.Attributes{}, err
	}

	others := models.Attributes{}
	for rows.Next() {
		var attribute models.Attribute

		err = rows.Scan(
			&attribute.Code,
			&attribute.Labels.EsCl,
			&attribute.Group,
			&attribute.Localizable,
			&attribute.Scopable,
			&attribute.SortOrder,
			&attribute.Unique,
			&attribute.UseableAsGridFilter,
			&attribute.Type,
		)
		if err != nil {
			fmt.Println("other")
			return models.Attributes{}, err
		}
		others = append(others, attribute)
	}
	attributes = append(attributes, inputText...)
	attributes = append(attributes, textArea...)
	attributes = append(attributes, numbers...)
	attributes = append(attributes, prices...)
	attributes = append(attributes, others...)
	return attributes, nil
}

func (m *magnetoRepo) GetAttributeType(code string) (string, error) {
	rows, err := m.db.Query(`Select Distinct
	case 
		when ea.backend_type = 'datetime' Or ea.frontend_input = 'date' then 'pim_catalog_date'
		when ea.frontend_input = 'boolean' then 'pim_catalog_boolean'
		when ea.frontend_input = 'select' then 'pim_catalog_simpleselect'
		when ea.frontend_input = 'multiselect' then 'pim_catalog_multiselect'
		when ea.frontend_input in ('gallery', 'media_image', 'image') then 'pim_catalog_image'
		when ea.backend_type = 'decimal' And ea.frontend_input in ('price', 'text') then 'pim_catalog_price_collection'
		when (ea.backend_type = 'int' and  ea.frontend_input = 'text') Or 
			(ea.backend_type = 'int' and  ea.frontend_input is null) Or
		 	(ea.backend_type = 'decimal' And ea.frontend_input = 'weight') then 'pim_catalog_number'
	 	When ea.backend_type = 'text'And ea.frontend_input = 'textarea' then 'pim_catalog_textarea'
	 	When (ea.backend_type = 'varchar' and ea.frontend_input in ('text', 'textarea')) Or    
			(ea.backend_type = 'varchar' and ea.frontend_input is null) Or 
			(ea.backend_type = 'text' and  ea.frontend_input = 'text') then 'pim_catalog_text'
	end as 'type'
From eav_attribute ea 
Where ea.attribute_code = '` + code + " ' ")
	if err != nil {
		return "", err
	}

	var attributeType any
	for rows.Next() {

		err = rows.Scan(
			&attributeType,
		)
		if err != nil {
			return "", err
		}
	}
	var resp string
	if attributeType == nil {
		resp = ""
	} else {
		resp = string(attributeType.([]uint8))
	}
	return resp, nil
}

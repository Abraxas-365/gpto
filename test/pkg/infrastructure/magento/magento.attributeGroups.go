package magento

import (
	"github/Abraxas-365/akeneo-connector/pkg/core/models"
)

func (m *magnetoRepo) GetAttributeGroups() (models.AttributeGroups, error) {

	rows, err := m.db.Query(`
Select Distinct 
    replace(attribute_group_code, '-', '_') as 'code', 
    min(sort_order) as 'sort_order', 
    attribute_group_name as 'label-es_CL'
From eav_attribute_group eag
Where replace(attribute_group_code, '-', '_') not in ('advanced_pricing', 'bundle_items', 'schedule_design_update')
Group By attribute_group_code
Order By sort_order;
`)

	if err != nil {
		return models.AttributeGroups{}, err
	}

	attributeGroups := models.AttributeGroups{}
	for rows.Next() {
		var attributeGroup models.AttributeGroup
		err = rows.Scan(&attributeGroup.Code, &attributeGroup.SortOrder, &attributeGroup.Labels.EsCl)
		if err != nil {
			return models.AttributeGroups{}, err
		}
		attributeGroups = append(attributeGroups, attributeGroup)
	}

	return attributeGroups, nil
}

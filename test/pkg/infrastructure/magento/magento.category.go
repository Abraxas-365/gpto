package magento

import (
	"github/Abraxas-365/akeneo-connector/pkg/core/models"
	"strconv"
	"strings"
)

func (m *magnetoRepo) GetCategories() (models.Categories, error) {

	rows, err := m.db.Query(`
Select Distinct 
    case cce.level
        when 1 then 'abcdin' 
        else CONCAT(snakeCase(ccev.value), '_', cce.entity_id) 
    end as code,
    case cce2.level 
        when 1 then 'abcdin'
        else CONCAT(snakeCase(ccev2.value), '_', cce2.entity_id)
    end as parent,
    trim(ccev.value) as 'label-es_CL'
From catalog_category_entity cce
Inner Join catalog_category_entity_varchar ccev
    On ccev.row_id = cce.row_id
    And ccev.attribute_id = getAttributeId('name', 'catalog_category')
    And ccev.store_id = 0
Left Join catalog_category_entity cce2
ON cce2.entity_id = cce.parent_id
Left Join catalog_category_entity_varchar ccev2
    On ccev2.row_id = cce2.row_id
    And ccev2.attribute_id = getAttributeId('name', 'catalog_category')
    And ccev2.store_id = 0
Where cce.level >= 2 # Root, ABCDIN
Order By cce.level, cce2.entity_id, cce.entity_id ASC;
        `)
	if err != nil {
		return models.Categories{}, err
	}
	categories := models.Categories{}
	for rows.Next() {
		var category models.Category
		err = rows.Scan(&category.Code, &category.Parent, &category.Labels.EsCl)
		if err != nil {
			return models.Categories{}, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (m *magnetoRepo) CategoryInicialLoad() error {
	if _, err := m.db.Exec(`INSERT into akeneo_connector_entities (import,code,entity_id,created_at)
Select Distinct 'category',case cce.level when 1 then 'abcdin' else
CONCAT(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REGEXP_REPLACE(REGEXP_REPLACE(LCASE(trim(ccev.value)), '[,|.|-]', ' '), '["|(|)|´]', ''),' ','_'),' ','_'),'__','_'),'á','a'),'é','e'),'í','i'),'ó','o'),'ú','u'),'ñ','n'),'&','y'),'_',cce.entity_id) end as code,cce.entity_id,NOW()From catalog_category_entity cce
Inner Join catalog_category_entity_varchar ccev
On ccev.row_id = cce.row_id
And ccev.attribute_id = (select attribute_id From eav_attribute where attribute_code = 'name' and entity_type_id = 3) and ccev.store_id = 0
Left Join catalog_category_entity cce2
ON cce2.entity_id = cce.parent_id
Left Join catalog_category_entity_varchar ccev2
On ccev2.row_id = cce2.row_id
And ccev2.attribute_id = (select attribute_id From eav_attribute where attribute_code = 'name' and entity_type_id = 3) and ccev2.store_id = 0
Where cce.level >=2
Order By cce.level, cce2.entity_id, cce.entity_id ASC;

`); err != nil {
		return err
	}

	return nil
}

func (m *magnetoRepo) GetCategoriesById(id int) ([]string, error) {
	rows, err := m.db.Query(`Select getCategories(` + strconv.Itoa(id) + ")")
	if err != nil {
		return []string{}, err
	}

	var categories string
	for rows.Next() {
		err = rows.Scan(&categories)
		if err != nil {
			return []string{}, err
		}
	}

	return strings.Split(categories, ","), nil
}

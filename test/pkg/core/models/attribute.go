package models

type Attribute struct {
	Code                string `json:"code"`
	Labels              Labels `json:"labels"`
	Group               string `json:"group"`
	Localizable         *int   `json:"localizable"`
	MaxCharacters       *int   `json:"max_characters"`
	Scopable            *int   `json:"scopable"`
	SortOrder           int    `json:"sort_order"`
	Type                string `json:"type"`
	Unique              *int   `json:"unique"`
	UseableAsGridFilter *int   `json:"useable_as_grid_filter"`
	WisingEnable        *int   `json:"wysiwyg_enabled"`
	DecimalAllowed      *int   `json:"decimals_allowed"`
	NegativeAllowed     *int   `json:"negative_allowed"`
}

type Attributes []Attribute

type AttributeAkeneo struct {
	Code                string `json:"code"`
	Labels              Labels `json:"labels"`
	Group               string `json:"group"`
	Localizable         any    `json:"localizable"`
	MaxCharacters       any    `json:"max_characters"`
	Scopable            any    `json:"scopable"`
	SortOrder           int    `json:"sort_order"`
	Type                string `json:"type"`
	Unique              any    `json:"unique"`
	UseableAsGridFilter any    `json:"useable_as_grid_filter"`
	WisingEnable        any    `json:"wysiwyg_enabled"`
	DecimalAllowed      any    `json:"decimals_allowed"`
	NegativeAllowed     any    `json:"negative_allowed"`
}
type AttributesAkeneo []AttributeAkeneo

func (a *Attributes) ToAkeneoStruct() AttributesAkeneo {
	attributesAkeneo := AttributesAkeneo{}
	for _, element := range *a {
		attributeAkeneo := AttributeAkeneo{
			Code:          element.Code,
			Labels:        element.Labels,
			Group:         element.Group,
			SortOrder:     element.SortOrder,
			Type:          element.Type,
			MaxCharacters: element.MaxCharacters,
		}
		if *element.Localizable == 1 {
			attributeAkeneo.Localizable = true
		} else if *element.Localizable == 0 {
			attributeAkeneo.Localizable = false
		}
		//
		if *element.Scopable == 1 {
			attributeAkeneo.Scopable = true
		} else if *element.Scopable == 0 {
			attributeAkeneo.Scopable = false
		}
		//
		if *element.Unique == 1 {
			attributeAkeneo.Unique = true
		} else if *element.Unique == 0 {
			attributeAkeneo.Unique = false
		}
		//
		if *element.UseableAsGridFilter == 1 {
			attributeAkeneo.UseableAsGridFilter = true
		} else if *element.UseableAsGridFilter == 0 {
			attributeAkeneo.UseableAsGridFilter = false
		}

		if element.WisingEnable != nil && *element.WisingEnable == 1 {
			attributeAkeneo.WisingEnable = true
		} else if element.WisingEnable != nil && *element.WisingEnable == 0 {
			attributeAkeneo.WisingEnable = false
		}

		if element.DecimalAllowed != nil && *element.DecimalAllowed == 1 {
			attributeAkeneo.DecimalAllowed = true
		} else if element.DecimalAllowed != nil && *element.DecimalAllowed == 0 {
			attributeAkeneo.DecimalAllowed = false
		}

		if element.NegativeAllowed != nil && *element.NegativeAllowed == 1 {
			attributeAkeneo.NegativeAllowed = true
		} else if element.NegativeAllowed != nil && *element.NegativeAllowed == 0 {
			attributeAkeneo.NegativeAllowed = false
		}

		attributesAkeneo = append(attributesAkeneo, attributeAkeneo)
	}
	return attributesAkeneo
}

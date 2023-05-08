package models

import (
	"regexp"
	"strings"
)

type OptionMagentoResponse struct {
	IsWysiwygEnabled          bool          `json:"is_wysiwyg_enabled"`
	IsHTMLAllowedOnFront      bool          `json:"is_html_allowed_on_front"`
	UsedForSortBy             bool          `json:"used_for_sort_by"`
	IsFilterable              bool          `json:"is_filterable"`
	IsFilterableInSearch      bool          `json:"is_filterable_in_search"`
	IsUsedInGrid              bool          `json:"is_used_in_grid"`
	IsVisibleInGrid           bool          `json:"is_visible_in_grid"`
	IsFilterableInGrid        bool          `json:"is_filterable_in_grid"`
	Position                  int           `json:"position"`
	ApplyTo                   []string      `json:"apply_to"`
	IsSearchable              string        `json:"is_searchable"`
	IsVisibleInAdvancedSearch string        `json:"is_visible_in_advanced_search"`
	IsComparable              string        `json:"is_comparable"`
	IsUsedForPromoRules       string        `json:"is_used_for_promo_rules"`
	IsVisibleOnFront          string        `json:"is_visible_on_front"`
	UsedInProductListing      string        `json:"used_in_product_listing"`
	IsVisible                 bool          `json:"is_visible"` //importante
	Scope                     string        `json:"scope"`
	AttributeID               int           `json:"attribute_id"`
	AttributeCode             string        `json:"attribute_code"` //importante
	FrontendInput             string        `json:"frontend_input"`
	EntityTypeID              string        `json:"entity_type_id"`
	IsRequired                bool          `json:"is_required"`
	Options                   OptionsMageto `json:"options"` //importante
	IsUserDefined             bool          `json:"is_user_defined"`
	DefaultFrontendLabel      string        `json:"default_frontend_label"`
	BackendType               string        `json:"backend_type"`
	SourceModel               string        `json:"source_model"`
	IsUnique                  string        `json:"is_unique"`
	// FrontendLabels            []struct {
	// 	StoreID int    `json:"store_id"`
	// 	Label   string `json:"label"`
	// } `json:"frontend_labels"`
}
type OptionsMagentoResponse []OptionMagentoResponse

type OptionMageto struct {
	Label string `json:"label"`
	Value string `json:"value"`
}
type OptionsMageto []OptionMageto

type AttributeOptionsValueMapAkeneo map[string]map[string]string

type OptionAkeneo struct {
	Code      string `json:"code"`
	Labels    Labels `json:"labels"`
	Attribute string `json:"attribute"`
	SortOrder int    `json:"sort_order"`
}

type OptionsAkeneo []OptionAkeneo

func (oas *OptionsAkeneo) AddOption(option OptionAkeneo) {

	if option.Labels.EsCl != "" {
		*oas = append(*oas, option)
	}
}

type OptionAkeneoBuilder interface {
	AddCode(string) OptionAkeneoBuilder
	AddLabels(string) OptionAkeneoBuilder
	AddAttribute(string) OptionAkeneoBuilder
	AddSortOrder(int) OptionAkeneoBuilder
	Build() OptionAkeneo
}

func (oa *OptionAkeneo) AddCode(code string) OptionAkeneoBuilder {
	if strings.TrimSpace(code) != "" {
		re := regexp.MustCompile("[^a-zA-Z0-9]+")
		oa.Code = re.ReplaceAllString(code, "_")

	} else {
		oa.Code = "0"
	}
	return oa
}

func (oa *OptionAkeneo) AddLabels(label string) OptionAkeneoBuilder {
	if strings.TrimSpace(label) == "" {
		oa.Labels.EsCl = ""
	} else {
		oa.Labels.EsCl = label
	}
	return oa
}
func (oa *OptionAkeneo) AddAttribute(attribute string) OptionAkeneoBuilder {
	oa.Attribute = attribute
	return oa
}
func (oa *OptionAkeneo) AddSortOrder(order int) OptionAkeneoBuilder {
	oa.SortOrder = order
	return oa
}
func (oa *OptionAkeneo) Build() OptionAkeneo {
	return OptionAkeneo{
		Code:      oa.Code,
		Labels:    oa.Labels,
		Attribute: oa.Attribute,
		SortOrder: oa.SortOrder,
	}
}
func NewOptionAkeneoBuilder() OptionAkeneoBuilder {
	return &OptionAkeneo{}
}

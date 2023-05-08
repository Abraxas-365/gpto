package models

type AttributeGroup struct {
	Code       string   `json:"code"`
	SortOrder  int      `json:"sort_order"`
	Attributes []string `json:"attributes"`
	Labels     Labels   `json:"labels"`
}
type AttributeGroups []AttributeGroup

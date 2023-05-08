package models

import "strings"

type Family struct {
	Code                string `json:"code"`
	Labels              Labels `json:"labels"`
	AttributeSetAsLabel string `json:"attribute_as_label"`
	Attributes          string `json:"attributes"`
}

type Families []Family

type FamilyAkeneo struct {
	Code                string   `json:"code"`
	Labels              Labels   `json:"labels"`
	AttributeSetAsLabel string   `json:"attribute_as_label"`
	Attributes          []string `json:"attributes"`
}

type FamiliesAkeneo []FamilyAkeneo

type FamilyAkeneoBuilder interface {
	AddCode(string) FamilyAkeneoBuilder
	AddLabels(Labels) FamilyAkeneoBuilder
	AddAttributeSetAsLabel(string) FamilyAkeneoBuilder
	AddAttributes(string) FamilyAkeneoBuilder
	Build() FamilyAkeneo
}

func (f *FamilyAkeneo) AddCode(code string) FamilyAkeneoBuilder {
	f.Code = code
	return f
}

func (f *FamilyAkeneo) AddLabels(labels Labels) FamilyAkeneoBuilder {
	f.Labels = labels
	return f
}

func (f *FamilyAkeneo) AddAttributeSetAsLabel(attributeSetAsLabel string) FamilyAkeneoBuilder {
	f.AttributeSetAsLabel = attributeSetAsLabel
	return f
}
func (f *FamilyAkeneo) AddAttributes(attributes string) FamilyAkeneoBuilder {
	attributesArray := strings.Split(attributes, ",")
	f.Attributes = attributesArray
	return f
}

func (f *FamilyAkeneo) Build() FamilyAkeneo {
	return FamilyAkeneo{
		Code:                f.Code,
		Labels:              f.Labels,
		AttributeSetAsLabel: f.AttributeSetAsLabel,
		Attributes:          f.Attributes,
	}
}

func NewAkeneoFamilyBuilder() FamilyAkeneo {
	return FamilyAkeneo{}
}

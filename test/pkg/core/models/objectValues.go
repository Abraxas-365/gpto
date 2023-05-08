package models

type StatusInt int

func (s StatusInt) ToBool() bool {
	if s == 1 {
		return true
	}
	return false
}

type StatusBool bool

func (s StatusBool) ToStringNumber() string {
	if s == true {
		return "1"
	}
	return "0"
}
func (s StatusBool) ToInt() int {
	if s == true {
		return 1
	}
	return 0
}

type CategoryLink struct {
	Position   int    `json:"position"`
	CategoryId string `json:"category_id"`
}
type CategoryLinks []CategoryLink

func (c *CategoryLinks) ExtractCategories() []string {
	var categories []string
	for _, elem := range *c {
		categories = append(categories, elem.CategoryId)
	}
	return categories
}

//Si se necesita agregar mas idiomas colocar abajo de esCL
type Labels struct {
	EsCl string `json:"es_CL"`
}

package models

type Category struct {
	Code   string `json:"code"`
	Parent string `json:"parent"`
	Labels Labels `json:"labels"`
}
type Categories []Category

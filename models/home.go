package models

import "go-blog/config"


type HomeResponse struct {
	config.Viewer
	Categorys []Category `json:"categorys,omitempty"`
	Posts     []PostMore `json:"posts,omitempty"`
	Total     int        `json:"total,omitempty"`
	Page      int        `json:"page,omitempty"`
	Pages     []int      `json:"pages,omitempty"`
	PageEnd   bool       `json:"page_end,omitempty"`
}

package system

import "gorm.io/gorm"

type CategoryModel struct {
	gorm.Model
	// 分类名称
	Name string `json:"name"`
	// 分类描述
	Desc string `json:"desc"`
	// 分类语言
	Lang string `json:"lang"`
}

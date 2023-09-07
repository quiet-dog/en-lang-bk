package system

import "gorm.io/gorm"

type CategoryModel struct {
	gorm.Model
	// 分类名称
	Name string `json:"name" gorm:"comment:分类名称"`
	// 分类描述
	Desc string `json:"desc" gorm:"comment:分类描述"`
	// 分类语言
	Lang string `json:"lang" gorm:"comment:分类语言"`
}

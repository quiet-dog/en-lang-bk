package system

import "gorm.io/gorm"

type WordModel struct {
	gorm.Model
	// 原本单词
	Content string `json:"content"`
	// 扩展批注内容
	Extend string `json:"extend"`
	// 归类
	Category *CategoryModel `json:"category"`
	// 归类ID
	CategoryID uint `json:"category_id"`
}

package system

import "gorm.io/gorm"

type WordModel struct {
	gorm.Model
	// 原本
	Content string `json:"content" gorm:"comment:原本单词"`
	// 扩展批注内容
	Extend string `json:"extend" gorm:"comment:扩展批注内容"`
	// 归类
	Category *CategoryModel `json:"category" gorm:"foreignKey:CategoryID;comment:归类"`
	// 归类ID
	CategoryID uint `json:"category_id" gorm:"comment:归类ID"`
	// 链接地址
	Link string `json:"link" gorm:"comment:链接地址"`
	// 意思
	Meaning string `json:"meaning" gorm:"comment:意思"`
	// 例句
	Sentence string `json:"sentence" gorm:"comment:例句"`
}

package system

import "gorm.io/gorm"

type ArticleModel struct {
	gorm.Model
	// 文章内容
	Content string `json:"content" gorm:"comment:文章内容"`
	// 标题
	Title string `json:"title" gorm:"comment:标题"`
}

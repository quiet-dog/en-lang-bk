package system

import "gorm.io/gorm"

type SentenceModel struct {
	gorm.Model
	// 原本句子
	Content string `json:"content" gorm:"comment:原本句子"`
	// 所属文章
	Article *ArticleModel `json:"article" gorm:"foreignKey:ArticleID;comment:所属文章"`
}

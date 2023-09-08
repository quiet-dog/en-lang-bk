package request

import (
	"en-lang-bk/model/common/request"
)

type ArticleCreate struct {
	// 文章内容
	Content string `json:"content" validate:"required"`
	// 标题
	Title string `json:"title" gorm:"comment:标题"`
}

type ArticleEdit struct {
	request.RequestID
	ArticleCreate
}

package request

import (
	"en-lang-bk/model/common/request"
)

type WordCreate struct {
	// 原本单词
	Content string `json:"content" validate:"required"`
	// 扩展批注内容
	Extend string `json:"extend"`
	// 归类ID
	CategoryID uint `json:"category" validate:"required"`
}

type WordEdit struct {
	request.RequestID
	WordCreate
}

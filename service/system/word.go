package system

import (
	"en-lang-bk/global"
	systemReq "en-lang-bk/model/system/request"

	"github.com/gin-gonic/gin"
)

type WordService struct{}

// 创建一个单词
func (w *WordService) CreateWord(c *gin.Context, word systemReq.WordCreate) (err error) {
	return global.DB.Create(&word).Error
}

func (w *WordService) EditWord(c *gin.Context, word systemReq.WordEdit) (err error) {
	return global.DB.Save(&word).Error
}

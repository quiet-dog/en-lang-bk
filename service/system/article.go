package system

import (
	"en-lang-bk/global"
	systemModel "en-lang-bk/model/system"
	systemReq "en-lang-bk/model/system/request"

	"github.com/gin-gonic/gin"
)

type ArticleService struct{}

// 创建一个文章
func (a *ArticleService) CreateArticle(c *gin.Context, req systemReq.ArticleCreate) (err error) {
	return global.DB.Model(&systemModel.ArticleModel{}).Create(&req).Error
}

// 编辑一个文章
func (a *ArticleService) EditArticle(c *gin.Context, req systemReq.ArticleEdit) (err error) {
	return global.DB.Model(&systemModel.ArticleModel{}).Save(&req).Error
}

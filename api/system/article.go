package system

import (
	"en-lang-bk/model/common/response"
	systemReq "en-lang-bk/model/system/request"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ArticleApi struct{}

// @Summary		添加文章
// @Description	添加文章
// @Tags			文章
// @Accept			application/json
// @Product		application/json
// @Param			data	body		systemReq.ArticleCreate	true	"data"
// @Success		200		{string}	string					"{"success":true,"data":{},"msg":"创建成功"}"
// @Router			/article/createArticle [post]
func (a *ArticleApi) CreateArticle(c *gin.Context) {
	var req systemReq.ArticleCreate
	_ = c.ShouldBindJSON(&req)

	validator := validator.New()
	if err := validator.Struct(req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := articleService.CreateArticle(c, req); err != nil {
		response.FailWithMessage("创建失败", c)
		return
	}

	response.OkWithMessage("创建成功", c)
}

// @Summary		编辑文章
// @Description	编辑文章
// @Tags			文章
// @Accept			application/json
// @Product		application/json
// @Param			data	body		systemReq.ArticleEdit	true	"data"
// @Success		200		{string}	string					"{"success":true,"data":{},"msg":"编辑成功"}"
// @Router			/article/editArticle [post]
func (a *ArticleApi) EditArticle(c *gin.Context) {
	var req systemReq.ArticleEdit
	_ = c.ShouldBindJSON(&req)

	validator := validator.New()
	if err := validator.Struct(req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := articleService.EditArticle(c, req); err != nil {
		response.FailWithMessage("编辑失败", c)
		return
	}

	response.OkWithMessage("编辑成功", c)
}

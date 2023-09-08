package system

import (
	"en-lang-bk/model/common/response"
	systemReq "en-lang-bk/model/system/request"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type WordApi struct{}

//	@Summary		创建词条
//	@Description	创建词条
//	@Tags			词条
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string					true	"token"
//	@Param			word			body		systemReq.WordCreate	true	"word"
//	@Success		200				{string}	string					"{"success":true,"data":{},"msg":"创建成功"}"
//	@Router			/word/createWord [post]
func (w *WordApi) CreateWord(c *gin.Context) {
	var word systemReq.WordCreate
	_ = c.ShouldBindJSON(&word)

	validate := validator.New()
	if err := validate.Struct(word); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := wordService.CreateWord(c, word); err != nil {
		response.FailWithMessage("创建失败", c)
		return
	}

	response.OkWithMessage("创建成功", c)
}

//	@Summary		编辑词条
//	@Description	编辑词条
//	@Tags			词条
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string				true	"token"
//	@Param			word			body		systemReq.WordEdit	true	"word"
//	@Success		200				{string}	string				"{"success":true,"data":{},"msg":"编辑成功"}"
//	@Router			/word/editWord [post]
func (w *WordApi) EditWord(c *gin.Context) {
	var word systemReq.WordEdit
	_ = c.ShouldBindJSON(&word)

	validate := validator.New()
	if err := validate.Struct(word); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := wordService.EditWord(c, word); err != nil {
		response.FailWithMessage("编辑失败", c)
		return
	}

	response.OkWithMessage("编辑成功", c)
}

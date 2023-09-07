package system

import "github.com/gin-gonic/gin"

type WordApi struct{}

// @Summary 创建词条
// @Description 创建词条
// @Tags 词条
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param word body Word true "word"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /word/createWord [post]
func (w *WordApi) CreateWord(c *gin.Context) {

}

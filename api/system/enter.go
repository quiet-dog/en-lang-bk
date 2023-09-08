package system

import (
	"en-lang-bk/service"
)

type ApiGroup struct {
	WordApi
}

var (
	wordService    = service.ServiceGroupApp.SystemServiceGroup.WordService
	articleService = service.ServiceGroupApp.SystemServiceGroup.ArticleService
)

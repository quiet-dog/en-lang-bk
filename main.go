package main

import (
	"en-lang-bk/core"
	"en-lang-bk/global"
)

func main() {

	global.VP = core.Viper()
	global.LOG = core.InitLogger()
	core.RunServer()
}

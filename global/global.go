package global

import (
	"en-lang-bk/config"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	RUNPATH = "."
	CONFIG  config.Server
	LOG     *logrus.Logger
	VP      *viper.Viper
	DB      *gorm.DB
)

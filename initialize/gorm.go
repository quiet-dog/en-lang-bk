package initialize

import (
	"en-lang-bk/global"
	"en-lang-bk/model/system"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return RegisterCallback(db.Debug())
}

func RegisterTables(db *gorm.DB) {

	// err := db.Migrator().CreateConstraint(&system.UserModel{}, "FkUserRoleId")

	db.AutoMigrate(
		system.WordModel{},
		system.MenuModel{},
		system.CategoryModel{},
		system.ArticleModel{},
		system.SentenceModel{},
	)

	global.LOG.Info("register table success")
}

func RegisterCallback(db *gorm.DB) *gorm.DB {

	return db
}

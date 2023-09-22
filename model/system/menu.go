package system

import "gorm.io/gorm"

type MenuModel struct {
	gorm.Model
	// 菜单名称
	Name string `json:"name" gorm:"comment:菜单名称"`
	// 菜单图标
	Icon string `json:"icon" gorm:"comment:菜单图标"`
	// 父级菜单
	Pid  uint       `json:"pid" gorm:"comment:父级菜单"`
	Menu *MenuModel `json:"menu" gorm:"foreignKey:Pid;comment:父级菜单"`
	// 菜单排序
	Sort uint `json:"sort" gorm:"comment:菜单排序"`
}

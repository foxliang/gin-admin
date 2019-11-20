package entity

import (
	"context"
	"github.com/LyricTian/gin-admin/internal/app/schema"
	"github.com/jinzhu/gorm"
)

// GetMenuActionDB 菜单动作管理
func GetMenuActionDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return getDBWithModel(ctx, defDB, MenuAction{})
}

// SchemaMenuAction 菜单动作管理
type SchemaMenuAction schema.MenuAction

// ToMenuAction 转换为菜单动作管理实体
func (a SchemaMenuAction) ToMenuAction() *MenuAction {
	item := &MenuAction{
		RecordID: &a.RecordID,
		MenuID:   &a.MenuID,
		Code:     &a.Code,
		Name:     &a.Name,
	}
	return item
}

// MenuAction 菜单动作管理实体
type MenuAction struct {
	Model
	RecordID *string `gorm:"column:record_id;size:36;index;"` // 记录ID
	MenuID   *string `gorm:"column:menu_id;size:36;index;"`   // 菜单ID
	Code     *string `gorm:"column:code;size:100;"`           // 动作编号
	Name     *string `gorm:"column:name;size:100;"`           // 动作名称
}

func (a MenuAction) String() string {
	return toString(a)
}

// TableName 表名
func (a MenuAction) TableName() string {
	return a.Model.TableName("menu_action")
}

// ToSchemaMenuAction 转换为菜单动作管理对象
func (a MenuAction) ToSchemaMenuAction() *schema.MenuAction {
	item := &schema.MenuAction{
		RecordID: *a.RecordID,
		MenuID:   *a.MenuID,
		Code:     *a.Code,
		Name:     *a.Name,
	}
	return item
}

// MenuActions 菜单动作管理列表
type MenuActions []*MenuAction

// ToSchemaMenuActions 转换为菜单动作管理对象列表
func (a MenuActions) ToSchemaMenuActions() []*schema.MenuAction {
	list := make([]*schema.MenuAction, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaMenuAction()
	}
	return list
}

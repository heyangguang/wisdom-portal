package models

import (
	"errors"
	"wisdom-portal/wisdom-portal/logger"
)

type UserGroup struct {
	BaseModel
	GroupName string `gorm:"unique;comment:'用户组名';not null" binding:"required" json:"group_name"`
	Remark    string `gorm:"type:text;comment:'备注'" json:"remark"`
	Users     []User `gorm:"many2many:user_usergroup" json:"users"`
}

// 添加用户组
func (userGroup *UserGroup) AddGroup() error {
	var uGroup UserGroup
	// 创建用户组
	if err := DB.Omit("id").Set("gorm:save_associations", false).Create(&userGroup).Error; err != nil {
		logger.Error("AddGroup    " + err.Error())
		return err
	}
	// 获取用户组对象
	if err := DB.Select("id").Where("group_name = ?", userGroup.GroupName).First(&uGroup).Error; err != nil {
		logger.Error("AddGroup    " + err.Error())
		return err
	}
	// 判断选择的用户对象是否存在
	for _, value := range userGroup.Users {
		if isOk := DB.Where("id = ?", value.ID).First(&User{}).RecordNotFound(); isOk {
			DB.Delete(&uGroup)
			logger.Error("AddGroup    选择的用户不存在")
			return errors.New("选择的用户不存在")
		}
	}
	// 添加用户对象
	if err := DB.Model(&uGroup).Association("Users").Append(&userGroup.Users).Error; err != nil {
		logger.Error("AddGroup    " + err.Error())
		return err
	}
	return nil
}

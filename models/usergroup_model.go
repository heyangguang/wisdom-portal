package models

import (
	"errors"
	"wisdom-portal/wisdom-portal/logger"
)

type UserGroup struct {
	BaseModel
	GroupName string    `binding:"required" json:"group_name"`
	Remark    string    `json:"remark"`
	Users     []UserObj `gorm:"MANY2MANY:user_usergroup;ASSOCIATION_JOINTABLE_FOREIGNKEY:user_id" binding:"required" json:"users"`
}

type UserObj struct {
	ID       int    `json:"id"`
	UserName string `json:"-"`
}

func (userObj *UserObj) TableName() string {
	return "user"
}

// 添加用户组
func (userGroup *UserGroup) AddGroup() error {
	var uGroup UserGroup

	// 创建用户组
	// DB.Omit("id")
	if err := DB.Set("gorm:save_associations", false).Create(&userGroup).Error; err != nil {
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

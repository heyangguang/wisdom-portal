package models

import (
	"fmt"
	wisdomPortal "wisdom-portal/wisdom-portal"
	"wisdom-portal/wisdom-portal/logger"
)

type User struct {
	BaseModel
	Name       string      `gorm:"unique;not null;comment:'姓名'" json:"name" binding:"required"`
	UserName   string      `gorm:"unique;not null;comment:'用户名' "json:"user_name" binding:"required"`
	PassWord   string      `gorm:"not null;comment:'密码'" json:"pass_word" binding:"required"`
	Phone      string      `gorm:"comment:'手机号'" json:"phone"`
	Email      string      `gorm:"comment:'电子邮箱'" json:"email"`
	Status     bool        `gorm:"comment:'用户状态'" json:"status"`
	Secret     string      `gorm:"comment:'双因子秘钥'" json:"secret"`
	UserGroups []UserGroup `gorm:"many2many:user_usergroup" json:"user_groups"`
}

type UserLogin struct {
	UserName string `json:"user_name" binding:"required"`
	PassWord string `json:"pass_word" binding:"required"`
	OtpCode  string `json:"otp_code" binding:"required"`
}

type PubCurrentUser struct {
	Name     string     `json:"name"`
	UserName string     `json:"user_name"`
	Id       string     `json:"user_id"`
	Roles    []RoleLine `json:"roles"`
}

func (pubCurrentUser PubCurrentUser) TableName() string {
	return "user"
}

// 获取获取当前用户信息
func (pubCurrentUser PubCurrentUser) GetPubCurrentUser(userName interface{}) (result PubCurrentUser, isExist bool) {
	if DB.Select("name, user_name, id").Where("user_name = ?", userName).Take(&pubCurrentUser).RecordNotFound() {
		return PubCurrentUser{}, false
	}
	e := LoadPolicyPerm()
	// 获取用户和用户组的全部权限
	userRoles := e.GetImplicitPermissionsForUser(userName.(string))
	for _, value := range userRoles {
		line := RoleLine{
			ObJName: value[1],
			ActName: value[2],
		}
		pubCurrentUser.Roles = append(pubCurrentUser.Roles, line)
	}
	logger.Debug(fmt.Sprint("GetPubCurrentUser: ", pubCurrentUser))
	return pubCurrentUser, true
}

// 检测用户名密码
func (u *UserLogin) CheckUserLogin() bool {
	var user User
	if DB.Select("pass_word").
		Where("user_name = ?", u.UserName).First(&user).RecordNotFound() {
		return false
	}
	if user.PassWord == wisdomPortal.String2md5(u.PassWord) {
		return true
	}
	return false
}

// 验证谷歌动态码Totp
func (u *UserLogin) CheckUserOtpCode() bool {
	var user User
	if DB.Select("secret").
		Where("user_name = ?", u.UserName).First(&user).RecordNotFound() {
		return false
	}
	if isOk, _ := NewGoogleAuth().VerifyCode(user.Secret, u.OtpCode); isOk {
		return true
	}
	return false
}

// 创建用户
func (u *User) AddUser(user User) error {
	if err := DB.Omit("id").Create(&user).Error; err != nil {
		return err
	}
	return nil
}

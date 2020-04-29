package models

import (
	wisdomPortal "wisdom-portal/wisdom-portal"
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
	OtpCode string `json:"otp_code" binding:"required"`
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

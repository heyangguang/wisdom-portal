package models

import (
	"fmt"
	wisdomPortal "wisdom-portal/wisdom-portal"
	"wisdom-portal/wisdom-portal/logger"
)

type User struct {
	BaseModel
	Name     string `json:"name" binding:"required"`
	UserName string `json:"user_name" binding:"required"`
	PassWord string `json:"pass_word" binding:"required"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Secret   string `json:"-"`
	Status   bool   `json:"-"`
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
	Rules    []RuleLine `json:"rules"`
}

func (u *User) TableName() string {
	return "user"
}

func (p *PubCurrentUser) TableName() string {
	return "user"
}

// 获取获取当前用户信息
func (p *PubCurrentUser) GetPubCurrentUser(userName interface{}) (result *PubCurrentUser, isExist bool) {
	if DB.Select("name, user_name, id").Where("user_name = ?", userName).Take(&p).RecordNotFound() {
		return &PubCurrentUser{}, false
	}
	e := LoadPolicyPerm()
	// 获取用户和用户组的全部权限
	userRoles := e.GetImplicitPermissionsForUser(userName.(string))
	for _, value := range userRoles {
		line := RuleLine{
			ObJName: value[1],
			ActName: value[2],
		}
		p.Rules = append(p.Rules, line)
	}
	logger.Debug(fmt.Sprint("GetPubCurrentUser: ", p))
	return p, true
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
	if err := DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

package models

import (
	"fmt"
	"wisdom-portal/schemas"
	wisdomPortal "wisdom-portal/wisdom-portal"
	"wisdom-portal/wisdom-portal/logger"
)

type User struct {
	BaseModel
	Name     string `json:"name" validate:"required" label:"name"`
	UserName string `json:"user_name" validate:"required,max=10,min=3,ValidationUserNameFormat" label:"user_name"`
	PassWord string `json:"pass_word" validate:"required,max=20,min=10" label:"pass_word"`
	Phone    string `json:"phone" validate:"required,ValidationPhoneFormat" label:"phone"`
	Email    string `json:"email" validate:"required,email" label:"email"`
	Secret   string `json:"-"`
	Status   bool   `json:"-"`
	RuleId   int    `json:"-"`
}

type UserLogin struct {
	UserName string `json:"user_name" validate:"required,max=10,min=3,ValidationUserNameFormat"`
	PassWord string `json:"pass_word" validate:"required,max=20,min=10" label:"pass_word"`
	OtpCode  string `json:"otp_code" validate:"required,len=6" label:"otp_code"`
}

// 当前用户信息返回结构体
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
	if isOk, _ := schemas.NewGoogleAuth().VerifyCode(user.Secret, u.OtpCode); isOk {
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

// 判断用户名是否存在
// true 找到了  false 未找到
func CheckUserName(userName string) bool {
	if isOk := DB.Where("user_name = ?", userName).Take(&User{}).RecordNotFound(); !isOk {
		logger.Debug("查询到user_name: " + userName)
		return true
	}
	return false
}

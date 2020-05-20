package user

import (
	"github.com/go-playground/validator/v10"
	"regexp"
	"strings"
	"wisdom-portal/wisdom-portal/forms"
)

func CustomValidations() {
	_ = forms.Validate.RegisterValidation("ValidationUserNameFormat", ValidationUserNameFormat)
	_ = forms.Validate.RegisterValidation("ValidationPhoneFormat", ValidationPhoneFormat)
}

// 自定义字段验证
func ValidationUserNameFormat(fl validator.FieldLevel) bool {
	if ok, _ := regexp.MatchString("^[a-zA-Z_]+$", fl.Field().String()); !ok {
		return false
	}
	return true
}

// 验证手机号
func ValidationPhoneFormat(fl validator.FieldLevel) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	reg := regexp.MustCompile(regular)
	if !reg.MatchString(fl.Field().String()) {
		return false
	}
	return true
}

// 修改验证字段错误返回值方法
func GetValidationError(err validator.ValidationErrors) []map[string]string {
	sliceErrs := forms.BaseFormValidationError(err)
	for _, value := range sliceErrs {
		if errValue, ok := value["user_name"]; ok {
			if strings.Contains(errValue, "ValidationUserNameFormat") {
				value["user_name"] = "user_name can only be English letters"
			}
		}
		if errValue, ok := value["phone"]; ok {
			if strings.Contains(errValue, "ValidationPhoneFormat") {
				value["phone"] = "Please enter the correct phone number"
			}
		}
	}
	return sliceErrs
}

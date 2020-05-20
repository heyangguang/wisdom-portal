package usergroup

import (
	"github.com/go-playground/validator/v10"
	"regexp"
	"strings"
	"wisdom-portal/wisdom-portal/forms"
)

func CustomValidations() {
	_ = forms.Validate.RegisterValidation("ValidationUserGroupNameFormat", ValidationUserGroupNameFormat)
}

// 验证用户组名
func ValidationUserGroupNameFormat(fl validator.FieldLevel) bool {
	if ok, _ := regexp.MatchString("^[a-zA-Z_]+$", fl.Field().String()); !ok {
		return false
	}
	return true
}

// 修改验证字段错误返回值方法
func GetValidationError(err validator.ValidationErrors) []map[string]string {
	sliceErrs := forms.BaseFormValidationError(err)
	for _, value := range sliceErrs {
		if errValue, ok := value["group_name"]; ok {
			if strings.Contains(errValue, "ValidationUserGroupNameFormat") {
				value["group_name"] = "group_name can only be English letters"
			}
		}
	}
	return sliceErrs
}

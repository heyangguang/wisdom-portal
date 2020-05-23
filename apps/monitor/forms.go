package monitor

import (
	"github.com/go-playground/validator/v10"
	"strings"
	"wisdom-portal/wisdom-portal/forms"
)

func CustomValidations() {
	_ = forms.Validate.RegisterValidation("ValidationNumFormat", ValidationNumFormat)
	_ = forms.Validate.RegisterValidation("ValidationAppTagFormat", ValidationAppTagFormat)
}

// 验证Num
func ValidationNumFormat(fl validator.FieldLevel) bool {
	if fl.Field().Int() == 10 || fl.Field().Int() == 20 || fl.Field().Int() == 40 || fl.Field().Int() == 60 {
		return true
	}
	return false
}

// 验证Tag
func ValidationAppTagFormat(fl validator.FieldLevel) bool {
	if fl.Field().String() == "MySQL" || fl.Field().String() == "ElasticSearch" {
		return true
	}
	return false
}

// 修改验证字段错误返回值方法
func GetValidationError(err validator.ValidationErrors) []map[string]string {
	sliceErrs := forms.BaseFormValidationError(err)
	for _, value := range sliceErrs {
		if errValue, ok := value["num"]; ok {
			if strings.Contains(errValue, "ValidationNumFormat") {
				value["num"] = "Please enter the correct num"
			}
		}
		if errValue, ok := value["app_tag"]; ok {
			if strings.Contains(errValue, "ValidationAppTagFormat") {
				value["app_tag"] = "Please enter the correct app_tag"
			}
		}
	}
	return sliceErrs
}

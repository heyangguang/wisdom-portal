package monitor

import (
	"github.com/go-playground/validator/v10"
	"strings"
	"time"
	"wisdom-portal/wisdom-portal/forms"
)

func CustomValidations() {
	_ = forms.Validate.RegisterValidation("ValidationNumFormat", ValidationNumFormat)
	_ = forms.Validate.RegisterValidation("ValidationAppTagFormat", ValidationAppTagFormat)
	_ = forms.Validate.RegisterValidation("ValidationIntervalFormat", ValidationIntervalFormat)
	_ = forms.Validate.RegisterValidation("ValidationTimeFormat", ValidationTimeFormat)
	_ = forms.Validate.RegisterValidation("ValidationTagFormat", ValidationTagFormat)
}

// 验证Num
func ValidationNumFormat(fl validator.FieldLevel) bool {
	if fl.Field().Int() == 10 || fl.Field().Int() == 20 || fl.Field().Int() == 40 || fl.Field().Int() == 60 {
		return true
	}
	return false
}

// 验证AppTag
func ValidationAppTagFormat(fl validator.FieldLevel) bool {
	if fl.Field().String() == "MySQL" || fl.Field().String() == "ElasticSearch" ||
		fl.Field().String() == "Kubernetes" || fl.Field().String() == "Kafka" {
		return true
	}
	return false
}

// 验证中间表Tag
func ValidationTagFormat(fl validator.FieldLevel) bool {
	if fl.Field().String() == "i" || fl.Field().String() == "u" {
		return true
	}
	return false
}

// 验证Time字段
func ValidationTimeFormat(fl validator.FieldLevel) bool {
	//fmt.Printf(fl.Field().String())
	if _, err := time.Parse("2006-01-02 15:04:05", fl.Field().String()); err != nil {
		return false
	}
	return true
}

// 验证质量检测平均时间
func ValidationIntervalFormat(fl validator.FieldLevel) bool {
	if fl.Field().String() == "1" || fl.Field().String() == "5" ||
		fl.Field().String() == "10" || fl.Field().String() == "20" ||
		fl.Field().String() == "40" || fl.Field().String() == "60" {
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
		if errValue, ok := value["interval"]; ok {
			if strings.Contains(errValue, "ValidationIntervalFormat") {
				value["interval"] = "Please enter the correct interval"
			}
		}
		if errValue, ok := value["time"]; ok {
			if strings.Contains(errValue, "ValidationTimeFormat") {
				value["time"] = "Please enter the correct time"
			}
		}
		if errValue, ok := value["tag"]; ok {
			if strings.Contains(errValue, "ValidationTagFormat") {
				value["tag"] = "Please enter the correct tag"
			}
		}
	}
	return sliceErrs
}

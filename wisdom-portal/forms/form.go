package forms

import (
	languageEn "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	"reflect"
)

type Option func()

var options []Option

// use a single instance , it caches struct info
var (
	Uni      *ut.UniversalTranslator
	Validate *validator.Validate
)

func InitFormValidate() {
	en := languageEn.New()
	Uni = ut.New(en)
	Validate = validator.New()
	RegisterTagFunc()
}

// 注册一个函数，获取struct tag里自定义的label作为字段名
func RegisterTagFunc() {
	Validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("label")
		return name
	})
}

// 注册自定义表单验证方法
func RegisterCustomValidationFunc(opts ...Option) {
	options = append(options, opts...)
}

// 初始化自定义表单方法
func InitCustomValidationFunc() {
	for _, opt := range options {
		opt()
	}
}

// 公共验证字段错误返回值方法
func BaseFormValidationError(err validator.ValidationErrors) []map[string]string {
	tans, _ := Uni.GetTranslator("en")
	_ = enTranslations.RegisterDefaultTranslations(Validate, tans)
	var sliceErrs []map[string]string
	for _, e := range err {
		sliceErrs = append(sliceErrs, map[string]string{e.Field(): e.Translate(tans)})
	}
	return sliceErrs
}

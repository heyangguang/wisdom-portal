package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	en2 "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	"net/http"
	"reflect"
	"regexp"
	"strings"
)

type Person struct {
	Age     int    `form:"age" validate:"required,gt=10" label:"age"`
	Name    string `form:"name" validate:"required,CustomValidationErrors" label:"name"`
	Address string `form:"address" json:"address" validate:"required" label:"address"`
}

// 公共验证字段错误返回值方法
func BaseError(err validator.ValidationErrors) []map[string]string {
	tans, _ := uni.GetTranslator("en")
	_ = enTranslations.RegisterDefaultTranslations(validate, tans)
	var sliceErrs []map[string]string
	for _, e := range err {
		sliceErrs = append(sliceErrs, map[string]string{e.Field(): e.Translate(tans)})
	}
	return sliceErrs
}

// 修改验证字段错误返回值方法
func (p *Person) GetError(err validator.ValidationErrors) []map[string]string {
	sliceErrs := BaseError(err)
	for _, value := range sliceErrs {
		if errValue, ok := value["name"]; ok {
			fmt.Println(errValue)
			if strings.Contains(errValue, "CustomValidationErrors") {
				value["name"] = "傻逼"
			}
		}
	}
	return sliceErrs
}

// use a single instance , it caches struct info
var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

func multiLangBindingHandler(c *gin.Context) {
	var person Person
	if err := c.ShouldBind(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if err := validate.Struct(person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": person.GetError(err.(validator.ValidationErrors)),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"personInfo": person,
	})
}

// 自定义字段验证
func CustomValidationErrors(fl validator.FieldLevel) bool {
	if ok, _ := regexp.MatchString("^[a-zA-Z]+$", fl.Field().String()); !ok {
		return false
	}
	return true
}

func main() {
	en := en2.New()
	uni = ut.New(en)
	validate = validator.New()

	//注册一个函数，获取struct tag里自定义的label作为字段名
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("label")
		return name
	})

	// 注册自定义函数
	_ = validate.RegisterValidation("CustomValidationErrors", CustomValidationErrors)

	router := gin.Default()
	router.GET("/testMultiLangBinding", multiLangBindingHandler)
	_ = router.Run(":9999")
}

package lang

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

var trans ut.Translator
var uni *ut.UniversalTranslator

var validate *validator.Validate

func InitTrans() (err error) {
	// 修改gin框架中的Validator引擎属性，实现自定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册一个获取json tag的自定义方法
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		zhT := zh.New() // 中文翻译器
		enT := en.New() // 英文翻译器

		// 第一个参数是备用（fallback）的语言环境
		// 后面的参数是应该支持的语言环境（支持多个）
		// uni := ut.New(zhT, zhT) 也是可以的
		uni = ut.New(enT, zhT, enT)
		trans, _ = uni.GetTranslator("en")
		err = enTranslations.RegisterDefaultTranslations(v, trans)
		validate = v
		return
	}
	return
}

func GetTrans(local string) ut.Translator {
	switch local {
	case "zh_CN":
		trans, _ = uni.GetTranslator("zh")
		_ = zhTranslations.RegisterDefaultTranslations(validate, trans)
	default:
		trans, _ = uni.GetTranslator("en")
		err := enTranslations.RegisterDefaultTranslations(validate, trans)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	return trans
}

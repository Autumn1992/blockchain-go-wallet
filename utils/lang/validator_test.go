package lang

import (
	"fmt"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"testing"
)

type TestUser struct {
	Name string `json:"name" validate:"required"`
	Age  int    `json:"age" validate:"required,gt=1"`
}

func TestValidator(t *testing.T) {
	zh := zh.New()
	uni := ut.New(zh, zh)

	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	trans, _ := uni.GetTranslator("zh")

	validate := validator.New()

	zh_translations.RegisterDefaultTranslations(validate, trans)
	user := &TestUser{
		Name: "123",
		Age:  0,
	}
	err := validate.Struct(user)
	if err != nil {

		// translate all error at once
		errs := err.(validator.ValidationErrors)

		// returns a map with key = namespace & value = translated error
		// NOTICE: 2 errors are returned and you'll see something surprising
		// translations are i18n aware!!!!
		// eg. '10 characters' vs '1 character'
		fmt.Println(errs.Translate(trans))
	}
}

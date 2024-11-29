package lang

import (
	"encoding/json"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"os"
)

type Translator struct {
	bundle *i18n.Bundle
}

var translator = &Translator{}

func init() {
	langFiles, err := os.ReadDir("./resource/lang/")
	if err != nil {
		fmt.Printf("InitTranslator() Error: %v", err)
	}
	translator.bundle = i18n.NewBundle(language.English)
	translator.bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	for _, langFile := range langFiles {
		if !langFile.IsDir() {
			translator.bundle.MustLoadMessageFile(fmt.Sprintf("./resource/lang/%s", langFile.Name()))
		}
	}
}

func (t *Translator) TranslateMessage(lang string, msg string) string {
	localizer := i18n.NewLocalizer(t.bundle, lang)
	m, err := localizer.LocalizeMessage(&i18n.Message{ID: msg})
	if err != nil {
		return msg
	}
	return m
}

func GetTranslation() *Translator {
	return translator
}

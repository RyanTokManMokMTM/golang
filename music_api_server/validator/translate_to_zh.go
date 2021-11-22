package validator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTw "github.com/go-playground/validator/v10/translations/zh_tw"
)

var (
	universal *ut.UniversalTranslator
	trans  ut.Translator
	Validate *validator.Validate
)

func init(){
	tw := zh_Hant_TW.New() //new zh_tw translation
	universal = ut.New(tw,tw) //new translator

	trans , _ = universal.GetTranslator("zh_tw")

	//setting gin binding validator
	Validate = binding.Validator.Engine().(*validator.Validate)
	//register translator to gin validator
	zhTw.RegisterDefaultTranslations(Validate,trans)
}

func Translate(err error) map[string][]string{
	result := make(map[string][]string)

	errs := err.(validator.ValidationErrors)
	for _,err:= range errs{
		result[err.Field()] = append(result[err.Field()],err.Translate(trans))
	}
	return  result
}


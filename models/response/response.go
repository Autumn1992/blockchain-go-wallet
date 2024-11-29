package response

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
	"reflect"
	"time"
	"walletserver/log"
	"walletserver/utils/lang"

	"github.com/gin-gonic/gin"
)

const (
	SUCCESS = 1
	FAILED  = -1 //通用错误
)

const (
	StatusInternalServerError = "Internal Server Error" // 服务器错误

	INVALID_AUTHORIZATION = "invalid_token" // 无效的token

	UNAUTHORIZED = "unauthorized" // 无权限

	PWD_NOT_TMATCH = "pwd_not_match" //两次密码不一致

	PWD_ERROR = "pwd_error" //密码错误

	PARAMETER_INCORRECT = "parameter_incorrect" // 参数无效

	INVALID_SESSION = "invalid_session" //无效的链接
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Ok(ctx *gin.Context, data interface{}) {
	result(ctx, SUCCESS, "ok", data)
}

func Fail(ctx *gin.Context, msg string) {
	//国际化
	l := ctx.GetHeader("lang")
	if l == "" {
		l = "en_US"
	}
	msg = lang.GetTranslation().TranslateMessage(l, msg)
	result(ctx, FAILED, msg, nil)
}

func FailWithError(ctx *gin.Context, err error) {
	//国际化
	l := ctx.GetHeader("lang")
	if l == "" {
		l = "en_US"
	}
	var errs validator.ValidationErrors
	ok := errors.As(err, &errs)
	if ok {
		msg := errs[0].Translate(lang.GetTrans(l))
		result(ctx, FAILED, msg, nil)
		return
	}

	msg := lang.GetTranslation().TranslateMessage(l, err.Error())
	result(ctx, FAILED, msg, nil)
}
func FailWithErrorWithCode(ctx *gin.Context, err error, code int) {
	//国际化
	l := ctx.GetHeader("lang")
	if l == "" {
		l = "en_US"
	}
	var errs validator.ValidationErrors
	ok := errors.As(err, &errs)
	if ok {
		msg := errs[0].Translate(lang.GetTrans(l))
		result(ctx, code, msg, nil)
		return
	}

	msg := lang.GetTranslation().TranslateMessage(l, err.Error())
	result(ctx, code, msg, nil)
}

func result(ctx *gin.Context, code int, msg string, data interface{}) {

	var start_time, _ = ctx.Get("reqstarttime")
	if start_time == nil {
		start_time = int64(0)
	}
	uid_tgid, exist := ctx.Get("uid_tgid")

	var printLog = ctx.GetBool("printLog")
	if printLog {
		if exist {
			log.Info(fmt.Sprintf("Response %s %s %d %s %s",
				"uid_tgid="+uid_tgid.(string),
				time.Now().Format("2006/01/02 - 15:04:05"),
				code,
				fmt.Sprintf("耗时%dms", time.Now().UnixMilli()-start_time.(int64)),
				msg), data)
		} else {
			log.Info(fmt.Sprintf("Response %s %d %s %s",
				time.Now().Format("2006/01/02 - 15:04:05"),
				code,
				fmt.Sprintf("耗时%dms", time.Now().UnixMilli()-start_time.(int64)),
				msg), data)
		}
	}
	ctx.JSON(http.StatusOK, Response{
		code,
		msg,
		data,
	})
}

func SimpleResult(ctx *gin.Context, code int, data interface{}) {
	var start_time, _ = ctx.Get("start_time")
	if start_time == nil {
		start_time = int64(0)
	}
	uid, exist := ctx.Get("uid")
	if !exist {
		uid, exist = ctx.Get("uid_tgid")
	}

	var datajson interface{}
	var kind = reflect.TypeOf(data).Kind()
	if kind == reflect.Struct || kind == reflect.Map {
		var bts, _ = json.Marshal(data)
		datajson = string(bts)
	} else {
		datajson = data
	}

	if exist {
		log.Info(fmt.Sprintf("SimpleResult Response %s %s %s",
			"uid="+uid.(string),
			time.Now().Format("2006/01/02 - 15:04:05"),
			fmt.Sprintf("耗时%dms", time.Now().UnixMilli()-start_time.(int64))), datajson)
	} else {
		log.Info(fmt.Sprintf("SimpleResult Response %s %s ",
			time.Now().Format("2006/01/02 - 15:04:05"),
			fmt.Sprintf("耗时%dms", time.Now().UnixMilli()-start_time.(int64))), datajson)
	}
	ctx.JSON(code, data)
}

func SimpleOk(ctx *gin.Context, data interface{}) {
	SimpleResult(ctx, http.StatusOK, data)
}

func XMLResult(ctx *gin.Context, data interface{}) {
	var start_time, _ = ctx.Get("start_time")
	if start_time == nil {
		start_time = int64(0)
	}
	uid, exist := ctx.Get("uid")
	if !exist {
		uid, exist = ctx.Get("uid_tgid")
	}
	if exist {
		log.Info(fmt.Sprintf("xmlResult Response %s %s %s",
			"uid="+uid.(string),
			time.Now().Format("2006/01/02 - 15:04:05"),
			fmt.Sprintf("耗时%dms", time.Now().UnixMilli()-start_time.(int64))), data)
	} else {
		log.Info(fmt.Sprintf("xmlResult Response %s %s ",
			time.Now().Format("2006/01/02 - 15:04:05"),
			fmt.Sprintf("耗时%dms", time.Now().UnixMilli()-start_time.(int64))), data)
	}
	ctx.XML(http.StatusOK, data)
}

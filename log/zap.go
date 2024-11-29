package log

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/onsi/ginkgo/reporters/stenographer/support/go-colorable"
	"io"
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
	"os"
	"strings"
)

var (
	defaultlogpath = "./logs" //日志路径
)

var logger *zap.SugaredLogger // 定义日志打印全局变量
var ErrorWriter io.Writer
var InfoWriter io.Writer

func init() {
	for _, arg := range os.Args {
		if strings.HasPrefix(arg, "--logpath=") {
			defaultlogpath = arg[len("--logpath="):]
		}
	}
	//fmt.Println("====================args=================")
	fmt.Println("defaultlogpath:", defaultlogpath)
	logger = newZapLogger(true, true, "info.log")
}

func getEncoder(isJSON bool) zapcore.Encoder {
	var encoderConfig = zap.NewDevelopmentEncoderConfig()
	if gin.Mode() == gin.ReleaseMode {
		encoderConfig = zap.NewProductionEncoderConfig()
	}
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	if isJSON {
		return zapcore.NewJSONEncoder(encoderConfig)
	}
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func newZapLogger(log2console bool, log2file bool, filename string) *zap.SugaredLogger {
	var cores []zapcore.Core
	//全部日志
	var logLevel = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return true
	})
	//var errLevel = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
	//	return lvl >= zapcore.ErrorLevel
	//})
	if gin.Mode() == gin.DebugMode {
		log2console = true
	}
	//log to console
	if log2console {
		// writer := zapcore.Lock(os.Stdout)
		//writer := zapcore.AddSync(colorable.NewColorableStdout())
		writer := zapcore.AddSync(colorable.NewColorableStdout())
		core := zapcore.NewCore(getEncoder(false), writer, logLevel) //日志已经自己解析并生成json 所以此处isJSON false
		cores = append(cores, core)
	}
	//log to file info debug warning
	if log2file {
		writer := zapcore.AddSync(&lumberjack.Logger{
			Filename:   defaultlogpath + "/" + filename,
			MaxSize:    100, //megabytes   //100M
			Compress:   true,
			MaxBackups: 30, //备份数
			MaxAge:     20, //days 保留20天
		})
		InfoWriter = writer

		errorWriter := zapcore.AddSync(&lumberjack.Logger{
			Filename:   defaultlogpath + "/error.log",
			MaxSize:    200, //megabytes   //100M
			Compress:   true,
			MaxBackups: 5,  //备份数
			MaxAge:     30, //days 保留20天
		})
		ErrorWriter = errorWriter
		core := zapcore.NewCore(getEncoder(false), writer, logLevel)
		cores = append(cores, core, zapcore.NewCore(getEncoder(false), errorWriter, zapcore.ErrorLevel))
	}

	combinedCore := zapcore.NewTee(cores...)

	// AddCallerSkip skips 2 number of callers, this is important else the file that gets
	// logged will always be the wrapped file. In our case zap.go
	l := zap.New(combinedCore). //zap.AddCallerSkip(1),
		//zap.AddCaller(),
		Sugar()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	return l
}

package log

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
)

const maxDepth = 8

func getDebugTrace() string {
	// 创建一个用于存储调用堆栈信息的切片
	pcs := make([]uintptr, maxDepth)
	// 获取当前调用栈信息，跳过前 2 层调用（0 是 runtime.Callers，1 是 printStack）
	depth := runtime.Callers(4, pcs)
	if depth == 0 {
		return ""
	}

	var tracemsg = ""
	// 遍历调用栈信息，打印每一层调用函数名和文件位置
	for i := 0; i < depth; i++ {
		var f = runtime.FuncForPC(pcs[i])
		var path, line = f.FileLine(pcs[i])
		var fileName = filepath.Base(path)
		if fileName == "zap.go" || !strings.Contains(path, "pay-server") {
			continue
		}
		tracemsg = tracemsg + fmt.Sprintf("%s:%d\n", fileName, line)
	}
	if len(tracemsg) > 1 {
		tracemsg = tracemsg[0 : len(tracemsg)-1] //去掉最后一个换行符
	}
	return tracemsg
}

func any2Json(data interface{}) string {
	var bts, err = json.Marshal(data)
	if err != nil {
		return fmt.Sprintf("%v", data)
	}
	return string(bts)
}

func toString(val interface{}) string {
	return fmt.Sprintf("%v", val)
}

func formatArgs(traceback bool, args []interface{}, level int) string {
	var argsList = make([]string, 0)
	if traceback {
		argsList = append(argsList, getDebugTrace())
	}
	sheetStr := strings.Repeat("\t", level-1)
	for idx, arg := range args {
		prefix := fmt.Sprintf("%s[%d] ", sheetStr, idx+1)
		switch v := reflect.ValueOf(arg); v.Kind() {
		case reflect.Struct, reflect.Map, reflect.Array, reflect.Slice:
			argsList = append(argsList, prefix+any2Json(arg))
		default:
			argsList = append(argsList, prefix+toString(arg))
		}
	}

	return strings.Join(argsList, "\n")
}

var mode = 2

// Log logs the provided arguments at provided level.
// Spaces are added between arguments when neither is a string.
func Log(lvl zapcore.Level, args ...interface{}) {
	if mode == 1 {
		logger.WithOptions(zap.AddCallerSkip(1)).Log(lvl, args...)
	} else {
		logger.Log(lvl, formatArgs(true, args, 1))
	}
}

// Debug logs the provided arguments at [DebugLevel].
// Spaces are added between arguments when neither is a string.
func Debug(args ...interface{}) {
	if mode == 1 {
		logger.WithOptions(zap.AddCallerSkip(1)).Debug(args...)
	} else {
		logger.Debug(formatArgs(true, args, 1))
	}
	if gin.Mode() == gin.TestMode {
		fmt.Print("[debug]")
		fmt.Println(args...)
	}
}

// Info logs the provided arguments at [InfoLevel].
// Spaces are added between arguments when neither is a string.
func Info(args ...interface{}) {
	if mode == 1 {
		logger.WithOptions(zap.AddCallerSkip(1)).Info(args...)
	} else {
		logger.Info(formatArgs(true, args, 1))
	}
	if gin.Mode() == gin.TestMode {
		fmt.Print("[info]")
		fmt.Println(args...)
	}
}

// Warn logs the provided arguments at [WarnLevel].
// Spaces are added between arguments when neither is a string.
func Warn(args ...interface{}) {
	if mode == 1 {
		logger.WithOptions(zap.AddCallerSkip(1)).Warn(args...)
	} else {
		logger.Warn(formatArgs(true, args, 1))
	}
	if gin.Mode() == gin.TestMode {
		fmt.Print("[warn]")
		fmt.Println(args...)
	}
}

// Error logs the provided arguments at [ErrorLevel].
// Spaces are added between arguments when neither is a string.
func Error(args ...interface{}) {
	if mode == 1 {
		logger.WithOptions(zap.AddCallerSkip(1)).Error(args...)
	} else {
		logger.Error(formatArgs(true, args, 1))
	}
	if gin.Mode() == gin.TestMode {
		fmt.Print("[error]")
		fmt.Println(args...)
	}
}

// DPanic logs the provided arguments at [DPanicLevel].
// In development, the logger then panics. (See [DPanicLevel] for details.)
// Spaces are added between arguments when neither is a string.
func DPanic(args ...interface{}) {
	logger.WithOptions(zap.AddCallerSkip(1)).DPanic(args...)
}

// Panic constructs a message with the provided arguments and panics.
// Spaces are added between arguments when neither is a string.
func Panic(args ...interface{}) {
	logger.WithOptions(zap.AddCallerSkip(1)).Panic(args...)
}

// Fatal constructs a message with the provided arguments and calls os.Exit.
// Spaces are added between arguments when neither is a string.
func Fatal(args ...interface{}) {
	//logger.WithOptions(zap.AddCallerSkip(1)).Fatal(args...)
	logger.Fatal(formatArgs(true, args, 1))
}

// Logf formats the message according to the format specifier
// and logs it at provided level.
func Logf(lvl zapcore.Level, template string, args ...interface{}) {
	//logger.WithOptions(zap.AddCallerSkip(1)).Logf(lvl, template, args...)
	Log(lvl, fmt.Sprintf(template, args...))
}

// Debugf formats the message according to the format specifier
// and logs it at [DebugLevel].
func Debugf(template string, args ...interface{}) {
	//logger.WithOptions(zap.AddCallerSkip(1)).Debugf(template, args...)
	Debug(fmt.Sprintf(template, args...))
}

// Infof formats the message according to the format specifier
// and logs it at [InfoLevel].
func Infof(template string, args ...interface{}) {
	//logger.WithOptions(zap.AddCallerSkip(1)).Infof(template, args...)
	Info(fmt.Sprintf(template, args...))
}

// Warnf formats the message according to the format specifier
// and logs it at [WarnLevel].
func Warnf(template string, args ...interface{}) {
	//logger.WithOptions(zap.AddCallerSkip(1)).Warnf(template, args...)
	Warn(fmt.Sprintf(template, args...))
}

// Errorf formats the message according to the format specifier
// and logs it at [ErrorLevel].
func Errorf(template string, args ...interface{}) {
	//logger.WithOptions(zap.AddCallerSkip(1)).Errorf(template, args...)
	Error(fmt.Sprintf(template, args...))
}

// DPanicf formats the message according to the format specifier
// and logs it at [DPanicLevel].
// In development, the logger then panics. (See [DPanicLevel] for details.)
func DPanicf(template string, args ...interface{}) {
	//logger.WithOptions(zap.AddCallerSkip(1)).DPanicf(template, args...)
	DPanic(fmt.Sprintf(template, args...))
}

// Panicf formats the message according to the format specifier
// and panics.
func Panicf(template string, args ...interface{}) {
	//logger.WithOptions(zap.AddCallerSkip(1)).Panicf(template, args...)
	Panic(fmt.Sprintf(template, args...))
}

// Fatalf formats the message according to the format specifier
// and calls os.Exit.
func Fatalf(template string, args ...interface{}) {
	//logger.WithOptions(zap.AddCallerSkip(1)).Fatalf(template, args...)
	Fatal(fmt.Sprintf(template, args...))
}

// Logw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Logw(lvl zapcore.Level, msg string, keysAndValues ...interface{}) {
	logger.WithOptions(zap.AddCallerSkip(1)).Logw(lvl, msg, keysAndValues...)
}

// Debugw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
//
// When debug-level logging is disabled, this is much faster than
//
//	s.With(keysAndValues).Debug(msg)
func Debugw(msg string, keysAndValues ...interface{}) {
	logger.WithOptions(zap.AddCallerSkip(1)).Debugw(msg, keysAndValues...)
}

// Infow logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Infow(msg string, keysAndValues ...interface{}) {
	logger.WithOptions(zap.AddCallerSkip(1)).Infow(msg, keysAndValues...)
}

// Warnw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Warnw(msg string, keysAndValues ...interface{}) {
	logger.WithOptions(zap.AddCallerSkip(1)).Warnw(msg, keysAndValues...)
}

// Errorw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Errorw(msg string, keysAndValues ...interface{}) {
	logger.WithOptions(zap.AddCallerSkip(1)).Errorw(msg, keysAndValues...)
}

// DPanicw logs a message with some additional context. In development, the
// logger then panics. (See DPanicLevel for details.) The variadic key-value
// pairs are treated as they are in With.
func DPanicw(msg string, keysAndValues ...interface{}) {
	logger.WithOptions(zap.AddCallerSkip(1)).DPanicw(msg, keysAndValues...)
}

// Panicw logs a message with some additional context, then panics. The
// variadic key-value pairs are treated as they are in With.
func Panicw(msg string, keysAndValues ...interface{}) {
	logger.WithOptions(zap.AddCallerSkip(1)).Panicw(msg, keysAndValues...)
}

// Fatalw logs a message with some additional context, then calls os.Exit. The
// variadic key-value pairs are treated as they are in With.
func Fatalw(msg string, keysAndValues ...interface{}) {
	logger.WithOptions(zap.AddCallerSkip(1)).Fatalw(msg, keysAndValues...)
}

// Logln logs a message at provided level.
// Spaces are always added between arguments.
func Logln(lvl zapcore.Level, args ...interface{}) {
	logger.WithOptions(zap.AddCallerSkip(1)).Logln(lvl, args...)
}

// Debugln logs a message at [DebugLevel].
// Spaces are always added between arguments.
func Debugln(args ...interface{}) {
	logger.WithOptions(zap.AddCallerSkip(1)).Debugln(args...)
}

// Infoln logs a message at [InfoLevel].
// Spaces are always added between arguments.
func Infoln(args ...interface{}) {
	logger.WithOptions(zap.AddCallerSkip(1)).Infoln(args...)
}

// Warnln logs a message at [WarnLevel].
// Spaces are always added between arguments.
func Warnln(args ...interface{}) {
	logger.WithOptions(zap.AddCallerSkip(1)).Warnln(args...)
}

// Errorln logs a message at [ErrorLevel].
// Spaces are always added between arguments.
func Errorln(args ...interface{}) {
	logger.WithOptions(zap.AddCallerSkip(1)).Errorln(args...)
}

// DPanicln logs a message at [DPanicLevel].
// In development, the logger then panics. (See [DPanicLevel] for details.)
// Spaces are always added between arguments.
func DPanicln(args ...interface{}) {
	logger.WithOptions(zap.AddCallerSkip(1)).DPanicln(args...)
}

// Panicln logs a message at [PanicLevel] and panics.
// Spaces are always added between arguments.
func Panicln(args ...interface{}) {
	logger.WithOptions(zap.AddCallerSkip(1)).Panicln(args...)
}

// Fatalln logs a message at [FatalLevel] and calls os.Exit.
// Spaces are always added between arguments.
func Fatalln(args ...interface{}) {
	logger.WithOptions(zap.AddCallerSkip(1)).Fatalln(args...)
}

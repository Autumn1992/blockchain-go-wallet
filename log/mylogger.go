package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type MyLogger struct {
	*zap.SugaredLogger
	LoggerName string
}

func NewLogger(name string) *MyLogger {
	var logFileName = name + ".log"
	var sugaredLogger = newZapLogger(false, true, name)
	return &MyLogger{sugaredLogger, logFileName}
}

func (logger *MyLogger) Close() {
	logger.Sync()
}

// Log logs the provided arguments at provided level.
// Spaces are added between arguments when neither is a string.
func (logger *MyLogger) Log(lvl zapcore.Level, args ...interface{}) {
	if mode == 1 {
		logger.SugaredLogger.WithOptions(zap.AddCallerSkip(1)).Log(lvl, args...)
	} else {
		logger.SugaredLogger.Log(lvl, formatArgs(true, args, 1))
	}
}

// Debug logs the provided arguments at [DebugLevel].
// Spaces are added between arguments when neither is a string.
func (logger *MyLogger) Debug(args ...interface{}) {
	if mode == 1 {
		logger.SugaredLogger.WithOptions(zap.AddCallerSkip(1)).Debug(args...)
	} else {
		logger.SugaredLogger.Debug(formatArgs(true, args, 1))
	}
}

// Info logs the provided arguments at [InfoLevel].
// Spaces are added between arguments when neither is a string.
func (logger *MyLogger) Info(args ...interface{}) {
	if mode == 1 {
		logger.SugaredLogger.WithOptions(zap.AddCallerSkip(1)).Info(args...)
	} else {
		logger.SugaredLogger.Info(formatArgs(true, args, 1))
	}
}

// Warn logs the provided arguments at [WarnLevel].
// Spaces are added between arguments when neither is a string.
func (logger *MyLogger) Warn(args ...interface{}) {
	if mode == 1 {
		logger.SugaredLogger.WithOptions(zap.AddCallerSkip(1)).Warn(args...)
	} else {
		logger.SugaredLogger.Warn(formatArgs(true, args, 1))
	}
}

// Error logs the provided arguments at [ErrorLevel].
// Spaces are added between arguments when neither is a string.
func (logger *MyLogger) Error(args ...interface{}) {
	if mode == 1 {
		logger.SugaredLogger.WithOptions(zap.AddCallerSkip(1)).Error(args...)
	} else {
		logger.SugaredLogger.Error(formatArgs(true, args, 1))
	}
}

// DPanic logs the provided arguments at [DPanicLevel].
// In development, the logger then panics. (See [DPanicLevel] for details.)
// Spaces are added between arguments when neither is a string.
func (logger *MyLogger) DPanic(args ...interface{}) {
	logger.SugaredLogger.WithOptions(zap.AddCallerSkip(1)).DPanic(args...)
}

// Panic constructs a message with the provided arguments and panics.
// Spaces are added between arguments when neither is a string.
func (logger *MyLogger) Panic(args ...interface{}) {
	logger.SugaredLogger.WithOptions(zap.AddCallerSkip(1)).Panic(args...)
}

// Fatal constructs a message with the provided arguments and calls os.Exit.
// Spaces are added between arguments when neither is a string.
func (logger *MyLogger) Fatal(args ...interface{}) {
	//logger.SugaredLogger.WithOptions(zap.AddCallerSkip(1)).Fatal(args...)
	logger.SugaredLogger.Fatal(formatArgs(true, args, 1))
}

// Logf formats the message according to the format specifier
// and logs it at provided level.
func (logger *MyLogger) Logf(lvl zapcore.Level, template string, args ...interface{}) {
	//logger.SugaredLogger.WithOptions(zap.AddCallerSkip(1)).Logf(lvl, template, args...)
	Log(lvl, fmt.Sprintf(template, args...))
}

// Debugf formats the message according to the format specifier
// and logs it at [DebugLevel].
func (logger *MyLogger) Debugf(template string, args ...interface{}) {
	//logger.SugaredLogger.WithOptions(zap.AddCallerSkip(1)).Debugf(template, args...)
	Debug(fmt.Sprintf(template, args...))
}

// Infof formats the message according to the format specifier
// and logs it at [InfoLevel].
func (logger *MyLogger) Infof(template string, args ...interface{}) {
	//logger.SugaredLogger.WithOptions(zap.AddCallerSkip(1)).Infof(template, args...)
	Info(fmt.Sprintf(template, args...))
}

// Warnf formats the message according to the format specifier
// and logs it at [WarnLevel].
func (logger *MyLogger) Warnf(template string, args ...interface{}) {
	//logger.SugaredLogger.WithOptions(zap.AddCallerSkip(1)).Warnf(template, args...)
	Warn(fmt.Sprintf(template, args...))
}

// Errorf formats the message according to the format specifier
// and logs it at [ErrorLevel].
func (logger *MyLogger) Errorf(template string, args ...interface{}) {
	//logger.SugaredLogger.WithOptions(zap.AddCallerSkip(1)).Errorf(template, args...)
	Error(fmt.Sprintf(template, args...))
}

// DPanicf formats the message according to the format specifier
// and logs it at [DPanicLevel].
// In development, the logger then panics. (See [DPanicLevel] for details.)
func (logger *MyLogger) DPanicf(template string, args ...interface{}) {
	//logger.SugaredLogger.WithOptions(zap.AddCallerSkip(1)).DPanicf(template, args...)
	DPanic(fmt.Sprintf(template, args...))
}

// Panicf formats the message according to the format specifier
// and panics.
func (logger *MyLogger) Panicf(template string, args ...interface{}) {
	//logger.SugaredLogger.WithOptions(zap.AddCallerSkip(1)).Panicf(template, args...)
	Panic(fmt.Sprintf(template, args...))
}

// Fatalf formats the message according to the format specifier
// and calls os.Exit.
func (logger *MyLogger) Fatalf(template string, args ...interface{}) {
	//logger.SugaredLogger.WithOptions(zap.AddCallerSkip(1)).Fatalf(template, args...)
	Fatal(fmt.Sprintf(template, args...))
}

// Logw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (logger *MyLogger) Logw(lvl zapcore.Level, msg string, keysAndValues ...interface{}) {
	logger.SugaredLogger.WithOptions(zap.AddCallerSkip(1)).Logw(lvl, msg, keysAndValues...)
}

// Debugw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
//
// When debug-level logging is disabled, this is much faster than
//
//	s.With(keysAndValues).Debug(msg)
func (logger *MyLogger) Debugw(msg string, keysAndValues ...interface{}) {
	logger.SugaredLogger.WithOptions(zap.AddCallerSkip(1)).Debugw(msg, keysAndValues...)
}

// Infow logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (logger *MyLogger) Infow(msg string, keysAndValues ...interface{}) {
	logger.SugaredLogger.WithOptions(zap.AddCallerSkip(1)).Infow(msg, keysAndValues...)
}

// Warnw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (logger *MyLogger) Warnw(msg string, keysAndValues ...interface{}) {
	logger.SugaredLogger.WithOptions(zap.AddCallerSkip(1)).Warnw(msg, keysAndValues...)
}

// Errorw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (logger *MyLogger) Errorw(msg string, keysAndValues ...interface{}) {
	logger.SugaredLogger.WithOptions(zap.AddCallerSkip(1)).Errorw(msg, keysAndValues...)
}

// DPanicw logs a message with some additional context. In development, the
// logger then panics. (See DPanicLevel for details.) The variadic key-value
// pairs are treated as they are in With.
func (logger *MyLogger) DPanicw(msg string, keysAndValues ...interface{}) {
	logger.SugaredLogger.WithOptions(zap.AddCallerSkip(1)).DPanicw(msg, keysAndValues...)
}

// Panicw logs a message with some additional context, then panics. The
// variadic key-value pairs are treated as they are in With.
func (logger *MyLogger) Panicw(msg string, keysAndValues ...interface{}) {
	logger.SugaredLogger.WithOptions(zap.AddCallerSkip(1)).Panicw(msg, keysAndValues...)
}

// Fatalw logs a message with some additional context, then calls os.Exit. The
// variadic key-value pairs are treated as they are in With.
func (logger *MyLogger) Fatalw(msg string, keysAndValues ...interface{}) {
	logger.SugaredLogger.WithOptions(zap.AddCallerSkip(1)).Fatalw(msg, keysAndValues...)
}

// Logln logs a message at provided level.
// Spaces are always added between arguments.
func (logger *MyLogger) Logln(lvl zapcore.Level, args ...interface{}) {
	logger.SugaredLogger.WithOptions(zap.AddCallerSkip(1)).Logln(lvl, args...)
}

// Debugln logs a message at [DebugLevel].
// Spaces are always added between arguments.
func (logger *MyLogger) Debugln(args ...interface{}) {
	logger.SugaredLogger.WithOptions(zap.AddCallerSkip(1)).Debugln(args...)
}

// Infoln logs a message at [InfoLevel].
// Spaces are always added between arguments.
func (logger *MyLogger) Infoln(args ...interface{}) {
	logger.SugaredLogger.WithOptions(zap.AddCallerSkip(1)).Infoln(args...)
}

// Warnln logs a message at [WarnLevel].
// Spaces are always added between arguments.
func (logger *MyLogger) Warnln(args ...interface{}) {
	logger.SugaredLogger.WithOptions(zap.AddCallerSkip(1)).Warnln(args...)
}

// Errorln logs a message at [ErrorLevel].
// Spaces are always added between arguments.
func (logger *MyLogger) Errorln(args ...interface{}) {
	logger.SugaredLogger.WithOptions(zap.AddCallerSkip(1)).Errorln(args...)
}

// DPanicln logs a message at [DPanicLevel].
// In development, the logger then panics. (See [DPanicLevel] for details.)
// Spaces are always added between arguments.
func (logger *MyLogger) DPanicln(args ...interface{}) {
	logger.SugaredLogger.WithOptions(zap.AddCallerSkip(1)).DPanicln(args...)
}

// Panicln logs a message at [PanicLevel] and panics.
// Spaces are always added between arguments.
func (logger *MyLogger) Panicln(args ...interface{}) {
	logger.SugaredLogger.WithOptions(zap.AddCallerSkip(1)).Panicln(args...)
}

// Fatalln logs a message at [FatalLevel] and calls os.Exit.
// Spaces are always added between arguments.
func (logger *MyLogger) Fatalln(args ...interface{}) {
	logger.SugaredLogger.WithOptions(zap.AddCallerSkip(1)).Fatalln(args...)
}

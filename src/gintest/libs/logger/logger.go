package logger

import (
	"fmt"
	"io"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// logger 日志实例
var logger *zap.Logger

// Init 日志初始化
func Init(logFilePath string) {
	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		EncodeLevel: zapcore.CapitalLevelEncoder,
		TimeKey:     "ts",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		CallerKey:    "file",
		EncodeCaller: zapcore.ShortCallerEncoder,
	})

	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.WarnLevel
	})

	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel
	})

	infoWriter := getWriter(logFilePath + "/info.log")
	warnWriter := getWriter(logFilePath + "/error.log")

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(warnWriter), warnLevel),
	)

	logger = zap.New(core)
	defer logger.Sync()
}

// getWriter 使用rotatelogs切割日志
func getWriter(filename string) io.Writer {
	hook, err := rotatelogs.New(
		// 实际文件名 XXXX.YYmmDD
		filename+".%Y%m%d",
		rotatelogs.WithLinkName(filename),
		// 保存1年日志
		rotatelogs.WithMaxAge(time.Hour*24*365),
		// 每天切割一次
		rotatelogs.WithRotationTime(time.Hour*24),
	)

	if err != nil {
		panic(err)
	}
	return hook
}

// Info 常规日志，使用zap.Field
func Info(msg string, args ...zap.Field) {
	logger.Info(msg, args...)
}

// Warn 告警日志，使用zap.Field
func Warn(msg string, args ...zap.Field) {
	logger.Warn(msg, args...)
}

// Error 错误日志，使用zap.Field
func Error(msg string, args ...zap.Field) {
	logger.Error(msg, args...)
}

// Debug 调试日志，使用zap.Field
func Debug(msg string, args ...zap.Field) {
	logger.Debug(msg, args...)
}

// Infof 常规日志，使用interface{}
func Infof(msg string, args ...interface{}) {
	logMsg := fmt.Sprintf(msg, args...)
	logger.Info(logMsg)
}

// Warnf 告警日志，使用interface{}
func Warnf(msg string, args ...interface{}) {
	logMsg := fmt.Sprintf(msg, args...)
	logger.Warn(logMsg)
}

// Errorf 错误日志，使用interface{}
func Errorf(msg string, args ...interface{}) {
	logMsg := fmt.Sprintf(msg, args...)
	logger.Error(logMsg)
}

// Debugf 调试日志，使用interface{}
func Debugf(msg string, args ...interface{}) {
	logMsg := fmt.Sprintf(msg, args...)
	logger.Debug(logMsg)
}

package logger

import (
	"io"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger 日志实例
var Logger *zap.Logger

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
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
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

	Logger = zap.New(core, zap.AddCaller())
	Logger.Info("logger start")
	defer Logger.Sync()
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

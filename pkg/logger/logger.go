package logger

import (
	"base_go_be/pkg/setting"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

type LogZap struct {
	*zap.Logger
}

func NewLogger(config setting.LoggerSetting) *LogZap {
	logLevel := config.LogLevel
	// debug-> info-> warn-> error-> fatal->panic
	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	default:
		level = zapcore.InfoLevel
	}
	encoder := getEncoderLogger()
	hook := lumberjack.Logger{
		Filename:   config.FileLogName,
		MaxSize:    config.MaxSize, // megabytes
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge, //days
		Compress:   config.Compress,
	}
	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), level)
	return &LogZap{zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))}
}

func getEncoderLogger() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	// 1735870446.0274096 -> 2025-01-03T09:14:06.027+0700
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// ts -> Time
	encoderConfig.TimeKey = "time"
	// from INFO
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// "caller":"cli/main.log.go:19"
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

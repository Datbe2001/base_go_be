package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func main() {
	encoder := getEncoderLogger()
	sync := getWriteSyncer()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())
	logger.Info("start", zap.Int("line", 1))
	logger.Error("start", zap.Int("line", 2))
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

func getWriteSyncer() zapcore.WriteSyncer {
	file, _ := os.OpenFile("./log/log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	syncer := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stdout)
	return zapcore.NewMultiWriteSyncer(syncConsole, syncer)
}

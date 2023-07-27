package main

import (
	"fmt"
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// log level
// Debug, Info, Warn, Error, DPanic, Panic, Fatal

func NewProduction() {
	var logger *zap.Logger
	logger, _ = zap.NewProduction()
	des := "这是一条日志数据"
	logger.Debug("debug", zap.String("des", des))
	logger.Info("info", zap.String("des", des))
	logger.Error("error", zap.String("des", des))
}

func NewDevelopment() {
	var logger *zap.Logger
	logger, _ = zap.NewDevelopment()
	des := "这是一条日志数据"
	logger.Debug("debug", zap.Any("des", des))
	logger.Info("info", zap.Any("des", des))
	logger.Error("error", zap.Any("des", des))
}

func NewExample() {
	var logger *zap.Logger
	logger = zap.NewExample()
	des := "这是一条日志数据"
	logger.Debug("debug", zap.String("des", des))
	logger.Info("info", zap.String("des", des))
	logger.Error("error", zap.String("des", des))
}

func SugarLogger() {
	var logger *zap.Logger
	logger, _ = zap.NewProduction()

	var sugarLogger *zap.SugaredLogger
	sugarLogger = logger.Sugar()
	defer sugarLogger.Sync()

	log := map[string]interface{}{
		"code": 10000,
		"msg":  "success",
		"data": "这是一条日志",
	}
	sugarLogger.Info(log)

	sugarLogger.Info("这是一条日志")
	sugarLogger.Infof("这是一条日志 code=%d", 10000)
}

func CustomLogger() {
	// Encoder
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder   // 修改时间编码器
	encoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder // 使用大写字母记录日志级别
	encoder := zapcore.NewConsoleEncoder(encoderCfg)
	// encoder := zapcore.NewJSONEncoder(encoderCfg)

	// WriterSyncer
	// file, _ := os.Create("./zap.log")
	file, err := os.OpenFile("./zap.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	ws := io.MultiWriter(file, os.Stdout)
	writeSyncer := zapcore.AddSync(ws)

	// Log Level
	// zapcore
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	// logger
	// 记录函数调用信息
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	// Sugar Logger
	sugarLogger := logger.Sugar()
	defer sugarLogger.Sync()

	sugarLogger.Info("这是一条日志")

	// err = errors.New("这是一个错误")
	if _, err = os.Open("test.log"); err != nil {
		sugarLogger.Errorf("这是一条日志 err: %v", err)
	}

	log := map[string]interface{}{
		"code": 10000,
		"msg":  "success",
		"data": "这是一条日志",
	}
	sugarLogger.Debug(log)
	sugarLogger.Debugf("%#v", log)
}

func main() {
	// NewProduction()
	// NewDevelopment()
	// NewExample()
	// SugarLogger()
	CustomLogger()
}

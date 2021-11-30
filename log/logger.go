package log

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"syscall"
)

const (
	BaseLog = "/var/log"
	AppLog  = BaseLog + "/app"
)

var LOG *zap.Logger

func init() {
	mask := syscall.Umask(0)
	defer syscall.Umask(mask)
	err := os.MkdirAll(AppLog, os.ModePerm)
	if err != nil {
		panic(fmt.Sprintf("Failed to create dir %v.", AppLog))
		return
	}
	LOG = createLogger()
}

func createLogger() *zap.Logger {
	// First, define our level-handling logic.
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	stdPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.InfoLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.DebugLevel
	})

	// write to the files or the stdout[stderr]
	fileDebug := getWriteSyncer(AppLog + "/debug.log")
	fileStd := getWriteSyncer(AppLog + "/app.log")
	fileError := getWriteSyncer(AppLog + "/error.log")
	consoleDebug := zapcore.Lock(os.Stdout)

	enc := zap.NewProductionEncoderConfig()
	enc.TimeKey = "time"
	enc.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000Z0700")
	fileEncoder := zapcore.NewJSONEncoder(enc)
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	// Join the outputs, encoders, and level-handling functions into
	// zapcore.Cores, then tee the four cores together.
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, fileError, highPriority),
		zapcore.NewCore(fileEncoder, fileStd, stdPriority),
		zapcore.NewCore(fileEncoder, fileDebug, lowPriority),
		zapcore.NewCore(consoleEncoder, consoleDebug, lowPriority),
	)

	// From a zapcore.Core, it's easy to construct a Logger.
	//Open development mode, stack trace
	caller := zap.AddCaller()
	//Open file and line number
	development := zap.Development()
	fields := zap.Fields(zap.String("app", "octopus"))
	logger := zap.New(core, caller, development, fields)

	// https://github.com/uber-go/zap/issues/1026
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil && !errors.Is(err, syscall.ENOTTY) {
			panic(err)
		}
	}(logger)
	return logger
}

func getWriteSyncer(path string) zapcore.WriteSyncer {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	return file
}

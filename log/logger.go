package log

import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)

const (
    BaseLog = "/var/log"
    AppLog  = BaseLog + "/app"
)

var LOG *zap.Logger

func init() {
    LOG = newLogger()
}

func newLogger() *zap.Logger {
    config := zap.Config{
        Level:             zap.NewAtomicLevelAt(zap.DebugLevel),
        Development:       false,
        DisableCaller:     false,
        DisableStacktrace: false,
        Sampling:          nil,
        Encoding:          "json",
        EncoderConfig: zapcore.EncoderConfig{
            MessageKey:       "msg",
            LevelKey:         "level",
            TimeKey:          "time",
            NameKey:          "logger",
            CallerKey:        "file",
            FunctionKey:      "func",
            StacktraceKey:    "stacktrace",
            LineEnding:       zapcore.DefaultLineEnding,
            EncodeLevel:      zapcore.LowercaseLevelEncoder,
            EncodeTime:       zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000Z0700"),
            EncodeDuration:   zapcore.SecondsDurationEncoder,
            EncodeCaller:     zapcore.ShortCallerEncoder,
            EncodeName:       zapcore.FullNameEncoder,
            ConsoleSeparator: "",
        },
        OutputPaths:      []string{"stdout", AppLog + "/octopus.log"},
        ErrorOutputPaths: []string{"stderr", AppLog + "/octopus-error.log"},
        InitialFields: map[string]interface{}{
            "app": "octopus",
        },
    }
    logger, err := config.Build()
    if err != nil {
        panic(err)
    }
    defer func(logger *zap.Logger) {
        err := logger.Sync()
        if err != nil {
            panic(err)
        }
    }(logger)
    return logger
}

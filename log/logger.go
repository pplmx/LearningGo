package log

import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "os"
)

const (
    BaseLog = "/var/log"
    AppLog  = BaseLog + "/app"
)

var LOG *zap.Logger

func init() {
    LOG = createLogger()
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

func createLogger() *zap.Logger {
    // First, define our level-handling logic.
    highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
        return lvl >= zapcore.ErrorLevel
    })
    lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
        return lvl < zapcore.ErrorLevel
    })

    // Assume that we have clients for two Kafka topics. The clients implement
    // zapcore.WriteSyncer and are safe for concurrent use. (If they only
    // implement io.Writer, we can use zapcore.AddSync to add a no-op Sync
    // method. If they're not safe for concurrent use, we can add a protecting
    // mutex with zapcore.Lock.)
    fileDebug := getWriteSyncer("/var/log/app/app.log")
    fileError := getWriteSyncer("/var/log/app/error.log")

    // High-priority output should also go to standard error, and low-priority
    // output should also go to standard out.
    consoleDebug := zapcore.Lock(os.Stdout)
    consoleError := zapcore.Lock(os.Stderr)

    // Optimize the Kafka output for machine consumption and the console output
    // for human operators.
    fileEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
    consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

    // Join the outputs, encoders, and level-handling functions into
    // zapcore.Cores, then tee the four cores together.
    core := zapcore.NewTee(
        zapcore.NewCore(fileEncoder, fileError, highPriority),
        zapcore.NewCore(consoleEncoder, consoleError, highPriority),
        zapcore.NewCore(fileEncoder, fileDebug, lowPriority),
        zapcore.NewCore(consoleEncoder, consoleDebug, lowPriority),
    )

    // From a zapcore.Core, it's easy to construct a Logger.
    logger := zap.New(core)
    defer func(logger *zap.Logger) {
        err := logger.Sync()
        if err != nil {
            panic(err)
        }
    }(logger)
    return logger
}

func getWriteSyncer(path string) zapcore.WriteSyncer {
    file, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 644)
    if err != nil {
        panic(err)
    }
    return zapcore.AddSync(file)
}

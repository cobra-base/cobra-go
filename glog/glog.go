package glog

import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "gopkg.in/natefinch/lumberjack.v2"
    "os"
)

type Config struct {
    LogName      string `json:"logName"`
    LogLevel     string `json:"logLevel"`
    LogDir       string `json:"logDir"`
    WriteJson    bool   `json:"writeJson"`
    WriteConsole bool   `json:"writeConsole"`
    MaxSize      int    `json:"maxSize"`
    MaxBackups   int    `json:"maxBackups"`
    MaxAge       int    `json:"maxAge"`
}

type Logger struct {
    // level string // debug,info,warn,error
    sugar *zap.SugaredLogger

    debugMode bool
}

var logger = &Logger{}

func Init(logConf *Config) error {
    level := zapcore.WarnLevel
    switch logConf.LogLevel {
    case "debug":
        level = zapcore.DebugLevel
    case "info":
        level = zapcore.InfoLevel
    case "warn":
        level = zapcore.WarnLevel
    case "error":
        level = zapcore.ErrorLevel
    }

    encoderConfig := zap.NewProductionEncoderConfig()
    encoderConfig.ConsoleSeparator = " "
    encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
    encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

    encoder := zapcore.NewConsoleEncoder(encoderConfig)
    if logConf.WriteJson {
        encoder = zapcore.NewJSONEncoder(encoderConfig)
    }

    writeSyncer := zapcore.AddSync(os.Stdout)
    coreConsole := zapcore.NewCore(encoder, writeSyncer, level)

    traceLogger := &lumberjack.Logger{
        Filename:   logConf.LogDir + "/" + logConf.LogName + ".log",
        MaxSize:    logConf.MaxSize, // MB
        MaxBackups: logConf.MaxBackups,
        MaxAge:     logConf.MaxAge,
        Compress:   false,
    }
    coreTrace := zapcore.NewCore(encoder, zapcore.AddSync(traceLogger), level)

    var core zapcore.Core
    if logConf.WriteConsole {
        core = zapcore.NewTee(coreConsole, coreTrace)
    } else {
        core = zapcore.NewTee(coreTrace)
    }

    l := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.DPanicLevel))
    logger.sugar = l.Sugar()

    if level == zapcore.DebugLevel {
        logger.debugMode = true
    }

    return nil
}

func DebugMode() bool {
    return logger.debugMode
}

func Debugw(msg string, keysAndValues ...interface{}) {
    logger.sugar.Debugw(msg, keysAndValues...)
}

func Infow(msg string, keysAndValues ...interface{}) {
    logger.sugar.Infow(msg, keysAndValues...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
    logger.sugar.Warnw(msg, keysAndValues...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
    logger.sugar.Errorw(msg, keysAndValues...)
}

package zaplog

import (
	"os"
	"time"

	lumberjack "gopkg.in/natefinch/lumberjack.v2"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Zaplog struct {
	Logger *zap.Logger
}

// Init 系统日志初始化
func (zaplog *Zaplog) Init() {
	logT1 := time.Now()
	logDate := logT1.Format("20060102")
	// lumberjack.Logger is already safe for concurrent use, so we don't need to
	// lock it.
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "log/logfile" + logDate + ".log",
		MaxSize:    1, // megabytes
		MaxBackups: 3,
		MaxAge:     1, // days
	})
	atom := zap.NewAtomicLevel()
	atom.SetLevel(zap.DebugLevel)

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderCfg.CallerKey = "caller"
	encoderCfg.EncodeCaller = zapcore.ShortCallerEncoder

	zaplog.Logger = zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		w,
		atom,
	), zap.AddCaller())

	consoleEncoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	})

	zaplog.Logger = zap.New(zapcore.NewTee(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderCfg),
			w,
			atom,
		),
		zapcore.NewCore(
			consoleEncoder,
			zapcore.AddSync(os.Stdout),
			atom,
		),
	), zap.AddCaller())
}

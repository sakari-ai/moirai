package log

import (
	"os"
	"time"

	"github.com/sakari-ai/moirai/log/field"
	"github.com/sakari-ai/moirai/log/mode"
	"go.elastic.co/apm/module/apmzap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	encodingJSON    = "json"
	encodingConsole = "console"
	stderr          = "stderr"
	keyTime         = "time"
	keyLevel        = "level"
	keyName         = "name"
	keyCaller       = "caller"
	keyMsg          = "msg"
	keyStacktrace   = "stacktrace"
)

type Logger struct {
	*zap.Logger
	Mode mode.Mode
}

var (
	l = New(withDevelopmentMode())
)

func init() {
	if os.Getenv("ENV") == "prod" {
		l = New(WithModeFromString("production"))
	}
}

func New(options ...Option) *Logger {
	logger := &Logger{}
	for _, option := range options {
		option(logger)
	}

	switch logger.Mode {
	case mode.Development:
		logger.Logger = newDevelopment()
	case mode.Production:
		logger.Logger = newProduction()
	}
	return logger
}

func Init(logger *Logger) {
	l = logger
}

func Close() {
	l.Sync()
}

func Debug(msg string, fields ...field.Field) {
	l.Debug(msg, fields...)
}

func Info(msg string, fields ...field.Field) {
	l.Info(msg, fields...)
}

func Warn(msg string, fields ...field.Field) {
	l.Warn(msg, fields...)
}

func Error(msg string, fields ...field.Field) {
	l.Error(msg, fields...)
}

func DPanic(msg string, fields ...field.Field) {
	l.DPanic(msg, fields...)
}

func Panic(msg string, fields ...field.Field) {
	l.Panic(msg, fields...)
}

func Fatal(msg string, fields ...field.Field) {
	l.Fatal(msg, fields...)
}

func newProduction() *zap.Logger {
	var (
		config  = newProductionConfig()
		options = []zap.Option{zap.AddCallerSkip(1)}
		l, _    = config.Build(options...)
	)
	return l
}

func newDevelopment() *zap.Logger {
	var (
		config  = newDevelopmentConfig()
		options = []zap.Option{zap.AddCallerSkip(1), zap.WrapCore((&apmzap.Core{}).WrapCore)}
		l, _    = config.Build(options...)
	)
	return l
}

func newProductionEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        keyTime,
		LevelKey:       keyLevel,
		NameKey:        keyName,
		CallerKey:      keyCaller,
		MessageKey:     keyMsg,
		StacktraceKey:  keyStacktrace,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     rfc3339TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func newProductionConfig() zap.Config {
	return zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         encodingJSON,
		EncoderConfig:    newProductionEncoderConfig(),
		OutputPaths:      []string{stderr},
		ErrorOutputPaths: []string{stderr},
	}
}

func newDevelopmentEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        keyTime,
		LevelKey:       keyLevel,
		NameKey:        keyName,
		CallerKey:      keyCaller,
		MessageKey:     keyMsg,
		StacktraceKey:  keyStacktrace,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     rfc3339TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func newDevelopmentConfig() zap.Config {
	return zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:      true,
		Encoding:         encodingConsole,
		EncoderConfig:    newDevelopmentEncoderConfig(),
		OutputPaths:      []string{stderr},
		ErrorOutputPaths: []string{stderr},
	}
}

func rfc3339TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(time.RFC3339))
}

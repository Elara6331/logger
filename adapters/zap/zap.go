package adapters

import (
	"fmt"
	"time"

	"go.arsenm.dev/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _ logger.Logger = &ZapLogger{}
var _ logger.LogBuilder = &ZapLogBuilder{}

type ZapLogger struct {
	Logger *zap.Logger
}

func New(z *zap.Logger) *ZapLogger {
	return &ZapLogger{z.WithOptions(zap.AddCallerSkip(1))}
}

// NoPanic is a no-op because Zap does not provide this functionality
func (zl *ZapLogger) NoPanic() {}

// NoExit is a no-op because Zap does not provide this functionality
func (zl *ZapLogger) NoExit() {}

// SetLevel sets the log level of the logger.
//
// Because zap only allows level increases, it can only increase
// the level. If the given level is lower than the current log level,
// SetLevel will be a no-op.
func (zl *ZapLogger) SetLevel(lvl logger.LogLevel) {
	var zlvl zapcore.Level
	switch lvl {
	case logger.LogLevelDebug:
		zlvl = zapcore.DebugLevel
	case logger.LogLevelInfo:
		zlvl = zapcore.InfoLevel
	case logger.LogLevelWarn:
		zlvl = zapcore.WarnLevel
	case logger.LogLevelError:
		zlvl = zapcore.WarnLevel
	case logger.LogLevelFatal:
		zlvl = zapcore.FatalLevel
	case logger.LogLevelPanic:
		zlvl = zapcore.PanicLevel
	}

	zl.Logger = zl.Logger.WithOptions(zap.IncreaseLevel(zlvl))
}

// Debug creates a new debug event with the given message
func (zl *ZapLogger) Debug(msg string) logger.LogBuilder {
	return &ZapLogBuilder{
		logger: zl.Logger,
		lvl:    zapcore.DebugLevel,
		msg:    msg,
	}
}

// Debugf creates a new debug event with the formatted message
func (zl *ZapLogger) Debugf(format string, v ...any) logger.LogBuilder {
	return &ZapLogBuilder{
		logger: zl.Logger,
		lvl:    zapcore.DebugLevel,
		msg:    fmt.Sprintf(format, v...),
	}
}

// Info creates a new info event with the given message
func (zl *ZapLogger) Info(msg string) logger.LogBuilder {
	return &ZapLogBuilder{
		logger: zl.Logger,
		lvl:    zapcore.InfoLevel,
		msg:    msg,
	}
}

// Infof creates a new info event with the formatted message
func (zl *ZapLogger) Infof(format string, v ...any) logger.LogBuilder {
	return &ZapLogBuilder{
		logger: zl.Logger,
		lvl:    zapcore.InfoLevel,
		msg:    fmt.Sprintf(format, v...),
	}
}

// Warn creates a new warn event with the given message
func (zl *ZapLogger) Warn(msg string) logger.LogBuilder {
	return &ZapLogBuilder{
		logger: zl.Logger,
		lvl:    zapcore.WarnLevel,
		msg:    msg,
	}
}

// Warnf creates a new warn event with the formatted message
func (zl *ZapLogger) Warnf(format string, v ...any) logger.LogBuilder {
	return &ZapLogBuilder{
		logger: zl.Logger,
		lvl:    zapcore.WarnLevel,
		msg:    fmt.Sprintf(format, v...),
	}
}

// Error creates a new error event with the given message
func (zl *ZapLogger) Error(msg string) logger.LogBuilder {
	return &ZapLogBuilder{
		logger: zl.Logger,
		lvl:    zapcore.ErrorLevel,
		msg:    msg,
	}
}

// Errorf creates a new error event with the formatted message
func (zl *ZapLogger) Errorf(format string, v ...any) logger.LogBuilder {
	return &ZapLogBuilder{
		logger: zl.Logger,
		lvl:    zapcore.ErrorLevel,
		msg:    fmt.Sprintf(format, v...),
	}
}

// Fatal creates a new fatal event with the given message
func (zl *ZapLogger) Fatal(msg string) logger.LogBuilder {
	return &ZapLogBuilder{
		logger: zl.Logger,
		lvl:    zapcore.FatalLevel,
		msg:    msg,
	}
}

// Fatalf creates a new fatal event with the formatted message
func (zl *ZapLogger) Fatalf(format string, v ...any) logger.LogBuilder {
	return &ZapLogBuilder{
		logger: zl.Logger,
		lvl:    zapcore.FatalLevel,
		msg:    fmt.Sprintf(format, v...),
	}
}

// Panic creates a new panic event with the given message
func (zl *ZapLogger) Panic(msg string) logger.LogBuilder {
	return &ZapLogBuilder{
		logger: zl.Logger,
		lvl:    zapcore.PanicLevel,
		msg:    msg,
	}
}

// Panicf creates a new panic event with the formatted message
func (zl *ZapLogger) Panicf(format string, v ...any) logger.LogBuilder {
	return &ZapLogBuilder{
		logger: zl.Logger,
		lvl:    zapcore.PanicLevel,
		msg:    fmt.Sprintf(format, v...),
	}
}

type ZapLogBuilder struct {
	logger *zap.Logger
	lvl    zapcore.Level
	msg    string
	fields []zap.Field
}

func NewZapLogBuilder(logger *zap.Logger) logger.LogBuilder {
	return &ZapLogBuilder{logger: logger}
}

func (b *ZapLogBuilder) Int(key string, value int) logger.LogBuilder {
	b.fields = append(b.fields, zap.Int(key, value))
	return b
}

func (b *ZapLogBuilder) Int8(key string, value int8) logger.LogBuilder {
	b.fields = append(b.fields, zap.Int8(key, value))
	return b
}

func (b *ZapLogBuilder) Int16(key string, value int16) logger.LogBuilder {
	b.fields = append(b.fields, zap.Int16(key, value))
	return b
}

func (b *ZapLogBuilder) Int32(key string, value int32) logger.LogBuilder {
	b.fields = append(b.fields, zap.Int32(key, value))
	return b
}

func (b *ZapLogBuilder) Int64(key string, value int64) logger.LogBuilder {
	b.fields = append(b.fields, zap.Int64(key, value))
	return b
}

func (b *ZapLogBuilder) Uint(key string, value uint) logger.LogBuilder {
	b.fields = append(b.fields, zap.Uint(key, value))
	return b
}

func (b *ZapLogBuilder) Uint8(key string, value uint8) logger.LogBuilder {
	b.fields = append(b.fields, zap.Uint8(key, value))
	return b
}

func (b *ZapLogBuilder) Uint16(key string, value uint16) logger.LogBuilder {
	b.fields = append(b.fields, zap.Uint16(key, value))
	return b
}

func (b *ZapLogBuilder) Uint32(key string, value uint32) logger.LogBuilder {
	b.fields = append(b.fields, zap.Uint32(key, value))
	return b
}

func (b *ZapLogBuilder) Uint64(key string, value uint64) logger.LogBuilder {
	b.fields = append(b.fields, zap.Uint64(key, value))
	return b
}

func (b *ZapLogBuilder) Float32(key string, value float32) logger.LogBuilder {
	b.fields = append(b.fields, zap.Float32(key, value))
	return b
}

func (b *ZapLogBuilder) Float64(key string, value float64) logger.LogBuilder {
	b.fields = append(b.fields, zap.Float64(key, value))
	return b
}

func (b *ZapLogBuilder) Stringer(key string, value fmt.Stringer) logger.LogBuilder {
	b.fields = append(b.fields, zap.Stringer(key, value))
	return b
}

func (b *ZapLogBuilder) Bytes(key string, value []byte) logger.LogBuilder {
	b.fields = append(b.fields, zap.Binary(key, value))
	return b
}

func (b *ZapLogBuilder) Timestamp() logger.LogBuilder {
	b.fields = append(b.fields, zap.Time("timestamp", time.Now()))
	return b
}

func (b *ZapLogBuilder) Bool(key string, value bool) logger.LogBuilder {
	b.fields = append(b.fields, zap.Bool(key, value))
	return b
}

func (b *ZapLogBuilder) Str(key string, value string) logger.LogBuilder {
	b.fields = append(b.fields, zap.String(key, value))
	return b
}

func (b *ZapLogBuilder) Any(key string, value any) logger.LogBuilder {
	b.fields = append(b.fields, zap.Any(key, value))
	return b
}

func (b *ZapLogBuilder) Err(err error) logger.LogBuilder {
	b.fields = append(b.fields, zap.Error(err))
	return b
}

func (b *ZapLogBuilder) Send() {
	b.logger.Log(b.lvl, b.msg, b.fields...)
}

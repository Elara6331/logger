package zerolog

import (
	"fmt"
	"time"

	"github.com/rs/zerolog"
	"go.elara.ws/logger"
)

var (
	_ logger.Logger     = &ZerologLogger{}
	_ logger.LogBuilder = &ZerologLogBuilder{}
)

type ZerologLogger struct {
	logger zerolog.Logger
}

func New(z zerolog.Logger) *ZerologLogger {
	return &ZerologLogger{z}
}

// NoPanic is a no-op because Zerolog does not provide this functionality
func (l *ZerologLogger) NoPanic() {}

// NoPanic is a no-op because Zerolog does not provide this functionality
func (l *ZerologLogger) NoExit() {}

// SetLevel sets the log level of the logger.
func (l *ZerologLogger) SetLevel(level logger.LogLevel) {
	switch level {
	case logger.LogLevelDebug:
		l.logger = l.logger.Level(zerolog.DebugLevel)
	case logger.LogLevelInfo:
		l.logger = l.logger.Level(zerolog.InfoLevel)
	case logger.LogLevelWarn:
		l.logger = l.logger.Level(zerolog.WarnLevel)
	case logger.LogLevelError:
		l.logger = l.logger.Level(zerolog.ErrorLevel)
	case logger.LogLevelFatal:
		l.logger = l.logger.Level(zerolog.FatalLevel)
	case logger.LogLevelPanic:
		l.logger = l.logger.Level(zerolog.PanicLevel)
	}
}

func (l *ZerologLogger) Debug(msg string) logger.LogBuilder {
	return &ZerologLogBuilder{
		logEvent: l.logger.Debug(),
		msg:      msg,
	}
}

func (l *ZerologLogger) Debugf(format string, a ...any) logger.LogBuilder {
	return &ZerologLogBuilder{
		logEvent: l.logger.Debug(),
		msg:      fmt.Sprintf(format, a...),
	}
}

func (l *ZerologLogger) Info(msg string) logger.LogBuilder {
	return &ZerologLogBuilder{
		logEvent: l.logger.Info(),
		msg:      msg,
	}
}

func (l *ZerologLogger) Infof(format string, a ...any) logger.LogBuilder {
	return &ZerologLogBuilder{
		logEvent: l.logger.Info(),
		msg:      fmt.Sprintf(format, a...),
	}
}

func (l *ZerologLogger) Warn(msg string) logger.LogBuilder {
	return &ZerologLogBuilder{
		logEvent: l.logger.Warn(),
		msg:      msg,
	}
}

func (l *ZerologLogger) Warnf(format string, a ...any) logger.LogBuilder {
	return &ZerologLogBuilder{
		logEvent: l.logger.Warn(),
		msg:      fmt.Sprintf(format, a...),
	}
}

func (l *ZerologLogger) Error(msg string) logger.LogBuilder {
	return &ZerologLogBuilder{
		logEvent: l.logger.Error(),
		msg:      msg,
	}
}

func (l *ZerologLogger) Errorf(format string, a ...any) logger.LogBuilder {
	return &ZerologLogBuilder{
		logEvent: l.logger.Error(),
		msg:      fmt.Sprintf(format, a...),
	}
}

func (l *ZerologLogger) Fatal(msg string) logger.LogBuilder {
	return &ZerologLogBuilder{
		logEvent: l.logger.Fatal(),
		msg:      msg,
	}
}

func (l *ZerologLogger) Fatalf(format string, a ...any) logger.LogBuilder {
	return &ZerologLogBuilder{
		logEvent: l.logger.Fatal(),
		msg:      fmt.Sprintf(format, a...),
	}
}

func (l *ZerologLogger) Panic(msg string) logger.LogBuilder {
	return &ZerologLogBuilder{
		logEvent: l.logger.Panic(),
		msg:      msg,
	}
}

func (l *ZerologLogger) Panicf(format string, a ...any) logger.LogBuilder {
	return &ZerologLogBuilder{
		logEvent: l.logger.Panic(),
		msg:      fmt.Sprintf(format, a...),
	}
}

type ZerologLogBuilder struct {
	logEvent *zerolog.Event
	msg      string
}

func (b *ZerologLogBuilder) Int(key string, value int) logger.LogBuilder {
	b.logEvent.Int(key, value)
	return b
}

func (b *ZerologLogBuilder) Int8(key string, value int8) logger.LogBuilder {
	b.logEvent.Int8(key, value)
	return b
}

func (b *ZerologLogBuilder) Int16(key string, value int16) logger.LogBuilder {
	b.logEvent.Int16(key, value)
	return b
}

func (b *ZerologLogBuilder) Int32(key string, value int32) logger.LogBuilder {
	b.logEvent.Int32(key, value)
	return b
}

func (b *ZerologLogBuilder) Int64(key string, value int64) logger.LogBuilder {
	b.logEvent.Int64(key, value)
	return b
}

func (b *ZerologLogBuilder) Uint(key string, value uint) logger.LogBuilder {
	b.logEvent.Uint(key, value)
	return b
}

func (b *ZerologLogBuilder) Uint8(key string, value uint8) logger.LogBuilder {
	b.logEvent.Uint8(key, value)
	return b
}

func (b *ZerologLogBuilder) Uint16(key string, value uint16) logger.LogBuilder {
	b.logEvent.Uint16(key, value)
	return b
}

func (b *ZerologLogBuilder) Uint32(key string, value uint32) logger.LogBuilder {
	b.logEvent.Uint32(key, value)
	return b
}

func (b *ZerologLogBuilder) Uint64(key string, value uint64) logger.LogBuilder {
	b.logEvent.Uint64(key, value)
	return b
}

func (b *ZerologLogBuilder) Float32(key string, value float32) logger.LogBuilder {
	b.logEvent.Float32(key, value)
	return b
}

func (b *ZerologLogBuilder) Float64(key string, value float64) logger.LogBuilder {
	b.logEvent.Float64(key, value)
	return b
}

func (b *ZerologLogBuilder) Stringer(key string, value fmt.Stringer) logger.LogBuilder {
	b.logEvent.Str(key, value.String())
	return b
}

func (b *ZerologLogBuilder) Bytes(key string, value []byte) logger.LogBuilder {
	b.logEvent.Bytes(key, value)
	return b
}

func (b *ZerologLogBuilder) Timestamp() logger.LogBuilder {
	b.logEvent.Time("timestamp", time.Now())
	return b
}

func (b *ZerologLogBuilder) Bool(key string, value bool) logger.LogBuilder {
	b.logEvent.Bool(key, value)
	return b
}

func (b *ZerologLogBuilder) Str(key string, value string) logger.LogBuilder {
	b.logEvent.Str(key, value)
	return b
}

func (b *ZerologLogBuilder) Any(key string, value any) logger.LogBuilder {
	b.logEvent.Interface(key, value)
	return b
}

func (b *ZerologLogBuilder) Err(err error) logger.LogBuilder {
	b.logEvent.Err(err)
	return b
}

func (b *ZerologLogBuilder) Send() {
	b.logEvent.Msg(b.msg)
}

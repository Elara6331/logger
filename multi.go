package logger

import (
	"fmt"
	"os"
)

var _ Logger = (*MultiLogger)(nil)

// MultiLogger implements the Logger interface by
// writing to multiple underlying loggers sequentially.
type MultiLogger struct {
	Loggers []Logger
	noPanic bool
	noExit  bool
}

// NewMulti creates and returns a new MultiLogger
func NewMulti(l ...Logger) *MultiLogger {
	for _, logger := range l {
		logger.NoPanic()
		logger.NoExit()
	}
	return &MultiLogger{Loggers: l}
}

// NoExit prevents the logger from exiting on fatal events
func (ml *MultiLogger) NoExit() {
	ml.noExit = true
}

// NoPanic prevents the logger from panicking on panic events
func (ml *MultiLogger) NoPanic() {
	ml.noPanic = true
}

// Debug creates a new debug event with the given message
func (ml *MultiLogger) Debug(msg string) LogBuilder {
	lbs := make([]LogBuilder, len(ml.Loggers))
	for index, logger := range ml.Loggers {
		lbs[index] = logger.Debug(msg)
	}
	return &MultiLogBuilder{ml, lbs, LogLevelDebug}
}

// Debugf creates a new debug event with the formatted message
func (ml *MultiLogger) Debugf(format string, v ...any) LogBuilder {
	lbs := make([]LogBuilder, len(ml.Loggers))
	for index, logger := range ml.Loggers {
		lbs[index] = logger.Debugf(format, v...)
	}
	return &MultiLogBuilder{ml, lbs, LogLevelDebug}
}

// Info creates a new info event with the given message
func (ml *MultiLogger) Info(msg string) LogBuilder {
	lbs := make([]LogBuilder, len(ml.Loggers))
	for index, logger := range ml.Loggers {
		lbs[index] = logger.Info(msg)
	}
	return &MultiLogBuilder{ml, lbs, LogLevelInfo}
}

// Infof creates a new info event with the formatted message
func (ml *MultiLogger) Infof(format string, v ...any) LogBuilder {
	lbs := make([]LogBuilder, len(ml.Loggers))
	for index, logger := range ml.Loggers {
		lbs[index] = logger.Infof(format, v...)
	}
	return &MultiLogBuilder{ml, lbs, LogLevelInfo}
}

// Warn creates a new warn event with the given message
func (ml *MultiLogger) Warn(msg string) LogBuilder {
	lbs := make([]LogBuilder, len(ml.Loggers))
	for index, logger := range ml.Loggers {
		lbs[index] = logger.Warn(msg)
	}
	return &MultiLogBuilder{ml, lbs, LogLevelWarn}
}

// Warnf creates a new warn event with the formatted message
func (ml *MultiLogger) Warnf(format string, v ...any) LogBuilder {
	lbs := make([]LogBuilder, len(ml.Loggers))
	for index, logger := range ml.Loggers {
		lbs[index] = logger.Warnf(format, v...)
	}
	return &MultiLogBuilder{ml, lbs, LogLevelWarn}
}

// Error creates a new error event with the given message
func (ml *MultiLogger) Error(msg string) LogBuilder {
	lbs := make([]LogBuilder, len(ml.Loggers))
	for index, logger := range ml.Loggers {
		lbs[index] = logger.Error(msg)
	}
	return &MultiLogBuilder{ml, lbs, LogLevelError}
}

// Errorf creates a new error event with the formatted message
func (ml *MultiLogger) Errorf(format string, v ...any) LogBuilder {
	lbs := make([]LogBuilder, len(ml.Loggers))
	for index, logger := range ml.Loggers {
		lbs[index] = logger.Errorf(format, v...)
	}
	return &MultiLogBuilder{ml, lbs, LogLevelError}
}

// Error creates a new error event with the given message
func (ml *MultiLogger) Fatal(msg string) LogBuilder {
	lbs := make([]LogBuilder, len(ml.Loggers))
	for index, logger := range ml.Loggers {
		lbs[index] = logger.Fatal(msg)
	}
	return &MultiLogBuilder{ml, lbs, LogLevelFatal}
}

// Errorf creates a new error event with the formatted message
func (ml *MultiLogger) Fatalf(format string, v ...any) LogBuilder {
	lbs := make([]LogBuilder, len(ml.Loggers))
	for index, logger := range ml.Loggers {
		lbs[index] = logger.Fatalf(format, v...)
	}
	return &MultiLogBuilder{ml, lbs, LogLevelFatal}
}

// Error creates a new error event with the given message
func (ml *MultiLogger) Panic(msg string) LogBuilder {
	lbs := make([]LogBuilder, len(ml.Loggers))
	for index, logger := range ml.Loggers {
		lbs[index] = logger.Panic(msg)
	}
	return &MultiLogBuilder{ml, lbs, LogLevelPanic}
}

// Errorf creates a new error event with the formatted message
func (ml *MultiLogger) Panicf(format string, v ...any) LogBuilder {
	lbs := make([]LogBuilder, len(ml.Loggers))
	for index, logger := range ml.Loggers {
		lbs[index] = logger.Panicf(format, v...)
	}
	return &MultiLogBuilder{ml, lbs, LogLevelPanic}
}

// MultiLogBuilder implements the LogBuilder interface
// by writing to multiple underlying LogBuilders sequentially.
type MultiLogBuilder struct {
	l   *MultiLogger
	lbs []LogBuilder
	lvl LogLevel
}

// Int adds an int field to the output
func (mlb *MultiLogBuilder) Int(key string, val int) LogBuilder {
	for _, lb := range mlb.lbs {
		lb.Int(key, val)
	}
	return mlb
}

// Int64 adds an int64 field to the output
func (mlb *MultiLogBuilder) Int64(key string, val int64) LogBuilder {
	for _, lb := range mlb.lbs {
		lb.Int64(key, val)
	}
	return mlb
}

// Int32 adds an int32 field to the output
func (mlb *MultiLogBuilder) Int32(key string, val int32) LogBuilder {
	for _, lb := range mlb.lbs {
		lb.Int32(key, val)
	}
	return mlb
}

// Int16 adds an int16 field to the output
func (mlb *MultiLogBuilder) Int16(key string, val int16) LogBuilder {
	for _, lb := range mlb.lbs {
		lb.Int16(key, val)
	}
	return mlb
}

// Int8 adds an int8 field to the output
func (mlb *MultiLogBuilder) Int8(key string, val int8) LogBuilder {
	for _, lb := range mlb.lbs {
		lb.Int8(key, val)
	}
	return mlb
}

// Uint adds a uint field to the output
func (mlb *MultiLogBuilder) Uint(key string, val uint) LogBuilder {
	for _, lb := range mlb.lbs {
		lb.Uint(key, val)
	}
	return mlb
}

// Uint64 adds a uint64 field to the output
func (mlb *MultiLogBuilder) Uint64(key string, val uint64) LogBuilder {
	for _, lb := range mlb.lbs {
		lb.Uint64(key, val)
	}
	return mlb
}

// Uint32 adds a uint32 field to the output
func (mlb *MultiLogBuilder) Uint32(key string, val uint32) LogBuilder {
	for _, lb := range mlb.lbs {
		lb.Uint32(key, val)
	}
	return mlb
}

// Uint16 adds a uint16 field to the output
func (mlb *MultiLogBuilder) Uint16(key string, val uint16) LogBuilder {
	for _, lb := range mlb.lbs {
		lb.Uint16(key, val)
	}
	return mlb
}

// Uint8 adds a uint8 field to the output
func (mlb *MultiLogBuilder) Uint8(key string, val uint8) LogBuilder {
	for _, lb := range mlb.lbs {
		lb.Uint8(key, val)
	}
	return mlb
}

// Float64 adds a float64 field to the output
func (mlb *MultiLogBuilder) Float64(key string, val float64) LogBuilder {
	for _, lb := range mlb.lbs {
		lb.Float64(key, val)
	}
	return mlb
}

// Float32 adds a float32 field to the output
func (mlb *MultiLogBuilder) Float32(key string, val float32) LogBuilder {
	for _, lb := range mlb.lbs {
		lb.Float32(key, val)
	}
	return mlb
}

// Stringer calls the String method of an fmt.Stringer
// and adds the resulting string as a field to the output
func (mlb *MultiLogBuilder) Stringer(key string, s fmt.Stringer) LogBuilder {
	for _, lb := range mlb.lbs {
		lb.Stringer(key, s)
	}
	return mlb
}

// Bytes writes base64-encoded bytes as a field to the output
func (mlb *MultiLogBuilder) Bytes(key string, b []byte) LogBuilder {
	for _, lb := range mlb.lbs {
		lb.Bytes(key, b)
	}
	return mlb
}

// Timestamp adds the time formatted as RFC3339Nano
// as a field to the output using the key "timestamp"
func (mlb *MultiLogBuilder) Timestamp() LogBuilder {
	for _, lb := range mlb.lbs {
		lb.Timestamp()
	}
	return mlb
}

// Bool adds a bool as a field to the output
func (mlb *MultiLogBuilder) Bool(key string, val bool) LogBuilder {
	for _, lb := range mlb.lbs {
		lb.Bool(key, val)
	}
	return mlb
}

// Str adds a string as a field to the output
func (mlb *MultiLogBuilder) Str(key, val string) LogBuilder {
	for _, lb := range mlb.lbs {
		lb.Str(key, val)
	}
	return mlb
}

// Any uses reflection to marshal any type and writes
// the result as a field to the output. This is much slower
// than the type-specific functions.
func (mlb *MultiLogBuilder) Any(key string, val any) LogBuilder {
	for _, lb := range mlb.lbs {
		lb.Any(key, val)
	}
	return mlb
}

// Err adds an error as a field to the output
func (mlb *MultiLogBuilder) Err(err error) LogBuilder {
	for _, lb := range mlb.lbs {
		lb.Err(err)
	}
	return mlb
}

// Send sends the event to the output.
//
// After calling send, do not use the event again.
func (mlb *MultiLogBuilder) Send() {
	for _, lb := range mlb.lbs {
		lb.Send()
	}
	if mlb.lvl == LogLevelFatal && !mlb.l.noExit {
		os.Exit(1)
	} else if mlb.lvl == LogLevelPanic && !mlb.l.noPanic {
		panic("")
	}
}

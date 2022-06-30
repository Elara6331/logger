package logger

import (
	"fmt"
)

var _ Logger = (*NopLogger)(nil)

// NopLogger implements the Logger interface
// using human-readable output for log messages.
type NopLogger struct{}

// NewNop creates and returns a new NopLogger
func NewNop() NopLogger {
	return NopLogger{}
}

// NoPanic prevents the logger from panicking on panic events
func (nl NopLogger) NoPanic() {}

// NoExit prevents the logger from exiting on fatal events
func (nl NopLogger) NoExit() {}

// Debug creates a new debug event with the given message
func (nl NopLogger) Debug(msg string) LogBuilder {
	return NopLogBuilder{}
}

// Debugf creates a new debug event with the formatted message
func (nl NopLogger) Debugf(format string, v ...any) LogBuilder {
	return NopLogBuilder{}
}

// Info creates a new info event with the given message
func (nl NopLogger) Info(msg string) LogBuilder {
	return NopLogBuilder{}
}

// Infof creates a new info event with the formatted message
func (nl NopLogger) Infof(format string, v ...any) LogBuilder {
	return NopLogBuilder{}
}

// Warn creates a new warn event with the given message
func (nl NopLogger) Warn(msg string) LogBuilder {
	return NopLogBuilder{}
}

// Warnf creates a new warn event with the formatted message
func (nl NopLogger) Warnf(format string, v ...any) LogBuilder {
	return NopLogBuilder{}
}

// Error creates a new error event with the given message
func (nl NopLogger) Error(msg string) LogBuilder {
	return NopLogBuilder{}
}

// Errorf creates a new error event with the formatted message
func (nl NopLogger) Errorf(format string, v ...any) LogBuilder {
	return NopLogBuilder{}
}

// Fatal creates a new fatal event with the given message
func (nl NopLogger) Fatal(msg string) LogBuilder {
	return NopLogBuilder{}
}

// Fatalf creates a new fatal event with the formatted message
func (nl NopLogger) Fatalf(format string, v ...any) LogBuilder {
	return NopLogBuilder{}
}

// Panic creates a new panic event with the given message
func (nl NopLogger) Panic(msg string) LogBuilder {
	return NopLogBuilder{}
}

// Panicf creates a new panic event with the formatted message
func (nl NopLogger) Panicf(format string, v ...any) LogBuilder {
	return NopLogBuilder{}
}

// NopLogBuilder implements the LogBuilder interface
// using human-readable output for log messages
type NopLogBuilder struct{}

// Int adds an int field to the output
func (nlb NopLogBuilder) Int(key string, val int) LogBuilder { return nlb }

// Int64 adds an int64 field to the output
func (nlb NopLogBuilder) Int64(key string, val int64) LogBuilder { return nlb }

// Int32 adds an int32 field to the output
func (nlb NopLogBuilder) Int32(key string, val int32) LogBuilder { return nlb }

// Int16 adds an int16 field to the output
func (nlb NopLogBuilder) Int16(key string, val int16) LogBuilder { return nlb }

// Int8 adds an int8 field to the output
func (nlb NopLogBuilder) Int8(key string, val int8) LogBuilder { return nlb }

// Uint adds a uint field to the output
func (nlb NopLogBuilder) Uint(key string, val uint) LogBuilder { return nlb }

// Uint64 adds a uint64 field to the output
func (nlb NopLogBuilder) Uint64(key string, val uint64) LogBuilder { return nlb }

// Uint32 adds a uint32 field to the output
func (nlb NopLogBuilder) Uint32(key string, val uint32) LogBuilder { return nlb }

// Uint16 adds a uint16 field to the output
func (nlb NopLogBuilder) Uint16(key string, val uint16) LogBuilder { return nlb }

// Uint8 adds a uint8 field to the output
func (nlb NopLogBuilder) Uint8(key string, val uint8) LogBuilder { return nlb }

// Float64 adds a float64 field to the output
func (nlb NopLogBuilder) Float64(key string, val float64) LogBuilder { return nlb }

// Float32 adds a float32 field to the output
func (nlb NopLogBuilder) Float32(key string, val float32) LogBuilder { return nlb }

// Stringer calls the String method of an fmt.Stringer
// and adds the resulting string as a field to the output
func (nlb NopLogBuilder) Stringer(key string, s fmt.Stringer) LogBuilder { return nlb }

// Bytes writes hex-encoded bytes as a field to the output
func (nlb NopLogBuilder) Bytes(key string, b []byte) LogBuilder { return nlb }

// Timestamp adds the time formatted as RFC3339Nano
// as a field to the output
func (nlb NopLogBuilder) Timestamp() LogBuilder { return nlb }

// Bool adds a bool as a field to the output
func (nlb NopLogBuilder) Bool(key string, val bool) LogBuilder { return nlb }

// Str adds a string as a field to the output
func (nlb NopLogBuilder) Str(key, val string) LogBuilder { return nlb }

// Any uses reflection to marshal any type and writes
// the result as a field to the output. This is much slower
// than the type-specific functions.
func (nlb NopLogBuilder) Any(key string, val any) LogBuilder { return nlb }

// Err adds an error as a field to the output
func (nlb NopLogBuilder) Err(err error) LogBuilder { return nlb }

// Send sends the event to the output.
//
// After calling send, do not use the event again.
func (nlb NopLogBuilder) Send() {}

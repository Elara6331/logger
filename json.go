package logger

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

var _ Logger = (*JSONLogger)(nil)

// JSONLogger implements the Logger interface
// using JSON for log messages.
type JSONLogger struct {
	Out   io.Writer
	Level LogLevel

	noPanic bool
	noExit  bool
}

// NewJSON creates and returns a new JSONLogger
// If the input writer is io.Discard, NopLogger
// will be returned.
func NewJSON(out io.Writer) *JSONLogger {
	return &JSONLogger{Out: out, Level: LogLevelInfo}
}

// NoPanic prevents the logger from panicking on panic events
func (jl *JSONLogger) NoPanic() {
	jl.noPanic = true
}

// NoExit prevents the logger from exiting on fatal events
func (jl *JSONLogger) NoExit() {
	jl.noExit = true
}

// Debug creates a new debug event with the given message
func (jl *JSONLogger) Debug(msg string) LogBuilder {
	return newJSONLogBuilder(jl, msg, LogLevelDebug)
}

// Debugf creates a new debug event with the formatted message
func (jl *JSONLogger) Debugf(format string, v ...any) LogBuilder {
	return newJSONLogBuilder(jl, fmt.Sprintf(format, v...), LogLevelDebug)
}

// Info creates a new info event with the given message
func (jl *JSONLogger) Info(msg string) LogBuilder {
	return newJSONLogBuilder(jl, msg, LogLevelInfo)
}

// Infof creates a new info event with the formatted message
func (jl *JSONLogger) Infof(format string, v ...any) LogBuilder {
	return newJSONLogBuilder(jl, fmt.Sprintf(format, v...), LogLevelInfo)
}

// Warn creates a new warn event with the given message
func (jl *JSONLogger) Warn(msg string) LogBuilder {
	return newJSONLogBuilder(jl, msg, LogLevelWarn)
}

// Warnf creates a new warn event with the formatted message
func (jl *JSONLogger) Warnf(format string, v ...any) LogBuilder {
	return newJSONLogBuilder(jl, fmt.Sprintf(format, v...), LogLevelWarn)
}

// Error creates a new error event with the given message
func (jl *JSONLogger) Error(msg string) LogBuilder {
	return newJSONLogBuilder(jl, msg, LogLevelError)
}

// Errorf creates a new error event with the formatted message
func (jl *JSONLogger) Errorf(format string, v ...any) LogBuilder {
	return newJSONLogBuilder(jl, fmt.Sprintf(format, v...), LogLevelError)
}

// Fatal creates a new fatal event with the given message
//
// When sent, fatal events will cause a call to os.Exit(1)
func (jl *JSONLogger) Fatal(msg string) LogBuilder {
	return newJSONLogBuilder(jl, msg, LogLevelFatal)
}

// Fatalf creates a new fatal event with the formatted message
//
// When sent, fatal events will cause a call to os.Exit(1)
func (jl *JSONLogger) Fatalf(format string, v ...any) LogBuilder {
	return newJSONLogBuilder(jl, fmt.Sprintf(format, v...), LogLevelFatal)
}

// Panic creates a new panic event with the given message
//
// When sent, panic events will cause a panic
func (jl *JSONLogger) Panic(msg string) LogBuilder {
	return newJSONLogBuilder(jl, msg, LogLevelPanic)
}

// Panicf creates a new panic event with the formatted message
//
// When sent, panic events will cause a panic
func (jl *JSONLogger) Panicf(format string, v ...any) LogBuilder {
	return newJSONLogBuilder(jl, fmt.Sprintf(format, v...), LogLevelPanic)
}

// JSONLogBuilder implements the LogBuilder interface
// using JSON for log messages
type JSONLogBuilder struct {
	l   *JSONLogger
	lvl LogLevel
	out writer
}

func newJSONLogBuilder(jl *JSONLogger, msg string, lvl LogLevel) LogBuilder {
	if jl.Out == io.Discard || lvl < jl.Level {
		return NopLogBuilder{}
	}
	lb := &JSONLogBuilder{
		out: writer{&bytes.Buffer{}, jl.Out},
		lvl: lvl,
		l:   jl,
	}
	lb.out.WriteString(`{"msg":"`)
	lb.out.WriteString(msg)
	lb.out.WriteString(`","level":"`)
	lb.out.WriteString(logLevelNames[lvl])
	lb.out.WriteByte('"')
	return lb
}

// writeKey writes a JSON key to the buffer
func (jlb *JSONLogBuilder) writeKey(k string) {
	jlb.out.WriteString(`,"`)
	jlb.out.WriteString(k)
	jlb.out.WriteString(`":`)
}

// Int adds an int field to the output
func (jlb *JSONLogBuilder) Int(key string, val int) LogBuilder {
	return jlb.Int64(key, int64(val))
}

// Int64 adds an int64 field to the output
func (jlb *JSONLogBuilder) Int64(key string, val int64) LogBuilder {
	jlb.writeKey(key)
	jlb.out.WriteString(strconv.FormatInt(val, 10))
	return jlb
}

// Int32 adds an int32 field to the output
func (jlb *JSONLogBuilder) Int32(key string, val int32) LogBuilder {
	return jlb.Int64(key, int64(val))
}

// Int16 adds an int16 field to the output
func (jlb *JSONLogBuilder) Int16(key string, val int16) LogBuilder {
	return jlb.Int64(key, int64(val))
}

// Int8 adds an int8 field to the output
func (jlb *JSONLogBuilder) Int8(key string, val int8) LogBuilder {
	return jlb.Int64(key, int64(val))
}

// Uint adds a uint field to the output
func (jlb *JSONLogBuilder) Uint(key string, val uint) LogBuilder {
	return jlb.Uint64(key, uint64(val))
}

// Uint64 adds a uint64 field to the output
func (jlb *JSONLogBuilder) Uint64(key string, val uint64) LogBuilder {
	jlb.writeKey(key)
	jlb.out.WriteString(strconv.FormatUint(val, 10))
	return jlb
}

// Uint32 adds a uint32 field to the output
func (jlb *JSONLogBuilder) Uint32(key string, val uint32) LogBuilder {
	return jlb.Uint64(key, uint64(val))
}

// Uint16 adds a uint16 field to the output
func (jlb *JSONLogBuilder) Uint16(key string, val uint16) LogBuilder {
	return jlb.Uint64(key, uint64(val))
}

// Uint8 adds a uint8 field to the output
func (jlb *JSONLogBuilder) Uint8(key string, val uint8) LogBuilder {
	return jlb.Uint64(key, uint64(val))
}

// float adds a float of specified bitsize to the output
func (jlb *JSONLogBuilder) float(key string, val float64, bitsize int) LogBuilder {
	jlb.writeKey(key)
	jlb.out.WriteString(strconv.FormatFloat(val, 'f', -1, bitsize))
	return jlb
}

// Float64 adds a float64 field to the output
func (jlb *JSONLogBuilder) Float64(key string, val float64) LogBuilder {
	return jlb.float(key, val, 64)
}

// Float32 adds a float32 field to the output
func (jlb *JSONLogBuilder) Float32(key string, val float32) LogBuilder {
	return jlb.float(key, float64(val), 32)
}

// Stringer calls the String method of an fmt.Stringer
// and adds the resulting string as a field to the output
func (jlb *JSONLogBuilder) Stringer(key string, s fmt.Stringer) LogBuilder {
	return jlb.Str(key, s.String())
}

// Bytes writes base64-encoded bytes as a field to the output
func (jlb *JSONLogBuilder) Bytes(key string, b []byte) LogBuilder {
	return jlb.Str(key, base64.StdEncoding.EncodeToString(b))
}

// Timestamp adds the time formatted as RFC3339Nano
// as a field to the output using the key "timestamp"
func (jlb *JSONLogBuilder) Timestamp() LogBuilder {
	return jlb.Str("timestamp", time.Now().Format(time.RFC3339Nano))
}

// Bool adds a bool as a field to the output
func (jlb *JSONLogBuilder) Bool(key string, val bool) LogBuilder {
	jlb.writeKey(key)
	if val {
		jlb.out.WriteString("true")
	} else {
		jlb.out.WriteString("false")
	}
	return jlb
}

// Str adds a string as a field to the output
func (jlb *JSONLogBuilder) Str(key, val string) LogBuilder {
	jlb.writeKey(key)
	jlb.out.WriteByte('"')
	jlb.out.WriteString(val)
	jlb.out.WriteByte('"')
	return jlb
}

// Any uses reflection to marshal any type and writes
// the result as a field to the output. This is much slower
// than the type-specific functions.
func (jlb *JSONLogBuilder) Any(key string, val any) LogBuilder {
	jlb.writeKey(key)
	data, err := json.Marshal(val)
	if err != nil {
		panic(err)
	}
	jlb.out.Write(data)
	return jlb
}

// Err adds an error as a field to the output
func (jlb *JSONLogBuilder) Err(err error) LogBuilder {
	return jlb.Str("error", err.Error())
}

// Send sends the event to the output.
//
// After calling send, do not use the event again.
func (jlb *JSONLogBuilder) Send() {
	jlb.out.WriteByte('}')
	jlb.out.Flush()
	if jlb.lvl == LogLevelFatal && !jlb.l.noExit {
		os.Exit(1)
	} else if jlb.lvl == LogLevelPanic && !jlb.l.noPanic {
		panic("")
	}
}

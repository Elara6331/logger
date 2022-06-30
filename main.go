package logger

import (
	"fmt"
)

// LogLevel represents a log level
type LogLevel uint8

// Log levels
const (
	LogLevelDebug LogLevel = iota
	LogLevelInfo
	LogLevelWarn
	LogLevelError
	LogLevelFatal
	LogLevelPanic
)

var logLevelNames = [...]string{
	LogLevelDebug: "debug",
	LogLevelInfo:  "info",
	LogLevelWarn:  "warn",
	LogLevelError: "error",
	LogLevelFatal: "fatal",
	LogLevelPanic: "panic",
}

// Logger represents a logger
type Logger interface {
	// NoPanic prevents the logger from panicking on panic events
	NoPanic()

	// NoExit prevents the logger from exiting on fatal events
	NoExit()

	// Debug creates a new debug event with the given message
	Debug(string) LogBuilder

	// Debugf creates a new debug event with the formatted message
	Debugf(string, ...any) LogBuilder

	// Info creates a new info event with the given message
	Info(string) LogBuilder

	// Infof creates a new info event with the formatted message
	Infof(string, ...any) LogBuilder

	// Warn creates a new warn event with the given message
	Warn(string) LogBuilder

	// Warnf creates a new warn event with the formatted message
	Warnf(string, ...any) LogBuilder

	// Error creates a new error event with the given message
	Error(string) LogBuilder

	// Errorf creates a new error event with the formatted message
	Errorf(string, ...any) LogBuilder

	// Fatal creates a new fatal event with the given message
	//
	// When sent, fatal events will cause a call to os.Exit(1)
	Fatal(string) LogBuilder

	// Fatalf creates a new fatal event with the formatted message
	//
	// When sent, fatal events will cause a call to os.Exit(1)
	Fatalf(string, ...any) LogBuilder

	// Panic creates a new panic event with the given message
	//
	// When sent, panic events will cause a panic
	Panic(string) LogBuilder

	// Panicf creates a new panic event with the formatted message
	//
	// When sent, panic events will cause a panic
	Panicf(string, ...any) LogBuilder
}

// LogBuilder represents a log event builder
type LogBuilder interface {
	// Int adds an int field to the output
	Int(string, int) LogBuilder
	// Int8 adds an int8 field to the output
	Int8(string, int8) LogBuilder
	// Int16 adds an int16 field to the output
	Int16(string, int16) LogBuilder
	// Int32 adds an int32 field to the output
	Int32(string, int32) LogBuilder
	// Int64 adds an int64 field to the output
	Int64(string, int64) LogBuilder
	// Uint adds a uint field to the output
	Uint(string, uint) LogBuilder
	// Uint8 adds a uint8 field to the output
	Uint8(string, uint8) LogBuilder
	// Uint16 adds a uint16 field to the output
	Uint16(string, uint16) LogBuilder
	// Uint32 adds a uint32 field to the output
	Uint32(string, uint32) LogBuilder
	// Uint64 adds a uint64 field to the output
	Uint64(string, uint64) LogBuilder
	// Float32 adds a float32 field to the output
	Float32(string, float32) LogBuilder
	// Float64 adds a float64 field to the output
	Float64(string, float64) LogBuilder
	// Stringer calls the String method of an fmt.Stringer
	// and adds the resulting string as a field to the output
	Stringer(string, fmt.Stringer) LogBuilder
	// Bytes adds []byte as a field to the output
	Bytes(string, []byte) LogBuilder
	// Timestamp adds the time formatted as RFC3339Nano
	// as a field to the output using the key "timestamp"
	Timestamp() LogBuilder
	// Bool adds a bool as a field to the output
	Bool(string, bool) LogBuilder
	// Str adds a string as a field to the output
	Str(string, string) LogBuilder
	// Any uses reflection to marshal any type and writes
	// the result as a field to the output. This is much slower
	// than the type-specific functions.
	Any(string, any) LogBuilder
	// Err adds an error as a field to the output
	Err(error) LogBuilder
	// Send sends the event to the output.
	//
	// After calling send, do not use the event again.
	Send()
}

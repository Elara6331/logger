package logger

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/gookit/color"
	"github.com/mattn/go-isatty"
)

var _ Logger = (*PrettyLogger)(nil)

// PrettyLogger implements the Logger interface
// using human-readable output for log messages.
type PrettyLogger struct {
	Out   io.Writer
	Level LogLevel

	TimeFormat string

	UseColor bool

	TimeColor color.Color
	MsgColor  color.Color
	KeyColor  color.Color

	DebugColor color.Color
	InfoColor  color.Color
	WarnColor  color.Color
	ErrColor   color.Color
	FatalColor color.Color
	PanicColor color.Color

	noPanic bool
	noExit  bool
}

// NewPretty creates and returns a new PrettyLogger.
// If the input writer is io.Discard, NopLogger
// will be returned.
func NewPretty(out io.Writer) *PrettyLogger {
	useColor := false
	if f, ok := out.(*os.File); ok {
		useColor = isatty.IsTerminal(f.Fd())
	}

	return &PrettyLogger{
		Out:   out,
		Level: LogLevelInfo,

		TimeFormat: time.Kitchen,

		UseColor: useColor,

		TimeColor: color.FgGray,
		MsgColor:  color.Normal,
		KeyColor:  color.FgCyan,

		DebugColor: color.FgYellow,
		InfoColor:  color.FgGreen,
		WarnColor:  color.FgRed,
		ErrColor:   color.FgLightRed,
		FatalColor: color.FgRed.Darken(),
		PanicColor: color.FgRed.Darken(),
	}
}

// NoPanic prevents the logger from panicking on panic events
func (pl *PrettyLogger) NoPanic() {
	pl.noPanic = true
}

// NoExit prevents the logger from exiting on fatal events
func (pl *PrettyLogger) NoExit() {
	pl.noExit = true
}

// Debug creates a new debug event with the given message
func (pl *PrettyLogger) Debug(msg string) LogBuilder {
	return newPrettyLogBuilder(pl, msg, LogLevelDebug)
}

// Debugf creates a new debug event with the formatted message
func (pl *PrettyLogger) Debugf(format string, v ...any) LogBuilder {
	return newPrettyLogBuilder(pl, fmt.Sprintf(format, v...), LogLevelDebug)
}

// Info creates a new info event with the given message
func (pl *PrettyLogger) Info(msg string) LogBuilder {
	return newPrettyLogBuilder(pl, msg, LogLevelInfo)
}

// Infof creates a new info event with the formatted message
func (pl *PrettyLogger) Infof(format string, v ...any) LogBuilder {
	return newPrettyLogBuilder(pl, fmt.Sprintf(format, v...), LogLevelInfo)
}

// Warn creates a new warn event with the given message
func (pl *PrettyLogger) Warn(msg string) LogBuilder {
	return newPrettyLogBuilder(pl, msg, LogLevelWarn)
}

// Warnf creates a new warn event with the formatted message
func (pl *PrettyLogger) Warnf(format string, v ...any) LogBuilder {
	return newPrettyLogBuilder(pl, fmt.Sprintf(format, v...), LogLevelWarn)
}

// Error creates a new error event with the given message
func (pl *PrettyLogger) Error(msg string) LogBuilder {
	return newPrettyLogBuilder(pl, msg, LogLevelError)
}

// Errorf creates a new error event with the formatted message
func (pl *PrettyLogger) Errorf(format string, v ...any) LogBuilder {
	return newPrettyLogBuilder(pl, fmt.Sprintf(format, v...), LogLevelError)
}

// Fatal creates a new fatal event with the given message
//
// When sent, fatal events will cause a call to os.Exit(1)
func (pl *PrettyLogger) Fatal(msg string) LogBuilder {
	return newPrettyLogBuilder(pl, msg, LogLevelFatal)
}

// Fatalf creates a new fatal event with the formatted message
//
// When sent, fatal events will cause a call to os.Exit(1)
func (pl *PrettyLogger) Fatalf(format string, v ...any) LogBuilder {
	return newPrettyLogBuilder(pl, fmt.Sprintf(format, v...), LogLevelFatal)
}

// Panic creates a new panic event with the given message
//
// When sent, panic events will cause a panic
func (pl *PrettyLogger) Panic(msg string) LogBuilder {
	return newPrettyLogBuilder(pl, msg, LogLevelPanic)
}

// Panicf creates a new panic event with the formatted message
//
// When sent, panic events will cause a panic
func (pl *PrettyLogger) Panicf(format string, v ...any) LogBuilder {
	return newPrettyLogBuilder(pl, fmt.Sprintf(format, v...), LogLevelPanic)
}

// PrettyLogBuilder implements the LogBuilder interface
// using human-readable output for log messages
type PrettyLogBuilder struct {
	l   *PrettyLogger
	lvl LogLevel
	out writer
}

func newPrettyLogBuilder(pl *PrettyLogger, msg string, lvl LogLevel) LogBuilder {
	if pl.Out == io.Discard || lvl < pl.Level {
		return NopLogBuilder{}
	}
	lb := &PrettyLogBuilder{
		l:   pl,
		out: writer{&bytes.Buffer{}, pl.Out},
		lvl: lvl,
	}
	lb.writeColor(lb.l.TimeColor, time.Now().Format(lb.l.TimeFormat))
	lb.out.WriteByte(' ')

	switch lvl {
	case LogLevelDebug:
		lb.writeColor(lb.l.DebugColor, "DBG")
	case LogLevelInfo:
		lb.writeColor(lb.l.InfoColor, "INF")
	case LogLevelWarn:
		lb.writeColor(lb.l.WarnColor, "WRN")
	case LogLevelError:
		lb.writeColor(lb.l.ErrColor, "ERR")
	case LogLevelFatal:
		lb.writeColor(lb.l.FatalColor, "FTL")
	case LogLevelPanic:
		lb.writeColor(lb.l.PanicColor, "PNC")
	}
	lb.out.WriteByte(' ')

	lb.writeColor(lb.l.MsgColor, msg)
	return lb
}

// writeKey writes a JSON key to the buffer
func (plb *PrettyLogBuilder) writeKey(k string) {
	plb.out.WriteByte(' ')
	plb.writeColor(plb.l.KeyColor, k)
	plb.out.WriteByte('=')
}

// Int adds an int field to the output
func (plb *PrettyLogBuilder) Int(key string, val int) LogBuilder {
	return plb.Int64(key, int64(val))
}

// Int64 adds an int64 field to the output
func (plb *PrettyLogBuilder) Int64(key string, val int64) LogBuilder {
	plb.writeKey(key)
	plb.out.WriteString(strconv.FormatInt(val, 10))
	return plb
}

// Int32 adds an int32 field to the output
func (plb *PrettyLogBuilder) Int32(key string, val int32) LogBuilder {
	return plb.Int64(key, int64(val))
}

// Int16 adds an int16 field to the output
func (plb *PrettyLogBuilder) Int16(key string, val int16) LogBuilder {
	return plb.Int64(key, int64(val))
}

// Int8 adds an int8 field to the output
func (plb *PrettyLogBuilder) Int8(key string, val int8) LogBuilder {
	return plb.Int64(key, int64(val))
}

// Uint adds a uint field to the output
func (plb *PrettyLogBuilder) Uint(key string, val uint) LogBuilder {
	return plb.Uint64(key, uint64(val))
}

// Uint64 adds a uint64 field to the output
func (plb *PrettyLogBuilder) Uint64(key string, val uint64) LogBuilder {
	plb.writeKey(key)
	plb.out.WriteString(strconv.FormatUint(val, 10))
	return plb
}

// Uint32 adds a uint32 field to the output
func (plb *PrettyLogBuilder) Uint32(key string, val uint32) LogBuilder {
	return plb.Uint64(key, uint64(val))
}

// Uint16 adds a uint16 field to the output
func (plb *PrettyLogBuilder) Uint16(key string, val uint16) LogBuilder {
	return plb.Uint64(key, uint64(val))
}

// Uint8 adds a uint8 field to the output
func (plb *PrettyLogBuilder) Uint8(key string, val uint8) LogBuilder {
	return plb.Uint64(key, uint64(val))
}

// float adds a float of specified bitsize to the output
func (plb *PrettyLogBuilder) float(key string, val float64, bitsize int) LogBuilder {
	plb.writeKey(key)
	plb.out.WriteString(strconv.FormatFloat(val, 'f', -1, bitsize))
	return plb
}

// Float64 adds a float64 field to the output
func (plb *PrettyLogBuilder) Float64(key string, val float64) LogBuilder {
	return plb.float(key, val, 64)
}

// Float32 adds a float32 field to the output
func (plb *PrettyLogBuilder) Float32(key string, val float32) LogBuilder {
	return plb.float(key, float64(val), 32)
}

// Stringer calls the String method of an fmt.Stringer
// and adds the resulting string as a field to the output
func (plb *PrettyLogBuilder) Stringer(key string, s fmt.Stringer) LogBuilder {
	return plb.Str(key, s.String())
}

// Bytes writes hex-encoded bytes as a field to the output
func (plb *PrettyLogBuilder) Bytes(key string, b []byte) LogBuilder {
	return plb.Str(key, hex.EncodeToString(b))
}

// Timestamp adds the time formatted as RFC3339Nano
// as a field to the output
func (plb *PrettyLogBuilder) Timestamp() LogBuilder {
	return plb.Str("timestamp", time.Now().Format(time.RFC3339Nano))
}

// Bool adds a bool as a field to the output
func (plb *PrettyLogBuilder) Bool(key string, val bool) LogBuilder {
	plb.writeKey(key)
	if val {
		plb.out.WriteString("true")
	} else {
		plb.out.WriteString("false")
	}
	return plb
}

// Str adds a string as a field to the output
func (plb *PrettyLogBuilder) Str(key, val string) LogBuilder {
	plb.writeKey(key)
	plb.out.WriteByte('"')
	plb.out.WriteString(val)
	plb.out.WriteByte('"')
	return plb
}

// Any uses reflection to marshal any type and writes
// the result as a field to the output. This is much slower
// than the type-specific functions.
func (plb *PrettyLogBuilder) Any(key string, val any) LogBuilder {
	plb.writeKey(key)
	data, err := json.Marshal(val)
	if err != nil {
		panic(err)
	}
	plb.out.Write(data)
	return plb
}

// Err adds an error as a field to the output
func (plb *PrettyLogBuilder) Err(err error) LogBuilder {
	plb.out.WriteByte(' ')
	plb.writeColor(plb.l.ErrColor, "error=")
	plb.writeColor(plb.l.ErrColor, `"`+err.Error()+`"`)
	return plb
}

// Send sends the event to the output.
//
// After calling send, do not use the event again.
func (plb *PrettyLogBuilder) Send() {
	plb.out.WriteByte('\n')
	plb.out.Flush()
	if plb.lvl == LogLevelFatal && !plb.l.noExit {
		os.Exit(1)
	} else if plb.lvl == LogLevelPanic && !plb.l.noPanic {
		panic("")
	}
}

// writeColor writes a string to the buffer using the given color
func (plb *PrettyLogBuilder) writeColor(c color.Color, s string) {
	if plb.l.UseColor {
		plb.out.WriteString(c.Text(s))
	} else {
		plb.out.WriteString(s)
	}
}

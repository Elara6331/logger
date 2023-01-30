package logrus

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"go.arsenm.dev/logger"
)

var (
	_ logger.Logger     = &LogrusLogger{}
	_ logger.LogBuilder = &LogrusLogBuilder{}
)

type LogrusLogger struct {
	logger *logrus.Logger
}

func New(l *logrus.Logger) *LogrusLogger {
	return &LogrusLogger{l}
}

// NoPanic is a no-op because logrus does not provide this functionality
func (ll *LogrusLogger) NoPanic() {}

// NoExit prevents the logger from exiting on fatal events
func (ll *LogrusLogger) NoExit() {
	ll.logger.ExitFunc = func(int) {}
}

// SetLevel sets the log level of the logger
func (ll *LogrusLogger) SetLevel(logger.LogLevel) {}

// Debug creates a new debug event with the given message
func (ll *LogrusLogger) Debug(msg string) logger.LogBuilder {
	return &LogrusLogBuilder{
		log: logrus.NewEntry(ll.logger),
		lvl: logrus.DebugLevel,
		msg: msg,
	}
}

// Debugf creates a new debug event with the formatted message
func (ll *LogrusLogger) Debugf(format string, v ...any) logger.LogBuilder {
	return &LogrusLogBuilder{
		log: logrus.NewEntry(ll.logger),
		lvl: logrus.DebugLevel,
		msg: fmt.Sprintf(format, v...),
	}
}

// Info creates a new info event with the given message
func (ll *LogrusLogger) Info(msg string) logger.LogBuilder {
	return &LogrusLogBuilder{
		log: logrus.NewEntry(ll.logger),
		lvl: logrus.InfoLevel,
		msg: msg,
	}
}

// Infof creates a new info event with the formatted message
func (ll *LogrusLogger) Infof(format string, v ...any) logger.LogBuilder {
	return &LogrusLogBuilder{
		log: logrus.NewEntry(ll.logger),
		lvl: logrus.InfoLevel,
		msg: fmt.Sprintf(format, v...),
	}
}

// Warn creates a new warn event with the given message
func (ll *LogrusLogger) Warn(msg string) logger.LogBuilder {
	return &LogrusLogBuilder{
		log: logrus.NewEntry(ll.logger),
		lvl: logrus.WarnLevel,
		msg: msg,
	}
}

// Warnf creates a new warn event with the formatted message
func (ll *LogrusLogger) Warnf(format string, v ...any) logger.LogBuilder {
	return &LogrusLogBuilder{
		log: logrus.NewEntry(ll.logger),
		lvl: logrus.WarnLevel,
		msg: fmt.Sprintf(format, v...),
	}
}

// Error creates a new error event with the given message
func (ll *LogrusLogger) Error(msg string) logger.LogBuilder {
	return &LogrusLogBuilder{
		log: logrus.NewEntry(ll.logger),
		lvl: logrus.ErrorLevel,
		msg: msg,
	}
}

// Errorf creates a new error event with the formatted message
func (ll *LogrusLogger) Errorf(format string, v ...any) logger.LogBuilder {
	return &LogrusLogBuilder{
		log: logrus.NewEntry(ll.logger),
		lvl: logrus.ErrorLevel,
		msg: fmt.Sprintf(format, v...),
	}
}

// Fatal creates a new fatal event with the given message
func (ll *LogrusLogger) Fatal(msg string) logger.LogBuilder {
	return &LogrusLogBuilder{
		log: logrus.NewEntry(ll.logger),
		lvl: logrus.FatalLevel,
		msg: msg,
	}
}

// Fatalf creates a new fatal event with the formatted message
func (ll *LogrusLogger) Fatalf(format string, v ...any) logger.LogBuilder {
	return &LogrusLogBuilder{
		log: logrus.NewEntry(ll.logger),
		lvl: logrus.FatalLevel,
		msg: fmt.Sprintf(format, v...),
	}
}

// Panic creates a new panic event with the given message
func (ll *LogrusLogger) Panic(msg string) logger.LogBuilder {
	return &LogrusLogBuilder{
		log: logrus.NewEntry(ll.logger),
		lvl: logrus.PanicLevel,
		msg: msg,
	}
}

// Panicf creates a new panic event with the formatted message
func (ll *LogrusLogger) Panicf(format string, v ...any) logger.LogBuilder {
	return &LogrusLogBuilder{
		log: logrus.NewEntry(ll.logger),
		lvl: logrus.PanicLevel,
		msg: fmt.Sprintf(format, v...),
	}
}

type LogrusLogBuilder struct {
	log *logrus.Entry
	lvl logrus.Level
	msg string
}

func (lb *LogrusLogBuilder) Int(name string, value int) logger.LogBuilder {
	lb.log = lb.log.WithField(name, value)
	return lb
}

func (lb *LogrusLogBuilder) Int8(name string, value int8) logger.LogBuilder {
	lb.log = lb.log.WithField(name, value)
	return lb
}

func (lb *LogrusLogBuilder) Int16(name string, value int16) logger.LogBuilder {
	lb.log = lb.log.WithField(name, value)
	return lb
}

func (lb *LogrusLogBuilder) Int32(name string, value int32) logger.LogBuilder {
	lb.log = lb.log.WithField(name, value)
	return lb
}

func (lb *LogrusLogBuilder) Int64(name string, value int64) logger.LogBuilder {
	lb.log = lb.log.WithField(name, value)
	return lb
}

func (lb *LogrusLogBuilder) Uint(name string, value uint) logger.LogBuilder {
	lb.log = lb.log.WithField(name, value)
	return lb
}

func (lb *LogrusLogBuilder) Uint8(name string, value uint8) logger.LogBuilder {
	lb.log = lb.log.WithField(name, value)
	return lb
}

func (lb *LogrusLogBuilder) Uint16(name string, value uint16) logger.LogBuilder {
	lb.log = lb.log.WithField(name, value)
	return lb
}

func (lb *LogrusLogBuilder) Uint32(name string, value uint32) logger.LogBuilder {
	lb.log = lb.log.WithField(name, value)
	return lb
}

func (lb *LogrusLogBuilder) Uint64(name string, value uint64) logger.LogBuilder {
	lb.log = lb.log.WithField(name, value)
	return lb
}

func (lb *LogrusLogBuilder) Float32(name string, value float32) logger.LogBuilder {
	lb.log = lb.log.WithField(name, value)
	return lb
}

func (lb *LogrusLogBuilder) Float64(name string, value float64) logger.LogBuilder {
	lb.log = lb.log.WithField(name, value)
	return lb
}

func (lb *LogrusLogBuilder) Stringer(name string, value fmt.Stringer) logger.LogBuilder {
	lb.log = lb.log.WithField(name, value.String())
	return lb
}

func (lb *LogrusLogBuilder) Bytes(name string, value []byte) logger.LogBuilder {
	lb.log = lb.log.WithField(name, value)
	return lb
}

func (lb *LogrusLogBuilder) Timestamp() logger.LogBuilder {
	lb.log = lb.log.WithTime(time.Now())
	return lb
}

func (lb *LogrusLogBuilder) Bool(name string, value bool) logger.LogBuilder {
	lb.log = lb.log.WithField(name, value)
	return lb
}

func (lb *LogrusLogBuilder) Str(name string, value string) logger.LogBuilder {
	lb.log = lb.log.WithField(name, value)
	return lb
}

func (lb *LogrusLogBuilder) Any(name string, value any) logger.LogBuilder {
	lb.log = lb.log.WithField(name, value)
	return lb
}

func (lb *LogrusLogBuilder) Err(value error) logger.LogBuilder {
	lb.log = lb.log.WithError(value)
	return lb
}

func (lb *LogrusLogBuilder) Send() {
	lb.log.Logln(lb.lvl, lb.msg)
}

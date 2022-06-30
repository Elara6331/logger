package logger_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"logger"
	"os"
	"runtime"
	"testing"
	"time"
)

var buf = &bytes.Buffer{}
var jsonlog = logger.NewJSON(buf)
var prettylog = logger.NewPretty(buf)

func TestJSON(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		jsonlog.Info("").Send()

		if got, want := getStr(), `{"msg":"","level":"info"}`; got != want {
			t.Errorf("got: %s, want: %s", got, want)
		}
	})

	t.Run("one-field", func(t *testing.T) {
		jsonlog.Info("Test").Int("n", 1234).Send()

		if got, want := getStr(), `{"msg":"Test","level":"info","n":1234}`; got != want {
			t.Errorf("got: %s, want: %s", got, want)
		}
	})

	t.Run("two-field", func(t *testing.T) {
		jsonlog.
			Info("Test").
			Int("n", 1234).
			Float32("pi", 3.14).
			Send()

		if got, want := getStr(), `{"msg":"Test","level":"info","n":1234,"pi":3.14}`; got != want {
			t.Errorf("got: %s, want: %s", got, want)
		}
	})

	t.Run("any", func(t *testing.T) {
		jsonlog.Info("Test").Any("any", nil).Send()

		if got, want := getStr(), `{"msg":"Test","level":"info","any":null}`; got != want {
			t.Errorf("got: %s, want: %s", got, want)
		}
	})

	t.Run("all", func(t *testing.T) {
		jsonlog.
			Info("All").
			Int("int", -1).
			Int8("int8", -1).
			Int16("int16", -1).
			Int32("int32", -1).
			Int64("int64", -1).
			Uint("uint", 1).
			Uint8("uint8", 1).
			Uint16("uint16", 1).
			Uint32("uint32", 1).
			Uint64("uint64", 1).
			Float32("float32", 3.14).
			Float64("float64", 6.28).
			Bool("bool", true).
			Str("string", "").
			Bytes("[]byte", []byte{0x12, 0x34, 0x56}).
			Stringer("stringer", time.Second).
			Any("any", nil).
			Err(errors.New("err")).
			Send()

		if got, want := getStr(), `{"msg":"All","level":"info","int":-1,"int8":-1,"int16":-1,"int32":-1,"int64":-1,"uint":1,"uint8":1,"uint16":1,"uint32":1,"uint64":1,"float32":3.14,"float64":6.28,"bool":true,"string":"","[]byte":"EjRW","stringer":"1s","any":null,"error":"err"}`; got != want {
			t.Errorf("got: %s, want: %s", got, want)
		}
	})

	t.Run("decode", func(t *testing.T) {
		var log struct {
			Message string `json:"msg"`
			Level   string `json:"level"`
			N       int    `json:"n"`
		}
		jsonlog.Info("Test").Int("n", 1234).Send()

		err := json.Unmarshal(buf.Bytes(), &log)
		if err != nil {
			t.Error(err)
		}
		buf.Reset()

		if log.Message != "Test" {
			t.Errorf("expected message Test, got %s", log.Message)
		}
		if log.Level != "info" {
			t.Errorf("expected level info, got %s", log.Level)
		}
		if log.N != 1234 {
			t.Errorf("expected n 1234, got %d", log.N)
		}
	})
}

func TestPretty(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		prettylog.Info("").Send()

		ctime := time.Now().Format(time.Kitchen)
		if got, want := getStr(), ctime+" INF \n"; got != want {
			t.Errorf("got: %s, want: %s", got, want)
		}
	})

	t.Run("one-field", func(t *testing.T) {
		prettylog.Info("Test").Int("n", 1234).Send()

		ctime := time.Now().Format(time.Kitchen)
		if got, want := getStr(), ctime+" INF Test n=1234\n"; got != want {
			t.Errorf("got: %s, want: %s", got, want)
		}
	})

	t.Run("two-field", func(t *testing.T) {
		prettylog.
			Info("Test").
			Int("n", 1234).
			Float32("pi", 3.14).
			Send()

		ctime := time.Now().Format(time.Kitchen)
		if got, want := getStr(), ctime+" INF Test n=1234 pi=3.14\n"; got != want {
			t.Errorf("got: %s, want: %s", got, want)
		}
	})

	t.Run("any", func(t *testing.T) {
		prettylog.Info("Test").Any("any", nil).Send()

		ctime := time.Now().Format(time.Kitchen)
		if got, want := getStr(), ctime+" INF Test any=null\n"; got != want {
			t.Errorf("got: %s, want: %s", got, want)
		}
	})

	t.Run("all", func(t *testing.T) {
		prettylog.
			Info("All").
			Int("int", -1).
			Int8("int8", -1).
			Int16("int16", -1).
			Int32("int32", -1).
			Int64("int64", -1).
			Uint("uint", 1).
			Uint8("uint8", 1).
			Uint16("uint16", 1).
			Uint32("uint32", 1).
			Uint64("uint64", 1).
			Float32("float32", 3.14).
			Float64("float64", 6.28).
			Bool("bool", true).
			Str("string", "").
			Bytes("[]byte", []byte{0x12, 0x34, 0x56}).
			Stringer("stringer", time.Second).
			Any("any", nil).
			Err(errors.New("err")).
			Send()

		ctime := time.Now().Format(time.Kitchen)
		if got, want := getStr(), ctime+` INF All int=-1 int8=-1 int16=-1 int32=-1 int64=-1 uint=1 uint8=1 uint16=1 uint32=1 uint64=1 float32=3.14 float64=6.28 bool=true string="" []byte="123456" stringer="1s" any=null error="err"`+"\n"; got != want {
			t.Errorf("got: %#v, want: %#v", got, want)
		}
	})
}

func BenchmarkJSON(b *testing.B) {
	devnull, err := os.Open(getNullPath())
	if err != nil {
		b.Fatal(err)
	}

	jsonlog.Out = devnull
	jsonlog.Level = logger.LogLevelDebug

	b.Run("one-field", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			jsonlog.Debug("Benchmark").Int("int", 1)
		}
	})

	b.Run("two-field", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			jsonlog.Debug("Benchmark").Int("int", 1).Float32("pi", 3.14)
		}
	})

	b.Run("all", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			jsonlog.
				Debug("All").
				Int("int", -1).
				Int8("int8", -1).
				Int16("int16", -1).
				Int32("int32", -1).
				Int64("int64", -1).
				Uint("uint", 1).
				Uint8("uint8", 1).
				Uint16("uint16", 1).
				Uint32("uint32", 1).
				Uint64("uint64", 1).
				Float32("float32", 3.14).
				Float64("float64", 6.28).
				Bool("bool", true).
				Str("string", "").
				Bytes("[]byte", []byte{0x12, 0x34, 0x56}).
				Stringer("stringer", time.Second).
				Any("any", nil).
				Err(errors.New("err")).
				Send()
		}
	})
}

func BenchmarkPretty(b *testing.B) {
	devnull, err := os.Open(getNullPath())
	if err != nil {
		b.Fatal(err)
	}

	prettylog.Out = devnull
	prettylog.Level = logger.LogLevelDebug

	b.Run("one-field", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			prettylog.Debug("Benchmark").Int("int", 1)
		}
	})

	b.Run("two-field", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			prettylog.Debug("Benchmark").Int("int", 1).Float32("pi", 3.14)
		}
	})

	b.Run("all", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			prettylog.
				Debug("All").
				Int("int", -1).
				Int8("int8", -1).
				Int16("int16", -1).
				Int32("int32", -1).
				Int64("int64", -1).
				Uint("uint", 1).
				Uint8("uint8", 1).
				Uint16("uint16", 1).
				Uint32("uint32", 1).
				Uint64("uint64", 1).
				Float32("float32", 3.14).
				Float64("float64", 6.28).
				Bool("bool", true).
				Str("string", "").
				Bytes("[]byte", []byte{0x12, 0x34, 0x56}).
				Stringer("stringer", time.Second).
				Any("any", nil).
				Err(errors.New("err")).
				Send()
		}
	})
}

func getStr() string {
	defer buf.Reset()
	return buf.String()
}

func getNullPath() string {
	if runtime.GOOS == "windows" {
		return "nul"
	} else {
		return "/dev/null"
	}
}

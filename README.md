# Logger

Logger is a simple and fast logger that can write both machine-readable and human-readable logs.

### Why?

I really like zerolog, and especially its `ConsoleLogger`. I used it in all my command-line applications, but it has an issue. Zerolog can only output in JSON or CBOR. It can't produce human-readable output such as what the `ConsoleLogger` outputs, so zerolog unmarshals JSON generated by the logger and then makes human-readable output out of it. In my opinion, that's incredibly wasteful, so I made a light structured logger that can write logs in whatever format I need. Logger can use any logging implementation that implements the `logger.Logger` interface.

### Benchmarks

Logger is very fast. Here are its benchmarks, done on my laptop:

```text
goos: linux
goarch: amd64
pkg: logger
cpu: 11th Gen Intel(R) Core(TM) i7-1185G7 @ 3.00GHz
BenchmarkJSON/one-field-8                4216245               294.0 ns/op           160 B/op          3 allocs/op
BenchmarkJSON/two-field-8                1939634               594.3 ns/op           188 B/op          5 allocs/op
BenchmarkJSON/all-8                       310526              3955 ns/op             752 B/op         21 allocs/op
BenchmarkPretty/one-field-8              1603789               658.5 ns/op           168 B/op          4 allocs/op
BenchmarkPretty/two-field-8              1388920               864.5 ns/op           200 B/op          6 allocs/op
BenchmarkPretty/all-8                     285554              3726 ns/op             760 B/op         22 allocs/op
```

To run the benchmarks yourself, simply clone this repo and run `go test -bench=.`. Keep in mind that they will be different, depending on what your computer's specs are.

### Example

```go
package main

import (
    "strconv"
    "os"

    "go.arsenm.dev/logger/log"
    "go.arsenm.dev/logger"
)

func init() {
    // Default logger is JSONLogger going to stderr
    log.Logger = logger.NewPretty(os.Stdout)
}

func main() {
    s := "hello"
    i, err := strconv.Atoi(s)
    if err != nil {
        log.Error("Couldn't convert to integer").
            Str("input", s).
            Err(err).
            Send()
    }
    log.Info("Converted to integer").
        Int("output", i).
        Send()
}
```

# Jog Logging Library

Jog is a logging library for Go (Golang) designed to provide a simple and efficient way to handle logging in your Go applications. It leverages the `zap` logging library for high-performance logging. (Basically it's a wrapper around `zap` with(out) some additional features.)

## Features

- Provides various logging levels: Debug, Info, Warn, Error, Fatal, Panic.
- Supports structured logging with key-value pairs.
- Easy-to-use interface for logging messages with different severity levels.
- Integration with `zap` for efficient and fast logging.

## Installation

To install Jog, you need to have Go installed on your system. You can then use `go get` to install the library:

```sh
go get github.com/oeasenet/jog
```

## Usage

Here is a basic example of how to use Jog in your Go application:

```go
package main

import (
    "github.com/oeasenet/jog"
)

func main() {
    // Initialize the logger (optional, default instance uses dev level)
    jog.New(jog.LevelProd)

    // Logging examples
    jog.Debug("This is a debug message")
    jog.Info("This is an info message")
    jog.Warn("This is a warning message")
    jog.Error("This is an error message")
    jog.Fatal("This is a fatal message")
}
```

## Logger Interface

The `Logger` interface defines the various logging methods available in Jog:

```go
type Logger interface {
    Debug(args ...any)
    Log(args ...any)
    Info(args ...any)
    Warn(args ...any)
    Error(args ...any)
    Fatal(args ...any)
    Panic(args ...any)

    Debugf(format string, args ...any)
    Logf(format string, args ...any)
    Infof(format string, args ...any)
    Warnf(format string, args ...any)
    Errorf(format string, args ...any)
    Fatalf(format string, args ...any)
    Panicf(format string, args ...any)

    Debugw(msg string, keysAndValues ...any)
    Infow(msg string, keysAndValues ...any)
    Warnw(msg string, keysAndValues ...any)
    Errorw(msg string, keysAndValues ...any)
    Fatalw(msg string, keysAndValues ...any)
    Panicw(msg string, keysAndValues ...any)

    GetZapLogger() *zap.Logger
    GetZapSugarLogger() *zap.SugaredLogger
    Close()
}
```

## Example

Here is a more comprehensive example that demonstrates various logging methods:

```go
package main

import (
    "github.com/oeasenet/jog"
)

func main() {
    // Initialize the logger
    logger := jog.New()

    // Plain log messages
    logger.Debug("Debug message")
    logger.Info("Info message")
    logger.Warn("Warning message")
    logger.Error("Error message")
    logger.Fatal("Fatal message")

    // Formatted log messages
    logger.Debugf("Debug message: %d", 1)
    logger.Infof("Info message: %s", "info")
    logger.Warnf("Warning message: %s", "warning")
    logger.Errorf("Error message: %s", "error")
    logger.Fatalf("Fatal message: %s", "fatal")

    // Structured log messages
    logger.Debugw("Debug message", "key1", "value1", "key2", "value2")
    logger.Infow("Info message", "key1", "value1", "key2", "value2")
    logger.Warnw("Warning message", "key1", "value1", "key2", "value2")
    logger.Errorw("Error message", "key1", "value1", "key2", "value2")
    logger.Fatalw("Fatal message", "key1", "value1", "key2", "value2")

    // Close the logger when done
    logger.Close()
}
```

## Testing

Jog includes a suite of tests to ensure the reliability and correctness of the logging functionality. You can run the tests using the `go test` command:

```sh
go test ./...
```

## License

This project is licensed under the Apache License 2.0. See the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## Acknowledgments

This library is built on top of the fantastic `zap` logging library. Kudos to the authors of `zap` for their excellent work.

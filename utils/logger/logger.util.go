package logger

type LoggerClient int

const (
	Zerolog LoggerClient = iota + 1
)

type LoggerConfig struct {
	DefaultLogLevel string
	LoggerClient    LoggerClient
}

type ILogger interface {
	Info(msg string)
	Infof(msg string, args ...interface{})

	Warn(msg string)
	Warnf(msg string, args ...interface{})

	Error(msg string)
	Errorf(msg string, args ...interface{})

	Fatal(msg string)
	Fatalf(msg string, args ...interface{})

	Debug(msg string)
	Debugf(msg string, args ...interface{})

	Panic(msg string)
	Panicf(msg string, args ...interface{})

	Trace(msg string)
	Tracef(msg string, args ...interface{})
}

type logger struct {
}

func GetLogger(config LoggerConfig) ILogger {

	switch config.LoggerClient {
	case Zerolog:
		return newZeroLogger()
	}
	panic("Invalid logger client")
}

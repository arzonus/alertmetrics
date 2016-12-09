package logger

type ILogger interface {
	Debug(args ...interface{})
	Error(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})
}

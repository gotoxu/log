package core

import "go.uber.org/zap"

// Logger 日志器接口
type Logger interface {
	Print(v ...interface{})
	Printf(format string, args ...interface{})
	Println(v ...interface{})
	Debug(v ...interface{})
	Debugf(format string, args ...interface{})
	Debugln(v ...interface{})
	Info(v ...interface{})
	Infof(format string, args ...interface{})
	Infoln(v ...interface{})
	Warn(v ...interface{})
	Warnf(format string, args ...interface{})
	Warnln(v ...interface{})
	Error(v ...interface{})
	Errorf(format string, args ...interface{})
	Errorln(v ...interface{})
	Fatal(v ...interface{})
	Fatalf(format string, args ...interface{})
	Fatalln(v ...interface{})
	Panic(v ...interface{})
	Panicf(format string, args ...interface{})
	Panicln(v ...interface{})
	With(key string, value interface{}) Logger
	Sync() error
}

// NewLogger 创建一个新的日志器，该日志器实现Logger接口
func NewLogger(log *zap.SugaredLogger) Logger {
	w := &wrapper{log: log}
	return w
}

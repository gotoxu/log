package core

import (
	"fmt"

	"go.uber.org/zap"
)

// Logger 是zap.Logger的包装器
type Logger struct {
	log *zap.Logger
}

func (l *Logger) Print(v ...interface{}) {
	l.log.Info(fmt.Sprint(v))
}

func (l *Logger) Printf(format string, args ...interface{}) {
	l.log.Info(fmt.Sprintf(format, args...))
}

func (l *Logger) Println(v ...interface{}) {
	l.log.Info(fmt.Sprintln(v))
}

func (l *Logger) Debug(v ...interface{}) {
	l.log.Debug(fmt.Sprint(v))
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.log.Debug(fmt.Sprintf(format, args...))
}

func (l *Logger) Debugln(v ...interface{}) {
	l.log.Debug(fmt.Sprintln(v))
}

func (l *Logger) Info(v ...interface{}) {
	l.log.Info(fmt.Sprint(v))
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.log.Info(fmt.Sprintf(format, args...))
}

func (l *Logger) Infoln(v ...interface{}) {
	l.log.Info(fmt.Sprintln(v))
}

func (l *Logger) Warn(v ...interface{}) {
	l.log.Warn(fmt.Sprint(v))
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.log.Warn(fmt.Sprintf(format, args...))
}

func (l *Logger) Warnln(v ...interface{}) {
	l.log.Warn(fmt.Sprintln(v))
}

func (l *Logger) Error(v ...interface{}) {
	l.log.Error(fmt.Sprint(v))
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.log.Error(fmt.Sprintf(format, args...))
}

func (l *Logger) Errorln(v ...interface{}) {
	l.log.Error(fmt.Sprintln(v))
}

func (l *Logger) Fatal(v ...interface{}) {
	l.log.Fatal(fmt.Sprint(v))
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.log.Fatal(fmt.Sprintf(format, args...))
}

func (l *Logger) Fatalln(v ...interface{}) {
	l.log.Fatal(fmt.Sprintln(v))
}

func (l *Logger) Panic(v ...interface{}) {
	l.log.Panic(fmt.Sprint(v))
}

func (l *Logger) Panicf(format string, args ...interface{}) {
	l.log.Panic(fmt.Sprintf(format, args...))
}

func (l *Logger) Panicln(v ...interface{}) {
	l.log.Panic(fmt.Sprintln(v))
}

func (l *Logger) With(key string, value interface{}) *Logger {
	return &Logger{l.log.With(zap.Any(key, value))}
}

func (l *Logger) Sync() error {
	return l.log.Sync()
}

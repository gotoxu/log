package core

import (
	"fmt"

	"go.uber.org/zap"
)

// wrapper 是一个zap logger的包装器
type wrapper struct {
	log *zap.SugaredLogger
}

func (l *wrapper) Print(v ...interface{}) {
	l.log.Info(fmt.Sprint(v))
}

func (l *wrapper) Printf(format string, args ...interface{}) {
	l.log.Info(fmt.Sprintf(format, args...))
}

func (l *wrapper) Println(v ...interface{}) {
	l.log.Info(fmt.Sprintln(v))
}

func (l *wrapper) Debug(v ...interface{}) {
	l.log.Debug(fmt.Sprint(v))
}

func (l *wrapper) Debugf(format string, args ...interface{}) {
	l.log.Debug(fmt.Sprintf(format, args...))
}

func (l *wrapper) Debugln(v ...interface{}) {
	l.log.Debug(fmt.Sprintln(v))
}

func (l *wrapper) Info(v ...interface{}) {
	l.log.Info(fmt.Sprint(v))
}

func (l *wrapper) Infof(format string, args ...interface{}) {
	l.log.Info(fmt.Sprintf(format, args...))
}

func (l *wrapper) Infoln(v ...interface{}) {
	l.log.Info(fmt.Sprintln(v))
}

func (l *wrapper) Warn(v ...interface{}) {
	l.log.Warn(fmt.Sprint(v))
}

func (l *wrapper) Warnf(format string, args ...interface{}) {
	l.log.Warn(fmt.Sprintf(format, args...))
}

func (l *wrapper) Warnln(v ...interface{}) {
	l.log.Warn(fmt.Sprintln(v))
}

func (l *wrapper) Error(v ...interface{}) {
	l.log.Error(fmt.Sprint(v))
}

func (l *wrapper) Errorf(format string, args ...interface{}) {
	l.log.Error(fmt.Sprintf(format, args...))
}

func (l *wrapper) Errorln(v ...interface{}) {
	l.log.Error(fmt.Sprintln(v))
}

func (l *wrapper) Fatal(v ...interface{}) {
	l.log.Fatal(fmt.Sprint(v))
}

func (l *wrapper) Fatalf(format string, args ...interface{}) {
	l.log.Fatal(fmt.Sprintf(format, args...))
}

func (l *wrapper) Fatalln(v ...interface{}) {
	l.log.Fatal(fmt.Sprintln(v))
}

func (l *wrapper) Panic(v ...interface{}) {
	l.log.Panic(fmt.Sprint(v))
}

func (l *wrapper) Panicf(format string, args ...interface{}) {
	l.log.Panic(fmt.Sprintf(format, args...))
}

func (l *wrapper) Panicln(v ...interface{}) {
	l.log.Panic(fmt.Sprintln(v))
}

func (l *wrapper) With(key string, value interface{}) Logger {
	return &wrapper{l.log.With(zap.Any(key, value))}
}

func (l *wrapper) Sync() error {
	return l.log.Sync()
}

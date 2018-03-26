package syslog

import (
	"fmt"

	"github.com/ycyz/log/core"
	"go.uber.org/zap"
)

type logger struct {
	log *zap.Logger
}

func (l *logger) Print(v ...interface{}) {
	l.log.Info(fmt.Sprint(v))
}

func (l *logger) Printf(format string, args ...interface{}) {
	l.log.Info(fmt.Sprintf(format, args...))
}

func (l *logger) Println(v ...interface{}) {
	l.log.Info(fmt.Sprintln(v))
}

func (l *logger) Debug(v ...interface{}) {
	l.log.Debug(fmt.Sprint(v))
}

func (l *logger) Debugf(format string, args ...interface{}) {
	l.log.Debug(fmt.Sprintf(format, args...))
}

func (l *logger) Debugln(v ...interface{}) {
	l.log.Debug(fmt.Sprintln(v))
}

func (l *logger) Info(v ...interface{}) {
	l.log.Info(fmt.Sprint(v))
}

func (l *logger) Infof(format string, args ...interface{}) {
	l.log.Info(fmt.Sprintf(format, args...))
}

func (l *logger) Infoln(v ...interface{}) {
	l.log.Info(fmt.Sprintln(v))
}

func (l *logger) Warn(v ...interface{}) {
	l.log.Warn(fmt.Sprint(v))
}

func (l *logger) Warnf(format string, args ...interface{}) {
	l.log.Warn(fmt.Sprintf(format, args...))
}

func (l *logger) Warnln(v ...interface{}) {
	l.log.Warn(fmt.Sprintln(v))
}

func (l *logger) Error(v ...interface{}) {
	l.log.Error(fmt.Sprint(v))
}

func (l *logger) Errorf(format string, args ...interface{}) {
	l.log.Error(fmt.Sprintf(format, args...))
}

func (l *logger) Errorln(v ...interface{}) {
	l.log.Error(fmt.Sprintln(v))
}

func (l *logger) Fatal(v ...interface{}) {
	l.log.Fatal(fmt.Sprint(v))
}

func (l *logger) Fatalf(format string, args ...interface{}) {
	l.log.Fatal(fmt.Sprintf(format, args...))
}

func (l *logger) Fatalln(v ...interface{}) {
	l.log.Fatal(fmt.Sprintln(v))
}

func (l *logger) Panic(v ...interface{}) {
	l.log.Panic(fmt.Sprint(v))
}

func (l *logger) Panicf(format string, args ...interface{}) {
	l.log.Panic(fmt.Sprintf(format, args...))
}

func (l *logger) Panicln(v ...interface{}) {
	l.log.Panic(fmt.Sprintln(v))
}

func (l *logger) With(key string, value interface{}) core.Logger {
	return &logger{l.log.With(zap.Any(key, value))}
}

func (l *logger) Sync() error {
	return l.log.Sync()
}

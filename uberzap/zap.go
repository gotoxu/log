package uberzap

import (
	"fmt"

	"github.com/gotoxu/log/core"
	"go.uber.org/zap"
)

// NewLogger 创建一个新的日志器，该日志器实现Logger接口
func NewLogger(log *zap.SugaredLogger) core.Logger {
	w := &wrapper{log: log}
	return w
}

// wrapper 是一个zap logger的包装器
type wrapper struct {
	log *zap.SugaredLogger
}

func (l *wrapper) Log(level core.Level, v ...interface{}) {
	if level == core.Debug {
		l.log.Debug(v...)
		return
	}

	if level == core.Info {
		l.log.Info(v...)
		return
	}

	if level == core.Warn {
		l.log.Warn(v...)
		return
	}

	if level == core.Error {
		l.log.Error(v...)
		return
	}

	if level == core.Panic {
		l.log.Panic(v...)
		return
	}

	l.log.Fatal(v...)
}

func (l *wrapper) Logf(level core.Level, format string, v ...interface{}) {
	if level == core.Debug {
		l.log.Debugf(format, v...)
		return
	}

	if level == core.Info {
		l.log.Infof(format, v...)
		return
	}

	if level == core.Warn {
		l.log.Warnf(format, v...)
		return
	}

	if level == core.Error {
		l.log.Errorf(format, v...)
		return
	}

	if level == core.Panic {
		l.log.Panicf(format, v...)
		return
	}

	l.log.Fatalf(format, v...)
}

func (l *wrapper) Logln(level core.Level, v ...interface{}) {
	if level == core.Debug {
		l.log.Debug(fmt.Sprintln(v...))
		return
	}

	if level == core.Info {
		l.log.Info(fmt.Sprintln(v...))
		return
	}

	if level == core.Warn {
		l.log.Warn(fmt.Sprintln(v...))
		return
	}

	if level == core.Error {
		l.log.Error(fmt.Sprintln(v...))
		return
	}

	if level == core.Panic {
		l.log.Panic(fmt.Sprintln(v...))
		return
	}

	l.log.Fatal(fmt.Sprintln(v...))
}

func (l *wrapper) With(key string, value interface{}) core.Logger {
	return &wrapper{l.log.With(zap.Any(key, value))}
}

func (l *wrapper) Sync() error {
	return l.log.Sync()
}

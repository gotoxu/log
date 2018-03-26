// +build !windows,!nacl,!plan9

package syslog

import (
	"log/syslog"
	"os"

	"go.uber.org/zap"

	"github.com/ycyz/log/core"
	"go.uber.org/zap/zapcore"
)

// NewSyncer 创建新的syslog同步器
func newSyncer() (*Syncer, error) {
	s := &Syncer{}
	return s, s.connect()
}

// New 新建syslog
func New(debugLevel bool, app string) (core.Logger, error) {
	enc := newEncoder(config{
		EncoderConfig: zapcore.EncoderConfig{
			NameKey:       "logger",
			CallerKey:     "caller",
			MessageKey:    "msg",
			StacktraceKey: "stacktrace",
			EncodeLevel:   zapcore.CapitalLevelEncoder,
			EncodeTime:    zapcore.ISO8601TimeEncoder,
			EncodeCaller:  core.DefaultCallerEncoder,
		},

		Facility: syslog.LOG_LOCAL0,
		Hostname: "localhost",
		PID:      os.Getpid(),
		App:      app,
	})

	sink, err := newSyncer()
	if err != nil {
		return nil, err
	}

	var atom zap.AtomicLevel
	if debugLevel {
		atom = zap.NewAtomicLevelAt(zap.DebugLevel)
	} else {
		atom = zap.NewAtomicLevelAt(zap.InfoLevel)
	}

	_log := zap.New(zapcore.NewCore(enc, zapcore.Lock(sink), atom))
	return &logger{log: _log}, nil
}

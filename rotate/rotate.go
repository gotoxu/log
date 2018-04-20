package rotate

import (
	"path/filepath"

	"github.com/gotoxu/log/core"
	"github.com/gotoxu/log/uberzap"
	"github.com/robfig/cron"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

// New 创建一个新的rotate file logger
func New(debugLevel bool, options ...Option) core.Logger {
	cfg := defaultConfig()
	for _, opt := range options {
		opt(cfg)
	}

	writer := newRollingFile(cfg)
	zl := newZap(debugLevel, writer)

	return uberzap.NewLogger(zl.Sugar())
}

func newZap(debugLevel bool, output zapcore.WriteSyncer) *zap.Logger {
	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeCaller = uberzap.DefaultCallerEncoder
	cfg.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncodeDuration = zapcore.SecondsDurationEncoder
	cfg.EncodeLevel = zapcore.LowercaseLevelEncoder

	var atom zap.AtomicLevel
	if debugLevel {
		atom = zap.NewAtomicLevelAt(zap.DebugLevel)
	} else {
		atom = zap.NewAtomicLevelAt(zap.InfoLevel)
	}

	opts := []zap.Option{zap.AddCaller()}
	opts = append(opts, zap.AddStacktrace(atom))

	encoder := zapcore.NewConsoleEncoder(cfg)
	return zap.New(zapcore.NewCore(encoder, output, atom), opts...)
}

func newRollingFile(config *Config) zapcore.WriteSyncer {
	lj := lumberjack.Logger{
		Filename:   filepath.Join(config.Directory, config.Filename),
		MaxSize:    config.MaxSize,
		MaxAge:     config.MaxAge,
		MaxBackups: config.MaxBackups,
		LocalTime:  true,
		Compress:   config.Compress,
	}

	c := cron.New()
	c.AddFunc("@daily", func() { lj.Rotate() })
	c.Start()

	return zapcore.AddSync(&lj)
}

package logstash

import (
	"github.com/gotoxu/log/core"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// New 创建一个logstash日志器
func New(key, host string, debugLevel bool, options ...Option) core.Logger {
	cfg := newConfig(key, host)
	for _, opt := range options {
		opt(cfg)
	}

	writer := newRedis(cfg)
	zl := newZap(debugLevel, writer)

	return core.NewLogger(zl.Sugar())
}

func newZap(debugLevel bool, output zapcore.WriteSyncer) *zap.Logger {
	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeCaller = core.DefaultCallerEncoder
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

	encoder := zapcore.NewJSONEncoder(cfg)
	return zap.New(zapcore.NewCore(encoder, output, atom), opts...)
}

func newRedis(cfg *Config) zapcore.WriteSyncer {
	r := &redisSyncer{
		config: cfg,
	}

	return zapcore.AddSync(r)
}

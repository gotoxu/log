package file

import (
	"errors"
	"os"
	"path/filepath"

	"go.uber.org/zap"

	"github.com/natefinch/lumberjack"
	"github.com/ycyz/log/core"
	"go.uber.org/zap/zapcore"
)

// New 创建一个新的rotate file logger
func New(debugLevel bool, app string) (core.Logger, error) {
	if app == "" {
		return nil, errors.New("Please specify the application name")
	}

	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	path := filepath.Join(wd, "logs", "app.log")
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   path,
		MaxSize:    500,
		MaxBackups: 3,
		MaxAge:     28,
	})

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

	core := zapcore.NewCore(zapcore.NewJSONEncoder(cfg), w, atom)
	l := zap.New(core).Sugar()

	return &logger{log: l}, nil
}

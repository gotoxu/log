package uberzap

import (
	"runtime"
	"strings"

	"go.uber.org/zap/zapcore"
)

// DefaultCallerEncoder 默认的zapcore.CallerEncoder
func DefaultCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(strings.Join([]string{caller.TrimmedPath(), runtime.FuncForPC(caller.PC).Name()}, ":"))
}

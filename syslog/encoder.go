// +build !windows,!nacl,!plan9

package syslog

import (
	"log/syslog"
	"os"
	"path"
	"strings"
	"time"

	"github.com/ycyz/log/syslog/internal"
	"github.com/ycyz/log/syslog/internal/bufferpool"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

const (
	nilValue        = "-"
	maxHostnameLen  = 255
	maxAppNameLen   = 48
	facilityMask    = 0xf8
	severityMask    = 0x07
	version         = 1
	timestampFormat = "2006-01-02T15:04:05.000000Z07:00"
)

type framing int

const (
	nonTransparentFraming framing = iota
	octetCountingFraming
	defaultFraming = nonTransparentFraming
)

type jsonEncoder interface {
	zapcore.Encoder
	zapcore.ArrayEncoder
}

type config struct {
	zapcore.EncoderConfig

	Framing  framing
	Facility syslog.Priority
	PID      int
	Hostname string
	App      string
}

type encoder struct {
	*config
	je jsonEncoder
}

func rfc5424CompliantASCIIMapper(r rune) rune {
	if r < 33 || r > 126 {
		return '_'
	}

	return r
}

func toRFC5424CompliantASCIIString(s string) string {
	return strings.Map(rfc5424CompliantASCIIMapper, s)
}

func newEncoder(cfg config) zapcore.Encoder {
	if cfg.Hostname == "" {
		hostname, _ := os.Hostname()
		cfg.Hostname = hostname
	}
	if cfg.Hostname == "" {
		cfg.Hostname = nilValue
	} else {
		hostname := toRFC5424CompliantASCIIString(cfg.Hostname)
		if len(hostname) > maxHostnameLen {
			hostname = hostname[:maxHostnameLen]
		}
		cfg.Hostname = hostname
	}

	if cfg.PID == 0 {
		cfg.PID = os.Getpid()
	}
	if cfg.App == "" {
		cfg.App = nilValue
	} else {
		app := cfg.App
		if len(app) > maxAppNameLen {
			app = path.Base(app)
		}
		if len(app) > maxAppNameLen {
			app = app[:maxAppNameLen]
		}
		app = toRFC5424CompliantASCIIString(app)
		cfg.App = app
	}

	cfg.EncoderConfig.LineEnding = "\n"
	je := zapcore.NewJSONEncoder(cfg.EncoderConfig).(jsonEncoder)
	return &encoder{
		config: &cfg,
		je:     je,
	}
}

func (enc *encoder) AddArray(key string, arr zapcore.ArrayMarshaler) error {
	return enc.je.AddArray(key, arr)
}

func (enc *encoder) AddObject(key string, obj zapcore.ObjectMarshaler) error {
	return enc.je.AddObject(key, obj)
}

func (enc *encoder) AddBinary(key string, val []byte) {
	enc.je.AddBinary(key, val)
}

func (enc *encoder) AddByteString(key string, val []byte)      { enc.je.AddByteString(key, val) }
func (enc *encoder) AddBool(key string, val bool)              { enc.je.AddBool(key, val) }
func (enc *encoder) AddComplex128(key string, val complex128)  { enc.je.AddComplex128(key, val) }
func (enc *encoder) AddDuration(key string, val time.Duration) { enc.je.AddDuration(key, val) }
func (enc *encoder) AddFloat64(key string, val float64)        { enc.je.AddFloat64(key, val) }
func (enc *encoder) AddInt64(key string, val int64)            { enc.je.AddInt64(key, val) }

func (enc *encoder) AddReflected(key string, obj interface{}) error {
	return enc.je.AddReflected(key, obj)
}

func (enc *encoder) OpenNamespace(key string)          { enc.je.OpenNamespace(key) }
func (enc *encoder) AddString(key, val string)         { enc.je.AddString(key, val) }
func (enc *encoder) AddTime(key string, val time.Time) { enc.je.AddTime(key, val) }
func (enc *encoder) AddUint64(key string, val uint64)  { enc.je.AddUint64(key, val) }

func (enc *encoder) AppendArray(arr zapcore.ArrayMarshaler) error {
	return enc.je.AppendArray(arr)
}

func (enc *encoder) AppendObject(obj zapcore.ObjectMarshaler) error {
	return enc.je.AppendObject(obj)
}

func (enc *encoder) AppendBool(val bool)              { enc.je.AppendBool(val) }
func (enc *encoder) AppendByteString(val []byte)      { enc.je.AppendByteString(val) }
func (enc *encoder) AppendComplex128(val complex128)  { enc.je.AppendComplex128(val) }
func (enc *encoder) AppendDuration(val time.Duration) { enc.je.AppendDuration(val) }
func (enc *encoder) AppendInt64(val int64)            { enc.je.AppendInt64(val) }

func (enc *encoder) AppendReflected(val interface{}) error {
	return enc.je.AppendReflected(val)
}

func (enc *encoder) AppendString(val string)            { enc.je.AppendString(val) }
func (enc *encoder) AppendTime(val time.Time)           { enc.je.AppendTime(val) }
func (enc *encoder) AppendUint64(val uint64)            { enc.je.AppendUint64(val) }
func (enc *encoder) AddComplex64(k string, v complex64) { enc.je.AddComplex64(k, v) }
func (enc *encoder) AddFloat32(k string, v float32)     { enc.je.AddFloat32(k, v) }
func (enc *encoder) AddInt(k string, v int)             { enc.je.AddInt(k, v) }
func (enc *encoder) AddInt32(k string, v int32)         { enc.je.AddInt32(k, v) }
func (enc *encoder) AddInt16(k string, v int16)         { enc.je.AddInt16(k, v) }
func (enc *encoder) AddInt8(k string, v int8)           { enc.je.AddInt8(k, v) }
func (enc *encoder) AddUint(k string, v uint)           { enc.je.AddUint(k, v) }
func (enc *encoder) AddUint32(k string, v uint32)       { enc.je.AddUint32(k, v) }
func (enc *encoder) AddUint16(k string, v uint16)       { enc.je.AddUint16(k, v) }
func (enc *encoder) AddUint8(k string, v uint8)         { enc.je.AddUint8(k, v) }
func (enc *encoder) AddUintptr(k string, v uintptr)     { enc.je.AddUintptr(k, v) }
func (enc *encoder) AppendComplex64(v complex64)        { enc.je.AppendComplex64(v) }
func (enc *encoder) AppendFloat64(v float64)            { enc.je.AppendFloat64(v) }
func (enc *encoder) AppendFloat32(v float32)            { enc.je.AppendFloat32(v) }
func (enc *encoder) AppendInt(v int)                    { enc.je.AppendInt(v) }
func (enc *encoder) AppendInt32(v int32)                { enc.je.AppendInt32(v) }
func (enc *encoder) AppendInt16(v int16)                { enc.je.AppendInt16(v) }
func (enc *encoder) AppendInt8(v int8)                  { enc.je.AppendInt8(v) }
func (enc *encoder) AppendUint(v uint)                  { enc.je.AppendUint(v) }
func (enc *encoder) AppendUint32(v uint32)              { enc.je.AppendUint32(v) }
func (enc *encoder) AppendUint16(v uint16)              { enc.je.AppendUint16(v) }
func (enc *encoder) AppendUint8(v uint8)                { enc.je.AppendUint8(v) }
func (enc *encoder) AppendUintptr(v uintptr)            { enc.je.AppendUintptr(v) }

func (enc *encoder) Clone() zapcore.Encoder {
	return enc.clone()
}

func (enc *encoder) clone() *encoder {
	clone := &encoder{
		config: enc.config,
		je:     enc.je.Clone().(jsonEncoder),
	}
	return clone
}

func (enc *encoder) EncodeEntry(ent zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	msg := bufferpool.Get()

	var p syslog.Priority
	switch ent.Level {
	case zapcore.FatalLevel:
		p = syslog.LOG_EMERG
	case zapcore.PanicLevel:
		p = syslog.LOG_CRIT
	case zapcore.DPanicLevel:
		p = syslog.LOG_CRIT
	case zapcore.ErrorLevel:
		p = syslog.LOG_ERR
	case zapcore.WarnLevel:
		p = syslog.LOG_WARNING
	case zapcore.InfoLevel:
		p = syslog.LOG_INFO
	case zapcore.DebugLevel:
		p = syslog.LOG_DEBUG
	}
	pr := int64((enc.Facility & facilityMask) | (p & severityMask))

	msg.AppendByte('<')
	msg.AppendInt(pr)
	msg.AppendByte('>')
	msg.AppendInt(version)

	ts := ent.Time.UTC().Format(timestampFormat)
	if ts == "" {
		ts = nilValue
	}
	msg.AppendByte(' ')
	msg.AppendString(ts)

	msg.AppendByte(' ')
	msg.AppendString(enc.Hostname)

	msg.AppendByte(' ')
	msg.AppendString(enc.Hostname)

	msg.AppendByte(' ')
	msg.AppendString(enc.App)

	msg.AppendByte(' ')
	msg.AppendInt(int64(enc.PID))

	msg.AppendString(" - -")

	json, err := enc.je.EncodeEntry(ent, fields)
	if json.Len() > 0 {
		msg.AppendString("\xef\xbb\xbf")
		bs := json.Bytes()
		if enc.Framing == octetCountingFraming {
			bs = bs[:len(bs)-1]
		}
		msg.AppendString(internal.BytesToString(bs))
	}

	if enc.Framing != octetCountingFraming {
		return msg, err
	}

	out := bufferpool.Get()
	out.AppendInt(int64(msg.Len()))
	out.AppendByte(' ')
	out.AppendString(internal.BytesToString(msg.Bytes()))
	msg.Free()
	return out, err
}

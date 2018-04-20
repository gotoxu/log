// +build linux darwin dragonfly freebsd netbsd openbsd solaris

package syslog

import (
	"bytes"
	"errors"
	"fmt"
	"log/syslog"
	"os"
	"strings"
	"sync"

	"github.com/gotoxu/log/core"
)

var bufferPool = sync.Pool{
	New: func() interface{} {
		bs := bytes.NewBufferString("")
		return bs
	},
}

func getBuffer() *bytes.Buffer {
	return bufferPool.Get().(*bytes.Buffer)
}

func putBuffer(buffer *bytes.Buffer) {
	buffer.Reset()
	bufferPool.Put(buffer)
}

// NewLogger 用来创建一个新的syslog日志器
func NewLogger(p Priority, facility, tag string) (core.Logger, error) {
	fPriority, err := facilityPriority(facility)
	if err != nil {
		return nil, err
	}

	priority := syslog.Priority(p) | fPriority
	l, err := newBuiltin(priority, tag)
	if err != nil {
		return nil, err
	}

	return &builtinLogger{l}, nil
}

type builtinLogger struct {
	*builtinWriter
}

func (b *builtinLogger) Log(level core.Level, v ...interface{}) {
	bs := getBuffer()
	defer putBuffer(bs)
	fmt.Fprint(bs, v...)

	if level == core.Debug {
		b.writeLevel(LOG_DEBUG, bs.Bytes())
		return
	}

	if level == core.Info {
		b.writeLevel(LOG_NOTICE, bs.Bytes())
		return
	}

	if level == core.Warn {
		b.writeLevel(LOG_WARNING, bs.Bytes())
		return
	}

	if level == core.Error {
		b.writeLevel(LOG_ERR, bs.Bytes())
		return
	}

	if level == core.Panic {
		b.writeLevel(LOG_CRIT, bs.Bytes())
		panic(errors.New(bs.String()))
	}

	if level == core.Fatal {
		b.writeLevel(LOG_EMERG, bs.Bytes())
		os.Exit(1)
	}
}

func (b *builtinLogger) Logf(level core.Level, format string, v ...interface{}) {
	bs := getBuffer()
	defer putBuffer(bs)
	fmt.Fprintf(bs, format, v...)

	if level == core.Debug {
		b.writeLevel(LOG_DEBUG, bs.Bytes())
		return
	}

	if level == core.Info {
		b.writeLevel(LOG_NOTICE, bs.Bytes())
		return
	}

	if level == core.Warn {
		b.writeLevel(LOG_WARNING, bs.Bytes())
		return
	}

	if level == core.Error {
		b.writeLevel(LOG_ERR, bs.Bytes())
		return
	}

	if level == core.Panic {
		b.writeLevel(LOG_CRIT, bs.Bytes())
		panic(errors.New(bs.String()))
	}

	if level == core.Fatal {
		b.writeLevel(LOG_EMERG, bs.Bytes())
		os.Exit(1)
	}
}

func (b *builtinLogger) Logln(level core.Level, v ...interface{}) {
	bs := getBuffer()
	defer putBuffer(bs)
	fmt.Fprintln(bs, v...)

	if level == core.Debug {
		b.writeLevel(LOG_DEBUG, bs.Bytes())
		return
	}

	if level == core.Info {
		b.writeLevel(LOG_NOTICE, bs.Bytes())
		return
	}

	if level == core.Warn {
		b.writeLevel(LOG_WARNING, bs.Bytes())
		return
	}

	if level == core.Error {
		b.writeLevel(LOG_ERR, bs.Bytes())
		return
	}

	if level == core.Panic {
		b.writeLevel(LOG_CRIT, bs.Bytes())
		panic(errors.New(bs.String()))
	}

	if level == core.Fatal {
		b.writeLevel(LOG_EMERG, bs.Bytes())
		os.Exit(1)
	}
}

func (b *builtinLogger) With(key string, value interface{}) core.Logger {
	return b
}

func (b *builtinLogger) Sync() error {
	return nil
}

// Priority maps to the syslog priority levels
type Priority int

// consts
const (
	LOG_EMERG Priority = iota
	LOG_ALERT
	LOG_CRIT
	LOG_ERR
	LOG_WARNING
	LOG_NOTICE
	LOG_INFO
	LOG_DEBUG
)

func (b *builtinLogger) writeLevel(p Priority, buf []byte) error {
	var err error
	m := string(buf)
	switch p {
	case LOG_EMERG:
		_, err = b.writeAndRetry(syslog.LOG_EMERG, m)
	case LOG_ALERT:
		_, err = b.writeAndRetry(syslog.LOG_ALERT, m)
	case LOG_CRIT:
		_, err = b.writeAndRetry(syslog.LOG_CRIT, m)
	case LOG_ERR:
		_, err = b.writeAndRetry(syslog.LOG_ERR, m)
	case LOG_WARNING:
		_, err = b.writeAndRetry(syslog.LOG_WARNING, m)
	case LOG_NOTICE:
		_, err = b.writeAndRetry(syslog.LOG_NOTICE, m)
	case LOG_INFO:
		_, err = b.writeAndRetry(syslog.LOG_INFO, m)
	case LOG_DEBUG:
		_, err = b.writeAndRetry(syslog.LOG_DEBUG, m)
	default:
		err = fmt.Errorf("Unknown priority: %v", p)
	}
	return err
}

func facilityPriority(facility string) (syslog.Priority, error) {
	facility = strings.ToUpper(facility)
	switch facility {
	case "KERN":
		return syslog.LOG_KERN, nil
	case "USER":
		return syslog.LOG_USER, nil
	case "MAIL":
		return syslog.LOG_MAIL, nil
	case "DAEMON":
		return syslog.LOG_DAEMON, nil
	case "AUTH":
		return syslog.LOG_AUTH, nil
	case "SYSLOG":
		return syslog.LOG_SYSLOG, nil
	case "LPR":
		return syslog.LOG_LPR, nil
	case "NEWS":
		return syslog.LOG_NEWS, nil
	case "UUCP":
		return syslog.LOG_UUCP, nil
	case "CRON":
		return syslog.LOG_CRON, nil
	case "AUTHPRIV":
		return syslog.LOG_AUTHPRIV, nil
	case "FTP":
		return syslog.LOG_FTP, nil
	case "LOCAL0":
		return syslog.LOG_LOCAL0, nil
	case "LOCAL1":
		return syslog.LOG_LOCAL1, nil
	case "LOCAL2":
		return syslog.LOG_LOCAL2, nil
	case "LOCAL3":
		return syslog.LOG_LOCAL3, nil
	case "LOCAL4":
		return syslog.LOG_LOCAL4, nil
	case "LOCAL5":
		return syslog.LOG_LOCAL5, nil
	case "LOCAL6":
		return syslog.LOG_LOCAL6, nil
	case "LOCAL7":
		return syslog.LOG_LOCAL7, nil
	default:
		return 0, fmt.Errorf("invalid syslog facility: %s", facility)
	}
}

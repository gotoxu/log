package syslog

import (
	"errors"
	"fmt"
	"log/syslog"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

const severityMask = 0x07
const facilityMask = 0xf8
const localDeadline = 20 * time.Millisecond
const remoteDeadline = 50 * time.Millisecond

// builtinWriter 用来连接到系统syslog服务
type builtinWriter struct {
	priority syslog.Priority
	tag      string
	hostname string

	mu   sync.Mutex
	conn serverConn
}

func newBuiltin(priority syslog.Priority, tag string) (*builtinWriter, error) {
	if priority < 0 || priority > syslog.LOG_LOCAL7|syslog.LOG_DEBUG {
		return nil, errors.New("log/syslog: invalid priority")
	}

	if tag == "" {
		tag = os.Args[0]
	}
	hostname, _ := os.Hostname()

	w := &builtinWriter{
		priority: priority,
		tag:      tag,
		hostname: hostname,
	}

	w.mu.Lock()
	defer w.mu.Unlock()

	err := w.connect()
	if err != nil {
		return nil, err
	}

	return w, nil
}

type serverConn interface {
	writeString(p syslog.Priority, hostname, tag, s, nl string) error
	close() error
}

type netConn struct {
	conn net.Conn
}

func (w *builtinWriter) Write(b []byte) (int, error) {
	return w.writeAndRetry(w.priority, string(b))
}

func (w *builtinWriter) Close() error {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.conn != nil {
		err := w.conn.close()
		w.conn = nil
		return err
	}

	return nil
}

func (w *builtinWriter) writeAndRetry(p syslog.Priority, s string) (int, error) {
	pr := (w.priority & facilityMask) | (p & severityMask)

	w.mu.Lock()
	defer w.mu.Unlock()

	if w.conn != nil {
		if n, err := w.write(pr, s); err == nil {
			return n, err
		}
	}

	if err := w.connect(); err != nil {
		return 0, err
	}

	return w.write(pr, s)
}

func (w *builtinWriter) connect() (err error) {
	if w.conn != nil {
		w.conn.close()
		w.conn = nil
	}

	w.conn, err = unixSyslog()
	if w.hostname == "" {
		w.hostname = "localhost"
	}

	return
}

func (w *builtinWriter) write(p syslog.Priority, msg string) (int, error) {
	nl := ""
	if !strings.HasSuffix(msg, "\n") {
		nl = "\n"
	}

	err := w.conn.writeString(p, w.hostname, w.tag, msg, nl)
	if err != nil {
		return 0, err
	}

	return len(msg), nil
}

func (n *netConn) writeString(p syslog.Priority, hostname, tag, msg, nl string) error {
	timestamp := time.Now().Format(time.Stamp)
	n.conn.SetWriteDeadline(time.Now().Add(localDeadline))
	_, err := fmt.Fprintf(n.conn, "<%d>%s %s[%d]: %s%s",
		p, timestamp,
		tag, os.Getpid(), msg, nl)
	return err
}

func (n *netConn) close() error {
	return n.conn.Close()
}

// unixSyslog 使用unix域套接字打开一个连接本机syslog守护进程的连接
func unixSyslog() (conn serverConn, err error) {
	logTypes := []string{"unixgram", "unix"}
	logPaths := []string{"/dev/log", "/var/run/syslog", "/var/run/log"}
	for _, network := range logTypes {
		for _, path := range logPaths {
			conn, err := net.DialTimeout(network, path, localDeadline)
			if err != nil {
				continue
			} else {
				return &netConn{conn: conn}, nil
			}
		}
	}
	return nil, errors.New("Unix syslog delivery error")
}

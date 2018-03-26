package syslog

import (
	"errors"
	"net"
	"time"
)

const localDeadline = 20 * time.Millisecond

// Syncer 表示unix syslog同步器
type Syncer struct {
	conn net.Conn
}

func (s *Syncer) Write(p []byte) (n int, err error) {
	if s.conn != nil {
		if n, err := s.conn.Write(p); err == nil {
			return n, err
		}
	}

	if err := s.connect(); err != nil {
		return 0, err
	}

	return s.conn.Write(p)
}

// Sync 是为了实现zapcore.WriteSyncer接口
func (s *Syncer) Sync() error {
	return nil
}

func (s *Syncer) connect() error {
	if s.conn != nil {
		s.conn.Close()
		s.conn = nil
	}

	var err error
	s.conn, err = unixSyslog()
	return err
}

func unixSyslog() (conn net.Conn, err error) {
	logTypes := []string{"unixgram", "unix"}
	logPaths := []string{"/dev/log", "/var/run/syslog", "/var/run/log"}
	for _, network := range logTypes {
		for _, path := range logPaths {
			conn, err = net.DialTimeout(network, path, localDeadline)
			if err != nil {
				continue
			} else {
				return
			}
		}
	}

	return nil, errors.New("Unix syslog delivery error")
}

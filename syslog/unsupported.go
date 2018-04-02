// +build windows plan9 nacl

package syslog

import (
	"fmt"

	"github.com/gotoxu/log/core"
)

// NewLogger 用来创建一个新的syslog日志器
func NewLogger(p Priority, facility, tag string) (core.Logger, error) {
	return nil, fmt.Errorf("Platform does not support syslog")
}

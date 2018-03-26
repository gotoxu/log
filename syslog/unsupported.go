// +build windows plan9

package syslog

import (
	"errors"

	"github.com/ycyz/log/core"
)

func newSyncer() (*Syncer, error) {
	return nil, errors.New("Platform does not support syslog")
}

func New(debugLevel bool, app string) (core.Logger, error) {
	return nil, errors.New("Platform does not support syslog")
}

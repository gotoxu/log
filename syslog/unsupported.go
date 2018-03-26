// +build windows plan9

package syslog

import (
	"fmt"
)

func newSyncer() (*Syncer, error) {
	return nil, fmt.Errorf("Platform does not support syslog")
}

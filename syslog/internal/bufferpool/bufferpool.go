package bufferpool

import (
	"go.uber.org/zap/buffer"
)

var (
	_pool = buffer.NewPool()

	// Get 提供zap bufferpool的实现
	Get = _pool.Get
)

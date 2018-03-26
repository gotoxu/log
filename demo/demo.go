package main

import (
	"github.com/ycyz/log/core"
	"github.com/ycyz/log/rotate"
)

func main() {
	fl := rotate.New(true)
	method1(fl)
}

func method1(logger core.Logger) {
	logger.Error("Test error: error message.")
}

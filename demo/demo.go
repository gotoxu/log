package main

import (
	"github.com/gotoxu/log/core"
	"github.com/gotoxu/log/rotate"
)

func main() {
	fl := rotate.New(true)
	method1(fl)
}

func method1(logger core.Logger) {
	logger.Log(core.Error, "Test error: error message.")
}

package main

import (
	"fmt"

	"github.com/ycyz/log/core"
	"github.com/ycyz/log/file"
)

func main() {
	fl, closer, err := file.New(true)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer closer.Close()

	method1(fl)
}

func method1(logger core.Logger) {
	logger.Error("Test error: error message.")
}

package main

import (
	"fmt"

	"github.com/ycyz/log/file"
)

func main() {
	fl, closer, err := file.New(true, "galaxy")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer closer.Close()

	fl.Warnln("test logger: 1")
}

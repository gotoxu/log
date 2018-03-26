package main

import (
	"fmt"

	"github.com/ycyz/log/file"
)

func main() {
	fl, err := file.New(true, "galaxy")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fl.Warnln("test logger: 1")
}

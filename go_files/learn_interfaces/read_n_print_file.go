package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	customArgs := os.Args[1:]
	filePath := customArgs[0]
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Errorf(err.Error())
		os.Exit(1)
	}
	io.Copy(os.Stdout, file)
}

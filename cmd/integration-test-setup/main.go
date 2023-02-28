package main

import (
	"flag"
	"log"
	"syscall"
)

func main() {
	var fileTTY string
	major := 5
	minor := 0

	flag.StringVar(&fileTTY, "file", "", "set the terminal device file name with path included")
	flag.Parse()

	if fileTTY == "" {
		log.Fatal("empty terminal device file name not allowed")
	}

	err := syscall.Mknod(fileTTY, syscall.S_IFCHR|0777, (major<<8)|minor)
	if err != nil {
		panic(err)
	}
}

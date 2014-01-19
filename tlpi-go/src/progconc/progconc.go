package main

import (
	. "arguments"
	. "check"
	"log"
	"syscall"
)

func main() {
	var numcalls int
	var err error
	args := Arguments()
	if !args.Exist() {
		numcalls = 10000000
	} else {
		numcalls, args, err = args.Int("n", 0, 0)
		Check(err)
	}

	log.Println("Calling syscall.Getppid()")

	for i := 0; i < numcalls; i++ {
		syscall.Getppid()
	}
}

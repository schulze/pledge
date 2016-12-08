package main

import (
	"fmt"
	"syscall"

	"github.com/schulze/pledge"
)

func main() {
	// syscall.Setrlimit(syscall.RLIMIT_CORE, &syscall.Rlimit{0, 0})
	err := pledge.Pledge("undefined promise", nil)
	if err == nil || err != syscall.EINVAL{
		fmt.Errorf("Got error %v, but should be EINVAL.\n", err)
	}
	fmt.Println("Pledged!")
}

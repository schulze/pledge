package main

import (
	"fmt"
	"syscall"

	"github.com/schulze/pledge"
)

func main() {
	syscall.Setrlimit(syscall.RLIMIT_CORE, &syscall.Rlimit{0, 0})
	err := pledge.Pledge("", nil)
	if err != nil {
		fmt.Errorf("err=%v, but should be nil.\n", err)
	}
	fmt.Println("Pledged!")
}

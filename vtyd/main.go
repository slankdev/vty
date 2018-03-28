package main

import (
	"fmt"
	"github.com/slankdev/vtyd"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("Starting quagga interpreter daemon")

	vty.Main()

	// Register clearn up function.
	sigs := make(chan os.Signal, 1)
	done := make(chan bool)

	signal.Ignore(syscall.SIGWINCH)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigs
		done <- true
	}()

	<-done
}

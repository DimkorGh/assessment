package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"assessment/internal/cli"
	"assessment/internal/core/wiring"
)

func main() {
	argsReader := cli.NewArgsReader()
	address, port, err := argsReader.ReadArgs()
	if err != nil {
		log.Fatal(err.Error())
	}

	container := wiring.NewContainer(address, port)
	container.InitializeDependencies()

	osSignal := make(chan os.Signal)
	signal.Notify(osSignal, os.Interrupt, syscall.SIGTERM)

	go container.Server.Start()

	<-osSignal
	container.Server.Shutdown()
}

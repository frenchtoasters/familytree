package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"

	"github.com/frenchtoasters/familytree/cmd"
)

var onlyOneSignalHandler = make(chan struct{})

func main() {
	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	ctx := setupSignalHandler()
	command := cmd.NewFamilyTreeCommand(ctx)
	if err := command.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// setupSignalHandler creates context carrying system signals. A context is returned
// which is cancelled on one of these signals. If a second signal is caught, the program
// is terminated with exit code 1.
func setupSignalHandler() context.Context {
	close(onlyOneSignalHandler) // panics when called twice

	var shutdownSignals = []os.Signal{os.Interrupt}
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 2)
	signal.Notify(c, shutdownSignals...)
	go func() {
		<-c
		cancel()
		<-c
		os.Exit(1) // second signal. Exit directly.
	}()

	return ctx
}

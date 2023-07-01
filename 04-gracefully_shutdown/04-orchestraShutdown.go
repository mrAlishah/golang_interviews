package main

import (
	"context"
	"github.com/stephenafamo/orchestra"
	"net/http"
	"os"
	"syscall"
	"time"
)

func main() {
	var srv = &http.Server{}

	// creates a player from a myWorker function
	workerPlayer := orchestra.PlayerFunc(myWorker)
	// A player from a server
	serverPlayer := orchestra.ServerPlayer{srv}

	// A conductor to control them all
	conductor := &orchestra.Conductor{
		Timeout: 5 * time.Second,
		Players: map[string]orchestra.Player{
			// the names are used to identify the players
			// both in logs and the returned errors
			"worker": workerPlayer,
			"server": serverPlayer,
		},
	}

	// Use the conductor as a Player
	err := orchestra.PlayUntilSignal(conductor, os.Interrupt, syscall.SIGTERM)
	if err != nil {
		panic(err)
	}
}

func myWorker(ctx context.Context) error {

	var shouldStop = false

	go func() {
		<-ctx.Done()
		shouldStop = true
	}()

	for !shouldStop {
		err := doSomethingRepeatedly()
		if err != nil {
			return err
		}
	}

	return nil
}

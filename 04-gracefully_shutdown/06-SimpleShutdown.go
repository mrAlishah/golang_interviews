package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	var wg sync.WaitGroup

	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
		<-signals
		log.Println("shutting down")
		cancel()
	}()

	wg.Add(1)
	go func() {
		if err := myWorker(ctx); err != nil {
			log.Println("err := myWorker(ctx)")
			cancel()
		}
		log.Println("myWorker wg.Done()")
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		if err := startServer(ctx); err != nil {
			log.Println(" err := startServer(ctx)")
			cancel()
		}
		log.Println("startServer wg.Done()")
		wg.Done()
	}()

	wg.Wait()
}

func myWorker(ctx context.Context) error {

	var shouldStop = false

	go func() {
		<-ctx.Done()
		shouldStop = true
	}()

	for !shouldStop {
		time.Sleep(time.Millisecond * 1000)
		err := doSomethingRepeatedly()
		if err != nil {
			return err
		}
	}

	return nil
}

func doSomethingRepeatedly() error {
	log.Println("doSomethingRepeatedly")
	return nil
}

func startServer(ctx context.Context) error {

	var srv http.Server

	go func() {
		<-ctx.Done() // Wait for the context to be done

		// Shutdown the server
		if err := srv.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server Shutdown: %v", err)
		}
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		return fmt.Errorf("HTTP server ListenAndServe: %w", err)
	}

	return nil
}

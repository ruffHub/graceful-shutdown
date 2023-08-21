package app

import (
	"fmt"
	"github.com/ruffHub/graceful-shutdown/internal/events"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func Run(period int) {
	var wg sync.WaitGroup

	done := make(chan bool)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-done:
				return
			default:
				events.Listen(period)
			}
		}
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	fmt.Println("\nReceived shutdown signal. Closing tasks...")
	close(done)

	wg.Wait()
	fmt.Println("Program gracefully exited.")
}

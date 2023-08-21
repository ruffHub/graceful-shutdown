package events

import (
	"fmt"
	"time"
)

const (
	message = "Last sync date: %q. Data received in period of %d seconds: 0 \n"
)

// Listen
// Emulates listening some events and return collected data by period in seconds passed as argument
// Period should be passed as an argument in seconds.
func Listen(period int) {
	time.Sleep(time.Second * time.Duration(period))

	fmt.Printf(message, time.Now(), period)
}

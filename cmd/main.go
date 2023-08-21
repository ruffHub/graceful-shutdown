package main

import (
	"fmt"
	"github.com/ruffHub/graceful-shutdown/internal/app"
	"os"
	"strconv"
)

const (
	usage = "Usage: <seconds> - Frequency in seconds"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println(usage)
		return
	}

	period, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println(usage)
		return
	}

	app.Run(period)
}

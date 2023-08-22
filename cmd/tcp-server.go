package main

import (
	"context"
	"github.com/ruffHub/graceful-shutdown/internal/tcp-server/server"
	"os/signal"
	"syscall"
)

const (
	HOST = "localhost"
	PORT = "9000"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	s := server.New(&server.Config{Host: HOST, Port: PORT})
	s.Run(ctx)
}

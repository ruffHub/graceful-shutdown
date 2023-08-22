package server

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"slices"
)

type Server struct {
	host string
	port string
}

type Config struct {
	Host string
	Port string
}

func New(config *Config) *Server {
	return &Server{
		host: config.Host,
		port: config.Port,
	}
}

func (server *Server) Run(ctx context.Context) {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", server.host, server.port))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Listening on port:", server.port)

	connCh := make(chan net.Conn)
	errCh := make(chan error)

	// Accepting connections
	go func(l net.Listener, c chan net.Conn, e chan error) {
		for {
			conn, err := l.Accept()
			if err != nil {
				e <- err
				return
			}
			c <- conn
		}

	}(listener, connCh, errCh)

	for {
		select {
		case <-ctx.Done():
			listener.Close()
			log.Println("\nListener closed")
		case conn := <-connCh:
			go handleRequest(conn)
		case err = <-errCh:
			listener.Close()
			log.Fatal(err)
		default:

		}
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		bytes, err := reader.ReadBytes('\n')
		if err != nil {
			return
		}

		slices.Reverse(bytes)
		bytes = append(bytes, '\n')
		conn.Write(bytes)
	}
}

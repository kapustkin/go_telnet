package internal

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type myConnection struct {
	Connect net.Conn
	Context context.Context
}

func initialConnection(c *config) (*myConnection, error) {
	dialer := &net.Dialer{}
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Duration(c.Timeout)*time.Second)
	connect, err := dialer.DialContext(ctx, "tcp", fmt.Sprintf("%s:%d", c.Address, c.Port))
	if err != nil {
		return nil, err
	}

	cancelC := make(chan os.Signal)
	signal.Notify(cancelC, syscall.SIGINT)
	go func() {
		<-cancelC
		log.Println("Выход по SIGINT")
		cancel()
	}()

	connection := myConnection{Connect: connect, Context: ctx}

	return &connection, nil
}

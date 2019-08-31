package internal

import (
	flag "github.com/spf13/pflag"
)

// Config конфигурация
type config struct {
	Address string
	Port    int
	Timeout int
}

func initConfig() *config {
	cfg := config{}
	flag.IntVar(&cfg.Port, "port", 5000, "remote port")
	flag.StringVar(&cfg.Address, "host", "localhost", "remote host")
	flag.IntVar(&cfg.Timeout, "timeout", 30, "connection timeout")
	flag.Parse()
	return &cfg
}

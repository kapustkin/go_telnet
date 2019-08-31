package main

import (
	"log"

	"github.com/kapustkin/go_telnet/internal"
)

func main() {
	err := internal.Run()
	if err != nil {
		log.Fatal(err)
	}
}

package internal

import (
	"bufio"
	"io"
	"log"
)

type scanner struct {
	reader  io.Reader
	content chan string
}

func newScanner(reader io.Reader) scanner {
	return scanner{reader, make(chan string)}
}

func (scan scanner) start() {
	go func() {
		scanner := bufio.NewScanner(scan.reader)
		for scanner.Scan() {
			scan.content <- scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			log.Println(err)
		}
	}()
}

func (scan scanner) text() <-chan string {
	return scan.content
}

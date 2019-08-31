package internal

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
)

func reader(ctx context.Context, connect net.Conn) {
	scanner := newScanner(connect)
	scanner.start()
	for {
		select {
		case <-ctx.Done():
			return
		case content := <-scanner.text():
			log.Println(content)
		}
	}
}

func writer(ctx context.Context, connect net.Conn) {
	scanner := newScanner(os.Stdin)
	scanner.start()
	for {
		select {
		case <-ctx.Done():
			return

		case text := <-scanner.text():
			text = fmt.Sprintf("%s\n", text)
			_, err := connect.Write([]byte(text))
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
}

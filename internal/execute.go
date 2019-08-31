package internal

import (
	"sync"
)

// Run запуск программы
func Run() error {
	c := initConfig()
	connect, err := initialConnection(c)
	if err != nil {
		return err
	}
	startTelnet(connect)
	return nil
}

func startTelnet(conn *myConnection) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		reader(conn.Context, conn.Connect)
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		writer(conn.Context, conn.Connect)
		wg.Done()
	}()

	wg.Wait()
	conn.Connect.Close()
}

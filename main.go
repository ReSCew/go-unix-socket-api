package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"

	log "github.com/wired-R/minilog"
)

const SOCK = "/tmp/unixsocket"

func main() {
	log.Level = 0

	keyChan := make(chan os.Signal, 1)
	signal.Notify(keyChan, os.Interrupt)
	defer func() {
		os.Remove(SOCK)
		log.Info("Exit...")
	}()

	go listener()

	for {
		select {
		case <-keyChan:
			fmt.Println("trl+c: Exit")
			return
		}
	}

}

func listener() {
	l, err := net.Listen("unix", SOCK)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}
		var buf [1024]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			panic(err)
		}
		log.Info(fmt.Sprintf("%s", string(buf[:n])))
		conn.Close()
	}
}

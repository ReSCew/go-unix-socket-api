package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/ReSCew/go-unix-socket-api/socket"
	log "github.com/wired-R/minilog"
)

func main() {
	log.Level = 0
	keyChan := make(chan os.Signal, 1)
	signal.Notify(keyChan, os.Interrupt)
	defer func() {
		socket.CloseSocket()
		log.Info("Exit...")
	}()
	log.Info("Try to start Listener.")
	go socket.Listen()
	for {
		select {
		case <-keyChan:
			fmt.Println("trl+c: Exit")
			return
		}
	}
}

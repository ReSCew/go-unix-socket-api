package socket

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	log "github.com/wired-R/minilog"
)

const SOCK = "/tmp/rescew-push.sock"

//Lisener creates and listen socket file.
func Listener() {
	if _, err := os.Stat(SOCK); err == nil {
		log.Warning("Socket file is exist, deleteing...")
		os.Remove(SOCK)
		log.Warning("Done.")
	}

	listener, err := net.Listen("unix", SOCK)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Listener successful started.")
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, syscall.SIGTERM)
	go func(ln net.Listener, c chan os.Signal) {
		sig := <-c
		log.Info(fmt.Sprintf("Caught signal %s: shutting down.", sig))
		ln.Close()
		os.Exit(0)
	}(listener, sigc)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(fmt.Sprintf("Accept error: ", err))
		}

		var buf [1024]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			log.Error(err)
		}
		log.Info(string(buf[:n]))
		go handleMessage(buf[:n])
		conn.Close()
	}
}

//CloseSocket delete sock file
func CloseSocket() {
	log.Info("Try to close socket..")
	if err := os.Remove(SOCK); err != nil {
		log.Error(err)
		return
	}
	log.Info("Socket successful closed")

}

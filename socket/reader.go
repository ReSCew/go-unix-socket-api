package socket

import (
	"bufio"
	"net"
	"os"

	log "github.com/wired-R/minilog"
)

const SOCK = "/tmp/rescew-push.sock"

//Listener creates and listen socket file.
func Listen() {
	if _, err := os.Stat(SOCK); err == nil {
		log.Warning("Socket file is exist, deleteing...")
		os.Remove(SOCK)
		log.Warning("Done.")
	}

	listener, err := net.Listen("unix", SOCK)
	if err != nil {
		log.Info(err)
	}
	log.Info("Listener succesfull started.")

	conns := clientConn(listener)
	for {
		go handleConn(<-conns)
	}

	log.Info("Listener successful started")

}

func clientConn(listener net.Listener) chan net.Conn {
	ch := make(chan net.Conn)
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Error(err)
				continue
			}
			ch <- conn
		}
	}()
	return ch
}

func handleConn(client net.Conn) {
	b := bufio.NewReader(client)
	for {
		line, err := b.ReadBytes('\n')
		if err != nil { // EOF, or worse
			break
		}
		client.Write(line)
		log.Info(string(line))
		handleMessage(line)
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

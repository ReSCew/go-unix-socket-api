package socket

import (
	"io"
	"net"
	"os"

	log "github.com/wired-R/minilog"
)

const SOCK = "/tmp/rescew-push.sock"

//Listener creates and listen socket file.
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
	defer listener.Close()

	log.Info("Listener successful started")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Error(err)
		}
		go read(conn)
		// defer conn.Close()
	}
}

func read(r io.Reader) {
	var buf [1024]byte
	for {

		n, err := r.Read(buf[:])
		if err != nil {
			log.Error(err)
			return
		}
		log.Info(string(buf[:n]))
		handleMessage(buf[:n])
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

package socket

import (
	"encoding/json"
	"net"

	log "github.com/wired-R/minilog"
)

const WRITE_SOCK = "/tmp/rescew-push-result.sock"

func write(answer *Answer) {

	c, err := net.Dial("unix", WRITE_SOCK)
	if err != nil {
		log.Error(err)
		return
	}
	log.Info(answer.Result.Message)
	b, _ := json.Marshal(answer)

	_, err = c.Write(b)
	if err != nil {
		log.Error(err)
	}
}

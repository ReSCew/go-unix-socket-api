package socket

import (
	"encoding/json"
)

type InputMessage struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Answer struct {
	ID     string  `json:"id"`
	Result *Result `json:"result"`
}
type Result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func handleMessage(buf []byte) {
	messages := parseMessage(&buf)
	for _, messages := range *messages {
		answer := makeAnswer(messages)
		go write(answer)
	}
}

func parseMessage(buf *[]byte) *[]InputMessage {
	messages := []InputMessage{}
	json.Unmarshal(*buf, &messages)
	return &messages
}

func makeAnswer(message InputMessage) *Answer {
	answer := Answer{
		ID: message.ID,
	}
	result := Result{}
	result.Code = 0
	result.Message = message.Name
	answer.Result = &result
	return &answer

}

// func wait() {
// 	// delay := time.Duration((int64(500 + rand.Intn(500-6000))))
// 	time.Sleep(500 * time.Millisecond)
// }

package main

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type jsonMessage struct {
	Type     string        `json:"type"`
	Message  string        `json:"message,omitempty"`
	Messages []chatMessage `json:"messages,omitempty"`
	Count    int           `json:"count,omitempty"`
	Error    string        `json:"error,omitempty"`
}
type chatMessage struct {
	Id   string `json:"id,omitempty"`
	Time int64  `json:"time,omitempty"`
	Body string `json:"body"`
}

func main() {
	//ctx := context.Background()
	endpointUrl := "wss://jch.irif.fr:8443/chat/ws"
	d := websocket.Dialer{}
	conn, _, err := d.Dial(endpointUrl, nil)
	//c, _, err := d.DialContext(ctx, endpointUrl, nil)

	if err != nil {
		log.Panicf("Dial failed: %#v\n", err)
	}
	message := jsonMessage{
		Type:    "post",
		Message: "time to sleep",
	}
	for {
		conn.WriteJSON(message)
		/**
		response := jsonMessage{}
		p := conn.ReadJSON(&response)
		if p != nil {
			log.Println(p)
			return
		}

		for i, m := range response.Messages {

			fmt.Println("Message number : ", i)
			fmt.Println(m.Body)
		}**/
		time.Sleep(2 * time.Second)
	}
	conn.Close()
}

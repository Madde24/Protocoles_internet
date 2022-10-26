package main

//how to use Last-Modified in Go?

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
	//"strconv"
)

type jsonMessage struct {
	Id   string `json:"id"`
	Time int64  `json:"time"`
	Body string `json:"body"`
}

func request(resp *http.Response, id string) string {
	if resp.StatusCode != 200 {
		return id
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		defer resp.Body.Close()
		log.Fatal(err)
	}
	var r []jsonMessage
	json.Unmarshal(body, &r)
	flag := 0
	if id == "" {
		flag = 1
	}
	for i, m := range r {
		if m.Id == id {
			flag = 1
			continue
		}
		if flag == 1 {
			fmt.Println("Message number : ", len(r)-i)
			fmt.Println("Id :", m.Id)
			//tm, _ := time.Parse("Jan 2, 2006 at 3:04pm (MST)", string(m.Time))
			fmt.Println("Time :", m.Time)
			fmt.Println("Body :", m.Body)
		}
	}
	return r[len(r)-1].Id
}

func main() {
	client := http.Client{}
	//func NewRequest(method, target string, body io.Reader) *http.Request
	req, err := http.NewRequest("GET", "https://jch.irif.fr:8082/chat/messages.json?count=4", nil)

	if err != nil {
		log.Fatal("aie", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("uie", err)
	}
	id := request(resp, "")
	etag := resp.Header.Get("ETag")
	for true {
		//we build a new request to get the new messages on the chat
		req, err := http.NewRequest("GET", "https://jch.irif.fr:8082/chat/messages.json?count=4", nil)
		if err != nil {
			log.Fatal("nooon", err)
		}

		//we add a header tag to specifiy on which condition we actually do the request
		req.Header.Add("if-None-Match", etag)
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal("nooon", err)
		}
		//the request will only work if the code is 200, otherwise the request won't go through
		id = request(resp, id)
		etag = resp.Header.Get("ETag")
		time.Sleep(5 * time.Second)
	}

	// https://jch.irif.fr:8082/chat/messages.json?count=4 affiche 4 derniers messages
}

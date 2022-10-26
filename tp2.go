package main

//how to use etags in Go?

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
	//"strconv"
)

/*Doc pour les différentes requêtes et StatusCode :
https://pkg.go.dev/net/http
*/

func request(resp *http.Response, start int) int {
	if resp.StatusCode != 200 {
		return start
	}
	defer resp.Body.Close()
	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		defer resp.Body.Close()
		log.Fatal(err)
	}
	responseString := string(responseData)
	line := strings.Split(responseString, "\n")
	for i := start; i < len(line)-1; i++ {
		req := "https://jch.irif.fr:8082/chat/" + line[i]
		resp1, err := http.Get(req)
		if err != nil {
			log.Fatal("encore nooon", err)
		}
		responseData1, err := io.ReadAll(resp1.Body)
		if err != nil {
			log.Fatal("non", err)
		}
		responseString1 := string(responseData1)
		fmt.Println(responseString1)
	}
	return len(line) - 1
}

/*
func main() {
	i := 0
	etag := ""
	for true {

		inm := http.Header{}.Get("If-None-Match")

		fmt.Println(string(inm[:]))
		if etag != string(inm) {
			etag = string(inm)
			fmt.Println(string(etag))

			resp, err := http.Get("https://jch.irif.fr:8082/chat/")
			if err != nil {
				print(err)
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				print(err)
			}
			response := strings.Split(string(body), "\n")
			for ; i < len(response)-1; i++ {
				message, err := http.Get("https://jch.irif.fr:8082/chat/" + response[i])
				if err != nil {
					print(err)
				}
				defer resp.Body.Close()
				body_message, err := ioutil.ReadAll(message.Body)
				if err != nil {
					print(err)
				}
				fmt.Println(string(body_message))

			}
		}
		time.Sleep(10 * time.Second)
	}
}*/

func main() {
	client := http.Client{}
	//func NewRequest(method, target string, body io.Reader) *http.Request
	req, err := http.NewRequest("GET", "https://jch.irif.fr:8082/chat/", nil)

	if err != nil {
		log.Fatal("aie", err)
	}

	//func (c *Client) Do(req *Request) (*Response, error)
	//Do sends an HTTP request and returns an HTTP response, following policy (such as redirects, cookies, auth) as configured on the client.
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("uie", err)
	}
	start := request(resp, 0)
	etag := resp.Header.Get("ETag")
	for true {
		//we build a new request to get the new messages on the chat
		req, err := http.NewRequest("GET", "https://jch.irif.fr:8082/chat/", nil)
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
		start = request(resp, start)
		etag = resp.Header.Get("ETag")
		time.Sleep(5 * time.Second)
	}
}

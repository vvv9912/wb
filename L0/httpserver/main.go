package main

import (
	"encoding/json"
	"fmt"
	"github.com/patrickmn/go-cache"
	"httpserver/message"
	"io"
	"net/http"
	"os"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}
func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}
func main() {

	//http.HandleFunc("/", getRoot)
	//http.HandleFunc("/hello", getHello)
	//
	//err := http.ListenAndServe("127.0.0.1:3333", nil)
	//_ = err

	//how use Cache
	var msg modelmessage.ModelMessage
	messageinjson, err := os.ReadFile("message/model.json")
	err = json.Unmarshal(messageinjson, &msg)
	if err != nil {
		fmt.Println("Error: Создание кэша\n ", err.Error())
		return
	}
	c := cache.New(cache.DefaultExpiration, 0)
	err = c.Add(msg.OrderUid, messageinjson, cache.DefaultExpiration)
	err = c.Add("1", messageinjson, cache.DefaultExpiration)
	if err != nil {
		fmt.Println("Error: Ошибка при добавлении кэша\n ", err.Error())
		return
	}
	getmsg, found := c.Get(msg.OrderUid)
	fmt.Println(found)
	getmsg2, found := c.Get("1")
	fmt.Println(found)
	getmsg3, found := c.Get("3")
	fmt.Println(found)
	c.Flush()
	getmsg, found = c.Get(msg.OrderUid)
	fmt.Println(found)
	getmsg2, found = c.Get("1")
	fmt.Println(found)
	getmsg3, found = c.Get("3")
	fmt.Println(found)
	_ = getmsg
	_ = getmsg2
	_ = getmsg3
}

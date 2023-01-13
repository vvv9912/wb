package main

import (
	modelmessage "client_produser/message"
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"os"
)

// nats-streaming-server -V
func main() {

	clusterID := "test-cluster"
	clientID := "Produser"

	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL("0.0.0.0:4222"))
	if err != nil {
		fmt.Printf(err.Error())
	}
	defer sc.Close()
	//test start
	messageinjson, err := os.ReadFile("message/model.json")
	var message modelmessage.ModelMessage
	err = json.Unmarshal(messageinjson, &message)
	if err != nil {
		fmt.Println("Error: Неправильное сообщение (Неудалось распарсить) \n ", err.Error())
		return
	}
	subject := "foo"
	err = sc.Publish(subject, []byte(messageinjson))

	if err != nil {
		fmt.Printf(err.Error())
		return
	}
}

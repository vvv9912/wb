package main

import (
	modelmessage "client_produser/message"
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"math/rand"
	"os"
	"testing"
)

// Формируется последовательность из сообщений: сначала посылается правильное сообщение, затем неправильное
func TestXxx(t *testing.T) {
	clusterID := "test-cluster"
	clientID := "Produser"
	NatsUrl := "0.0.0.0:4222"
	subject := "foo"
	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL(NatsUrl))
	if err != nil {
		fmt.Printf(err.Error())
	}
	defer func(sc stan.Conn) {
		err := sc.Close()
		if err != nil {
			fmt.Printf(err.Error()) //todo
		}
	}(sc)
	messageinjson, err := os.ReadFile("message/model.json")
	var message modelmessage.ModelMessage
	err = json.Unmarshal(messageinjson, &message)
	if err != nil {
		fmt.Println("Error: Неправильное сообщение (Неудалось распарсить) \n ", err.Error())
		return
	}

	for i := 1; ; i++ {
		if (i % 2) == 0 {
			message.OrderUid = fmt.Sprintf("%d", i)
			fmt.Println(message.OrderUid)
			messageinjson, err = json.Marshal(message)
			if err != nil {
				fmt.Println("Error: ", err.Error())
				return
			}
		} else {
			messageinjson = []byte(fmt.Sprintf("%d", rand.Int()))
		}

		err = sc.Publish(subject, []byte(messageinjson))
		if err != nil {
			fmt.Printf(err.Error())
			return
		}
		fmt.Println(i)
	}

}

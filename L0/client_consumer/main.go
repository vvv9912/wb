package main

import (
	"fmt"
	"github.com/nats-io/stan.go"
	"time"
)

func main() {
	clusterID := "test-cluster"
	clientID := "Consumer"

	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL("0.0.0.0:4222"))
	if err != nil {
		fmt.Printf(err.Error())
	}
	defer sc.Close()
	fmt.Printf("Старт")

	sub, err := sc.Subscribe("foo", func(m *stan.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})
	//
	//stan.StartAtSequence(10) //start position
	if err != nil {
		fmt.Printf("error in subscribe: %v", err)
		return
	}
	time.Sleep(10 + time.Second)
	sub.Unsubscribe()

}

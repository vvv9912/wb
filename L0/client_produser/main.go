package main

import (
	"fmt"
	"github.com/nats-io/stan.go"
	"strconv"
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

	subject := "foo"
	for i := 1; ; i++ {
		msg := "РФРmsg " + strconv.Itoa(i)
		err := sc.Publish(subject, []byte(msg))
		if err != nil {
			fmt.Printf(err.Error())
			return
		}
		//time.Sleep(1 + time.Second)
		fmt.Printf(msg)
		_ = sc
	}

}

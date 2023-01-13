package main

import (
	"client_consumer/message"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/nats-io/stan.go"
	"github.com/patrickmn/go-cache"
	"io"
	"net/http"
)

func getCache(c *cache.Cache, key string) (interface{}, bool) {
	data, found := c.Get(key)
	return data, found
}

type Cache struct {
	cache *cache.Cache
}

func (cache Cache) getHttpAnswer(w http.ResponseWriter, r *http.Request) {
	key := r.RequestURI
	data, found := getCache(cache.cache, key[4:])
	if found {
		io.WriteString(w, string(data.([]byte))+"\n")
	} else {
		io.WriteString(w, "no found\n")
	}
}
func main() {
	//host := "localhost"
	//port := "5432"
	//database := "postgres"
	//username := "postgres"
	//password := "postgres"
	clusterID := "test-cluster"
	clientID := "Consumer"

	var c Cache
	c.cache = cache.New(cache.DefaultExpiration, 0)
	http.HandleFunc("/id/", c.getHttpAnswer)
	var db tdb
	db.openDB()
	defer db.db.Close()
	row := db.db.QueryRow("SELECT COUNT(*) FROM  message")
	var count int
	err := row.Scan(&count)
	if err != nil {
		fmt.Printf(err.Error()) //todo
		return
	}
	go db.bdtocache(c, count)
	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL("0.0.0.0:4222"))
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	defer sc.Close()
	fmt.Printf("Старт")

	sub, err := sc.Subscribe("foo", func(m *stan.Msg) {
		rowsAffected, err := db.addDB(m.Data)
		if err != nil {
			fmt.Printf("\nerr to bd: \n", err.Error())
		} else if err == nil {
			fmt.Printf("add: %d \n", rowsAffected)
			var msg modelmessage.ModelMessage
			err = json.Unmarshal(m.Data, &msg)
			if err != nil {
				fmt.Println("Error: Неправильное сообщение (Неудалось распарсить) \n ", err.Error())
				return
			}
			err = c.cache.Add(msg.OrderUid, m.Data, cache.DefaultExpiration) //todo
			if err != nil {
				fmt.Println(err.Error())
				return
			}

		}
	})
	defer sub.Unsubscribe()
	if err != nil {
		fmt.Printf("error in subscribe: %v", err)
		return
	}
	err = http.ListenAndServe("127.0.0.1:3333", nil)
	if err != nil {
		fmt.Printf(err.Error()) //to do
		return
	}
}

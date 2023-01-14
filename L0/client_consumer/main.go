package main

import (
	"client_consumer/message"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/nats-io/stan.go"
	"github.com/patrickmn/go-cache"
	"io"
	"net/http"
)

type Cache struct {
	cache *cache.Cache
}

func (c Cache) getCache(key string) (interface{}, bool) {
	data, found := c.cache.Get(key)
	return data, found
}
func (c Cache) getHttpAnswer(w http.ResponseWriter, r *http.Request) {
	key := r.RequestURI
	data, found := c.getCache(key[4:])
	if found {
		_, err := io.WriteString(w, string(data.([]byte))+"\n")
		if err != nil {
			return
		}
	} else {
		_, err := io.WriteString(w, "no found\n")
		if err != nil {
			return
		}
	}
}
func main() {

	clusterID := "test-cluster"
	clientID := "Consumer"
	NatsUrl := "0.0.0.0:4222"
	httpUrl := "127.0.0.1:3333"

	var c Cache
	c.cache = cache.New(cache.DefaultExpiration, 0)
	http.HandleFunc("/id/", c.getHttpAnswer)
	var db tdb
	db.openDB()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Printf(err.Error()) //todo
		}
	}(db.db)
	row := db.db.QueryRow("SELECT COUNT(*) FROM  message")
	var count int
	err := row.Scan(&count)
	if err != nil {
		fmt.Printf(err.Error()) //todo
		return
	}
	go db.bdtocache(c, count)
	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL(NatsUrl))
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	defer func(sc stan.Conn) {
		err := sc.Close()
		if err != nil {
			fmt.Printf(err.Error()) //todo
		}
	}(sc)
	fmt.Printf("Старт")

	sub, err := sc.Subscribe("foo", func(m *stan.Msg) {
		rowsAffected, err := db.addDB(m.Data)
		if err != nil {
			fmt.Printf("err to bd: \n", err.Error())
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
	defer func(sub stan.Subscription) {
		err := sub.Unsubscribe()
		if err != nil {
			fmt.Printf(err.Error()) //todo
		}
	}(sub)
	if err != nil {
		fmt.Printf("error in subscribe: %v", err)
		return
	}
	err = http.ListenAndServe(httpUrl, nil)
	if err != nil {
		fmt.Printf(err.Error()) //to do
		return
	}
}

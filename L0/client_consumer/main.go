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

	var c Cache
	c.cache = cache.New(cache.DefaultExpiration, 0)
	http.HandleFunc("/id/", c.getHttpAnswer)

	connStr := "postgres://postgres:postgres@localhost/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Printf(err.Error()) //to do
		return
	}
	// Получаем сколько в таблице зн-ий для переписывания их
	// https://www.calhoun.io/querying-for-a-single-record-using-gos-database-sql-package/
	row := db.QueryRow("SELECT COUNT(*) FROM  message")
	var count int
	err = row.Scan(&count)
	if err != nil {
		fmt.Printf(err.Error()) //todo
		return
	}
	defer db.Close()
	go bdtocache(db, c, count)

	//host := "localhost"
	//port := "5432"
	//database := "postgres"
	//username := "postgres"
	//password := "postgres"

	clusterID := "test-cluster"
	clientID := "Consumer"

	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL("0.0.0.0:4222"))
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	defer sc.Close()
	fmt.Printf("Старт")
	//yes := 0
	//no := 0

	sub, err := sc.Subscribe("foo", func(m *stan.Msg) {
		rowsAffected, err := addDB(m.Data, db)
		//time.Sleep(2 + time.Second)
		if err != nil {
			fmt.Printf("\nerr to bd: \n", err.Error())
			//	no++
		} else if err == nil {
			fmt.Printf("add: %d \n", rowsAffected)
			//yes++
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
	//
	//stan.StartAtSequence(10) //start position
	if err != nil {
		fmt.Printf("error in subscribe: %v", err)
		return
	}
	//fmt.Printf("yes:%d", yes)
	//fmt.Printf("no:%d", no)
	err = http.ListenAndServe("127.0.0.1:3333", nil)
	if err != nil {
		fmt.Printf(err.Error()) //to do
		return
	}
	//fmt.Scanf(" ")
	//
	//fmt.Printf("yes:%d", yes)
	//fmt.Printf("no:%d", no)
}

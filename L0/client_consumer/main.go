package main

import (
	//_ "./message/modelMessage"
	"client_consumer/message"
	"database/sql"
	"encoding/json"
	"errors"
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
func bdtocache(db *sql.DB, c Cache, count int) {
	for i := 1; i <= count; i++ {
		//
		order_uid, data, err := getDB(i, db)
		if err != nil {
			fmt.Printf(err.Error()) //todo
			break
		}
		err = c.cache.Add(order_uid, data, cache.DefaultExpiration)
		if err != nil {
			fmt.Printf(err.Error()) //todo
			break
		}
	}
}
func addDB(messagetojson []byte, db *sql.DB) (int64, error) {
	var message modelmessage.ModelMessage
	err := json.Unmarshal(messagetojson, &message)
	if err != nil {
		fmt.Println("Error: Неправильное сообщение (Неудалось распарсить) \n ", err.Error())
		return 0, err
	}
	rows, err := db.Query("select order_uid from message")
	if err != nil {
		fmt.Println("Error 3:", err.Error())
		return 0, err
	}
	defer rows.Close()
	for rows.Next() {
		var order_uid string
		err := rows.Scan(&order_uid)
		if err != nil {
			fmt.Println("Error 4: ", err.Error())
			return 0, err
		}
		if order_uid == message.OrderUid {
			//fmt.Println("Error: такой OrderUid существует\n")
			err2 := errors.New("Такой OrderUid существует")
			return 0, err2
		}
	}
	result, err := db.Exec("insert into message (order_uid, data) values ($1,$2)", message.OrderUid, messagetojson)
	if err != nil {
		fmt.Println("Error 5:", err.Error())
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Error 6:", err.Error())
		return 0, err
	}
	//fmt.Println(rowsAffected)
	return rowsAffected, nil
}

func getDB(row int, db *sql.DB) (string, []byte, error) {
	if row < 1 {
		err2 := errors.New("Неправильно задан столбец")
		return "", nil, err2
	}
	var order_uid string
	var data []byte

	rows, err := db.Query("select order_uid, data from message")
	if err != nil {
		fmt.Println("Error 3:", err.Error())
		return "", nil, err
	}
	defer rows.Close()
	for rows.Next() {
		row--
		if row == 0 {
			err := rows.Scan(&order_uid, &data)
			if err != nil {
				fmt.Println("Error 4: ", err.Error()) //todo
				return "", nil, err
			} else {
				return order_uid, data, nil
			}
		}
	}
	return "", nil, err
}

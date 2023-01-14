package main

import (
	modelmessage "client_consumer/message"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/patrickmn/go-cache"
)

type tdb struct {
	db *sql.DB
}

// go get github.com/lib/pq
func (db *tdb) openDB() {
	connStr := "postgres://postgres:postgres@localhost/postgres?sslmode=disable"
	dbb, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Printf(err.Error()) //to do
		return
	}
	db.db = dbb
}
func (db tdb) bdtocache(c Cache, count int) {
	for i := 1; i <= count; i++ {
		//
		order_uid, data, err := db.getDB(i)
		if err != nil {
			fmt.Printf(err.Error()) //todo
			return
		}
		err = c.cache.Add(order_uid, data, cache.DefaultExpiration)
		if err != nil {
			fmt.Printf(err.Error()) //todo
			return
		}
	}
}
func (db tdb) addDB(messagetojson []byte) (int64, error) {
	var message modelmessage.ModelMessage
	err := json.Unmarshal(messagetojson, &message)
	if err != nil {
		fmt.Println(err.Error()) //todo
		return 0, err
	}
	rows, err := db.db.Query("select order_uid from message")
	if err != nil {
		fmt.Println(err.Error()) //todo
		return 0, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println(err.Error()) //todo
		}
	}(rows)
	result, err := db.db.Exec("insert into message (order_uid, data) values ($1,$2)", message.OrderUid, messagetojson)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}
	return rowsAffected, nil
}

func (db tdb) getDB(row int) (string, []byte, error) {
	if row < 1 {
		err2 := errors.New("Неправильно задан столбец")
		return "", nil, err2
	}
	var order_uid string
	var data []byte

	rows, err := db.db.Query("select order_uid, data from message")
	if err != nil {
		fmt.Println(err.Error()) //todo
		return "", nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println(err.Error()) //todo
		}
	}(rows)
	for rows.Next() {
		row--
		if row == 0 {
			err := rows.Scan(&order_uid, &data)
			if err != nil {
				fmt.Println(err.Error())
				return "", nil, err
			} else {
				return order_uid, data, nil
			}
		}
	}
	return "", nil, err
}

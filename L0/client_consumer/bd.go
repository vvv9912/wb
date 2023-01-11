package main

import (
	modelmessage "client_consumer/message"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/patrickmn/go-cache"
)

// go get github.com/lib/pq
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

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
	"time"
)

func main() {
	connStr := "postgres://postgres:postgres@localhost/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	defer db.Close()
	//var message modelmessage.ModelMessage
	//plan, err := os.ReadFile("message/model.json")
	//checkErrDb(err, 1)
	//err = json.Unmarshal(plan, &message)
	//checkErrDb(err, 2)
	///////////

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
	yes := 0
	no := 0

	sub, err := sc.Subscribe("foo", func(m *stan.Msg) {
		rowsAffected, err := addDB(m.Data)
		time.Sleep(2 + time.Second)
		if err != nil {
			fmt.Printf("\nerr to bd: \n", err.Error())
			no++
		} else if err == nil {

			fmt.Printf("add: %d \n", rowsAffected)
			yes++
		}

		//fmt.Printf("Received a message: %s\n", string(m.Data))
	})
	//
	//stan.StartAtSequence(10) //start position
	if err != nil {
		fmt.Printf("error in subscribe: %v", err)
		return
	}
	fmt.Printf("yes:%d", yes)
	fmt.Printf("no:%d", no)
	fmt.Scanf(" ")
	sub.Unsubscribe()
	fmt.Printf("yes:%d", yes)
	fmt.Printf("no:%d", no)
}

func addDB(messagetojson []byte) (int64, error) {
	var message modelmessage.ModelMessage
	err := json.Unmarshal(messagetojson, &message)
	if err != nil {
		fmt.Println("Error: Неправильное сообщение (Неудалось распарсить) \n ", err.Error())
		return 0, err
	}
	connStr := "postgres://postgres:postgres@localhost/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	defer db.Close()
	if err != nil {
		fmt.Println("Error 1:", err.Error())
		return 0, err
	}

	//order_uid text,
	//	data JSONB
	err = db.Ping()
	if err != nil {
		fmt.Println("Error 2:", err.Error())
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
			err2 := errors.New("Error: такой OrderUid существует:")
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

//
//func get(messagetojson []byte) (int64, error) {
//	connStr := "postgres://postgres:postgres@localhost/postgres?sslmode=disable"
//	db, err := sql.Open("postgres", connStr)
//	defer db.Close()
//	if err != nil {
//		fmt.Println("Error 1:", err.Error())
//		return 0, err
//	}
//// получаем сколько записано => добавляем до того кол-ва
//	rows, err := db.Query("select order_uid from message")
//	if err != nil {
//		fmt.Println("Error 3:", err.Error())
//		return 0, err
//	}
//	defer rows.Close()
//	for rows.Next() {
//		var order_uid string
//		err := rows.Scan(&order_uid)
//		if err != nil {
//			fmt.Println("Error 4: ", err.Error())
//			return 0, err
//		}
//		if order_uid == message.OrderUid {
//			//fmt.Println("Error: такой OrderUid существует\n")
//			err2 := errors.New("Error: такой OrderUid существует:")
//			return 0, err2
//		}
//	}
//
//	result, err := db.Exec("insert into message (order_uid, data) values ($1,$2)", message.OrderUid, messagetojson)
//	if err != nil {
//		fmt.Println("Error 5:", err.Error())
//		return 0, err
//	}
//
//	rowsAffected, err := result.RowsAffected()
//	if err != nil {
//		fmt.Println("Error 6:", err.Error())
//		return 0, err
//	}
//	//fmt.Println(rowsAffected)
//	return rowsAffected, nil
//}

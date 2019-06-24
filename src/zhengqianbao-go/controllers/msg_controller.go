package controllers

import (
	"database/sql"
	"fmt"

	"../global"
	"../models"
	_ "github.com/lib/pq"
)

type Message_Interface interface {
	SelectMessage(messageID string) (message *models.Message, ok bool)

	GetCount(userID string) (count int, ok bool)

	GetAllMessage(userID string) (messages []models.Message, ok bool)

	InsertMessage(message *models.Message) (ok bool)

	ReadMessage(messageID string, userID string, state int) (ok bool)

	DeleteMessage(messageID string, userID string) (ok bool)

	GetUnReadMessage(userID string) (messages []models.Message, ok bool)

	GetReadMessage(userID string) (messages []models.Message, ok bool)
}

func (r *DBRepository) GetAllMessage(userID string) (messages []models.Message, ok bool) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		global.Host, global.Port, global.User, global.Password, global.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	rows, err := db.Query("select * from t_msg where receiver=$1", userID)

	for rows.Next() {
		messageObj := models.Message{}
		err = rows.Scan(&messageObj.MsgID, &messageObj.State, &messageObj.Receiver, &messageObj.Time,
			&messageObj.Title, &messageObj.Content)
		if err != nil {
			fmt.Printf("could not find message, %v", err)
			db.Close()
			return nil, false
		}
		messages = append(messages, messageObj)
	}

	db.Close()
	return messages, true
}

func (r *DBRepository) GetReadMessage(userID string) (messages []models.Message, ok bool) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		global.Host, global.Port, global.User, global.Password, global.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	rows, err := db.Query("select * from t_msg where receiver=$1", userID)

	for rows.Next() {
		messageObj := models.Message{}
		err = rows.Scan(&messageObj.MsgID, &messageObj.State, &messageObj.Receiver, &messageObj.Time,
			&messageObj.Title, &messageObj.Content)
		if err != nil {
			fmt.Printf("could not find message, %v", err)
			db.Close()
			return nil, false
		}
		if messageObj.State == 1 {
			messages = append(messages, messageObj)
		}
	}

	db.Close()
	return messages, true
}

func (r *DBRepository) GetUnReadMessage(userID string) (messages []models.Message, ok bool) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		global.Host, global.Port, global.User, global.Password, global.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	rows, err := db.Query("select * from t_msg where receiver=$1", userID)

	for rows.Next() {
		messageObj := models.Message{}
		err = rows.Scan(&messageObj.MsgID, &messageObj.State, &messageObj.Receiver, &messageObj.Time,
			&messageObj.Title, &messageObj.Content)
		if err != nil {
			fmt.Printf("could not find message, %v", err)
			db.Close()
			return nil, false
		}
		if messageObj.State == 0 {
			messages = append(messages, messageObj)
		}
	}

	db.Close()
	return messages, true
}

func (r *DBRepository) SelectMessage(messageID string) (message *models.Message, ok bool) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		global.Host, global.Port, global.User, global.Password, global.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	rows, err := db.Query("select * from t_msg where msgID=$1", messageID)

	messageObj := models.Message{}

	rows.Next()
	err = rows.Scan(&messageObj.MsgID, &messageObj.State, &messageObj.Receiver, &messageObj.Time,
		&messageObj.Title, &messageObj.Content)

	if err != nil {
		fmt.Printf("could not find message, %v", err)
		db.Close()
		return nil, false
	}

	db.Close()
	return &messageObj, true
}

func (r *DBRepository) GetCount(userID string) (count int64, ok bool) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		global.Host, global.Port, global.User, global.Password, global.Dbname)
	db, err := sql.Open("postgres", psqlInfo)

	rows, err := db.Query("select * from t_msg where receiver=$1", userID)
	count = 0
	for rows.Next() {
		messageObj := models.Message{}
		err = rows.Scan(&messageObj.MsgID, &messageObj.State, &messageObj.Receiver, &messageObj.Time,
			&messageObj.Title, &messageObj.Content)
		if err != nil {
			fmt.Printf("could not find message, %v", err)
			db.Close()
			return 0, false
		}
		if messageObj.State == 0 {
			count += 1
		}
	}

	db.Close()
	return count, true
}

func (r *DBRepository) InsertMessage(message *models.Message) (ok bool) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		global.Host, global.Port, global.User, global.Password, global.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	ok = true

	if err != nil {
		fmt.Printf("could not InsertQFormat, %v", err)
		ok = false
	}

	stmt, err := db.Prepare("insert into t_msg (msgID, state, receiver, time, title, content) values($1, $2, $3, $4, $5, $6)")
	if err != nil {
		fmt.Printf("could not Insert Mesage, %v", err)
		ok = false
	}

	_, err = stmt.Exec(message.MsgID, message.State, message.Receiver, message.Time, message.Title, message.Content)

	if err != nil {
		fmt.Printf("could not insert mesage, %v", err)
		ok = false
	}

	db.Close()

	return ok
}

func (r *DBRepository) ReadMessage(messageID string, userID string, state int) (ok bool) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		global.Host, global.Port, global.User, global.Password, global.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	ok = true

	stmt, err := db.Prepare("update t_msg set state=$1 WHERE msgID=$2 and receiver=$3")
	if err != nil {
		fmt.Println("fail to update:%v", err)
		ok = false
	}
	_, err = stmt.Exec(state, messageID, userID)
	if err != nil {
		fmt.Println(err)
		ok = false
	}
	db.Close()
	return ok
}

func (r *DBRepository) DeleteMessage(messageID string, userID string) (ok bool) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		global.Host, global.Port, global.User, global.Password, global.Dbname)
	db, err := sql.Open("postgres", psqlInfo)

	ok = true
	stmt, err := db.Prepare("delete from t_msg where msgID=$1 and receiver=$2")
	if err != nil {
		fmt.Println("delete error: %v", err)
		ok = false
	}
	_, err = stmt.Exec(messageID, userID)
	if err != nil {
		ok = false
		fmt.Println("delete error: %v", err)
	}

	db.Close()
	return ok
}

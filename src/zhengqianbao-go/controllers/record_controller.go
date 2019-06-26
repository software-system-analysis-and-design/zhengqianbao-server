package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"../global"
	"../models"
	_ "github.com/lib/pq"
)

type Record_Interface interface {
	// 查询记录
	QueryRecord(taskID string, userID string) (ok bool)

	// 获取记录
	SelectRecord(taskID string, userID string) (record *models.Record, ok bool)

	// 新建记录
	InsertRecord(record *models.Record) (ok bool)

	// 更新记录
	UpdateRecord(taskID string, userID string, record *models.Record) (ok bool)

	// 删除记录
	DeleteRecord(taskID string, userID string) (ok bool)

	// 获取所有记录
	SelectAllRecords(taskID string) (records []models.Record, ok bool)
}

func (r *DBRepository) SelectRecord(taskID string, userID string) (record *models.Record, ok bool) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		global.Host, global.Port, global.User, global.Password, global.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	rows, err := db.Query("select * from t_record where taskId=$1 and userID=$2", taskID, userID)

	recordObj := models.Record{}
	var chooseData string

	rows.Next()
	err = rows.Scan(&recordObj.TaskID, &recordObj.UserID, &chooseData)

	if err != nil {
		fmt.Printf("could not find record, %v", err)
		db.Close()
		return nil, false
	}

	err = json.Unmarshal([]byte(chooseData), &recordObj.Data)

	if err != nil {
		fmt.Printf("could not Unmarshal json, %v", err)
		db.Close()
		return nil, false
	}

	db.Close()
	return &recordObj, true
}

func (r *DBRepository) SelectAllRecords(taskID string) (records []models.Record, ok bool) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		global.Host, global.Port, global.User, global.Password, global.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	rows, err := db.Query("select * from t_record where taskId=$1", taskID)

	for rows.Next() {
		recordObj := models.Record{}
		var chooseData string
		err = rows.Scan(&recordObj.TaskID, &recordObj.UserID, &chooseData)
		if err != nil {
			fmt.Printf("could not find record, %v", err)
			db.Close()
			return nil, false
		}
		err = json.Unmarshal([]byte(chooseData), &recordObj.Data)
		if err != nil {
			fmt.Printf("could not Unmarshal json, %v", err)
			db.Close()
			return nil, false
		}
		records = append(records, recordObj)
	}

	db.Close()
	return records, true
}

func (r *DBRepository) InsertRecord(record *models.Record) (ok bool) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		global.Host, global.Port, global.User, global.Password, global.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	ok = true

	if err != nil {
		fmt.Printf("could not InsertQFormat, %v", err)
		ok = false
	}

	stmt, err := db.Prepare("insert into t_record (taskID, userID, data) values($1, $2, $3)")
	if err != nil {
		fmt.Printf("could not InsertQFormat, %v", err)
		ok = false
	}

	jsons, _ := json.Marshal(record.Data)

	_, err = stmt.Exec(record.TaskID, record.UserID, string(jsons))

	if err != nil {
		fmt.Printf("could not insert qFormat, %v", err)
		ok = false
	}

	db.Close()

	return ok
}

func (r *DBRepository) UpdateRecord(taskID string, userID string, record *models.Record) (ok bool) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		global.Host, global.Port, global.User, global.Password, global.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	ok = true

	stmt, err := db.Prepare("update t_record set data=$1 WHERE taskID=$2 and userID=$3")
	if err != nil {
		ok = false
	}

	jsons, _ := json.Marshal(record.Data)

	_, err = stmt.Exec(string(jsons), record.TaskID, record.UserID)
	if err != nil {
		fmt.Println(err)
		ok = false
	}
	db.Close()
	return ok
}

func (r *DBRepository) DeleteRecord(taskID string, userID string) (ok bool) {
	ok = r.QueryRecord(taskID, userID)
	if ok == false {
		return ok
	}
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		global.Host, global.Port, global.User, global.Password, global.Dbname)
	db, err := sql.Open("postgres", psqlInfo)

	stmt, err := db.Prepare("delete from t_record where taskid=$1 and userID=$2")
	if err != nil {
		ok = false
	}
	_, err = stmt.Exec(taskID, userID)
	if err != nil {
		ok = false
	}

	db.Close()
	return ok
}

func (r *DBRepository) QueryRecord(taskID string, userID string) (ok bool) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		global.Host, global.Port, global.User, global.Password, global.Dbname)
	db, err := sql.Open("postgres", psqlInfo)

	var count int64
	err = db.QueryRow("select count(*) from t_record where taskid=$1 and userid=$2", taskID, userID).Scan(&count)

	ok = true
	if err != nil || count == 0 {
		fmt.Printf("could not query, %v", err)
		ok = false
	}

	db.Close()
	return ok
}

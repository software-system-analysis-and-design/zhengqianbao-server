package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"../global"
	"../models"
	_ "github.com/lib/pq"
)

type QFormat_Interface interface {
	QueryQFormat(id string) (ok bool)

	SelectQFormat(id string) (qFormat *models.QuestionnaireFormat, ok bool)

	InsertQFormat(qFormat *models.QuestionnaireFormat) (ok bool, id string)

	UpdateQFormat(id string, qFormat *models.QuestionnaireFormat) (updatedQFormat *models.QuestionnaireFormat, ok bool)

	TrashQFormat(id string, inTrash int) (ok bool)

	DeleteQFormat(id string) (ok bool)

	SelectAllQFormats() (taskPreviews []models.TaskPreview, ok bool)
}

func (r *DBRepository) QueryQFormat(id string) (ok bool) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		global.Host, global.Port, global.User, global.Password, global.Dbname)
	db, err := sql.Open("postgres", psqlInfo)

	var count int64
	err = db.QueryRow("select count(*) from t_qformat where taskid=$1", id).Scan(&count)

	ok = true
	if err != nil || count == 0 {
		fmt.Printf("could not query, %v", err)
		ok = false
	}

	db.Close()
	return ok
}

func (r *DBRepository) SelectQFormat(id string) (qFormat *models.QuestionnaireFormat, ok bool) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		global.Host, global.Port, global.User, global.Password, global.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	rows, err := db.Query("select * from t_qformat where taskId=$1", id)

	qFormatObj := models.QuestionnaireFormat{}
	var chooseData string

	rows.Next()
	err = rows.Scan(&qFormatObj.TaskID, &qFormatObj.TaskName, &qFormatObj.InTrash, &qFormatObj.TaskType, &qFormatObj.Creator, &qFormatObj.Description,
		&qFormatObj.Money, &qFormatObj.Number, &qFormatObj.FinishedNumber, &qFormatObj.PublishTime, &qFormatObj.EndTime, &chooseData)

	if err != nil {
		fmt.Printf("could not find questionaire, %v", err)
		db.Close()
		return nil, false
	}

	err = json.Unmarshal([]byte(chooseData), &qFormatObj.ChooseData)

	if err != nil {
		fmt.Printf("could not Unmarshal json, %v", err)
		db.Close()
		return nil, false
	}

	db.Close()
	return &qFormatObj, true
}

func (r *DBRepository) InsertQFormat(qFormat *models.QuestionnaireFormat) (ok bool, id string) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		global.Host, global.Port, global.User, global.Password, global.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	ok = true

	if err != nil {
		fmt.Printf("could not InsertQFormat, %v", err)
		ok = false
	}

	stmt, err := db.Prepare("insert into t_qformat (taskID, taskName, inTrash, taskType, creator, description, money, number, finishedNumber, publishTime, " +
		"endTime, chooseData) values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)")
	if err != nil {
		fmt.Printf("could not InsertQFormat, %v", err)
		ok = false
	}

	jsons, _ := json.Marshal(qFormat.ChooseData)

	_, err = stmt.Exec(qFormat.TaskID, qFormat.TaskName, qFormat.InTrash, qFormat.TaskType, qFormat.Creator, qFormat.Description, qFormat.Money,
		qFormat.Number, qFormat.FinishedNumber, qFormat.PublishTime, qFormat.EndTime, string(jsons))

	if err != nil {
		fmt.Printf("could not insert qFormat, %v", err)
		ok = false
	}

	db.Close()

	if ok == false {
		return ok, ""
	} else {
		return ok, qFormat.TaskID
	}

}

func (r *DBRepository) UpdateQFormat(id string, qFormat *models.QuestionnaireFormat) (updatedQFormat *models.QuestionnaireFormat, ok bool) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		global.Host, global.Port, global.User, global.Password, global.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	ok = true

	stmt, err := db.Prepare("update t_qformat set taskName=$1, taskType=$2, creator=$3, description=$4, money=$5, " +
		"number=$6, finishedNumber=$7, publishTime=$8, endTime=$9, ChooseData=$10 WHERE taskID=$11")
	if err != nil {
		ok = false
	}

	jsons, _ := json.Marshal(qFormat.ChooseData)

	_, err = stmt.Exec(qFormat.TaskName, qFormat.TaskType, qFormat.Creator, qFormat.Description, qFormat.Money,
		qFormat.Number, qFormat.FinishedNumber, qFormat.PublishTime, qFormat.EndTime, string(jsons), qFormat.TaskID)
	if err != nil {
		ok = false
	}
	db.Close()
	qFormat.TaskID = id
	return qFormat, ok
}

func (r *DBRepository) DeleteQFormat(id string) (ok bool) {
	ok = r.QueryQFormat(id)
	if ok == false {
		return ok
	}
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		global.Host, global.Port, global.User, global.Password, global.Dbname)
	db, err := sql.Open("postgres", psqlInfo)

	stmt, err := db.Prepare("delete from t_qformat where taskid=$1")
	if err != nil {
		ok = false
	}
	_, err = stmt.Exec(id)
	if err != nil {
		ok = false
	}

	db.Close()
	return ok
}

func (r *DBRepository) SelectAllQFormats() (taskPreviews []models.TaskPreview, ok bool) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		global.Host, global.Port, global.User, global.Password, global.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Printf("could not insert qFormat, %v", err)
		db.Close()
		return nil, false
	}

	rows, err := db.Query("select * from t_qformat")

	if err != nil {
		fmt.Printf("could not select qFormat, %v", err)
		db.Close()
		return nil, false
	}

	for rows.Next() {
		var tPreviewObj = models.TaskPreview{}
		var chooseData string
		var state int
		err = rows.Scan(&tPreviewObj.TaskID, &tPreviewObj.TaskName, &tPreviewObj.InTrash, &tPreviewObj.TaskType, &tPreviewObj.Creator, &tPreviewObj.Description,
			&tPreviewObj.Money, &tPreviewObj.Number, &tPreviewObj.FinishedNumber, &tPreviewObj.PublishTime, &tPreviewObj.EndTime, &chooseData)
		if err != nil {
			fmt.Printf("could not scan rows, %v", err)
			db.Close()
			return nil, false
		}
		if state == 0 {
			taskPreviews = append(taskPreviews, tPreviewObj)
		}

	}

	db.Close()
	return taskPreviews, true
}

func (r *DBRepository) TrashQFormat(id string, inTrash int) (ok bool) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		global.Host, global.Port, global.User, global.Password, global.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	ok = true

	stmt, err := db.Prepare("update t_qformat set inTrash=$1 WHERE taskID=$2")
	if err != nil {
		ok = false
	}

	_, err = stmt.Exec(inTrash, id)
	if err != nil {
		ok = false
	}
	db.Close()
	return ok
}

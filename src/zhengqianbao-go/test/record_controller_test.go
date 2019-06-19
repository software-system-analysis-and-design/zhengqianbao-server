package test

import (
	"encoding/json"
	"fmt"
	"testing"

	"../controllers"
	"../models"
)

func TestInsertRecord(t *testing.T) {
	dbInstance := controllers.GetDBInstance()

	record := models.Record{TaskID: "20190611qiuxy23", UserID: "qiuxy23"}

	var chooseData []models.Answer

	answer1 := models.Answer{ID: 1, Type: 0, Content: "简答题的答案"}
	answer2 := models.Answer{ID: 2, Type: 1, Content: "A"}
	answer3 := models.Answer{ID: 3, Type: 2, Content: "A&B&C"}
	chooseData = append(chooseData, answer1)
	chooseData = append(chooseData, answer2)
	chooseData = append(chooseData, answer3)
	record.Data = chooseData

	ok := dbInstance.InsertRecord(&record)

	if !ok {
		t.Errorf("Test Insert FAIL")
	}

}

func TestQueryRecord(t *testing.T) {
	dbInstance := controllers.GetDBInstance()
	ok1 := dbInstance.QueryRecord("20190611qiuxy23", "qiuxy23")
	if !ok1 {
		t.Errorf("Test Insert FAIL")
	}

	ok2 := dbInstance.QueryRecord("20190611qiuxy22", "qiuxy23")
	if ok2 {
		t.Errorf("Test Insert FAIL")
	}
}

func TestSelectRecord(t *testing.T) {
	dbInstance := controllers.GetDBInstance()

	recordObj1, ok1 := dbInstance.SelectRecord("20190611qiuxy23", "qiuxy23")

	if ok1 {
		jsons, _ := json.Marshal(recordObj1)
		fmt.Println(string(jsons))
	} else {
		t.Errorf("Test Insert FAIL")
	}

	recordObj2, ok2 := dbInstance.SelectRecord("20190611qiuxy22", "qiuxy23")
	if ok2 {
		t.Errorf("Test Insert FAIL")
		fmt.Println(recordObj2)
	}
}

func TestDeleteRecord(t *testing.T) {
	dbInstance := controllers.GetDBInstance()
	ok1 := dbInstance.DeleteRecord("20190611qiuxy23", "qiuxy23")
	if !ok1 {
		t.Errorf("Test Insert FAIL")
	}

	ok2 := dbInstance.QueryRecord("20190611qiuxy22", "qiuxy23")
	if ok2 {
		t.Errorf("Test Insert FAIL")
	}
}

func TestUpdateRecord(t *testing.T) {
	dbInstance := controllers.GetDBInstance()
	record, _ := dbInstance.SelectRecord("20190611qiuxy23", "qiuxy23")
	answer4 := models.Answer{ID: 4, Type: 0, Content: "简答题4的答案"}
	record.Data = append(record.Data, answer4)

	ok := dbInstance.UpdateRecord("20190611qiuxy23", "qiuxy23", record)

	if !ok {
		t.Errorf("Test Insert FAIL")
	}

	updatedRecords, ok := dbInstance.SelectRecord("20190611qiuxy23", "qiuxy23")

	if !ok {
		t.Errorf("Test Insert FAIL")
	} else {
		fmt.Println(updatedRecords)
	}
}

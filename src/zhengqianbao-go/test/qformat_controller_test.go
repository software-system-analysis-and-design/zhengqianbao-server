package test

import (
	"encoding/json"
	"fmt"
	"testing"

	"../controllers"
	"../models"
)

func TestInsertQFormat(t *testing.T) {
	dbInstance := controllers.GetDBInstance()

	qFormat1 := models.QuestionnaireFormat{TaskName: "任务1", TaskID: "20190611qiuxy22", InTrash: 0, TaskType: "问卷", Creator: "qiuxy23",
		Description: "描述1", Money: 100, Number: 100, FinishedNumber: 50, PublishTime: "2019-06-12", EndTime: "2019-06-13"}

	var chooseData []models.Question

	question1 := models.Question{TitleNum: 1, ID: 1, Title: "question1", Required: true}
	var datacontent1 []models.Option
	question1.DataContent = datacontent1
	chooseData = append(chooseData, question1)

	question2 := models.Question{TitleNum: 2, ID: 2, Title: "question2", Required: true}
	var datacontent2 []models.Option
	option2 := models.Option{ID: 1, Content: "option2"}
	datacontent2 = append(datacontent2, option2)
	question2.DataContent = datacontent2
	chooseData = append(chooseData, question2)

	question3 := models.Question{TitleNum: 3, ID: 3, Title: "question3", Required: false}
	var datacontent3 []models.Option
	option3_1 := models.Option{ID: 1, Content: "option3_1"}
	option3_2 := models.Option{ID: 2, Content: "option3_2"}
	option3_3 := models.Option{ID: 3, Content: "option3_3"}
	datacontent3 = append(datacontent3, option3_1)
	datacontent3 = append(datacontent3, option3_2)
	datacontent3 = append(datacontent3, option3_3)
	question3.DataContent = datacontent3
	chooseData = append(chooseData, question3)

	qFormat1.ChooseData = chooseData

	ok, id := dbInstance.InsertQFormat(&qFormat1)

	if ok {
		fmt.Println("id: ", id)
	} else {
		t.Errorf("Test Insert FAIL")
	}

}

func TestQueryQFormat(t *testing.T) {
	dbInstance := controllers.GetDBInstance()
	ok1 := dbInstance.QueryQFormat("20190611qiuxy23")
	if !ok1 {
		t.Errorf("Test Insert FAIL")
	}

	ok2 := dbInstance.QueryQFormat("20190611qiuxy22")
	if ok2 {
		t.Errorf("Test Insert FAIL")
	}
}

func TestSelectQFormat(t *testing.T) {
	dbInstance := controllers.GetDBInstance()

	qFormatObj1, ok1 := dbInstance.SelectQFormat("20190611qiuxy23")

	if ok1 {
		jsons, _ := json.Marshal(qFormatObj1)
		fmt.Println(string(jsons))
	} else {
		t.Errorf("Test Insert FAIL")
	}

	qFormatObj2, ok2 := dbInstance.SelectQFormat("20190611qiuxy22")
	if ok2 {
		t.Errorf("Test Insert FAIL")
		fmt.Println(qFormatObj2)
	}
}

func TestDeleteQFormat(t *testing.T) {
	dbInstance := controllers.GetDBInstance()
	ok1 := dbInstance.DeleteQFormat("20190611qiuxy23")
	if !ok1 {
		t.Errorf("Test Insert FAIL")
	}

	ok2 := dbInstance.QueryQFormat("20190611qiuxy22")
	if ok2 {
		t.Errorf("Test Insert FAIL")
	}
}

func TestSelectAllQFormats(t *testing.T) {
	dbInstance := controllers.GetDBInstance()
	allQFormats, ok := dbInstance.SelectAllQFormats()
	if ok {
		jsons, _ := json.Marshal(allQFormats)
		fmt.Println(string(jsons))
	} else {
		t.Errorf("Test Insert FAIL")
	}

}

func TestUpdateQFormats(t *testing.T) {
	dbInstance := controllers.GetDBInstance()
	qFormats, _ := dbInstance.SelectQFormat("20190611qiuxy23")
	qFormats.TaskName = "task1"

	_, ok := dbInstance.UpdateQFormat(qFormats.TaskID, qFormats)

	if !ok {
		t.Errorf("Test Insert FAIL")
	}

	updatedQFormats, ok := dbInstance.SelectQFormat("20190611qiuxy23")

	if !ok || updatedQFormats.TaskName != "task1" {
		t.Errorf("Test Insert FAIL")
	}

}

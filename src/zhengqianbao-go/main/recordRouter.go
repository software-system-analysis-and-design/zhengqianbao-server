package main

import (
	"encoding/json"
	"fmt"

	"../controllers"
	"../helper"
	"../models"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
)

func RecordCreateHandler(ctx iris.Context) {
	userMsg := ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)

	phone := userMsg["phone"].(string)

	data := ctx.FormValue("data")
	var recordObj models.Record

	err := json.Unmarshal([]byte(data), &recordObj)
	if err != nil {
		response := helper.Gene_Response{
			Code: 400,
			Msg:  "json字符串解析错误！",
		}
		ctx.JSON(response)
		return
	}
	if phone != recordObj.UserID {
		response := helper.Gene_Response{
			Code: 400,
			Msg:  "创建者的ID和token所对应的用户ID不一致！",
		}
		ctx.JSON(response)
		return
	}

	dbInstance := controllers.GetDBInstance()
	ok := dbInstance.QueryQFormat(recordObj.TaskID)

	if !ok {
		response := helper.Gene_Response{
			Code: 400,
			Msg:  "该问卷不存在！",
		}
		ctx.JSON(response)
		return
	}

	ok = dbInstance.InsertRecord(&recordObj)

	if !ok {
		response := helper.Gene_Response{
			Code: 400,
			Msg:  "插入失败！",
		}
		ctx.JSON(response)
		return
	} else {
		dbInstance.AddOneQFormats(recordObj.TaskID)
		response := helper.Gene_Response{
			Code: 200,
			Msg:  "插入成功！",
		}
		ctx.JSON(response)
		return
	}
}

func RecordSelectHandler(ctx iris.Context) {
	userMsg := ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)
	phoneID := userMsg["phone"].(string)
	taskID := ctx.FormValue("taskID")
	dbInstance := controllers.GetDBInstance()
	recordObj, ok := dbInstance.SelectRecord(taskID, phoneID)
	if ok {
		ctx.JSON(recordObj)
		return
	} else {
		response := helper.Gene_Response{
			Code: 400,
			Msg:  "获取任务失败！",
		}
		ctx.JSON(response)
		return
	}
}

func RecordGetAllHandler(ctx iris.Context) {
	userMsg := ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)
	phoneID := userMsg["phone"].(string)
	taskID := ctx.FormValue("taskID")
	dbInstance := controllers.GetDBInstance()

	qFormatObj, ok := dbInstance.SelectQFormat(taskID)

	if !ok || qFormatObj.Creator != phoneID {
		fmt.Println(qFormatObj.Creator)
		fmt.Println(phoneID)
		response := helper.Gene_Response{
			Code: 400,
			Msg:  "该问卷不存在！",
		}
		ctx.JSON(response)
		return
	}

	recordObj, ok := dbInstance.SelectAllRecords(taskID)
	if ok {
		ctx.JSON(recordObj)
		return
	} else {
		response := helper.Gene_Response{
			Code: 400,
			Msg:  "获取任务失败！",
		}
		ctx.JSON(response)
		return
	}
}

func RecordUpdateHandler(ctx iris.Context) {
	userMsg := ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)
	phoneID := userMsg["phone"].(string)

	data := ctx.FormValue("data")
	var recordObj models.Record

	err := json.Unmarshal([]byte(data), &recordObj)
	if err != nil {
		response := helper.Gene_Response{
			Code: 400,
			Msg:  "json字符串解析错误！",
		}
		ctx.JSON(response)
		return
	}
	if phoneID != recordObj.UserID {
		response := helper.Gene_Response{
			Code: 400,
			Msg:  "创建者的ID和token所对应的用户ID不一致！",
		}
		ctx.JSON(response)
		return
	}

	dbInstance := controllers.GetDBInstance()

	ok := dbInstance.QueryRecord(recordObj.TaskID, recordObj.UserID)
	if !ok {
		response := helper.Gene_Response{
			Code: 400,
			Msg:  "更新失败，该记录不存在！",
		}
		ctx.JSON(response)
		return
	}

	ok = dbInstance.UpdateRecord(recordObj.TaskID, recordObj.UserID, &recordObj)

	if !ok {
		response := helper.Gene_Response{
			Code: 400,
			Msg:  "更新失败！",
		}
		ctx.JSON(response)
		return
	} else {
		response := helper.Gene_Response{
			Code: 200,
			Msg:  "更新成功！",
		}
		ctx.JSON(response)
		return
	}
}

func RecordDeleteHandler(ctx iris.Context) {
	userMsg := ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)
	phone := userMsg["phone"].(string)

	dbInstance := controllers.GetDBInstance()
	taskID := ctx.FormValue("taskID")
	ok := dbInstance.DeleteRecord(taskID, phone)

	if !ok {
		response := helper.Gene_Response{
			Code: 400,
			Msg:  "删除失败，该记录不存在！",
		}
		ctx.JSON(response)
		return
	} else {
		response := helper.Gene_Response{
			Code: 200,
			Msg:  "删除成功！",
		}
		ctx.JSON(response)
		return
	}
}

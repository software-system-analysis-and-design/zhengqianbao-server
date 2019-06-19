package main

import (
	"encoding/json"
	"strconv"
	"time"

	"../controllers"
	"../helper"
	"../models"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
)

func QPreviewHandler(ctx iris.Context) {
	dbInstance := controllers.GetDBInstance()
	previews, _ := dbInstance.SelectAllQFormats()
	ctx.JSON(previews)

}

func QCreateHandler(ctx iris.Context) {

	userMsg := ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)

	phone := userMsg["phone"].(string)

	data := ctx.FormValue("data")
	var qFormatObj models.QuestionnaireFormat

	err := json.Unmarshal([]byte(data), &qFormatObj)
	if err != nil {
		response := helper.Gene_Response{
			Code: 400,
			Msg:  "json字符串解析错误！",
		}
		ctx.JSON(response)
		return
	}
	if phone != qFormatObj.Creator {
		response := helper.Gene_Response{
			Code: 400,
			Msg:  "创建者的ID和token所对应的用户ID不一致！",
		}
		ctx.JSON(response)
		return
	}

	qFormatObj.TaskID = strconv.FormatInt(time.Now().Unix(), 10) + phone
	dbInstance := controllers.GetDBInstance()
	ok, id := dbInstance.InsertQFormat(&qFormatObj)

	if !ok {
		response := helper.Gene_Response{
			Code: 400,
			Msg:  "插入失败！",
		}
		ctx.JSON(response)
		return
	} else {
		response_data := map[string]string{"id": id}
		response := helper.Data_Response{
			Code: 200,
			Msg:  "插入成功！",
			Data: response_data,
		}
		ctx.JSON(response)
		return
	}
}

func QTokenHandler(ctx iris.Context) {
	userMsg := ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)
	phone := userMsg["phone"].(string)

	data := ctx.FormValue("data")
	var qFormatObj models.QuestionnaireFormat

	err := json.Unmarshal([]byte(data), qFormatObj)
	if err != nil {
		response := helper.Gene_Response{
			Code: 400,
			Msg:  "json字符串解析错误！",
		}
		ctx.JSON(response)
		return
	}
	if phone != qFormatObj.Creator {
		response := helper.Gene_Response{
			Code: 400,
			Msg:  "创建者的ID和token所对应的用户ID不一致！",
		}
		ctx.JSON(response)
		return
	}
}

func QSelectHandler(ctx iris.Context) {

	id := ctx.FormValue("id")
	dbInstance := controllers.GetDBInstance()
	qFormatObj, ok := dbInstance.SelectQFormat(id)
	if ok {
		ctx.JSON(qFormatObj)
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

func QUpdateHandler(ctx iris.Context) {
	userMsg := ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)
	phone := userMsg["phone"].(string)

	data := ctx.FormValue("data")
	var qFormatObj models.QuestionnaireFormat

	err := json.Unmarshal([]byte(data), &qFormatObj)
	if err != nil {
		response := helper.Gene_Response{
			Code: 400,
			Msg:  "json字符串解析错误！",
		}
		ctx.JSON(response)
		return
	}
	if phone != qFormatObj.Creator {
		response := helper.Gene_Response{
			Code: 400,
			Msg:  "创建者的ID和token所对应的用户ID不一致！",
		}
		ctx.JSON(response)
		return
	}

	dbInstance := controllers.GetDBInstance()

	ok := dbInstance.QueryQFormat(qFormatObj.TaskID)
	if !ok {
		response := helper.Gene_Response{
			Code: 400,
			Msg:  "更新失败，该id不存在！",
		}
		ctx.JSON(response)
		return
	}

	updatedQFormat, ok := dbInstance.UpdateQFormat(qFormatObj.TaskID, &qFormatObj)

	if !ok {
		response := helper.Gene_Response{
			Code: 400,
			Msg:  "更新失败！",
		}
		ctx.JSON(response)
		return
	} else {
		response_data := map[string]string{"id": updatedQFormat.TaskID}
		response := helper.Data_Response{
			Code: 200,
			Msg:  "更新成功！",
			Data: response_data,
		}
		ctx.JSON(response)
		return
	}
}

func QDeleteHandler(ctx iris.Context) {
	userMsg := ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)
	phone := userMsg["phone"].(string)

	dbInstance := controllers.GetDBInstance()
	id := ctx.FormValue("id")
	qFormatObj, ok := dbInstance.SelectQFormat(id)

	if !ok {
		response := helper.Gene_Response{
			Code: 400,
			Msg:  "删除失败，id不存在！",
		}
		ctx.JSON(response)
		return
	}
	if phone != qFormatObj.Creator {
		response := helper.Gene_Response{
			Code: 400,
			Msg:  "创建者的ID和token所对应的用户ID不一致！",
		}
		ctx.JSON(response)
		return
	}

	ok = dbInstance.DeleteQFormat(id)
	if !ok {
		response := helper.Gene_Response{
			Code: 400,
			Msg:  "删除失败，id不存在！",
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

func QTrashHandler(ctx iris.Context) {
	userMsg := ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)
	phone := userMsg["phone"].(string)

	id := ctx.FormValue("id")
	inTrash, _ := strconv.Atoi(ctx.FormValue("inTrash"))
	dbInstance := controllers.GetDBInstance()

	qFormatObj, ok := dbInstance.SelectQFormat(id)

	if !ok {
		response := helper.Gene_Response{
			Code: 400,
			Msg:  "ID不存在，执行失败！",
		}
		ctx.JSON(response)
		return
	}

	if qFormatObj.Creator != phone {
		response := helper.Gene_Response{
			Code: 400,
			Msg:  "创建者的ID和token所对应的用户ID不一致！",
		}
		ctx.JSON(response)
		return
	}
	ok = dbInstance.TrashQFormat(id, inTrash)
	if !ok {
		response := helper.Gene_Response{
			Code: 400,
			Msg:  "更新失败！",
		}
		ctx.JSON(response)
		return
	}
	response := helper.Gene_Response{
		Code: 200,
		Msg:  "更新成功！",
	}
	ctx.JSON(response)
	return

}

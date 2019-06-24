package main

import (
	"strconv"

	"../controllers"
	"../helper"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
)

func GetMsgHandler(ctx iris.Context) {

	userMsg := ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)

	phone := userMsg["phone"].(string)

	msgID := ctx.FormValue("msgID")

	dbInstance := controllers.GetDBInstance()
	messageObj, ok := dbInstance.SelectMessage(msgID)

	if !ok {
		response := helper.Gene_Response{
			Code: 400,
			Msg:  "消息不存在！",
		}
		ctx.JSON(response)
		return
	}

	if messageObj.Receiver != phone {
		response := helper.Gene_Response{
			Code: 400,
			Msg:  "接收者的ID和token所对应的用户ID不一致！",
		}
		ctx.JSON(response)
		return
	}

	ctx.JSON(messageObj)
	return
}

func GetMsgCount(ctx iris.Context) {
	userMsg := ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)

	phone := userMsg["phone"].(string)

	dbInstance := controllers.GetDBInstance()
	count, ok := dbInstance.GetCount(phone)

	if !ok {
		response := helper.Gene_Response{
			Code: 400,
			Msg:  "获取失败",
		}
		ctx.JSON(response)
		return
	}

	response := helper.Gene_Response{
		Code: 200,
		Msg:  strconv.FormatInt(count, 10),
	}
	ctx.JSON(response)
	return

}

func ReadMessage(ctx iris.Context) {
	userMsg := ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)

	phone := userMsg["phone"].(string)

	msgID := ctx.FormValue("msgID")
	state, _ := strconv.Atoi(ctx.FormValue("state"))

	dbInstance := controllers.GetDBInstance()

	ok := dbInstance.ReadMessage(msgID, phone, state)

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

func DeleteMessage(ctx iris.Context) {
	userMsg := ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)

	phone := userMsg["phone"].(string)

	msgID := ctx.FormValue("msgID")

	dbInstance := controllers.GetDBInstance()

	ok := dbInstance.DeleteMessage(msgID, phone)

	if !ok {
		response := helper.Gene_Response{
			Code: 400,
			Msg:  "删除失败！",
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

func GetAllMessages(ctx iris.Context) {
	userMsg := ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)

	phone := userMsg["phone"].(string)

	dbInstance := controllers.GetDBInstance()
	messages, ok := dbInstance.GetAllMessage(phone)

	if !ok {
		response := helper.Gene_Response{
			Code: 400,
			Msg:  "查询失败！",
		}
		ctx.JSON(response)
		return
	}

	ctx.JSON(messages)
	return

}

func GetReadMessages(ctx iris.Context) {
	userMsg := ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)

	phone := userMsg["phone"].(string)

	dbInstance := controllers.GetDBInstance()
	messages, ok := dbInstance.GetReadMessage(phone)

	if !ok {
		response := helper.Gene_Response{
			Code: 400,
			Msg:  "查询失败！",
		}
		ctx.JSON(response)
		return
	}

	ctx.JSON(messages)
	return

}

func GetUnReadMessages(ctx iris.Context) {
	userMsg := ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)

	phone := userMsg["phone"].(string)

	dbInstance := controllers.GetDBInstance()
	messages, ok := dbInstance.GetUnReadMessage(phone)

	if !ok {
		response := helper.Gene_Response{
			Code: 400,
			Msg:  "查询失败！",
		}
		ctx.JSON(response)
		return
	}

	ctx.JSON(messages)
	return

}

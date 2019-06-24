package main

import (
	"fmt"
	"strconv"
	"time"

	"../controllers"
	"../models"
)

func CreateMsgBackend() {

	/*
		Intrash:
		0 正常
		1 回收站
		2 待发布
		3 已到期
		4 已完成
	*/

	for {
		time.Sleep(time.Millisecond * 1000)
		dbInstance := controllers.GetDBInstance()
		previews, _ := dbInstance.SelectAllQFormats()

		for _, preview := range previews {

			if preview.InTrash != 0 {
				continue
			}

			if preview.Number == preview.FinishedNumber {
				CreateMsgFinish(preview)
				dbInstance.TrashQFormat(preview.TaskID, 4)
				continue
			}

			timeTemplate := "2006-01-02T15:04"
			stamp, err := time.ParseInLocation(timeTemplate, preview.EndTime, time.Local)
			if err == nil {
				if stamp.Unix() < time.Now().Unix() {
					creator := preview.Creator
					money, _ := strconv.Atoi(preview.Money)
					remain := (preview.Number - preview.FinishedNumber) * money
					userObj, _ := dbInstance.SelectUser(creator)
					userObj.Remain += remain
					dbInstance.UpdateMoney(creator, userObj.Remain)

					CreateMsgEndTime(preview, userObj, remain)
					dbInstance.TrashQFormat(preview.TaskID, 3)
					continue
				}
			}
		}

	}
}

func CreateMsgEndTime(preview models.TaskPreview, userObj *models.User, remain int) {
	msgID := strconv.FormatInt(time.Now().Unix(), 10) + userObj.Phone

	timeTemplate := "2006-01-02T15:04"
	timestamp := time.Now().Unix()
	t_str := (time.Unix(timestamp, 0).Format(timeTemplate))
	content := "您创建的问卷（编号为 " + preview.TaskID + ", 问卷名为 " + preview.TaskName + "）已过设置的结束时间，剩下的预付金额（" +
		strconv.Itoa(remain) + " tokens）已退还到原账号，请查收！"
	msg := models.Message{MsgID: msgID, State: 0, Receiver: userObj.Phone, Time: t_str, Title: "任务到期", Content: content}
	dbInstance := controllers.GetDBInstance()
	ok := dbInstance.InsertMessage(&msg)
	if !ok {
		fmt.Println("消息新建失败")
	}
}

func CreateMsgFinish(preview models.TaskPreview) {
	msgID := strconv.FormatInt(time.Now().Unix(), 10) + preview.Creator

	timeTemplate := "2006-01-02T15:04"
	timestamp := time.Now().Unix()
	t_str := (time.Unix(timestamp, 0).Format(timeTemplate))
	content := "您创建的问卷（编号为 " + preview.TaskID + ", 问卷名为 " + preview.TaskName + "）已完成！"
	msg := models.Message{MsgID: msgID, State: 0, Receiver: preview.Creator, Time: t_str, Title: "任务完成", Content: content}
	dbInstance := controllers.GetDBInstance()
	ok := dbInstance.InsertMessage(&msg)
	if !ok {
		fmt.Println("消息新建失败")
	}
}

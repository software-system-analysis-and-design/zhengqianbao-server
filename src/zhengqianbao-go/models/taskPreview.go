package models

// TaskPreview is our sample data structure.
// which could wrap by embedding the models.TaskPreview or
// declare new fields instead butwe will use this models
// as the only one TaskPreview model in our application,
// for the shake of simplicty.

type TaskPreview struct {
	TaskName       string `json:"taskName"`
	TaskID         string `json:"taskID"`
	InTrash        int    `json:"inTrash"`
	TaskType       string `json:"taskType"`
	Creator        string `json:"creator"`
	Description    string `json:"description"`
	Money          string `json:"money"`
	Number         int    `json:"number"`
	FinishedNumber int    `json:"finishedNumber"`
	PublishTime    string `json:"publishTime"`
	EndTime        string `json:"endTime"`
}

/*
任务预览字段：
{
  ["task1" :{
    taskName: "任务名"
    taskID: "任务ID"
    taskType: "问卷"     // 还有委托任务，暂时先做问卷
	number: 100         // 任务要求个数
	money: 100, // 问卷字段
    finishedNumber: 30 // 已完成的个数
    publishTime: ""    // 任务发布时间 yyyy-mm-dd-hh-mm
    endTime: ""       // 任务截止时间
  },
  ]
}
*/

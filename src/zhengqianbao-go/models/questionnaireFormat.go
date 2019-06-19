package models

// QuestionnaireFormat is our sample data structure.
// which could wrap by embedding the models.QuestionnaireFormat or
// declare new fields instead butwe will use this models
// as the only one QuestionnaireFormat model in our application,
// for the shake of simplicty.
type QuestionnaireFormat struct {
	TaskName       string     `json:"taskName"`
	TaskID         string     `json:"taskID"`
	InTrash        int        `json:"inTrash"`
	TaskType       string     `json:"taskType"`
	Creator        string     `json:"creator"`
	Description    string     `json:"description"`
	Money          int        `json:"money"`
	Number         int        `json:"number"`
	FinishedNumber int        `json:"finishedNumber"`
	PublishTime    string     `json:"publishTime"`
	EndTime        string     `json:"endTime"`
	ChooseData     []Question `json:"chooseData"`
}

/*
问卷字段设计：
state = {
  taskName: "问卷名"
  taskID: int
  taskType: "问卷"     // 还有委托任务，暂时先做问卷
  creator: phone
  description: "问卷描述"
  money: 100, // 问卷字段
  number: 100, //  问卷设置的份数, 默认 -1 则视为份数不限
  publishTime: "发布时间"  // 年月日时分，如果不设置，则需要手动发布
  endTime: "截止时间"    // 年月日时分， 如果不设置，则需要手动截止
*/

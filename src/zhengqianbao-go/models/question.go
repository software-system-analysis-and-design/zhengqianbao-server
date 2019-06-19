package models

// Question is our sample data structure.
// which could wrap by embedding the models.Question or
// declare new fields instead butwe will use this models
// as the only one Question model in our application,
// for the shake of simplicty.
type Question struct {
	TitleNum    int      `json:"titleNum"`
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	DataType    int      `json:"dataType"`
	Required    bool     `json:"required"`
	DataContent []Option `json:"dataContent"`
}

/*
	// 问答题的字段格式
	{
		titleNum: 1,   // 题目的编号
		id: 1           // 题目的 id, 在进行map渲染时，作为唯一的key标识
		title: "这是问答题的问题"
		dataType: 1,  （1 表示是问答题)
		required: 1    (1 表示是必选题目， 0非必选题)
		dataContent: []
	}

	// 单选题的字段格式
    {
      titleNum: 2
      id: 2
      title: "这是一道单选题"
      dataType: 2  (2表示是单选题)
      required: 0
      dataContent:[
		1: {
			id: 1, //  选项的id
			content: "选项的内容"
		},
		2: {
			id: 2,
			content: "2113"
		}
      ]
	}


	// 多选题的字段格式
    {
      titleNum: 3
      id: 3
      title: "多选题目"
      dataType: 3
      dataContent:[
        1: {
			id: 1,
			content: "asdasd"
		},
		2: {
			id: 2,
			content: "sdfsd"
		}
      ]
    }
*/

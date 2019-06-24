package models

// Message is our sample data structure.
// which could wrap by embedding the models.Message or
// declare new fields instead butwe will use this models
// as the only one Message model in our application,
// for the shake of simplicty.
type Message struct {
	MsgID    string `json:"msgID"`
	State    int    `json:"state"`
	Receiver string `json:"receiver"`
	Time     string `json:"time"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}

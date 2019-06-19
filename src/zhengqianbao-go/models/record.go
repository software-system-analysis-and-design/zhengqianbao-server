package models

// Answer is our sample data structure.
// which could wrap by embedding the models.Answer or
// declare new fields instead butwe will use this models
// as the only one Answer model in our application,
// for the shake of simplicty.
type Record struct {
	TaskID string   `json:"taskID"`
	UserID string   `json:"userID"`
	Data   []Answer `json:"data"`
}

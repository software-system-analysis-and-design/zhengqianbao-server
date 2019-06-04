package models

// User is our sample data structure.
// which could wrap by embedding the models.User or
// declare new fields instead butwe will use this models
// as the only one User model in our application,
// for the shake of simplicty.
type User struct {
	Phone       string `json:"phone"`
	Iscow       bool   `json:"iscow"`
	Name        string `json:"name"`
	Password    string `json:"password"`
	Gender      string `json:"gender"`
	Age         int    `json:"age"`
	University  string `json:"university"`
	Company     string `json:"company"`
	Description string `json:"description"`
	Class       int    `json:"class"`
}

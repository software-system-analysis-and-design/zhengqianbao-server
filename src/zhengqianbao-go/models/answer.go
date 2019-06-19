package models

// Answer is our sample data structure.
// which could wrap by embedding the models.Answer or
// declare new fields instead butwe will use this models
// as the only one Answer model in our application,
// for the shake of simplicty.
type Answer struct {
	ID      int    `json:"id"`
	Type    int    `json:"type"`
	Content string `json:"content"`
}

package models

// Option is our sample data structure.
// which could wrap by embedding the models.Option or
// declare new fields instead butwe will use this models
// as the only one Option model in our application,
// for the shake of simplicty.
type Option struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}

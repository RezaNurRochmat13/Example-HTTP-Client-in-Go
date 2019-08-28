package model

type Post []struct {
	UserID int    `json:"userId"`
	IDPost int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

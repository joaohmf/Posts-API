package models

type Post struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Date    string `json:"date"`
}

type RequestPost struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Date    string `json:"date"`
}

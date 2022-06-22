package models

type Todo struct {
	ID     int `json:"id"`
	Text   string `json:"text"`
	UserID int `json:"user_id"`
}

type NewTodo struct {
	Text string `json:"text"`
}

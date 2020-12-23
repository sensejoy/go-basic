package model

type User struct {
	UserId int64  `json:"user_id" db:"id"`
	Name   string `json:"name" db:"name"`
	Age    int    `json:"age"`
	Email  string `json:"email"`
}

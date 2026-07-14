package userhandler

import "time"

type User struct {
	ID uint `json : "id"`
	Name string `josn : "name"`
	Psd string `json : "psd"`
	CreatedAt time.Time `json : "created_at"`
	UpdatedAt time.Time `json : "updated_at"`
}

func (User) TableName() string {
	return "用户"
}
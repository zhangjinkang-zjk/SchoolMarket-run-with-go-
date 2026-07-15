package model

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey;autoincrement" json:"id"`
	Name      string    `gorm:"type:varchar(20);not null" json:"name"`
	Psd       string    `gorm:"type:varchar(30);not null" json:"psd"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}

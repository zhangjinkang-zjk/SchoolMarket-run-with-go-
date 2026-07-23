package model

import "time"

type Cart struct {
	ID        uint      `gorm:"primaryKey;autoincrement" json:"id"`
	UserID    uint      `gorm:"index" json:"user_id"`
	GoodsID   uint      `gorm:"index" json:"goods_id"`
	Num       uint      `json:"num"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	User  User  `gorm:"foreignKey:UserID" json:"-"`
	Goods Goods `gorm:"foreignKey:GoodsID" json:"-"`
}

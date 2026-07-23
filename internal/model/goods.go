package model

import "time"

type Goods struct {
	ID uint `gorm:"primarykey;autoincrement" json:"goodsid"`
	Name string `gorm:"type:varchar(50)" json:"goodsname"`
	Num uint `json:"goodsnum"`
	Price uint `json:"goodsprice"`
	UserID uint `gorm:"index" json:"user_id"`
	User User `gorm:"foreignKey:UserID" json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
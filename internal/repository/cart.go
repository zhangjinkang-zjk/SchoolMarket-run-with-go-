package repository

import (
	"SchoolMarket-run-with-go-/internal/model"
	"gorm.io/gorm"
)

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{db: db}
}

func (r *CartRepository) CreateCart(cart *model.Cart) (*model.Cart, error) {
	return cart, r.db.Create(cart).Error
}

func (r *CartRepository) FindByUserId(userID uint) ([]model.Cart, error) {
	var carts []model.Cart
	result := r.db.Where("user_id = ?", userID).Preload("Goods").Find(&carts)
	return carts, result.Error
}

func (r *CartRepository) FindAim(id uint) (*model.Cart, error) {
	var cart model.Cart
	result := r.db.First(&cart, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return &cart, nil
}

func (r *CartRepository) UpdateCart(cart *model.Cart) (*model.Cart, error) {
	return cart, r.db.Save(cart).Error
}

func (r *CartRepository) DeleteCart(id uint) (int64, error) {
	result := r.db.Delete(&model.Cart{}, id)
	return result.RowsAffected, result.Error
}

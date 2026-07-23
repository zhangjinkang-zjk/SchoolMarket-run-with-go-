package repository

import (
	"SchoolMarket-run-with-go-/internal/model"

	"gorm.io/gorm"
)

type GoodsRepository struct {
	db *gorm.DB
}

func NewGoodsRepository(db *gorm.DB) *GoodsRepository {
	return &GoodsRepository{db: db}
}

func (r *GoodsRepository) CreateGoods(goods *model.Goods) (*model.Goods, error) {
	return goods, r.db.Create(goods).Error
}

func (r *GoodsRepository) FindAim(id uint) (*model.Goods, error) {
	var goods model.Goods
	result := r.db.Preload("User").First(&goods, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return &goods, nil
}

func (r *GoodsRepository) FindAll() ([]model.Goods, error) {
	var goods []model.Goods
	result := r.db.Preload("User").Find(&goods)
	return goods, result.Error
}

func (r *GoodsRepository) UpdateGoods(goods *model.Goods) (*model.Goods, error) {
	return goods, r.db.Save(goods).Error
}

func (r *GoodsRepository) DeleteGoods(id uint) (int64, error) {
	result := r.db.Delete(&model.Goods{}, id)
	return result.RowsAffected, result.Error
}

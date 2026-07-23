package userservice

import (
	"errors"

	"SchoolMarket-run-with-go-/internal/cache"
	"SchoolMarket-run-with-go-/internal/model"
	"SchoolMarket-run-with-go-/internal/repository"
)

type CreateGoodsRequest struct {
	Name  string
	Num   uint
	Price uint
}

type UpdateGoodsRequest struct {
	ID      uint
	Name    *string
	Num     *uint
	Price   *uint
	UserID  uint
}

type GoodsService interface {
	CreateGoods(req CreateGoodsRequest, userID uint) (*model.Goods, error)
	GetGoodsById(id uint) (*model.Goods, error)
	GetAllGoods() ([]model.Goods, error)
	UpdateGoods(req UpdateGoodsRequest) (*model.Goods, error)
	DeleteGoods(id uint) error
}

type goodsService struct {
	repo *repository.GoodsRepository
}

func NewGoodsService(repo *repository.GoodsRepository) GoodsService {
	return &goodsService{repo: repo}
}

func (s *goodsService) CreateGoods(req CreateGoodsRequest, userID uint) (*model.Goods, error) {
	if req.Name == "" {
		return nil, errors.New("goods name is required")
	}
	goods := &model.Goods{
		Name:   req.Name,
		Num:    req.Num,
		Price:  req.Price,
		UserID: userID,
	}
	goods, err := s.repo.CreateGoods(goods)
	if err == nil {
		cache.DelGoodsAll()
	}
	return goods, err
}

func (s *goodsService) GetGoodsById(id uint) (*model.Goods, error) {
	if id == 0 {
		return nil, errors.New("invalid goods id")
	}
	goods, err := s.repo.FindAim(id)
	if err != nil {
		return nil, err
	}
	if goods == nil {
		return nil, errors.New("goods not found")
	}
	return goods, nil
}

func (s *goodsService) GetAllGoods() ([]model.Goods, error) {
	goods, err := cache.GetGoodsAll()
	if err == nil {
		return goods, nil
	}
	goods, err = s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	cache.SetGoodsAll(goods)
	return goods, nil
}

func (s *goodsService) UpdateGoods(req UpdateGoodsRequest) (*model.Goods, error) {
	if req.ID == 0 {
		return nil, errors.New("invalid goods id")
	}
	goods, err := s.repo.FindAim(req.ID)
	if err != nil {
		return nil, err
	}
	if goods == nil {
		return nil, errors.New("goods not found")
	}
	if req.Name != nil {
		goods.Name = *req.Name
	}
	if req.Num != nil {
		goods.Num = *req.Num
	}
	if req.Price != nil {
		goods.Price = *req.Price
	}
	return s.repo.UpdateGoods(goods)
}

func (s *goodsService) DeleteGoods(id uint) error {
	if id == 0 {
		return errors.New("invalid goods id")
	}
	rows, err := s.repo.DeleteGoods(id)
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("goods not found")
	}
	cache.DelGoodsAll()
	return nil
}

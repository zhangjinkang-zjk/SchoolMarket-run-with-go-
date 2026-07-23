package userservice

import (
	"errors"

	"SchoolMarket-run-with-go-/internal/model"
	"SchoolMarket-run-with-go-/internal/repository"
)

type AddCartRequest struct {
	GoodsID uint
	Num     uint
}

type UpdateCartRequest struct {
	ID  uint
	Num *uint
}

type CartService interface {
	AddCart(req AddCartRequest, userID uint) (*model.Cart, error)
	GetCart(userID uint) ([]model.Cart, error)
	UpdateCart(req UpdateCartRequest, userID uint) (*model.Cart, error)
	DeleteCart(id, userID uint) error
}

type cartService struct {
	cartRepo  *repository.CartRepository
	goodsRepo *repository.GoodsRepository
}

func NewCartService(cartRepo *repository.CartRepository, goodsRepo *repository.GoodsRepository) CartService {
	return &cartService{cartRepo: cartRepo, goodsRepo: goodsRepo}
}

func (s *cartService) AddCart(req AddCartRequest, userID uint) (*model.Cart, error) {
	if req.GoodsID == 0 || req.Num == 0 {
		return nil, errors.New("invalid params")
	}
	goods, err := s.goodsRepo.FindAim(req.GoodsID)
	if err != nil {
		return nil, err
	}
	if goods == nil {
		return nil, errors.New("goods not found")
	}
	cart := &model.Cart{
		UserID:  userID,
		GoodsID: req.GoodsID,
		Num:     req.Num,
	}
	return s.cartRepo.CreateCart(cart)
}

func (s *cartService) GetCart(userID uint) ([]model.Cart, error) {
	return s.cartRepo.FindByUserId(userID)
}

func (s *cartService) UpdateCart(req UpdateCartRequest, userID uint) (*model.Cart, error) {
	if req.ID == 0 {
		return nil, errors.New("invalid cart id")
	}
	cart, err := s.cartRepo.FindAim(req.ID)
	if err != nil {
		return nil, err
	}
	if cart == nil || cart.UserID != userID {
		return nil, errors.New("cart not found")
	}
	if req.Num != nil {
		cart.Num = *req.Num
	}
	return s.cartRepo.UpdateCart(cart)
}

func (s *cartService) DeleteCart(id, userID uint) error {
	if id == 0 {
		return errors.New("invalid cart id")
	}
	cart, err := s.cartRepo.FindAim(id)
	if err != nil {
		return err
	}
	if cart == nil || cart.UserID != userID {
		return errors.New("cart not found")
	}
	_, err = s.cartRepo.DeleteCart(id)
	return err
}

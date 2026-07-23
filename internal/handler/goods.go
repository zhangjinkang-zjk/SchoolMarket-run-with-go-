package userhandler

import (
	"strconv"

	userservice "SchoolMarket-run-with-go-/internal/service"

	"github.com/gin-gonic/gin"
)

type GoodsHandler struct {
	serv userservice.GoodsService
}

func NewGoodsHandler(serv userservice.GoodsService) *GoodsHandler {
	return &GoodsHandler{serv: serv}
}

type CreateGoodsRequest struct {
	Name  string `json:"name"`
	Num   uint   `json:"num"`
	Price uint   `json:"price"`
}

type UpdateGoodsRequest struct {
	ID    uint    `json:"id"`
	Name  *string `json:"name"`
	Num   *uint   `json:"num"`
	Price *uint   `json:"price"`
}

func (h *GoodsHandler) CreateGoods(c *gin.Context) {
	var req CreateGoodsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	userID, _ := c.Get("user_id")
	goods, err := h.serv.CreateGoods(userservice.CreateGoodsRequest{
		Name:  req.Name,
		Num:   req.Num,
		Price: req.Price,
	}, userID.(uint))
	if err != nil {
		c.JSON(400, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{
			"id":         goods.ID,
			"name":       goods.Name,
			"num":        goods.Num,
			"price":      goods.Price,
			"user_id":    goods.UserID,
			"created_at": goods.CreatedAt,
			"updated_at": goods.UpdatedAt,
		},
	})
}

func (h *GoodsHandler) GetGoodsById(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"code": 400, "msg": "invalid goods id"})
		return
	}
	goods, err := h.serv.GetGoodsById(uint(id))
	if err != nil {
		c.JSON(400, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{
			"id":         goods.ID,
			"name":       goods.Name,
			"num":        goods.Num,
			"price":      goods.Price,
			"user_id":    goods.UserID,
			"created_at": goods.CreatedAt,
			"updated_at": goods.UpdatedAt,
		},
	})
}

func (h *GoodsHandler) GetAllGoods(c *gin.Context) {
	goods, err := h.serv.GetAllGoods()
	if err != nil {
		c.JSON(400, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	data := make([]gin.H, 0, len(goods))
	for _, g := range goods {
		data = append(data, gin.H{
			"id":         g.ID,
			"name":       g.Name,
			"num":        g.Num,
			"price":      g.Price,
			"user_id":    g.UserID,
			"created_at": g.CreatedAt,
			"updated_at": g.UpdatedAt,
		})
	}
	c.JSON(200, gin.H{"code": 200, "data": data})
}

func (h *GoodsHandler) UpdateGoods(c *gin.Context) {
	var req UpdateGoodsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	goods, err := h.serv.UpdateGoods(userservice.UpdateGoodsRequest{
		ID:    req.ID,
		Name:  req.Name,
		Num:   req.Num,
		Price: req.Price,
	})
	if err != nil {
		c.JSON(400, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{
			"id":         goods.ID,
			"name":       goods.Name,
			"num":        goods.Num,
			"price":      goods.Price,
			"user_id":    goods.UserID,
			"created_at": goods.CreatedAt,
			"updated_at": goods.UpdatedAt,
		},
	})
}

func (h *GoodsHandler) DeleteGoods(c *gin.Context) {
	var req struct {
		ID uint `json:"id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	if err := h.serv.DeleteGoods(req.ID); err != nil {
		c.JSON(400, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 200, "msg": "deleted"})
}

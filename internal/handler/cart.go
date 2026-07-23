package userhandler

import (
	userservice "SchoolMarket-run-with-go-/internal/service"

	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	serv userservice.CartService
}

func NewCartHandler(serv userservice.CartService) *CartHandler {
	return &CartHandler{serv: serv}
}

type AddCartRequest struct {
	GoodsID uint `json:"goods_id"`
	Num     uint `json:"num"`
}

type UpdateCartRequest struct {
	ID  uint  `json:"id"`
	Num *uint `json:"num"`
}

func (h *CartHandler) AddCart(c *gin.Context) {
	var req AddCartRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	userID, _ := c.Get("user_id")
	cart, err := h.serv.AddCart(userservice.AddCartRequest{
		GoodsID: req.GoodsID,
		Num:     req.Num,
	}, userID.(uint))
	if err != nil {
		c.JSON(400, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{
			"id":         cart.ID,
			"goods_id":   cart.GoodsID,
			"num":        cart.Num,
			"created_at": cart.CreatedAt,
		},
	})
}

func (h *CartHandler) GetCart(c *gin.Context) {
	userID, _ := c.Get("user_id")
	carts, err := h.serv.GetCart(userID.(uint))
	if err != nil {
		c.JSON(400, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	data := make([]gin.H, 0, len(carts))
	for _, cart := range carts {
		item := gin.H{
			"id":      cart.ID,
			"goods_id": cart.GoodsID,
			"num":     cart.Num,
		}
		if cart.Goods.ID != 0 {
			item["goods_name"] = cart.Goods.Name
			item["goods_price"] = cart.Goods.Price
		}
		data = append(data, item)
	}
	c.JSON(200, gin.H{"code": 200, "data": data})
}

func (h *CartHandler) UpdateCart(c *gin.Context) {
	var req UpdateCartRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	userID, _ := c.Get("user_id")
	cart, err := h.serv.UpdateCart(userservice.UpdateCartRequest{
		ID:  req.ID,
		Num: req.Num,
	}, userID.(uint))
	if err != nil {
		c.JSON(400, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{
			"id":       cart.ID,
			"goods_id": cart.GoodsID,
			"num":      cart.Num,
		},
	})
}

func (h *CartHandler) DeleteCart(c *gin.Context) {
	var req struct {
		ID uint `json:"id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	userID, _ := c.Get("user_id")
	if err := h.serv.DeleteCart(req.ID, userID.(uint)); err != nil {
		c.JSON(400, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 200, "msg": "deleted"})
}

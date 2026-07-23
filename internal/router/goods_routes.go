package userrouter

import (
	userhandler "SchoolMarket-run-with-go-/internal/handler"
	"SchoolMarket-run-with-go-/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterGoodsRouter(r *gin.Engine, h *userhandler.GoodsHandler) {
	auth := r.Group("/goods")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.POST("/create", h.CreateGoods)
		auth.GET("/find/aim", h.GetGoodsById)
		auth.GET("/find/all", h.GetAllGoods)
		auth.PUT("/update", h.UpdateGoods)
		auth.DELETE("/delete", h.DeleteGoods)
	}
}

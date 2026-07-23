package userrouter

import (
	userhandler "SchoolMarket-run-with-go-/internal/handler"
	"SchoolMarket-run-with-go-/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterCartRouter(r *gin.Engine, h *userhandler.CartHandler) {
	auth := r.Group("/cart")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.POST("/add", h.AddCart)
		auth.GET("/get", h.GetCart)
		auth.PUT("/update", h.UpdateCart)
		auth.DELETE("/delete", h.DeleteCart)
	}
}

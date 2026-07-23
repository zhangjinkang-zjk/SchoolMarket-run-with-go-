package userrouter

import (
	userhandler "SchoolMarket-run-with-go-/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterUserRouter(r *gin.Engine, h *userhandler.UserHandler) {
	users := r.Group("/users")
	users.POST("/create", h.CreateUser)
	users.POST("/login", h.Login)
}

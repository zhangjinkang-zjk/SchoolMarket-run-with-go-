package userrouter

import (
	userhandler "SchoolMarket-run-with-go-/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine, h *userhandler.UserHandler) {
	users := r.Group("/users")

	users.POST("/create", h.CreateUser)
	users.GET("/find/aim", h.GetUserById)
	users.GET("/find/all", h.GetAllUser)
	users.PUT("/update", h.UpdateUser)
	users.DELETE("/delete", h.Delete)
}

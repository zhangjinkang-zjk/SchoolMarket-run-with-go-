package main

import (
	"SchoolMarket-run-with-go-/config"
	"SchoolMarket-run-with-go-/internal/database"
	userhandler "SchoolMarket-run-with-go-/internal/handler"
	"SchoolMarket-run-with-go-/internal/repository"
	userrouter "SchoolMarket-run-with-go-/internal/router"
	userservice "SchoolMarket-run-with-go-/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	db := database.InitDB(cfg)

	userRepo := repository.NewUserRepository(db)
	userServ := userservice.NewUserService(userRepo)
	userHandler := userhandler.NewUserHandler(userServ)

	r := gin.Default()

	userrouter.RegisterRouter(r, userHandler)

	r.Run(":" + cfg.PORT)
}

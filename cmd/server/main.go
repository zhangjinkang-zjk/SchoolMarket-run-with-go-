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

	// User
	userRepo := repository.NewUserRepository(db)
	userServ := userservice.NewUserService(userRepo)
	userHandler := userhandler.NewUserHandler(userServ)

	// Goods
	goodsRepo := repository.NewGoodsRepository(db)
	goodsServ := userservice.NewGoodsService(goodsRepo)
	goodsHandler := userhandler.NewGoodsHandler(goodsServ)

	// Cart
	cartRepo := repository.NewCartRepository(db)
	cartServ := userservice.NewCartService(cartRepo, goodsRepo)
	cartHandler := userhandler.NewCartHandler(cartServ)

	_ = database.InitRedis(cfg)

	r := gin.Default()

	r.StaticFile("/", "./web/index.html")
	userrouter.RegisterUserRouter(r, userHandler)
	userrouter.RegisterGoodsRouter(r, goodsHandler)
	userrouter.RegisterCartRouter(r, cartHandler)

	r.Run(":" + cfg.PORT)
}

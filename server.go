package main

import (
	"main.go/config"
	bDelivery "main.go/features/book/delivery"
	bRepo "main.go/features/book/repository"
	bServices "main.go/features/book/services"
	uDelivery "main.go/features/user/delivery"
	uRepo "main.go/features/user/repository"
	uServices "main.go/features/user/services"
	"main.go/utils/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.NewConfig()
	db := database.InitDB(cfg)
	database.MigrateDB(db)
	uRepo := uRepo.New(db)
	bRepo := bRepo.New(db)

	uService := uServices.New(uRepo)
	bService := bServices.New(bRepo)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	uDelivery.New(e, uService)
	bDelivery.New(e, bService)

	e.Logger.Fatal(e.Start(":8000"))
}

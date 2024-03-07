package main

import (
	"auth/hexagonal/internal/adapters/handler"
	"auth/hexagonal/internal/adapters/repositorys"
	"auth/hexagonal/internal/core/domain"
	"auth/hexagonal/internal/core/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	authService *services.AuthService
	beerService *services.BeerService
)

func main() {
	// Connection to PostgreSQL
	dsn := "user=postgres password=testpass123 dbname=go-hex port=3500 sslmode=disable TimeZone=Asia/Bangkok"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&domain.User{}, &domain.Beer{})

	store := repositorys.NewDB(db)

	authService = services.NewAuthService(store)
	beerService = services.NewBeerService(store)

	InitRoute()
}

func InitRoute() {
	router := fiber.New()
	router.Use(cors.New())

	router.Static("/uploads", "./uploads")

	v1 := router.Group("/api")

	authHandler := handler.NewAuthHandler(*authService)
	v1.Post("/signup", authHandler.SignUp)
	v1.Post("/signin", authHandler.SignIn)

	beerHandler := handler.NewBeerHandler(*beerService)
	v1.Post("/order", beerHandler.SaveBeer)
	v1.Get("/order", beerHandler.ReadBeers)
	v1.Get("/order/:id", beerHandler.ReadBeer)
	v1.Delete("/order/:id", beerHandler.DeleteBeer)
	v1.Put("/order/:id", beerHandler.UpdateBeer)

	router.Listen(":8080")
}

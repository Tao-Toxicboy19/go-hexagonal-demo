package main

import (
	"auth/hexagonal/adapters/handler"
	"auth/hexagonal/adapters/repositorys"
	"auth/hexagonal/core/domain"
	"auth/hexagonal/core/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	authService  *services.AuthService
	orderService *services.OrderService
)

func main() {
	// Connection to PostgreSQL
	dsn := "user=postgres password=testpass123 dbname=go-hex port=3500 sslmode=disable TimeZone=Asia/Bangkok"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&domain.User{}, &domain.Order{})

	store := repositorys.NewDB(db)

	authService = services.NewAuthService(store)
	orderService = services.NewOrderService(store)

	InitRoute()
}

func InitRoute() {
	router := fiber.New()
	v1 := router.Group("/api")

	authHandler := handler.NewAuthHandler(*authService)
	v1.Post("/signup", authHandler.SignUp)
	v1.Post("/signin", authHandler.SignIn)

	orderHandler := handler.NewOrderHandler(*orderService)
	v1.Post("/order", orderHandler.SaveOrder)
	v1.Get("/order", orderHandler.ReadOrders)
	router.Listen(":8080")
}

package main

import (
	"auth/hexagonal/internal/adapters/handler"
	"auth/hexagonal/internal/adapters/repositorys"
	"auth/hexagonal/internal/core/domain"
	"auth/hexagonal/internal/core/services"

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
	v1.Get("/order/:id", orderHandler.ReadOrder)
	v1.Delete("/order/:id", orderHandler.DeleteOrder)
	v1.Put("/order/:id", orderHandler.UpdateOrder)
	router.Listen(":8080")
}

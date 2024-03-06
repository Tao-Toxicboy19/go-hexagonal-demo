package main

import (
	"auth/hexagonal/internal/adapters/handler"
	"auth/hexagonal/internal/adapters/repositorys"
	"auth/hexagonal/internal/core/domain"
	"auth/hexagonal/internal/core/services"
	"fmt"
	"io"
	"time"

	"os"

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

func uploadHandler(c *fiber.Ctx) error {
	// รับไฟล์จาก form data
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	// ดึงข้อมูลไฟล์ที่อัปโหลด
	files := form.File["files"]

	// วนลูปตรวจสอบและบันทึกไฟล์
	for _, file := range files {
		// อ่านไฟล์
		src, err := file.Open()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		defer src.Close()

		currentTime := time.Now()
		newFilename := fmt.Sprintf("%d%02d%02d_%02d%02d%02d_%s",
			currentTime.Year(), currentTime.Month(), currentTime.Day(),
			currentTime.Hour(), currentTime.Minute(), currentTime.Second(),
			file.Filename)

		fmt.Println(newFilename)
		// สร้างไฟล์ปลายทาง
		dst, err := os.Create("./uploads/" + newFilename)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		defer dst.Close()

		// คัดลอกข้อมูลไฟล์
		if _, err = io.Copy(dst, src); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
	}

	// ส่งข้อความเมื่ออัปโหลดเสร็จสิ้น
	return c.SendString("Upload successful")
}

func InitRoute() {
	router := fiber.New()
	router.Use(cors.New())

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

	v1.Post("/upload", uploadHandler)

	router.Listen(":8080")
}

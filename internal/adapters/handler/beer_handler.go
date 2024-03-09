package handler

import (
	"auth/hexagonal/internal/core/domain"
	"auth/hexagonal/internal/core/services"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type BeerHandler struct {
	service services.BeerService
}

func NewBeerHandler(BeerService services.BeerService) *BeerHandler {
	return &BeerHandler{service: BeerService}
}

func (h *BeerHandler) SaveBeer(c *fiber.Ctx) error {
	// รับข้อมูลจาก FormData
	form, err := c.MultipartForm()
	if err != nil {
		return HandlerError(c, fiber.StatusBadRequest, err)
	}

	// ดึงข้อมูล price และ stock จาก FormData
	priceStr := form.Value["price"][0]
	stockStr := form.Value["stock"][0]

	// แปลงข้อมูลเป็น float64
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return HandlerError(c, fiber.StatusBadRequest, err)
	}

	stock, err := strconv.ParseFloat(stockStr, 64)
	if err != nil {
		return HandlerError(c, fiber.StatusBadRequest, err)
	}

	var beer domain.Beer
	if err := c.BodyParser(&beer); err != nil {
		return HandlerError(c, fiber.StatusBadRequest, err)
	}

	if err := Validate(c); err != nil {
		return HandlerError(c, fiber.StatusBadRequest, err)
	}

	// รับไฟล์จาก FormData
	files := form.File["files"]
	for _, file := range files {
		// อ่านไฟล์
		src, err := file.Open()
		if err != nil {
			return HandlerError(c, fiber.StatusInternalServerError, err)
		}
		defer src.Close()

		currentTime := time.Now()
		newFilename := fmt.Sprintf("%d%02d%02d_%02d%02d%02d_%s",
			currentTime.Year(), currentTime.Month(), currentTime.Day(),
			currentTime.Hour(), currentTime.Minute(), currentTime.Second(),
			file.Filename)

		// ตรวจสอบว่าแปลงข้อมูลได้สำเร็จหรือไม่
		beer.Price = price
		beer.Stock = stock
		beer.Image = newFilename

		if err := h.service.SaveBeer(&beer); err != nil {
			return HandlerError(c, fiber.StatusBadRequest, err)
		}

		dst, err := os.Create("./uploads/" + newFilename)
		if err != nil {
			return HandlerError(c, fiber.StatusInternalServerError, err)
		}
		defer dst.Close()

		if _, err = io.Copy(dst, src); err != nil {
			return HandlerError(c, fiber.StatusInternalServerError, err)
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "ok"})
}

func (h *BeerHandler) ReadBeers(c *fiber.Ctx) error {
	beers, err := h.service.ReadBeers()
	if err != nil {
		return HandlerError(c, fiber.StatusBadRequest, err)
	}

	if err := Validate(c); err != nil {
		return HandlerError(c, fiber.StatusBadRequest, err)
	}

	return c.Status(fiber.StatusOK).JSON(beers)
}

func (h *BeerHandler) ReadByUserId(c *fiber.Ctx) error {
	id := c.Params("id")
	beers, err := h.service.ReadByUserId(id)
	if err != nil {
		return HandlerError(c, fiber.StatusBadRequest, err)
	}

	if err := Validate(c); err != nil {
		return HandlerError(c, fiber.StatusBadRequest, err)
	}

	return c.Status(fiber.StatusOK).JSON(beers)
}

func (h *BeerHandler) ReadBeer(c *fiber.Ctx) error {
	id := c.Params("id")
	beer, err := h.service.ReadBeer(id)
	if err != nil {
		return HandlerError(c, fiber.StatusBadRequest, err)
	}

	if err := Validate(c); err != nil {
		return HandlerError(c, fiber.StatusBadRequest, err)
	}

	return c.Status(fiber.StatusOK).JSON(beer)
}

func (h *BeerHandler) DeleteBeer(c *fiber.Ctx) error {
	id := c.Params("id")
	fmt.Println(id)
	if err := Validate(c); err != nil {
		return HandlerError(c, fiber.StatusBadRequest, err)
	}

	if err := h.service.DeleteBeer(id); err != nil {
		return HandlerError(c, fiber.StatusBadRequest, err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "ok"})
}

func (h *BeerHandler) UpdateBeer(c *fiber.Ctx) error {

	// รับข้อมูลจาก FormData
	id := c.Params("id")
	form, err := c.MultipartForm()
	if err != nil {
		return HandlerError(c, fiber.StatusBadRequest, err)
	}

	// ดึงข้อมูล Beer จาก FormData
	var beer domain.Beer
	if err := c.BodyParser(&beer); err != nil {
		return HandlerError(c, fiber.StatusBadRequest, err)
	}
	beer.ID = id

	if err := Validate(c); err != nil {
		return HandlerError(c, fiber.StatusBadRequest, err)
	}

	if len(form.File) == 0 {
		if err := h.service.UpdateBeer(id, &beer); err != nil {
			return HandlerError(c, fiber.StatusBadRequest, err)
		}
		return c.Status(fiber.StatusCreated).JSON(beer)

	}

	// รับไฟล์จาก FormData
	files := form.File["files"]
	for _, file := range files {
		// อ่านไฟล์
		src, err := file.Open()
		if err != nil {
			return HandlerError(c, fiber.StatusInternalServerError, err)
		}
		defer src.Close()

		currentTime := time.Now()
		newFilename := fmt.Sprintf("%d%02d%02d_%02d%02d%02d_%s",
			currentTime.Year(), currentTime.Month(), currentTime.Day(),
			currentTime.Hour(), currentTime.Minute(), currentTime.Second(),
			file.Filename)

		beer.Image = newFilename

		if err := h.service.UpdateBeer(id, &beer); err != nil {
			return HandlerError(c, fiber.StatusBadRequest, err)
		}

		dst, err := os.Create("./uploads/" + newFilename)
		if err != nil {
			return HandlerError(c, fiber.StatusInternalServerError, err)
		}
		defer dst.Close()

		if _, err = io.Copy(dst, src); err != nil {
			return HandlerError(c, fiber.StatusInternalServerError, err)
		}
	}
	return c.Status(fiber.StatusCreated).JSON(beer)
}

// func (h *BeerHandler) UpdateBeer(c *fiber.Ctx) error {
// 	var beer domain.Beer
// 	id := c.Params("id")

// 	if err := c.BodyParser(&beer); err != nil {
// 		return HandlerError(c, fiber.StatusBadRequest, err)
// 	}

// 	if err := Validate(c); err != nil {
// 		return HandlerError(c, fiber.StatusBadRequest, err)
// 	}

// 	if err := h.service.UpdateBeer(id, &beer); err != nil {
// 		return HandlerError(c, fiber.StatusBadRequest, err)
// 	}

// 	return c.Status(fiber.StatusCreated).JSON(beer)
// }

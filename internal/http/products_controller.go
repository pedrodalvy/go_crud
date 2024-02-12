package http

import (
	"github.com/gofiber/fiber/v2"
	"go_crud/internal/http/dto"
	"go_crud/internal/products"
	"gorm.io/gorm"
)

type ProductsController struct {
	service *products.Service
}

func NewProductsController(db *gorm.DB) *ProductsController {
	return &ProductsController{
		service: products.NewService(db),
	}
}

func (pc *ProductsController) GetProducts(c *fiber.Ctx) error {
	p, err := pc.service.GetProducts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"status": "error", "message": "Error trying to get products"})
	}

	return c.JSON(p)
}

func (pc *ProductsController) FindProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	p, err := pc.service.FindProduct(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"status": "error", "message": "Error trying to find product"})
	}

	return c.JSON(p)
}

func (pc *ProductsController) CreateProduct(c *fiber.Ctx) error {
	input := new(dto.CreateProductDTO)
	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"status": "error", "message": "Error on request"})
	}

	p, err := pc.service.CreateProduct(input.Name)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"status": "error", "message": "Error trying to create product"})
	}

	return c.Status(fiber.StatusCreated).JSON(p)
}

func (pc *ProductsController) UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	input := new(dto.UpdateProductDTO)
	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"status": "error", "message": "Error on request"})
	}

	err := pc.service.UpdateProduct(id, input.Name)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"status": "error", "message": "Error trying to update product"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (pc *ProductsController) DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	err := pc.service.DeleteProduct(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"status": "error", "message": "Error trying to delete product"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

package products

import (
	"github.com/gofiber/fiber/v2"
	dto2 "go_crud/internal/products/dto"
	"gorm.io/gorm"
)

type Controller struct {
	service *Service
}

func NewController(db *gorm.DB) *Controller {
	return &Controller{
		service: NewService(db),
	}
}

func (c *Controller) GetProducts(ctx *fiber.Ctx) error {
	p, err := c.service.GetProducts()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"status": "error", "message": "Error trying to get products"})
	}

	return ctx.JSON(p)
}

func (c *Controller) FindProduct(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	p, err := c.service.FindProduct(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"status": "error", "message": "Error trying to find product"})
	}

	return ctx.JSON(p)
}

func (c *Controller) CreateProduct(ctx *fiber.Ctx) error {
	input := new(dto2.CreateProductDTO)
	if err := ctx.BodyParser(input); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"status": "error", "message": "Error on request"})
	}

	p, err := c.service.CreateProduct(input.Name)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"status": "error", "message": "Error trying to create product"})
	}

	return ctx.Status(fiber.StatusCreated).JSON(p)
}

func (c *Controller) UpdateProduct(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	input := new(dto2.UpdateProductDTO)
	if err := ctx.BodyParser(input); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"status": "error", "message": "Error on request"})
	}

	err := c.service.UpdateProduct(id, input.Name)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"status": "error", "message": "Error trying to update product"})
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

func (c *Controller) DeleteProduct(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	err := c.service.DeleteProduct(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"status": "error", "message": "Error trying to delete product"})
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

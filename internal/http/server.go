package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go_crud/internal/products"
	"gorm.io/gorm"
)

type Server struct {
	db *gorm.DB
}

func NewServer(db *gorm.DB) *Server {
	return &Server{
		db: db,
	}
}

func (s *Server) Start(port string) error {
	app := fiber.New()
	app.Use(logger.New())

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	api := app.Group("/api")

	productsController := products.NewController(s.db)

	api.Get("/products", productsController.GetProducts)
	api.Get("/products/:id", productsController.FindProduct)
	api.Post("/products", productsController.CreateProduct)
	api.Put("/products/:id", productsController.UpdateProduct)
	api.Delete("/products/:id", productsController.DeleteProduct)

	return app.Listen(port)
}

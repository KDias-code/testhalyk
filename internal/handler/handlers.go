package handler

import (
	"github.com/ansrivas/fiberprometheus/v2"
	"test/internal/service"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service service.Service
}

func NewHandler(s service.Service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) InitRoutes() *fiber.App {
	router := fiber.New()

	router.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{"message": "pong"})
	})

	prometheus := fiberprometheus.New("my-service-name")
	prometheus.RegisterAt(router, "/api")
	router.Use(prometheus.Middleware)

	router.Get("/author", h.GetAuthors)
	router.Get("/author/:id", h.GetAuthorById)
	router.Post("/author", h.CreateAuthor)
	router.Patch("/author/:id", h.UpdateAuthor)
	router.Delete("/author/:id", h.DeleteAuthor)

	router.Get("/author/:id/books", h.AuthorBooks)

	router.Get("/book/:id", h.GetBookById)
	router.Post("/book", h.CreateBook)
	router.Patch("/book/:id", h.UpdateBook)
	router.Delete("/book/:id", h.DeleteBook)

	router.Get("/reader/:id", h.GetReaderById)
	router.Post("/reader", h.CreateReader)
	router.Patch("/reader/:id", h.UpdateReader)
	router.Delete("/reader/:id", h.DeleteAuthor)
	router.Post("/reader/:id", h.TakeBook)

	router.Get("/reader/:id/books", h.ReaderBooks)

	return router
}

package handler

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
	"test/internal/entity"
	customError "test/pkg/error"
)

func (h *Handler) GetBookById(c *fiber.Ctx) error {
	bookID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return serverError(c, fiber.StatusBadRequest, err.Error())
	} else if bookID < 0 {
		return serverError(c, fiber.StatusBadRequest, "negative id number")
	}

	book, err := h.service.Book.GetBookById(bookID)
	//log.Println(book, err)
	if err != nil {
		if errors.Is(err, customError.ErrNothingToFound) {
			return serverError(c, fiber.StatusOK, errors.Unwrap(err).Error())
		}
		return serverError(c, fiber.StatusInternalServerError, errors.Unwrap(err).Error())
	}

	return c.Status(fiber.StatusOK).JSON(book)
}

func (h *Handler) CreateBook(c *fiber.Ctx) error {
	var book entity.Book

	if err := c.BodyParser(&book); err != nil {
		return serverError(c, fiber.StatusInternalServerError, err.Error())
	}

	if err := h.service.Book.CreateBook(book); err != nil {

		if errors.Is(err, customError.ErrEmptyFields) {
			log.Println(err)
			return serverError(c, fiber.StatusBadRequest, errors.Unwrap(err).Error())
		}

		log.Println(err)
		return serverError(c, fiber.StatusInternalServerError, errors.Unwrap(err).Error())
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (h *Handler) UpdateBook(c *fiber.Ctx) error {
	bookID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return serverError(c, fiber.StatusBadRequest, err.Error())
	} else if bookID < 0 {
		return serverError(c, fiber.StatusBadRequest, "negative id number")
	}

	var book entity.Book
	if err := c.BodyParser(&book); err != nil {
		return serverError(c, fiber.StatusInternalServerError, err.Error())
	}

	book.ID = bookID

	if err := h.service.Book.UpdateBook(book); err != nil {

		if errors.Is(err, customError.ErrNothingToUpdate) {
			log.Println(err)
			return serverError(c, fiber.StatusOK, errors.Unwrap(err).Error())
		} else if errors.Is(err, customError.ErrNothingToFound) {
			log.Println(err)
			return serverError(c, fiber.StatusOK, errors.Unwrap(err).Error())
		}

		log.Println(err)
		return serverError(c, fiber.StatusInternalServerError, errors.Unwrap(err).Error())
	}

	return c.Status(fiber.StatusOK).JSON(response{Message: "updated"})

}

func (h *Handler) DeleteBook(c *fiber.Ctx) error {
	bookID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return serverError(c, fiber.StatusBadRequest, err.Error())
	} else if bookID < 0 {
		return serverError(c, fiber.StatusBadRequest, "negative id number")
	}

	return c.Status(fiber.StatusOK).JSON(response{Message: "deleted"})
}

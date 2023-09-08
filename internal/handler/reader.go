package handler

import (
	"errors"
	"log"
	"strconv"
	"test/internal/entity"
	customError "test/pkg/error"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetReaderById(c *fiber.Ctx) error {

	readerID, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return serverError(c, fiber.StatusBadRequest, err.Error())
	} else if readerID < 0 {
		return serverError(c, fiber.StatusBadRequest, "negative id number")
	}

	reader, err := h.service.Reader.GetReaderById(readerID)
	if err != nil {

		if errors.Is(err, customError.ErrNothingToFound) {
			log.Println(err)
			return serverError(c, fiber.StatusOK, errors.Unwrap(err).Error())
		}

		log.Println(err)
		return serverError(c, fiber.StatusInternalServerError, errors.Unwrap(err).Error())

	}

	return c.Status(fiber.StatusOK).JSON(reader)
}

func (h *Handler) CreateReader(c *fiber.Ctx) error {

	var reader entity.Reader

	if err := c.BodyParser(&reader); err != nil {
		return serverError(c, fiber.StatusInternalServerError, errors.Unwrap(err).Error())
	}

	if err := h.service.Reader.CreateReader(reader); err != nil {

		if errors.Is(err, customError.ErrEmptyFields) {
			log.Println(err)
			return serverError(c, fiber.StatusBadRequest, errors.Unwrap(err).Error())
		}

		log.Println(err)
		return serverError(c, fiber.StatusInternalServerError, errors.Unwrap(err).Error())
	}

	return c.Status(fiber.StatusCreated).JSON(response{Message: "created"})
}

func (h *Handler) UpdateReader(c *fiber.Ctx) error {
	readerID, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return serverError(c, fiber.StatusBadRequest, err.Error())
	} else if readerID < 0 {
		return serverError(c, fiber.StatusBadRequest, "negative id number")
	}

	var reader entity.Reader
	if err := c.BodyParser(&reader); err != nil {
		return serverError(c, fiber.StatusInternalServerError, err.Error())
	}

	reader.ID = readerID

	if err := h.service.Reader.UpdateReader(reader); err != nil {

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

func (h *Handler) DeleteReader(c *fiber.Ctx) error {
	readerID, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return serverError(c, fiber.StatusBadRequest, err.Error())
	} else if readerID < 0 {
		return serverError(c, fiber.StatusBadRequest, "negative id number")
	}

	if err := h.service.Reader.DeleteReader(readerID); err != nil {
		return serverError(c, fiber.StatusInternalServerError, errors.Unwrap(err).Error())
	}

	return c.Status(fiber.StatusOK).JSON(response{Message: "deleted"})
}

func (h *Handler) TakeBook(c *fiber.Ctx) error {
	readerID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return serverError(c, fiber.StatusBadRequest, err.Error())
	} else if readerID < 0 {
		return serverError(c, fiber.StatusBadRequest, "negative id number")
	}

	var book entity.Book
	if err := c.BodyParser(&book); err != nil {
		log.Println(err)
		return serverError(c, fiber.StatusInternalServerError, err.Error())
	}

	if err := h.service.Reader.TakeBook(readerID, book.ID); err != nil {

		if errors.Is(err, customError.ErrNothingToFound) {
			log.Println(err)
			return serverError(c, fiber.StatusBadRequest, errors.Unwrap(err).Error())
		}

		log.Println(err)
		return serverError(c, fiber.StatusInternalServerError, errors.Unwrap(err).Error())
	}

	return c.Status(fiber.StatusOK).JSON(response{Message: "taken"})
}

func (h *Handler) ReaderBooks(c *fiber.Ctx) error {
	readerID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return serverError(c, fiber.StatusBadRequest, err.Error())
	} else if readerID < 0 {
		return serverError(c, fiber.StatusBadRequest, "negative id number")
	}

	bookList, err := h.service.Reader.GetReaderBookList(readerID)
	if err != nil {

		if errors.Is(err, customError.ErrNothingToFound) {
			log.Println(err)
			return serverError(c, fiber.StatusOK, errors.Unwrap(err).Error())
		}

		log.Println(err)
		return serverError(c, fiber.StatusInternalServerError, errors.Unwrap(err).Error())
	}

	return c.Status(fiber.StatusOK).JSON(bookList)
}

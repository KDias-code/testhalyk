package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
	"test/internal/domain/entity"
	errs "test/pkg/error"
)

func (h *Handler) CreateAuthor(c *fiber.Ctx) error {

	fmt.Println(string(c.Body()))

	author := new(entity.Author)

	if err := c.BodyParser(author); err != nil {
		log.Println("Failed to parse request body", "error", err)
		return serverError(c, fiber.StatusInternalServerError, errs.ErrServer.Error())
	}

	data, _ := json.Marshal(author)
	fmt.Println(string(data))

	if err := h.service.CreateAuthor(*author); err != nil {

		if errors.Is(err, errs.ErrEmptyFields) {
			log.Println(err)
			return serverError(c, fiber.StatusBadRequest, errs.ErrEmptyFields.Error())
		}

		log.Println(err)
		return serverError(c, fiber.StatusInternalServerError, errs.ErrServer.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(response{Message: "created"})
}

func (h *Handler) GetAuthors(c *fiber.Ctx) error {
	authors, err := h.service.GetAuthors()

	if err != nil {
		if errors.Is(err, errs.ErrNothingToFound) {
			log.Println(err)
			return serverError(c, fiber.StatusOK, errors.Unwrap(err).Error())
		}
		log.Println(err)
		return serverError(c, fiber.StatusInternalServerError, errs.ErrServer.Error())
	}

	return c.Status(fiber.StatusOK).JSON(authors)
}

func (h *Handler) GetAuthorById(c *fiber.Ctx) error {
	authorID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Println(err)
		return serverError(c, fiber.StatusBadRequest, errs.ErrInvalidId.Error())
	} else if authorID < 0 {
		return serverError(c, fiber.StatusBadRequest, errs.ErrInvalidId.Error())
	}

	author, err := h.service.GetAuthorById(c.Context(), authorID)
	if err != nil {

		if errors.Is(err, errs.ErrNothingToFound) {
			log.Println(err)
			return serverError(c, fiber.StatusOK, errs.ErrNothingToFound.Error())
		}

		log.Println(err)
		return serverError(c, fiber.StatusInternalServerError, errs.ErrServer.Error())
	}

	return c.Status(fiber.StatusOK).JSON(author)
}

func (h *Handler) DeleteAuthor(c *fiber.Ctx) error {
	authorID, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		log.Println(err)
		return serverError(c, fiber.StatusBadRequest, errs.ErrInvalidId.Error())
	} else if authorID < 0 {
		return serverError(c, fiber.StatusBadRequest, errs.ErrInvalidId.Error())
	}

	if err := h.service.DeleteAuthor(authorID); err != nil {
		log.Println(err)
		return serverError(c, fiber.StatusInternalServerError, errs.ErrServer.Error())
	}

	return c.Status(fiber.StatusOK).JSON(response{Message: "deleted"})
}

func (h *Handler) UpdateAuthor(c *fiber.Ctx) error {
	authorID, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		log.Println(err)
		return serverError(c, fiber.StatusBadRequest, errs.ErrInvalidId.Error())
	} else if authorID < 0 {
		return serverError(c, fiber.StatusBadRequest, errs.ErrInvalidId.Error())
	}

	var author entity.Author

	if err := c.BodyParser(&author); err != nil {
		log.Println(err)
		return serverError(c, fiber.StatusInternalServerError, errs.ErrServer.Error())
	}

	author.ID = authorID

	if err := h.service.UpdateAuthor(author); err != nil {

		if errors.Is(err, errs.ErrNothingToUpdate) {
			log.Println(err)
			return serverError(c, fiber.StatusOK, errors.Unwrap(err).Error())
		} else if errors.Is(err, errs.ErrNothingToFound) {
			log.Println(err)
			return serverError(c, fiber.StatusOK, errors.Unwrap(err).Error())
		}

		log.Println(err)
		return serverError(c, fiber.StatusInternalServerError, errs.ErrServer.Error())
	}

	return c.Status(fiber.StatusOK).JSON(response{Message: "updated"})
}

func (h *Handler) AuthorBooks(c *fiber.Ctx) error {
	authorID, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		log.Println(err)
		return serverError(c, fiber.StatusBadRequest, errs.ErrInvalidId.Error())
	} else if authorID < 0 {
		return serverError(c, fiber.StatusBadRequest, errs.ErrInvalidId.Error())
	}

	books, err := h.service.Author.GetAuthorBooks(authorID)

	if err != nil {
		if errors.Is(err, errs.ErrNothingToFound) {
			log.Println(err)
			return serverError(c, fiber.StatusOK, errors.Unwrap(err).Error())
		}

		log.Println(err)
		return serverError(c, fiber.StatusInternalServerError, errs.ErrServer.Error())
	}

	return c.Status(fiber.StatusOK).JSON(books)
}

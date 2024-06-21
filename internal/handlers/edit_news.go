package handlers

import (
	"strconv"
	"zero_api/internal/database"
	"zero_api/internal/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

type NewsUpdateRequest struct {
	Title      *string   `json:"Title" validate:"omitempty,max=255"`
	Content    *string   `json:"Content" validate:"omitempty"`
	Categories *[]uint64 `json:"Categories" validate:"omitempty,dive,gt=0"`
}

func EditNews(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("Id"))
	if err != nil {
		return utils.HandleError(c, fiber.StatusBadRequest, "Invalid news ID")
	}

	var request NewsUpdateRequest
	if err := c.BodyParser(&request); err != nil {
		return utils.HandleError(c, fiber.StatusBadRequest, err.Error())
	}

	if err := validate.Struct(&request); err != nil {
		return utils.HandleError(c, fiber.StatusBadRequest, "Validation failed: "+err.Error())
	}

	tx, err := database.DB.Begin()
	if err != nil {
		return utils.HandleError(c, fiber.StatusInternalServerError, err.Error())
	}

	if request.Title != nil {
		_, err := tx.Exec("UPDATE news SET title = $1 WHERE id = $2", *request.Title, id)
		if err != nil {
			tx.Rollback()
			return utils.HandleError(c, fiber.StatusInternalServerError, err.Error())
		}
	}

	if request.Content != nil {
		_, err := tx.Exec("UPDATE news SET content = $1 WHERE id = $2", *request.Content, id)
		if err != nil {
			tx.Rollback()
			return utils.HandleError(c, fiber.StatusInternalServerError, err.Error())
		}
	}

	if request.Categories != nil {
		_, err := tx.Exec("DELETE FROM news_category WHERE news_id = $1", id)
		if err != nil {
			tx.Rollback()
			return utils.HandleError(c, fiber.StatusInternalServerError, err.Error())
		}

		for _, categoryID := range *request.Categories {
			_, err := tx.Exec("INSERT INTO news_category (news_id, category_id) VALUES ($1, $2)", id, categoryID)
			if err != nil {
				tx.Rollback()
				return utils.HandleError(c, fiber.StatusInternalServerError, err.Error())
			}
		}
	}

	if err := tx.Commit(); err != nil {
		return utils.HandleError(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.HandleSuccess(c, fiber.StatusOK, fiber.Map{
		"message": "News updated successfully",
	})
}

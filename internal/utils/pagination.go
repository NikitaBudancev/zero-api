package utils

import (
	"github.com/gofiber/fiber/v2"

	"strconv"
)

type PaginationMeta struct {
	CurrentPage  int `json:"currentPage"`
	TotalPages   int `json:"totalPages"`
	PageSize     int `json:"pageSize"`
	TotalRecords int `json:"totalRecords"`
}

func GetPaginationParams(c *fiber.Ctx) (int, int, int) {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.Query("limit", "10"))
	if err != nil || limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	return page, limit, offset
}

func CalculateTotalPages(totalRecords, limit int) int {
	return (totalRecords + limit - 1) / limit
}

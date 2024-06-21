package handlers

import (
	"database/sql"
	"zero_api/internal/database"
	"zero_api/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
)

type NewsResponse struct {
	Success bool           `json:"success"`
	News    []NewsDetails  `json:"news"`
	Meta    PaginationMeta `json:"meta"`
}

type NewsDetails struct {
	ID         int64   `json:"id"`
	Title      string  `json:"title"`
	Content    string  `json:"content"`
	Categories []int64 `json:"categories"`
}

type PaginationMeta struct {
	CurrentPage  int `json:"currentPage"`
	TotalPages   int `json:"totalPages"`
	PageSize     int `json:"pageSize"`
	TotalRecords int `json:"totalRecords"`
}

func GetNews(c *fiber.Ctx) error {
	page, limit, offset := utils.GetPaginationParams(c)

	totalRecords, err := database.CountRecords("SELECT COUNT(*) FROM news")
	if err != nil {
		return utils.HandleError(c, fiber.StatusInternalServerError, err.Error())
	}

	totalPages := utils.CalculateTotalPages(totalRecords, limit)

	query := `
		SELECT n.id, n.title, n.content, ARRAY_AGG(nc.category_id) AS categories
		FROM news n
		LEFT JOIN news_category nc ON n.id = nc.news_id
		GROUP BY n.id
		ORDER BY n.id
		LIMIT $1 OFFSET $2
	`

	rows, err := database.DB.Query(query, limit, offset)
	if err != nil {
		return utils.HandleError(c, fiber.StatusInternalServerError, err.Error())
	}
	defer rows.Close()

	var newsList []NewsDetails
	for rows.Next() {
		var news NewsDetails
		var categories []sql.NullInt64

		if err := rows.Scan(&news.ID, &news.Title, &news.Content, pq.Array(&categories)); err != nil {
			return utils.HandleError(c, fiber.StatusInternalServerError, err.Error())
		}

		for _, category := range categories {
			if category.Valid {
				news.Categories = append(news.Categories, category.Int64)
			}
		}

		newsList = append(newsList, news)
	}

	if err := rows.Err(); err != nil {
		return utils.HandleError(c, fiber.StatusInternalServerError, err.Error())
	}

	response := NewsResponse{
		Success: true,
		News:    newsList,
		Meta: PaginationMeta{
			CurrentPage:  page,
			TotalPages:   totalPages,
			PageSize:     limit,
			TotalRecords: totalRecords,
		},
	}

	return utils.HandleSuccess(c, fiber.StatusOK, response)
}

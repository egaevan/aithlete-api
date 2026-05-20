package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Transaction Transaction `json:"transaction"`
	Data        any         `json:"data"`
	Meta        *Meta       `json:"meta,omitempty"`
}

type Transaction struct {
	StatusCode  string `json:"status_code"`
	StatusDesc  string `json:"status_desc"`
	HTTPCode    int    `json:"-"`
}

type Meta struct {
	Total      int `json:"total"`
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalPages int `json:"total_pages"`
}

func Success(c echo.Context, httpCode int, statusCode string, statusDesc string, data any) error {
	return c.JSON(httpCode, Response{
		Transaction: Transaction{
			StatusCode: statusCode,
			StatusDesc: statusDesc,
		},
		Data: data,
	})
}

func SuccessWithMeta(c echo.Context, httpCode int, statusCode string, statusDesc string, data any, meta *Meta) error {
	return c.JSON(httpCode, Response{
		Transaction: Transaction{
			StatusCode: statusCode,
			StatusDesc: statusDesc,
		},
		Data: data,
		Meta: meta,
	})
}

func Error(c echo.Context, httpCode int, statusCode string, statusDesc string) error {
	return c.JSON(httpCode, Response{
		Transaction: Transaction{
			StatusCode: statusCode,
			StatusDesc: statusDesc,
		},
		Data: map[string]any{},
	})
}

func BadRequest(c echo.Context, message string) error {
	return Error(c, http.StatusBadRequest, "40000", message)
}

func InternalServerError(c echo.Context, message string) error {
	return Error(c, http.StatusInternalServerError, "50000", message)
}

func NotFound(c echo.Context, message string) error {
	return Error(c, http.StatusNotFound, "40400", message)
}

func NewMeta(total, page, limit int) *Meta {
	totalPages := (total + limit - 1) / limit
	if totalPages < 1 {
		totalPages = 1
	}
	return &Meta{
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
	}
}

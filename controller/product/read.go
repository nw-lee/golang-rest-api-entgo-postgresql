package product

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/quavious/golang-docker-forum/ent"
)

func ReadMany(ctx *context.Context, conn *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		page, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Bad Request")
		}
		products, err := conn.Product.Query().Offset((page - 1) * perPage).Limit(perPage).All(*ctx)
		if err != nil {
			return c.String(http.StatusBadRequest, "Server ERROR")
		}

		filter := time.Now()
		var filtered []*ent.Product

		for _, product := range products {
			if filter.Before(product.ExpiredAt) {
				filtered = append(filtered, product)
			}
		}
		counted, err := conn.Product.Query().Count(*ctx)
		if err != nil {
			return c.String(http.StatusBadRequest, "Bad Request")
		}
		maxPage := (counted / perPage) + 1
		return c.JSON(http.StatusOK, echo.Map{
			"status":      true,
			"products":    filtered,
			"maxPage":     maxPage,
			"currentPage": page,
		})
	}
}

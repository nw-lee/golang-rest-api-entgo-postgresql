package product

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/quavious/golang-docker-forum/ent"
	"github.com/quavious/golang-docker-forum/ent/product"
)

func SearchMany(ctx *context.Context, conn *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		page, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Bad Request")
		}
		keyword := c.Param("search")

		productBuilder := conn.Product.Query().Order(
			ent.Desc(product.FieldExpiredAt),
		).Where(product.NameContains(keyword))
		products, err := productBuilder.Offset((page - 1) * perPage).Limit(perPage).All(*ctx)
		if err != nil {
			return c.String(http.StatusBadRequest, "Server ERROR")
		}
		counted, err := productBuilder.Count(*ctx)
		if err != nil {
			return c.String(http.StatusBadRequest, "Bad Request")
		}
		maxPage := (counted / perPage) + 1
		return c.JSON(http.StatusOK, echo.Map{
			"status":      true,
			"products":    products,
			"maxPage":     maxPage,
			"currentPage": page,
		})
	}
}

package product

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/quavious/golang-docker-forum/ent"
	"github.com/quavious/golang-docker-forum/ent/category"
	"github.com/quavious/golang-docker-forum/ent/product"
)

func ReadCategory(ctx *context.Context, conn *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		page, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Bad Request")
		}
		keyword, err := strconv.Atoi(c.Param("category"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Bad Request")
		}

		products, err := conn.Category.Query().Where(
			category.IDEQ(keyword),
		).QueryContains().Order(
			ent.Desc(product.FieldExpiredAt),
		).Offset((page - 1) * perPage).Limit(perPage).All(*ctx)

		if err != nil {
			return c.String(http.StatusBadRequest, "Server ERROR")
		}
		counted, err := conn.Category.Query().Where(
			category.IDEQ(keyword),
		).QueryContains().Count(*ctx)
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

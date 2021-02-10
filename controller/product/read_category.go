package product

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/quavious/golang-docker-forum/ent"
	"github.com/quavious/golang-docker-forum/ent/category"
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
		// products, err := conn.Category.Query().Where(category.IDEQ(keyword)).QueryContains().Offset((page - 1) * perPage).Limit(perPage).All(*ctx)

		productBuilder := conn.Product.Query().WithBelongs(
			func(cq *ent.CategoryQuery) {
				cq.Where(category.IDEQ(keyword)).Only(*ctx)
			},
		)
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

package post

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/quavious/golang-docker-forum/ent"
	"github.com/quavious/golang-docker-forum/ent/post"
)

func SearchMany(ctx *context.Context, conn *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		param := c.Param("param")
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Bad Request")
		}
		posts, err := conn.Post.Query().Where(post.TitleContains(param)).Offset((id - 1) * perPage).Limit(perPage).All(*ctx)
		if err != nil {
			return c.String(http.StatusBadRequest, "Bad Request")
		}

		counted, err := conn.Post.Query().Where(post.TitleContains(param)).Count(*ctx)
		if err != nil {
			return c.String(http.StatusBadRequest, "Bad Request")
		}
		maxPage := (counted / perPage) + 1
		return c.JSON(http.StatusOK, echo.Map{
			"status":      true,
			"response":    posts,
			"currentPage": id,
			"maxPage":     maxPage,
		})
	}
}

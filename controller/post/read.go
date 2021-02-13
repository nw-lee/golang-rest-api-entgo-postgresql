package post

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/quavious/golang-docker-forum/ent"
	"github.com/quavious/golang-docker-forum/ent/post"
	"github.com/quavious/golang-docker-forum/model"
)

func ReadMany(ctx *context.Context, conn *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Bad Request")
		}
		posts, err := conn.Post.Query().Order(
			ent.Desc(post.FieldCreatedAt),
		).Offset((id - 1) * perPage).Limit(perPage).All(*ctx)
		if err != nil {
			return c.String(http.StatusBadRequest, "Bad Request")
		}
		counted, err := conn.Post.Query().Count(*ctx)
		if err != nil {
			return c.String(http.StatusBadRequest, "Bad Request")
		}
		maxPage := (counted / perPage) + 1
		for _, post := range posts {
			model.HideInfoPost(post)
		}
		return c.JSON(http.StatusOK, echo.Map{
			"status":      true,
			"response":    posts,
			"currentPage": id,
			"maxPage":     maxPage,
		})
	}
}

func ReadOne(ctx *context.Context, conn *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Println(err)
			return c.String(http.StatusBadRequest, "Bad Request")
		}
		post, err := conn.Post.Query().Where(post.IDEQ(id)).WithComments().Only(*ctx)
		if err != nil {
			log.Println(err)
			return c.String(http.StatusBadRequest, "Bad Request")
		}
		model.HideInfoPost(post)
		for _, cmt := range post.Edges.Comments {
			model.HideInfoComment(cmt)
		}
		return c.JSON(http.StatusOK, echo.Map{
			"status":   true,
			"response": post,
		})
	}
}

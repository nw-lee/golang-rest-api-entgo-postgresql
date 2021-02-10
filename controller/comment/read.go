package comment

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
		postID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Println(err)
			return c.String(http.StatusBadRequest, "Bad Request")
		}
		comments, err := conn.Post.Query().Where(post.IDEQ(postID)).QueryComments().All(*ctx)
		if err != nil {
			log.Println(err)
			return c.String(http.StatusBadRequest, "Bad Request")
		}
		for _, comment := range comments {
			model.HideInfoComment(comment)
		}
		return c.JSON(http.StatusOK, echo.Map{
			"status":   true,
			"response": comments,
		})
	}
}

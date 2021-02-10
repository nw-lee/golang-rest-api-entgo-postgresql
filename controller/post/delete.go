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

func DeleteOne(ctx *context.Context, conn *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		p := new(model.Post)
		if err := c.Bind(p); err != nil {
			log.Println(err)
			return c.String(http.StatusBadRequest, "Bad Request")
		}
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Println(err)
			return c.String(http.StatusBadRequest, "Bad Request")
		}

		hash, err := conn.Post.Query().Where(post.IDEQ(id)).Select(post.FieldPassword).Only(*ctx)
		if err != nil {
			log.Println(err)
			return c.String(http.StatusBadRequest, "Bad Request")
		}
		err = p.PasswordCheck(hash.Password)
		if err != nil {
			log.Println(err)
			return c.String(http.StatusBadRequest, "Bad Request")
		}

		err = conn.Post.DeleteOneID(id).Exec(*ctx)
		if err != nil {
			log.Println(err)
			return c.String(http.StatusBadRequest, "Bad Request")
		}
		return c.JSON(http.StatusOK, echo.Map{
			"status":   true,
			"response": "Post Updated",
		})
	}
}

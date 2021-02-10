package comment

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/quavious/golang-docker-forum/ent"
	"github.com/quavious/golang-docker-forum/model"
)

func CreateOne(ctx *context.Context, conn *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		postID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Println(err)
			return c.String(http.StatusBadRequest, "Bad Request")
		}
		cmt := new(model.Post)
		if err := c.Bind(cmt); err != nil {
			log.Println(err)
			return c.String(http.StatusBadRequest, "Bad Request")
		}
		hash, err := cmt.PasswordHash()
		if err != nil {
			log.Println(err)
			return c.String(http.StatusBadRequest, "Bad Request")
		}
		userIP := c.RealIP()
		_, err = conn.Comment.Create().SetPostID(postID).SetUsername(cmt.Username).SetUserIP(userIP).SetContent(cmt.Content).SetPassword(hash).SetCreatedAt(time.Now()).Save(*ctx)
		if err != nil {
			log.Println(err)
			return c.String(http.StatusBadRequest, "Bad Request")
		}
		return c.JSON(http.StatusOK, echo.Map{
			"status":   true,
			"response": "A Comment Created",
		})
	}
}

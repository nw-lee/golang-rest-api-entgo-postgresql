package post

import (
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/quavious/golang-docker-forum/ent"
	"github.com/quavious/golang-docker-forum/model"
)

func CreateOne(ctx *context.Context, conn *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		p := new(model.Post)
		if err := c.Bind(p); err != nil {
			log.Println(err)
			return c.String(http.StatusBadRequest, "Bad Request")
		}
		hash, err := p.PasswordHash()
		if err != nil {
			log.Println(err)
			return c.String(http.StatusBadRequest, "Bad Request")
		}
		userIP := c.RealIP()
		_, err = conn.Post.Create().SetTitle(p.Title).SetContent(p.Content).SetPassword(hash).SetUsername(p.Username).SetUserIP(userIP).Save(*ctx)
		if err != nil {
			log.Println(err)
			return c.String(http.StatusBadRequest, "Bad Request")
		}
		return c.JSON(http.StatusOK, echo.Map{
			"status":   true,
			"response": "Post Created",
		})
	}
}

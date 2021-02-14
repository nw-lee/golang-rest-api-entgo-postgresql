package category

import (
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/quavious/golang-docker-forum/ent"
)

func CategoryMenu(ctx *context.Context, conn *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		menu, err := conn.Category.Query().All(*ctx)
		if err != nil {
			log.Println(err)
			return c.String(http.StatusBadRequest, "Bad Request")
		}
		return c.JSON(http.StatusOK, echo.Map{
			"status": true,
			"menu":   menu,
		})
	}
}

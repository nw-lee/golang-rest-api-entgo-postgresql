package comment

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/quavious/golang-docker-forum/ent"
	"github.com/quavious/golang-docker-forum/ent/comment"
	"github.com/quavious/golang-docker-forum/model"
)

func DeleteOne(ctx *context.Context, conn *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		commentID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Println(err)
			return c.String(http.StatusBadRequest, "Bad Request")
		}
		pwd := new(model.Post)
		if err := c.Bind(pwd); err != nil {
			log.Println(err)
			return c.String(http.StatusBadRequest, "Bad Request")
		}
		cmt, err := conn.Comment.Query().Where(comment.IDEQ(commentID)).Select(comment.FieldPassword).Only(*ctx)
		if err != nil {
			log.Println(err)
			return c.String(http.StatusBadRequest, "Bad Request")
		}
		err = pwd.PasswordCheck(cmt.Password)
		if err != nil {
			log.Println(err)
			return c.String(http.StatusBadRequest, "Bad Request")
		}
		err = conn.Comment.DeleteOneID(commentID).Exec(*ctx)
		if err != nil {
			log.Println(err)
			return c.String(http.StatusBadRequest, "Bad Request")
		}
		return c.JSON(http.StatusOK, echo.Map{
			"status":   true,
			"response": "The Comment Removed",
		})
	}
}

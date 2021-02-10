package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"github.com/quavious/golang-docker-forum/controller/comment"
	"github.com/quavious/golang-docker-forum/controller/post"
	"github.com/quavious/golang-docker-forum/controller/product"
	"github.com/quavious/golang-docker-forum/ent"
)

func main() {
	godotenv.Load()
	conn, err := ent.Open("postgres", "host="+os.Getenv("DB_HOST")+" port="+os.Getenv("DB_PORT")+" user="+os.Getenv("DB_USER")+" dbname="+os.Getenv("DB_NAME")+" password="+os.Getenv("DB_PASSWORD")+" sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()
	ctx := context.Background()
	if err := conn.Schema.Create(ctx); err != nil {
		log.Fatalln(err)
	}

	app := echo.New()
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	app.GET("/post/:id", post.ReadMany(&ctx, conn))
	app.GET("/post/search/title/:param/:id", post.SearchMany(&ctx, conn))
	app.GET("/post/id/:id", post.ReadOne(&ctx, conn))
	app.POST("/post", post.CreateOne(&ctx, conn))
	app.POST("/post/id/:id/update", post.UpdateOne(&ctx, conn))
	app.POST("/post/id/:id/delete", post.DeleteOne(&ctx, conn))

	app.GET("/comment/post/:id", comment.ReadMany(&ctx, conn))
	app.POST("/comment/post/:id/create", comment.CreateOne(&ctx, conn))
	app.POST("/comment/:id/delete", comment.DeleteOne(&ctx, conn))

	app.GET("/product/:id", product.ReadMany(&ctx, conn))
	app.GET("/product/category/:category/:id", product.ReadCategory(&ctx, conn))
	app.GET("/product/search/:search/:id", product.SearchMany(&ctx, conn))

	app.GET("/check", func(c echo.Context) error {
		return c.String(http.StatusOK, "Check Route")
	})

	app.Logger.Fatal(app.Start("0.0.0.0:5000"))
}

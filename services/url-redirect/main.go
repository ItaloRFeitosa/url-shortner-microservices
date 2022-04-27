package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/italorfeitosa/url-shortner-mvp/libs/encoder"
	"github.com/italorfeitosa/url-shortner-mvp/libs/environment"
	"github.com/italorfeitosa/url-shortner-mvp/services/url-redirect/event_handler"
	"github.com/italorfeitosa/url-shortner-mvp/services/url-redirect/handler"
	"github.com/italorfeitosa/url-shortner-mvp/services/url-redirect/infra"
	"github.com/italorfeitosa/url-shortner-mvp/services/url-redirect/service"
)

func main() {
	repo := infra.NewDynamoDBRepository()
	e := encoder.NewHashEncoder(environment.HashEncodeSalt)
	s := service.NewService(repo, e)

	a := event_handler.NewAddShortURLHandler(s)
	b := infra.NewSQSBroker()
	b.Subscribe("short_url.created", a)
	go b.Init(context.Background())

	r := gin.Default()
	h := handler.NewHandler(s)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "it's running",
		})
	})

	r.GET("/:slug", h.Redirect)

	r.Run(":3004")
}

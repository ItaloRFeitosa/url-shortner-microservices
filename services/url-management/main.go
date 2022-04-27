package main

import (
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/italorfeitosa/url-shortner-mvp/libs/encoder"
	"github.com/italorfeitosa/url-shortner-mvp/libs/environment"
	"github.com/italorfeitosa/url-shortner-mvp/services/url-management/handler"
	"github.com/italorfeitosa/url-shortner-mvp/services/url-management/infra"
	"github.com/italorfeitosa/url-shortner-mvp/services/url-management/service"
)

func main() {
	r := gin.Default()
	SetupValidator()
	h := SetupHandler()

	r.Use(UserMiddleware())
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "it's running",
		})
	})

	r.POST("/urls", h.ShortenURL)
	r.GET("/urls", h.ListUserURLs)
	r.GET("/urls/:id", h.GetURLStats)

	r.Run(":3003")
}

func UserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.GetHeader("user-id")
		if userId == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing user-id header"})
			c.Abort()
			return
		}
		traceId := c.GetHeader("trace-id")

		c.Set("user-id", userId)
		c.Set("trace-id", traceId)

		c.Next()
	}
}

func SetupValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}
}

func SetupHandler() *handler.Handler {
	repo := infra.NewMySQLRepository()
	encoder := encoder.NewHashEncoder(environment.HashEncodeSalt)
	publisher := infra.NewSNSBroker()
	service := service.NewService(repo, encoder, publisher)
	return handler.NewHandler(service)
}

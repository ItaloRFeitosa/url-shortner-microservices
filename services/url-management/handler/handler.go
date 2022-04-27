package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/italorfeitosa/url-shortner-mvp/libs/http_helper"
	"github.com/italorfeitosa/url-shortner-mvp/services/url-management/dto"
	"github.com/italorfeitosa/url-shortner-mvp/services/url-management/service"
)

type Handler struct {
	s service.Service
}

func NewHandler(s service.Service) *Handler {
	return &Handler{s}
}

func (h *Handler) ShortenURL(c *gin.Context) {
	d := dto.ShortenURLDTO{}

	if err := c.ShouldBindJSON(&d); err != nil {
		http_helper.ValidationErrorResponse(c, err)
		return
	}

	d.UserID = c.MustGet("user-id").(string)

	url, err := h.s.ShortenURL(c.Request.Context(), d)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, url)
}

func (h *Handler) ListUserURLs(c *gin.Context) {
	urls, err := h.s.ListUserURLs(c.Request.Context(), c.MustGet("user-id").(string))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, urls)
}

func (h *Handler) GetURLStats(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "needs to be implemented"})
}

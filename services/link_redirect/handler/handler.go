package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/italorfeitosa/url-shortner-mvp/services/link_redirect/dto"
	"github.com/italorfeitosa/url-shortner-mvp/services/link_redirect/service"
)

type Handler struct {
	s service.Service
}

func NewHandler(s service.Service) *Handler {
	return &Handler{s}
}

func (h *Handler) Redirect(c *gin.Context) {
	dto := dto.RedirectDTO{}

	if err := c.ShouldBindUri(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	url, err := h.s.GetRedirectURL(c.Request.Context(), dto.Slug)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusFound, url)
}

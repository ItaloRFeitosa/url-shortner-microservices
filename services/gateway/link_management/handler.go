package link_management

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	pb "github.com/italorfeitosa/url-shortner-mvp/services/link_management/external/grpc/proto"
)

type Handler struct {
	c pb.LinkManagementServiceClient
}

func NewHandler(c pb.LinkManagementServiceClient) *Handler {
	return &Handler{c}
}

type CreateLinkRequest struct {
	RedirectTo  string `json:"redirect_to"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type LinkData struct {
	Slug        string `json:"slug"`
	RedirectTo  string `json:"redirect_to"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}

func (h *Handler) CreateLink(ctx *gin.Context) {
	req := CreateLinkRequest{}

	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.c.CreateLink(ctx.Request.Context(), &pb.CreateLinkRequest{
		WorkspaceId: ctx.MustGet("workspace-id").(string),
		RedirectTo:  req.RedirectTo,
		Title:       req.Title,
		Description: req.Description,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	if res.Error {
		ctx.AbortWithError(http.StatusBadRequest, errors.New(res.Message))
		return
	}

	ctx.JSON(http.StatusCreated, LinkData{
		Slug:        res.Data.Slug,
		Title:       res.Data.Title,
		Description: res.Data.Description,
		CreatedAt:   res.Data.CreatedAt,
	})
}

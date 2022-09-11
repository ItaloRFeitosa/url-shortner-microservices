package grpc

import (
	"context"
	"time"

	"github.com/italorfeitosa/url-shortner-mvp/services/link_management/core"
	pb "github.com/italorfeitosa/url-shortner-mvp/services/link_management/external/grpc/proto"
)

type Handler struct {
	pb.UnimplementedLinkManagementServiceServer
	u core.LinkUseCase
}

func NewHandler(usecase core.LinkUseCase) *Handler {
	return &Handler{
		u: usecase,
	}
}

func (h *Handler) CreateLink(ctx context.Context, req *pb.CreateLinkRequest) (*pb.LinkReply, error) {
	link, err := h.u.Create(ctx, core.CreateLinkInput{
		WorkspaceID: req.WorkspaceId,
		RedirectTo:  req.RedirectTo,
		Title:       req.Title,
		Description: req.Description,
	})

	if err != nil {
		return &pb.LinkReply{
			Message: err.Error(),
			Data:    nil,
			Error:   true,
		}, nil
	}

	return &pb.LinkReply{
		Message: "link created succesfully.",
		Data: &pb.LinkData{
			Id:          int64(link.ID),
			WorkspaceId: link.WorkspaceID,
			Slug:        link.Slug,
			RedirectTo:  link.RedirectTo,
			Title:       link.Title,
			Description: link.Description,
			CreatedAt:   link.CreatedAt.Format(time.RFC3339),
		},
		Error: false,
	}, nil
}

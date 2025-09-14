package shortener

import (
	"context"

	"github.com/EmreZURNACI/url-shortener/app"
	"github.com/EmreZURNACI/url-shortener/domain"
)

type GetURLRequest struct {
	Address domain.Address `json:"address"`
}
type GetURLResponse struct {
	Address domain.Address `json:"address"`
}
type GetURLHandler struct {
	repository app.Repository
}

func NewGetURLHandler(repo app.Repository) *GetURLHandler {
	return &GetURLHandler{
		repository: repo,
	}
}

func (h *GetURLHandler) Handle(ctx context.Context, req *GetURLRequest) (*GetURLResponse, error) {

	address, err := h.repository.GetURL(ctx, req.Address)
	if err != nil {
		return nil, err
	}

	return &GetURLResponse{
		Address: *address,
	}, nil

}

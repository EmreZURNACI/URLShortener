package shortener

import (
	"context"

	"github.com/EmreZURNACI/url-shortener/app"
	"github.com/EmreZURNACI/url-shortener/domain"
)

type GetShortURLRequest struct {
	Address domain.Address `json:"address"`
}
type GetShortURLResponse struct {
	Address domain.Address `json:"address"`
}
type GetShortURLHandler struct {
	repository app.Repository
}

func NewGetShortURLHandler(repo app.Repository) *GetShortURLHandler {
	return &GetShortURLHandler{
		repository: repo,
	}
}

func (h *GetShortURLHandler) Handle(ctx context.Context, req *GetShortURLRequest) (*GetShortURLResponse, error) {

	address, err := h.repository.GetShortURL(ctx, req.Address)
	if err != nil {
		return nil, err
	}

	return &GetShortURLResponse{
		Address: *address,
	}, nil

}

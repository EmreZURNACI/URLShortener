package shortener

import (
	"context"
	"crypto/rand"
	"math/big"

	"github.com/EmreZURNACI/url-shortener/app"
	"github.com/EmreZURNACI/url-shortener/domain"
	"github.com/google/uuid"
)

type CreateURLRequest struct {
	Address domain.Address `json:"address"`
}
type CreateURLResponse struct {
	ShortURL *string `json:"short_url"`
}
type CreateURLHandler struct {
	repository app.Repository
}

func NewCreateURLHandler(repo app.Repository) *CreateURLHandler {
	return &CreateURLHandler{
		repository: repo,
	}
}

func (h *CreateURLHandler) Handle(ctx context.Context, req *CreateURLRequest) (*CreateURLResponse, error) {

	uuid, err := uuid.NewV7()

	if err != nil {
		return nil, err
	}

	shortLink, err := randomBase62()
	if err != nil {
		return nil, err
	}
	address := domain.Address{
		UUID:     uuid.String(),
		URL:      req.Address.URL,
		ShortURL: shortLink,
	}

	url, err := h.repository.CreateURL(ctx, address)
	if err != nil {
		return nil, err
	}

	return &CreateURLResponse{
		ShortURL: url,
	}, nil

}

const base62chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomBase62() (string, error) {
	b := make([]byte, 10)
	for i := 0; i < 10; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(base62chars))))
		if err != nil {
			return "", err
		}
		b[i] = base62chars[num.Int64()]
	}
	return string(b), nil
}

package app

import (
	"context"

	"github.com/EmreZURNACI/url-shortener/domain"
)

type Repository interface {
	GetURL(ctx context.Context, address domain.Address) (*domain.Address, error)
	GetShortURL(ctx context.Context, address domain.Address) (*domain.Address, error)
	CreateURL(ctx context.Context, address domain.Address) (*string, error)
}

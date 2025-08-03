package typeservice

import (
	"log/slog"
	"passwordbot/internal/storage/db"
)

type TypeService struct {
	repo db.TypeCred
	log  *slog.Logger
}

const op = "services.type"

func New(
	repo db.TypeCred,
	log *slog.Logger,
) *TypeService {
	return &TypeService{
		repo: repo,
		log:  log,
	}
}

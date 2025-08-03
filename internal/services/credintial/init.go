package credintialservice

import (
	"log/slog"
	"passwordbot/internal/storage/db"
)

type CredintialService struct {
	log  *slog.Logger
	repo db.Credintial
}

const op = "services.credintial"

func New(
	repo db.Credintial,
	log *slog.Logger,
) *CredintialService {
	return &CredintialService{
		repo: repo,
		log:  log,
	}
}

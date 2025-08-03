package credintialstorage

import (
	"log/slog"

	"gorm.io/gorm"
)

type CredintialStorage struct {
	db  *gorm.DB
	log *slog.Logger
}

const op = "storage.db.sqlite.credintial"

func New(
	db *gorm.DB,
	log *slog.Logger,
) *CredintialStorage {
	return &CredintialStorage{
		db:  db,
		log: log,
	}
}

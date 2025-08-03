package typestorage

import (
	"log/slog"

	"gorm.io/gorm"
)

type TypeStorage struct {
	db  *gorm.DB
	log *slog.Logger
}

const op = "storage.db.sqlite.type"

func New(
	db *gorm.DB,
	log *slog.Logger,
) *TypeStorage {
	return &TypeStorage{
		db:  db,
		log: log,
	}
}

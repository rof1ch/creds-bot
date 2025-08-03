package db

import (
	"log/slog"
	"passwordbot/internal/domain/models"
	credintialstorage "passwordbot/internal/storage/db/sqlite/credintial"
	typestorage "passwordbot/internal/storage/db/sqlite/type"

	"gorm.io/gorm"
)

type Credintial interface {
	ById(credId uint) (models.Credintial, error)
	Create(input models.Credintial) (uint, error)
	Delete(credId uint) error
	List(typeId uint) ([]models.Credintial, error)
	ByUserId(userId int64) ([]models.Credintial, error)
}

type TypeCred interface {
	Create(input models.Type) (uint, error)
	Delete(credId uint) error
	List(userId int64) ([]models.Type, error)
    ById(typeID uint) (models.Type, error)
}

type Storage struct {
	Credintial
	TypeCred
}

func New(
	db *gorm.DB,
	log *slog.Logger,
) *Storage {
	return &Storage{
		Credintial: credintialstorage.New(db, log),
		TypeCred:   typestorage.New(db, log),
	}
}

package services

import (
	"log/slog"
	"passwordbot/internal/domain/dto"
	"passwordbot/internal/domain/models"
	credintialservice "passwordbot/internal/services/credintial"
	typeservice "passwordbot/internal/services/type"
	"passwordbot/internal/storage/db"
)

type Credintial interface {
	Create(input dto.CredintialInput) error
	ByTypeId(typeId uint) ([]models.Credintial, error)
	ById(credId uint, key string) (dto.Credintial, error)
    ByUserId(userId int64) ([]models.Credintial, error)
    Delete(credId uint) error
}

type TypeCred interface {
	Create(name, icon string, userId int64) error
	List(userID int64) ([]models.Type, error)
	ById(typeId uint) (models.Type, error)
	Delete(typeId uint) error
}

type Services struct {
	TypeCred
	Credintial
}

func New(
	log *slog.Logger,
	repo *db.Storage,
) *Services {
	return &Services{
		TypeCred:   typeservice.New(repo.TypeCred, log),
		Credintial: credintialservice.New(repo.Credintial, log),
	}
}

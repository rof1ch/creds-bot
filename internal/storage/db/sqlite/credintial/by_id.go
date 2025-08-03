package credintialstorage

import (
	"errors"
	"passwordbot/internal/domain/models"
	"passwordbot/pkg/logger/sl"

	"gorm.io/gorm"
)

func (r *CredintialStorage) ById(credId uint) (models.Credintial, error) {
	log := r.log.With(sl.Op(op, "ById"))

	var cred models.Credintial

	err := r.db.
		Where("id = ?", credId).
        Preload("Type").
		First(&cred).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Credintial{}, ErrCredNotFound
		}
		log.Error(ErrGetByIdCred.Error(), sl.Err(err))
		return models.Credintial{}, ErrGetByIdCred
	}

	return cred, nil
}

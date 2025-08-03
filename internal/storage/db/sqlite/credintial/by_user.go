package credintialstorage

import (
	"passwordbot/internal/domain/models"
	"passwordbot/pkg/logger/sl"
)

func (r *CredintialStorage) ByUserId(userId int64) ([]models.Credintial, error) {
	log := r.log.With(sl.Op(op, "ByUserId"))

	var creds []models.Credintial

	err := r.db.Where("user_id = ?", userId).Find(&creds).Error
	if err != nil {
		log.Error(ErrGetListByTypeId.Error(), sl.Err(err))
		return nil, ErrGetListByTypeId
	}

	return creds, nil
}

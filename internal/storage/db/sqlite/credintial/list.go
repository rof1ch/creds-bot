package credintialstorage

import (
	"passwordbot/internal/domain/models"
	"passwordbot/pkg/logger/sl"
)

func (r *CredintialStorage) List(typeId uint) ([]models.Credintial, error) {
	log := r.log.With(sl.Op(op, "List"))

	var creds []models.Credintial

	err := r.db.Where("type_id = ?", typeId).Find(&creds).Error
	if err != nil {
		log.Error(ErrGetListByTypeId.Error(), sl.Err(err))
		return nil, ErrGetListByTypeId
	}

	return creds, nil
}
